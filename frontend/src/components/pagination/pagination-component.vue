<script setup lang="ts">
import { ChevronLeft, ChevronRight } from "lucide-vue-next";

const props = defineProps<{
  currentPage: number;
  totalItems: number;
  perPage: number;
}>();

const emit = defineEmits<{
  "update:currentPage": [value: number];
  "update:perPage": [value: number];
}>();

function goPrev() {
  if (props.currentPage > 1) emit("update:currentPage", props.currentPage - 1);
}

function goNext() {
  const maxPage = Math.ceil(props.totalItems / props.perPage);
  if (props.currentPage < maxPage)
    emit("update:currentPage", props.currentPage + 1);
}

function rangeLabel() {
  if (props.totalItems === 0) return "0 - 0 dari 0";
  const start = (props.currentPage - 1) * props.perPage + 1;
  const end = Math.min(props.currentPage * props.perPage, props.totalItems);
  return `${start} - ${end} dari ${props.totalItems}`;
}
</script>

<template>
  <div
    class="flex items-center justify-end gap-6 px-4 py-3 text-sm text-[#6B5D45]">
    <div class="flex items-center gap-2">
      <span>Per halaman:</span>
      <select
        :value="perPage"
        @change="
          emit(
            'update:perPage',
            Number(($event.target as HTMLSelectElement).value)
          )
        "
        class="border border-[#D9CBB0] rounded-lg px-2 py-1 bg-white">
        <option :value="10">10</option>
        <option :value="25">25</option>
        <option :value="50">50</option>
      </select>
    </div>
    <span>{{ rangeLabel() }}</span>
    <div class="flex items-center gap-1">
      <button
        @click="goPrev"
        :disabled="currentPage === 1"
        class="w-7 h-7 flex items-center justify-center rounded-lg hover:bg-[#F5EFE3] disabled:opacity-30 disabled:cursor-not-allowed transition">
        <ChevronLeft :size="16" />
      </button>
      <button
        @click="goNext"
        :disabled="currentPage * perPage >= totalItems"
        class="w-7 h-7 flex items-center justify-center rounded-lg hover:bg-[#F5EFE3] disabled:opacity-30 disabled:cursor-not-allowed transition">
        <ChevronRight :size="16" />
      </button>
    </div>
  </div>
</template>
