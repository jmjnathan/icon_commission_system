<script setup lang="ts">
import { ref } from "vue";
import { useAuth } from "../composables/useAuth";
import InputTextComponent from "../components/input-text/input-text-component.vue";
import InputPasswordComponent from "../components/input-text/input-password-component.vue";
import logo from "../assets/logo.png";

const username = ref("");
const password = ref("");
const { login, isLoading, errorMessage } = useAuth();

const validationError = ref("");

async function handleSubmit() {
  validationError.value = "";

  if (!username.value.trim() && !password.value.trim()) {
    validationError.value = "Username dan password wajib diisi";
    return;
  }
  if (!username.value.trim()) {
    validationError.value = "Username wajib diisi";
    return;
  }
  if (!password.value.trim()) {
    validationError.value = "Password wajib diisi";
    return;
  }

  await login(username.value, password.value);
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-[#F5EFE3] px-4">
    <div
      class="w-full max-w-lg rounded-3xl bg-white border border-[#E8DCC7] shadow-xl px-10 py-10">
      <!-- Brand -->
      <div class="text-center mb-8">
        <div class="h-20 w-20 mx-auto mb-4 rounded-full overflow-hidden">
          <img
            :src="logo"
            alt="Dominic's Art"
            class="w-full h-full object-cover" />
        </div>
        <h1 class="text-3xl font-serif text-[#7A1F2B]">Dominic's Art</h1>
        <p class="italic text-[#8A7A5C] mt-1">Sub Tutela Matris</p>
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-4">
        <InputTextComponent
          label="Username"
          v-model="username"
          required
          placeholder="Masukkan username" />

        <InputPasswordComponent
          label="Password"
          v-model="password"
          required
          placeholder="Masukkan password" />

        <p
          v-if="validationError || errorMessage"
          class="text-sm text-[#7A1F2B] bg-[#7A1F2B]/10 px-3 py-2 rounded-lg">
          {{ validationError || errorMessage }}
        </p>

        <button
          type="submit"
          :disabled="isLoading"
          class="w-full py-2.5 bg-[#7A1F2B] text-[#F5EFE3] rounded-lg font-medium hover:bg-[#5F1621] transition disabled:opacity-50">
          {{ isLoading ? "Memproses..." : "Masuk" }}
        </button>
      </form>
    </div>
  </div>
</template>
