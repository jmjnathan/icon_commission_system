<script setup lang="ts">
interface Column {
  key: string;
  label: string;
  width?: string;
  align?: "left" | "center" | "right";
}

defineProps<{
  columns: Column[];
  items: any[];
  isLoading?: boolean;
  emptyMessage?: string;
  rowKey?: string;
}>();

function headerAlignClass(align?: string) {
  if (align === "center") return "text-center";
  if (align === "right") return "text-right";
  return "text-left";
}

function cellJustifyClass(align?: string) {
  if (align === "center") return "justify-center text-center";
  if (align === "right") return "justify-end text-right";
  return "justify-start text-left";
}
</script>

<template>
  <div class="bg-white border border-[#E5D9BF] rounded-xl overflow-hidden">
    <p v-if="isLoading" class="p-6 text-sm text-[#8A7A5C]">Memuat...</p>

    <p v-else-if="items.length === 0" class="p-6 text-sm text-[#8A7A5C]">
      {{ emptyMessage || "Belum ada data." }}
    </p>

    <table v-else class="w-full text-sm">
      <thead
        class="bg-[#C9A24B]/10 text-[#8A6D1F] text-xs uppercase tracking-wide">
        <tr>
          <th
            v-for="col in columns"
            :key="col.key"
            class="px-4 py-3"
            :class="headerAlignClass(col.align)"
            :style="col.width ? { width: col.width } : {}">
            {{ col.label }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="(item, index) in items"
          :key="rowKey ? item[rowKey] : index"
          class="border-t border-[#E5D9BF] hover:bg-[#F5EFE3]/50">
          <td v-for="col in columns" :key="col.key" class="px-4 py-3">
            <div class="flex items-center" :class="cellJustifyClass(col.align)">
              <slot :name="col.key" :item="item">
                {{ item[col.key] }}
              </slot>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
