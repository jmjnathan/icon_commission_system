<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, nextTick, watch } from "vue";
import { MoreVertical, X } from "lucide-vue-next";
import type { Component } from "vue";

interface MenuItem {
  label: string;
  icon?: Component;
  access?: boolean;
  danger?: boolean;
  command?: () => void;
}

const props = defineProps<{
  items: MenuItem[];
}>();

const isOpen = ref(false);
const triggerRef = ref<HTMLButtonElement | null>(null);
const menuRef = ref<HTMLDivElement | null>(null);
const menuPosition = ref({ top: 0, left: 0 });
let closeTimeout: ReturnType<typeof setTimeout> | null = null;

const visibleItems = computed(() =>
  props.items.filter((item) =>
    Object.hasOwn(item, "access") ? item.access : true
  )
);

async function openMenu() {
  if (closeTimeout) clearTimeout(closeTimeout);
  isOpen.value = true;
  await nextTick();
  calculatePosition();
}

function closeMenu() {
  if (closeTimeout) clearTimeout(closeTimeout);
  isOpen.value = false;
}

function calculatePosition() {
  if (!triggerRef.value) return;

  const rect = triggerRef.value.getBoundingClientRect();
  const menuWidth = 176;
  const spaceOnRight = window.innerWidth - rect.right;
  const spaceOnLeft = rect.left;

  let left: number;
  if (spaceOnRight >= menuWidth) {
    left = rect.left + window.scrollX;
  } else if (spaceOnLeft >= menuWidth) {
    left = rect.right + window.scrollX - menuWidth;
  } else {
    left = window.scrollX + window.innerWidth - menuWidth - 8;
  }

  menuPosition.value = {
    top: rect.bottom + window.scrollY + 4,
    left,
  };
}

function handleItemClick(item: MenuItem) {
  closeMenu();
  item.command?.();
}

function handleClickOutside(event: MouseEvent) {
  if (
    isOpen.value &&
    triggerRef.value &&
    menuRef.value &&
    !triggerRef.value.contains(event.target as Node) &&
    !menuRef.value.contains(event.target as Node)
  ) {
    closeMenu();
  }
}

const BUFFER = 12;

function isPointInRect(x: number, y: number, rect: DOMRect) {
  return (
    x >= rect.left - BUFFER &&
    x <= rect.right + BUFFER &&
    y >= rect.top - BUFFER &&
    y <= rect.bottom + BUFFER
  );
}

function handleMouseMove(event: MouseEvent) {
  if (!isOpen.value || !triggerRef.value || !menuRef.value) return;

  const triggerRect = triggerRef.value.getBoundingClientRect();
  const menuRect = menuRef.value.getBoundingClientRect();

  const insideTrigger = isPointInRect(
    event.clientX,
    event.clientY,
    triggerRect
  );
  const insideMenu = isPointInRect(event.clientX, event.clientY, menuRect);

  if (insideTrigger || insideMenu) {
    if (closeTimeout) clearTimeout(closeTimeout);
  } else {
    scheduleClose();
  }
}

function scheduleClose() {
  if (closeTimeout) clearTimeout(closeTimeout);
  closeTimeout = setTimeout(() => {
    isOpen.value = false;
  }, 200);
}

watch(isOpen, (open) => {
  if (open) {
    document.addEventListener("mousemove", handleMouseMove);
  } else {
    document.removeEventListener("mousemove", handleMouseMove);
  }
});

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
  window.addEventListener("scroll", closeMenu, true);
});

onUnmounted(() => {
  if (closeTimeout) clearTimeout(closeTimeout);
  document.removeEventListener("click", handleClickOutside);
  document.removeEventListener("mousemove", handleMouseMove);
  window.removeEventListener("scroll", closeMenu, true);
});
</script>

<template>
  <div v-if="visibleItems.length > 0" class="inline-block">
    <button
      ref="triggerRef"
      @mouseenter="openMenu"
      @click="openMenu"
      type="button"
      class="w-9 h-9 flex items-center justify-center rounded-full text-[#8A7A5C] hover:bg-[#C9A24B]/15 hover:text-[#8A6D1F] transition">
      <X v-if="isOpen" :size="18" />
      <MoreVertical v-else :size="18" />
    </button>

    <Teleport to="body">
      <div
        v-if="isOpen"
        ref="menuRef"
        class="fixed w-44 bg-white rounded-lg shadow-lg border border-[#E5D9BF] py-1 z-50"
        :style="{
          top: `${menuPosition.top}px`,
          left: `${menuPosition.left}px`,
        }">
        <button
          v-for="(item, index) in visibleItems"
          :key="index"
          @click="handleItemClick(item)"
          type="button"
          class="w-full flex items-center gap-2 px-4 py-2 text-sm text-[#3A2E1F] hover:bg-[#F5EFE3] text-left transition"
          :class="item.danger ? 'text-[#B23A32] hover:bg-[#B23A32]/10' : ''">
          <component :is="item.icon" v-if="item.icon" :size="16" />
          <span>{{ item.label }}</span>
        </button>
      </div>
    </Teleport>
  </div>
</template>
