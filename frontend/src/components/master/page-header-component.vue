<script setup lang="ts">
import { Plus, Info } from "lucide-vue-next";

defineProps<{
  title: string;
  searchLabel: string;
  searchPlaceholder?: string;
}>();

const search = defineModel<string>("search", { default: "" });
const statusFilter = defineModel<"all" | "active" | "inactive">(
  "statusFilter",
  { default: "all" }
);

const emit = defineEmits<{
  create: [];
}>();
</script>

<template>
  <div class="bg-white border border-[#E5D9BF] rounded-xl p-6 mb-6">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-xl font-semibold text-[#3A2E1F] flex items-center gap-2">
        {{ title }}
      </h1>
      <button
        @click="emit('create')"
        class="flex items-center gap-2 px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition">
        <Plus :size="16" />
        Tambah
      </button>
    </div>

    <div class="flex items-end justify-between gap-6 flex-wrap">
      <div class="flex-1 min-w-60">
        <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
          {{ searchLabel }}
        </label>
        <input
          v-model="search"
          type="text"
          :placeholder="searchPlaceholder || 'Ketik untuk mencari...'"
          class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition" />
      </div>

      <div>
        <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
          >Status</label
        >
        <div class="flex items-center gap-4 h-10.5">
          <label
            class="flex items-center gap-2 text-sm text-[#3A2E1F] cursor-pointer">
            <input
              type="radio"
              value="all"
              v-model="statusFilter"
              class="accent-[#C9A24B]" />
            Semua
          </label>
          <label
            class="flex items-center gap-2 text-sm text-[#3A2E1F] cursor-pointer">
            <input
              type="radio"
              value="active"
              v-model="statusFilter"
              class="accent-[#C9A24B]" />
            Aktif
          </label>
          <label
            class="flex items-center gap-2 text-sm text-[#3A2E1F] cursor-pointer">
            <input
              type="radio"
              value="inactive"
              v-model="statusFilter"
              class="accent-[#C9A24B]" />
            Non-Aktif
          </label>
        </div>
      </div>
    </div>

    <div
      class="mt-5 flex items-start gap-3 bg-[#F5EFE3] border-l-4 border-[#C9A24B] rounded-lg px-4 py-3">
      <Info :size="16" class="text-[#8A6D1F] mt-0.5 shrink-0" />
      <p class="text-sm text-[#6B5D45]">
        Data akan otomatis diperbarui setiap kali filter diubah.
      </p>
    </div>
  </div>
</template>
