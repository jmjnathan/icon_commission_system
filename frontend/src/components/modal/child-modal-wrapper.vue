<script setup lang="ts">
import { watch, onUnmounted } from "vue";
import { X } from "lucide-vue-next";

const props = withDefaults(
  defineProps<{
    visible: boolean;
    headerTitle?: string;
    widthClass?: string;
    dismissableMask?: boolean;
    position?: "center" | "top" | "bottom";
    maximized?: boolean;
  }>(),
  {
    headerTitle: "",
    widthClass: "w-full max-w-2xl",
    dismissableMask: true,
    position: "center",
    maximized: false,
  }
);

const emit = defineEmits<{
  hide: [];
}>();

function close() {
  emit("hide");
}

function onMaskClick() {
  if (props.dismissableMask) close();
}
watch(
  () => props.visible,
  (val) => {
    document.body.style.overflow = val ? "hidden" : "";
  },
  { immediate: true }
);

onUnmounted(() => {
  document.body.style.overflow = "";
});

const positionClass = {
  center: "items-center justify-center",
  top: "items-start justify-center pt-12",
  bottom: "items-end justify-center pb-12",
};
</script>

<template>
  <Teleport to="body">
    <Transition name="modal-fade">
      <div
        v-if="visible"
        class="fixed inset-0 z-50 flex"
        :class="positionClass[position]">
        <!-- Mask -->
        <div class="absolute inset-0 bg-black/40" @click="onMaskClick"></div>

        <!-- Content -->
        <div
          class="relative bg-white rounded-xl shadow-xl flex flex-col overflow-hidden"
          :class="
            maximized
              ? 'w-full h-full rounded-none'
              : `${widthClass} max-h-[85vh]`
          ">
          <!-- Header -->
          <div
            class="flex items-center justify-between px-6 py-4 border-b border-[#E5D9BF] shrink-0">
            <h2 class="text-lg font-semibold text-[#3A2E1F]">
              {{ headerTitle }}
            </h2>
            <div class="flex items-center gap-3">
              <slot name="rightComponent" />
              <button
                @click="close"
                class="w-8 h-8 flex items-center justify-center rounded-lg text-[#8A7A5C] hover:bg-[#E5D9BF]/50 hover:text-[#3A2E1F] transition">
                <X :size="18" />
              </button>
            </div>
          </div>

          <!-- Body -->
          <div class="flex-1 overflow-y-auto px-6 py-5">
            <slot />
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.2s ease;
}
.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
</style>
