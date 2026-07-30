<script setup lang="ts">
import { useToast } from "../../composables/etc/useToast";
import { CheckCircle2, XCircle, Info, X } from "lucide-vue-next";

const { toasts, dismiss } = useToast();

const styleMap = {
  success: { bg: "bg-[#4C8C5B]", icon: CheckCircle2 },
  error: { bg: "bg-[#B23A32]", icon: XCircle },
  info: { bg: "bg-[#3B6FA8]", icon: Info },
};
</script>

<template>
  <Teleport to="body">
    <div class="fixed top-4 right-4 z-100 flex flex-col gap-2 w-80">
      <TransitionGroup name="toast-slide">
        <div
          v-for="toast in toasts"
          :key="toast.id"
          class="flex items-start gap-3 text-white rounded-lg shadow-lg px-4 py-3"
          :class="styleMap[toast.type].bg">
          <component
            :is="styleMap[toast.type].icon"
            :size="18"
            class="shrink-0 mt-0.5" />
          <p class="text-sm flex-1">{{ toast.message }}</p>
          <button
            @click="dismiss(toast.id)"
            class="shrink-0 text-white/70 hover:text-white transition">
            <X :size="16" />
          </button>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<style scoped>
.toast-slide-enter-active,
.toast-slide-leave-active {
  transition: all 0.25s ease;
}
.toast-slide-enter-from {
  opacity: 0;
  transform: translateX(30px);
}
.toast-slide-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>
