<!-- components/input-text/input-password-component.vue -->
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
    debounceMs?: number;
  }>(),
  {
    label: "",
    placeholder: "Masukkan password",
    required: false,
    disabled: false,
    readOnly: false,
    debounceMs: 300,
  }
);

const emit = defineEmits<{
  "update:modelValue": [value: string];
}>();

const tempValue = ref(props.modelValue);
const isVisible = ref(false);

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
  const val = target.value;
  tempValue.value = val;
  debouncedEmit(val);
}

function handleBlur() {
  debouncedEmit.flush();
}

function toggleVisibility() {
  isVisible.value = !isVisible.value;
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
    <div class="relative">
      <input
        :value="tempValue"
        :type="isVisible ? 'text' : 'password'"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readOnly"
        :maxlength="maxLength"
        @input="handleInput"
        @blur="handleBlur"
        class="w-full px-4 py-2.5 pr-11 bg-white border rounded-lg text-sm text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 transition disabled:bg-[#F5EFE3] disabled:cursor-not-allowed"
        :class="
          errorMessage
            ? 'border-[#B23A32] focus:ring-[#B23A32]/30 focus:border-[#B23A32]'
            : 'border-[#D9CBB0] focus:ring-[#C9A24B]/40 focus:border-[#C9A24B]'
        " />

      <button
        type="button"
        tabindex="-1"
        @click="toggleVisibility"
        :disabled="disabled"
        class="absolute right-3 top-1/2 -translate-y-1/2 text-[#8A7A5C] hover:text-[#B23A32] transition disabled:opacity-50 disabled:cursor-not-allowed">
        <svg
          v-if="isVisible"
          xmlns="http://www.w3.org/2000/svg"
          class="h-4.5 w-4.5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          stroke-width="1.8">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
        </svg>
        <svg
          v-else
          xmlns="http://www.w3.org/2000/svg"
          class="h-4.5 w-4.5"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          stroke-width="1.8">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M3.98 8.223A10.477 10.477 0 001.934 12c1.292 4.338 5.31 7.5 10.066 7.5.993 0 1.953-.138 2.863-.395M6.228 6.228A10.451 10.451 0 0112 4.5c4.756 0 8.773 3.162 10.065 7.498a10.523 10.523 0 01-4.293 5.774M6.228 6.228L3 3m3.228 3.228l3.65 3.65m7.894 7.894L21 21m-3.228-3.228l-3.65-3.65m0 0a3 3 0 10-4.243-4.243m4.242 4.242L9.88 9.88" />
        </svg>
      </button>
    </div>

    <p v-if="errorMessage" class="text-xs text-[#B23A32] mt-1">
      {{ errorMessage }}
    </p>
    <p v-else-if="infoMessage" class="text-xs text-[#8A7A5C] mt-1">
      {{ infoMessage }}
    </p>
  </div>
</template>
