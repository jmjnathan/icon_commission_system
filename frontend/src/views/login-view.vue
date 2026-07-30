<script setup lang="ts">
import { ref } from "vue";
import { useAuth } from "../composables/useAuth";

const username = ref("");
const password = ref("");
const { login, isLoading, errorMessage } = useAuth();
console.log("LoginView loaded");
function handleSubmit() {
  login(username.value, password.value);
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-[#F5EFE3] px-4">
    <div class="w-full max-w-sm">
      <div class="flex justify-center mb-8">
        <div class="relative w-16 h-16">
          <div
            class="absolute inset-0 rounded-full border-2 border-[#B08D3F]"></div>
          <div
            class="absolute inset-2 rounded-full border border-[#B08D3F]/50"></div>
          <div class="absolute inset-0 flex items-center justify-center">
            <div class="w-3 h-3 rounded-full bg-[#7A1F2B]"></div>
          </div>
        </div>
      </div>

      <h1 class="text-center font-serif text-2xl text-[#3A2E1F] mb-1">
        Dominic's Art
      </h1>
      <p class="text-center text-sm text-[#8A7A5C] mb-8">Manage Your Project</p>

      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Username
          </label>
          <input
            v-model="username"
            type="text"
            required
            class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 focus:ring-[#B08D3F]/40 focus:border-[#B08D3F] transition"
            placeholder="Masukkan username" />
        </div>

        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Password
          </label>
          <input
            v-model="password"
            type="password"
            required
            class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 focus:ring-[#B08D3F]/40 focus:border-[#B08D3F] transition"
            placeholder="Masukkan password" />
        </div>

        <p
          v-if="errorMessage"
          class="text-sm text-[#7A1F2B] bg-[#7A1F2B]/10 px-3 py-2 rounded-lg">
          {{ errorMessage }}
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
