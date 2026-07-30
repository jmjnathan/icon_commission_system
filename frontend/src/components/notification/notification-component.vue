<script setup lang="ts">
import { ref, computed } from "vue";
import { AlertTriangle, X } from "lucide-vue-next";

interface UrgentItem {
  id: number;
  clientName: string;
  subject: string;
  daysLeft: number;
  status: string;
}

const props = defineProps<{
  items: UrgentItem[];
}>();

const dismissedIds = ref<Set<number>>(new Set());

const visibleItems = computed(() =>
  props.items.filter((item) => !dismissedIds.value.has(item.id))
);

function dismiss(id: number) {
  dismissedIds.value.add(id);
  dismissedIds.value = new Set(dismissedIds.value);
}

function dismissAll() {
  dismissedIds.value = new Set(props.items.map((i) => i.id));
}
</script>

<template>
  <div
    v-if="visibleItems.length > 0"
    class="bg-white border border-[#D9534F]/30 rounded-xl p-4 mb-8">
    <div class="flex items-center justify-between mb-3">
      <h3 class="flex items-center gap-2 text-sm font-semibold text-[#B23A32]">
        Reminder
      </h3>
      <button
        v-if="visibleItems.length > 1"
        @click="dismissAll"
        class="text-xs text-[#8A7A5C] hover:text-[#3A2E1F] transition">
        Tutup semua
      </button>
    </div>

    <ul class="space-y-2">
      <li
        v-for="item in visibleItems"
        :key="item.id"
        class="flex items-start justify-between gap-3 text-sm text-[#3A2E1F] bg-[#D9534F]/5 rounded-lg px-3 py-2">
        <div class="flex items-start gap-2">
          <AlertTriangle :size="14" class="mt-0.5 shrink-0 text-[#B23A32]" />
          <span>
            Komisi ({{ item.clientName }} - {{ item.subject }}) tersisa
            <strong>{{ item.daysLeft }} hari</strong> lagi. Status saat ini:
            <span class="font-bold text-red-700">{{
              item.status.toUpperCase()
            }}</span>
          </span>
        </div>
        <button
          @click="dismiss(item.id)"
          class="shrink-0 w-5 h-5 flex items-center justify-center rounded-full text-[#B23A32]/60 hover:text-[#B23A32] hover:bg-[#B23A32]/10 transition"
          title="Tutup notifikasi ini">
          <X :size="14" />
        </button>
      </li>
    </ul>
  </div>
</template>
