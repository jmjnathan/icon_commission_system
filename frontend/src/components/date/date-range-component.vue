<script setup lang="ts">
interface Props {
  label?: string;
  startDate: string;
  endDate: string;
  required?: boolean;
  disabled?: boolean;
  errorMessage?: string;
}

withDefaults(defineProps<Props>(), {
  label: "Rentang Tanggal",
  required: false,
  disabled: false,
  errorMessage: "",
});

const emit = defineEmits<{
  (e: "update:startDate", value: string): void;
  (e: "update:endDate", value: string): void;
}>();
</script>

<template>
  <div class="w-full">
    <label
      class="block mb-2 text-sm font-medium"
      :class="errorMessage ? 'text-red-500' : 'text-[#6B5D45]'">
      {{ label }}
      <span v-if="required" class="text-red-500">*</span>
    </label>

    <div class="flex items-center gap-3">
      <input
        :value="startDate"
        @input="
          emit('update:startDate', ($event.target as HTMLInputElement).value)
        "
        type="date"
        :disabled="disabled"
        class="w-full h-12 rounded-2xl border border-[#E8DCC7] bg-white px-4 text-sm text-[#3A2E1F] transition-all duration-200 focus:outline-none focus:border-[#7A1F2B] focus:ring-4 focus:ring-[#7A1F2B]/10 disabled:bg-gray-100 disabled:cursor-not-allowed"
        :class="{ 'border-red-500': errorMessage }" />

      <span class="text-[#8A7A5C] font-semibold">—</span>

      <input
        :value="endDate"
        @input="
          emit('update:endDate', ($event.target as HTMLInputElement).value)
        "
        type="date"
        :disabled="disabled"
        class="w-full h-12 rounded-2xl border border-[#E8DCC7] bg-white px-4 text-sm text-[#3A2E1F] transition-all duration-200 focus:outline-none focus:border-[#7A1F2B] focus:ring-4 focus:ring-[#7A1F2B]/10 disabled:bg-gray-100 disabled:cursor-not-allowed"
        :class="{ 'border-red-500': errorMessage }" />
    </div>

    <p v-if="errorMessage" class="mt-1 text-xs text-red-500">
      {{ errorMessage }}
    </p>
  </div>
</template>
