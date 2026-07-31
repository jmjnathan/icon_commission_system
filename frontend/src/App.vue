<script setup lang="ts">
import ToastContainer from "./components/toast/toast-component.vue";
import ConfirmModal from "./components/modal/confirm-modal.vue";
</script>

<template>
  <div class="min-h-screen bg-[#F5EFE3]">
    <router-view />
    <ToastContainer />
    <ConfirmModal />
  </div>
</template>
