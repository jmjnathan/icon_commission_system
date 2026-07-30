<script setup lang="ts">
import { useConfirm } from "../../composables/etc/useConfirm";

const { isVisible, options, handleConfirm, handleCancel } = useConfirm();
</script>

<template>
  <Teleport to="body">
    <Transition name="confirm-fade">
      <div
        v-if="isVisible"
        class="fixed inset-0 z-110 flex items-center justify-center">
        <div class="absolute inset-0 bg-black/40" @click="handleCancel"></div>

        <div
          class="relative bg-[#FAF6EC] rounded-xl shadow-xl w-full max-w-sm mx-4 p-6">
          <div class="flex items-start gap-3 mb-4">
            <div class="flex-1 pt-1">
              <h3 class="text-base font-semibold text-[#3A2E1F]">
                {{ options.title }}
              </h3>
              <p class="text-sm text-[#6B5D45] mt-1">{{ options.message }}</p>
            </div>
          </div>

          <div class="flex justify-end gap-3">
            <button
              @click="handleCancel"
              class="px-4 py-2 text-sm font-medium text-[#6B5D45] rounded-lg hover:bg-[#E5D9BF]/50 transition">
              {{ options.cancelLabel }}
            </button>
            <button
              @click="handleConfirm"
              class="px-4 py-2 text-sm font-medium text-white rounded-lg transition"
              :class="
                options.danger
                  ? 'bg-[#B23A32] hover:bg-[#8F2E27]'
                  : 'bg-[#7A1F2B] hover:bg-[#5F1621]'
              ">
              {{ options.confirmLabel }}
            </button>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.confirm-fade-enter-active,
.confirm-fade-leave-active {
  transition: opacity 0.2s ease;
}
.confirm-fade-enter-from,
.confirm-fade-leave-to {
  opacity: 0;
}
</style>
