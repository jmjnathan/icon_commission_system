<script setup>
import { computed, ref } from "vue";

const props = defineProps({
  label: { type: String, default: "Tambah" },
  icon: { type: String, default: "+" },
  severity: { type: String, default: "primary" },
  buttonType: { type: String, default: "solid" },
  type: { type: String, default: "button" },
  disabled: { type: Boolean, default: false },
  tooltip: { type: String, default: "" },
  onClick: { type: Function, default: null },
});

// internal loading state, gak perlu dioper dari luar lagi
const internalLoading = ref(false);

const severityColorMap = {
  primary: {
    solid: "bg-blue-600 text-white hover:bg-blue-700",
    border: "border-blue-600 text-blue-600 hover:bg-blue-50",
  },
  success: {
    solid: "bg-green-600 text-white hover:bg-green-700",
    border: "border-green-600 text-green-600 hover:bg-green-50",
  },
  danger: {
    solid: "bg-red-600 text-white hover:bg-red-700",
    border: "border-red-600 text-red-600 hover:bg-red-50",
  },
  warning: {
    solid: "bg-yellow-500 text-white hover:bg-yellow-600",
    border: "border-yellow-500 text-yellow-600 hover:bg-yellow-50",
  },
  info: {
    solid: "bg-sky-500 text-white hover:bg-sky-600",
    border: "border-sky-500 text-sky-600 hover:bg-sky-50",
  },
  secondary: {
    solid: "bg-gray-500 text-white hover:bg-gray-600",
    border: "border-gray-400 text-gray-600 hover:bg-gray-50",
  },
};

const buttonClasses = computed(() => {
  const colors = severityColorMap[props.severity] || severityColorMap.primary;
  const base =
    "inline-flex items-center justify-center gap-2 px-4 py-2 font-medium transition disabled:opacity-50 disabled:cursor-not-allowed";

  if (props.buttonType === "text")
    return `${base} bg-transparent ${
      colors.border.split(" ")[1]
    } hover:bg-gray-50 rounded`;
  if (props.buttonType === "outlined")
    return `${base} bg-transparent border ${colors.border} rounded`;
  if (props.buttonType === "rounded")
    return `${base} ${colors.solid} rounded-full`;
  return `${base} ${colors.solid} rounded`;
});

async function handleClick(event) {
  if (!props.onClick || internalLoading.value) return;

  const result = props.onClick(event);

  // kalau function-nya async (balikin Promise), tunggu dulu, tampilin loading
  if (result instanceof Promise) {
    internalLoading.value = true;
    try {
      await result;
    } finally {
      internalLoading.value = false;
    }
  }
  // kalau function-nya sync biasa, gak perlu loading state sama sekali
}
</script>

<template>
  <div class="relative inline-block group">
    <button
      :type="type"
      :disabled="disabled || internalLoading"
      :class="buttonClasses"
      class="w-full"
      @click="handleClick">
      <span
        v-if="internalLoading"
        class="animate-spin inline-block w-4 h-4 border-2 border-current border-t-transparent rounded-full"></span>
      <span v-else-if="icon">{{ icon }}</span>
      <span>{{ label }}</span>
    </button>
    <div
      v-if="tooltip"
      class="absolute left-1/2 -translate-x-1/2 top-full mt-1 px-2 py-1 bg-gray-800 text-white text-xs rounded opacity-0 group-hover:opacity-100 pointer-events-none whitespace-nowrap transition z-10">
      {{ tooltip }}
    </div>
  </div>
</template>
