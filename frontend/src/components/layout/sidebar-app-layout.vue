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
  Notebook,
  BriefcaseBusiness,
  Ruler,
  PaintBucket,
  PersonStanding,
  WarehouseIcon,
  ChartAreaIcon,
  ShoppingCartIcon,
  HandIcon,
} from "lucide-vue-next";
import logo from "../../assets/logo.png";
import { Chart } from "chart.js";

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
    name: "Arus Kas Keluar",
    path: "/cashflow/cash-out",
    icon: BriefcaseBusiness,
  },
  {
    name: "Laporan Keuangan",
    path: "/cashflow/report",
    icon: BriefcaseBusiness,
  },
];

const inventoryItem = [
  {
    name: "Gudang",
    path: "inventory/current-stock",
    icon: WarehouseIcon,
  },
];

const purchasingItem = [
  {
    name: "Pemesanan Bahan Baku",
    path: "purchasing/purchase-order",
    icon: ShoppingCartIcon,
  },
  {
    name: "Penerimaan Barang",
    path: "purchasing/receive-goods",
    icon: HandIcon,
  },
];

const masterItems = [
  { name: "Ukuran", path: "/master/size", icon: Ruler },
  { name: "Bahan Media", path: "/master/material", icon: Package },
  { name: "Component", path: "/master/material-components", icon: Package },
  { name: "Gaya Lukisan", path: "/master/styles", icon: PaintBucket },
  { name: "Objek Lukisan", path: "/master/saints", icon: PersonStanding },
];
</script>

<template>
  <div class="min-h-screen flex bg-[#F5EFE3]">
    <!-- Sidebar -->
    <aside
      class="fixed top-0 left-0 h-screen w-64 bg-[#0F1C2E] text-[#D9CBB0] flex flex-col print:hidden z-50">
      <div class="px-6 py-5 border-b border-white/10">
        <div class="flex items-center gap-3">
          <img
            :src="logo"
            alt="Dominic's Art"
            class="w-10 h-10 object-contain shrink-0" />

          <div class="min-w-0">
            <h1 class="text-[#C9A24B] font-bold text-lg leading-tight">
              DOMINIC'S ART
            </h1>
            <p class="mt-1 text-[11px] italic text-[#A6926A]">
              Sub Tutela Matris
            </p>
          </div>
        </div>
      </div>

      <nav class="flex-1 overflow-y-auto px-3 py-4 space-y-1">
        <p
          class="px-3 mt-4 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mb-2">
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
          class="px-3 mt-4 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mb-2">
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
          class="px-3 mt-4 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mb-2">
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
          class="px-3 mt-4 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mb-2">
          Pengadaan Stock
        </p>
        <RouterLink
          v-for="item in purchasingItem"
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
          class="px-3 mt-4 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mb-2">
          Gudang
        </p>
        <RouterLink
          v-for="item in inventoryItem"
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
          class="px-3 mt-4 text-xs text-[#8A7A5C] font-medium uppercase tracking-wider mt-6 mb-2">
          Master Data
        </p>
        <div v-if="masterOpen" class="space-y-1">
          <RouterLink
            v-for="item in masterItems"
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
        </div>
      </nav>

      <div class="px-3 mt-4 py-4 border-t border-white/10">
        <button
          @click="logout"
          class="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-[#B0A588] hover:bg-white/5 hover:text-[#D9CBB0] transition">
          <LogOut :size="18" />
          Log Out
        </button>
      </div>
    </aside>

    <!-- Main content -->
    <div class="flex-1 flex flex-col min-w-0 ml-64 print:ml-0">
      <main class="flex-1 overflow-y-auto p-8">
        <slot />
      </main>
    </div>
  </div>
</template>
