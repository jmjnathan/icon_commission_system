<script setup lang="ts">
import { ref, watch } from "vue";
import { debounce } from "lodash-es";

const props = withDefaults(
  defineProps<{
    label?: string;
    placeholder?: string;
    modelValue: string;
    required?: boolean;
    errorMessage?: string;
    infoMessage?: string;
    disabled?: boolean;
    readOnly?: boolean;
    maxLength?: number;
    autoUppercase?: boolean;
    debounceMs?: number;
  }>(),
  {
    label: "",
    placeholder: "Masukkan teks",
    required: false,
    disabled: false,
    readOnly: false,
    autoUppercase: false,
    debounceMs: 300,
  }
);

const emit = defineEmits<{
  "update:modelValue": [value: string];
}>();

const tempValue = ref(props.modelValue);

const debouncedEmit = debounce((val: string) => {
  emit("update:modelValue", val);
}, props.debounceMs);

watch(
  () => props.modelValue,
  (val) => {
    if (val !== tempValue.value) tempValue.value = val;
  }
);

function handleInput(e: Event) {
  const target = e.target as HTMLInputElement;
  let val = target.value;
  if (props.autoUppercase) val = val.toUpperCase();
  tempValue.value = val;
  debouncedEmit(val);
}

function handleBlur() {
  debouncedEmit.flush();
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
    <input
      :value="tempValue"
      type="text"
      :placeholder="placeholder"
      :disabled="disabled"
      :readonly="readOnly"
      :maxlength="maxLength"
      @input="handleInput"
      @blur="handleBlur"
      class="w-full px-4 py-2.5 bg-white border rounded-lg text-sm text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 transition disabled:bg-[#F5EFE3] disabled:cursor-not-allowed"
      :class="
        errorMessage
          ? 'border-[#B23A32] focus:ring-[#B23A32]/30 focus:border-[#B23A32]'
          : 'border-[#D9CBB0] focus:ring-[#C9A24B]/40 focus:border-[#C9A24B]'
      " />
    <p v-if="errorMessage" class="text-xs text-[#B23A32] mt-1">
      {{ errorMessage }}
    </p>
    <p v-else-if="infoMessage" class="text-xs text-[#8A7A5C] mt-1">
      {{ infoMessage }}
    </p>
  </div>
</template>
