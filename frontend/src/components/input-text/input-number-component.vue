<script setup lang="ts">
import { ref, watch } from "vue";

const props = withDefaults(
  defineProps<{
    label?: string;
    placeholder?: string;
    modelValue: number | null;
    required?: boolean;
    errorMessage?: string;
    infoMessage?: string;
    disabled?: boolean;
    readOnly?: boolean;
    prefix?: string;
    suffix?: string;
    min?: number;
    max?: number | null;
    maxFractionDigits?: number;
  }>(),
  {
    label: "",
    placeholder: "0",
    required: false,
    disabled: false,
    readOnly: false,
    min: 0,
    max: null,
    maxFractionDigits: 2,
  }
);

const emit = defineEmits<{
  "update:modelValue": [value: number | null];
}>();

function formatNumber(num: number) {
  return num.toLocaleString("id-ID", {
    maximumFractionDigits: props.maxFractionDigits,
  });
}

const displayValue = ref(
  typeof props.modelValue === "number" ? formatNumber(props.modelValue) : ""
);

watch(
  () => props.modelValue,
  (val) => {
    displayValue.value = typeof val === "number" ? formatNumber(val) : "";
  }
);

function parseRaw(raw: string): number | null {
  // Format Indonesia: titik = ribuan, koma = desimal
  const cleaned = raw.replace(/\./g, "").replace(",", ".");
  const num = parseFloat(cleaned);
  return isNaN(num) ? null : num;
}

function handleInput(e: Event) {
  const target = e.target as HTMLInputElement;
  displayValue.value = target.value;
}

function handleBlur() {
  const num = parseRaw(displayValue.value);
  if (num === null) {
    displayValue.value = "";
    emit("update:modelValue", null);
    return;
  }
  const clamped = props.max !== null ? Math.min(num, props.max) : num;
  const finalVal = Math.max(clamped, props.min);
  displayValue.value = formatNumber(finalVal);
  emit("update:modelValue", finalVal);
}
</script>

<template>
  <div>
    <label
      v-if="label"
      class="block text-xs font-medium mb-1.5"
      :class="errorMessage ? 'text-[#B23A32]' : 'text-[#6B5D45]'">
      {{ label }} <span v-if="required" class="text-[#B23A32]">*</span>
    </label>
    <div
      class="flex"
      :class="[
        prefix || suffix ? 'rounded-lg border overflow-hidden' : '',
        errorMessage ? 'border-[#B23A32]' : 'border-[#D9CBB0]',
      ]">
      <span
        v-if="prefix"
        class="flex items-center px-3 bg-[#F5EFE3] text-sm text-[#6B5D45] border-r border-[#D9CBB0]">
        {{ prefix }}
      </span>
      <input
        :value="displayValue"
        type="text"
        inputmode="decimal"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readOnly"
        @input="handleInput"
        @blur="handleBlur"
        class="w-full px-4 py-2.5 bg-white text-sm text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 transition disabled:bg-[#F5EFE3] disabled:cursor-not-allowed"
        :class="[
          prefix || suffix ? '' : 'rounded-lg border',
          errorMessage ? 'focus:ring-[#B23A32]/30' : 'focus:ring-[#C9A24B]/40',
          !(prefix || suffix) &&
            (errorMessage ? 'border-[#B23A32]' : 'border-[#D9CBB0]'),
        ]" />
      <span
        v-if="suffix"
        class="flex items-center px-3 bg-[#F5EFE3] text-sm text-[#6B5D45] border-l border-[#D9CBB0]">
        {{ suffix }}
      </span>
    </div>
    <p v-if="errorMessage" class="text-xs text-[#B23A32] mt-1">
      {{ errorMessage }}
    </p>
    <p v-else-if="infoMessage" class="text-xs text-[#8A7A5C] mt-1">
      {{ infoMessage }}
    </p>
  </div>
</template>
