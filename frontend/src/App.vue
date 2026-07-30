<script setup lang="ts">
import ToastContainer from "./components/toast/toast-component.vue";
import ConfirmModal from "./components/modal/confirm-modal.vue";
</script>

<template>
  <router-view />
  <ToastContainer />
  <ConfirmModal />
</template>
