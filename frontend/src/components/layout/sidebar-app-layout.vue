<script setup lang="ts">
import { RouterLink, useRoute } from "vue-router";
import { useAuth } from "../../composables/useAuth";
import { ref } from "vue";
import {
  LayoutDashboard,
  Briefcase,
  User,
  Package,
  LogOut,
  ChevronDown,
  Notebook,
  BriefcaseBusiness,
} from "lucide-vue-next";

const route = useRoute();
const { logout } = useAuth();

const masterOpen = ref(true);

const navItems = [{ name: "Dashboard", path: "/", icon: LayoutDashboard }];

const transactionItems = [
  { name: "Pemesanan", path: "/transaction/commissions", icon: Briefcase },
  {
    name: "Riwayat",
    path: "/transaction/history-transaction",
    icon: Notebook,
  },
  { name: "Client", path: "/transaction/clients", icon: User },
];

const cashflowItem = [
  {
    name: "Laporan Keuangan",
    path: "/cashflow/report",
    icon: BriefcaseBusiness,
  },
];

const masterItems = [
  { name: "Ukuran", path: "/master/size" },
  { name: "Material", path: "/master/material" },
  { name: "Gaya Lukisan", path: "/master/styles" },
  { name: "Objek Lukisan", path: "/master/saints" },
];
</script>

<template>
  <div class="min-h-screen flex bg-[#F5EFE3]">
    <!-- Sidebar -->
    <aside class="w-64 bg-[#0F1C2E] text-[#D9CBB0] flex flex-col shrink-0">
      <div class="px-6 py-6 border-b border-white/10">
        <h1 class="text-[#C9A24B] font-semibold text-lg tracking-wide">
          DOMINIC'S ART
        </h1>
      </div>

      <nav class="flex-1 px-3 py-4 space-y-1 overflow-y-auto">
        <p
          class="px-3 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mb-2">
          Navigasi
        </p>
        <RouterLink
          v-for="item in navItems"
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition"
          :class="
            route.path === item.path
              ? 'bg-[#1E2E44] text-[#C9A24B] font-medium'
              : 'text-[#B0A588] hover:bg-white/5 hover:text-[#D9CBB0]'
          ">
          <component :is="item.icon" :size="18" />
          {{ item.name }}
        </RouterLink>

        <p
          class="px-3 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mb-2">
          Transaksi
        </p>
        <RouterLink
          v-for="item in transactionItems"
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition"
          :class="
            route.path === item.path
              ? 'bg-[#1E2E44] text-[#C9A24B] font-medium'
              : 'text-[#B0A588] hover:bg-white/5 hover:text-[#D9CBB0]'
          ">
          <component :is="item.icon" :size="18" />
          {{ item.name }}
        </RouterLink>
        <p
          class="px-3 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mb-2">
          Arus Kas
        </p>
        <RouterLink
          v-for="item in cashflowItem"
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition"
          :class="
            route.path === item.path
              ? 'bg-[#1E2E44] text-[#C9A24B] font-medium'
              : 'text-[#B0A588] hover:bg-white/5 hover:text-[#D9CBB0]'
          ">
          <component :is="item.icon" :size="18" />
          {{ item.name }}
        </RouterLink>

        <p
          class="px-3 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mt-6 mb-2">
          Master Data
        </p>
        <button
          @click="masterOpen = !masterOpen"
          class="w-full flex items-center justify-between px-3 py-2.5 rounded-lg text-sm text-[#B0A588] hover:bg-white/5 hover:text-[#D9CBB0] transition">
          <span class="flex items-center gap-3">
            <Package :size="18" />
            Data Master
          </span>
          <ChevronDown
            :size="14"
            class="transition-transform"
            :class="masterOpen ? 'rotate-180' : ''" />
        </button>
        <div v-if="masterOpen" class="pl-4 space-y-1">
          <RouterLink
            v-for="item in masterItems"
            :key="item.path"
            :to="item.path"
            class="block px-3 py-2 rounded-lg text-sm transition"
            :class="
              route.path === item.path
                ? 'bg-[#1E2E44] text-[#C9A24B] font-medium'
                : 'text-[#B0A588] hover:bg-white/5 hover:text-[#D9CBB0]'
            ">
            {{ item.name }}
          </RouterLink>
        </div>
      </nav>

      <div class="px-3 py-4 border-t border-white/10">
        <button
          @click="logout"
          class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-[#B0A588] hover:bg-white/5 hover:text-[#D9CBB0] transition">
          <LogOut :size="18" />
          Log Out
        </button>
      </div>
    </aside>

    <!-- Main content -->
    <div class="flex-1 flex flex-col min-w-0">
      <main class="flex-1 overflow-y-auto p-8">
        <slot />
      </main>
    </div>
  </div>
</template>
