<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import {
  TrendingUp,
  TrendingDown,
  Wallet,
  Calendar,
  Printer,
} from "lucide-vue-next";
import AppLayout from "../../../components/layout/sidebar-app-layout.vue";
import DataTablesComponent, {
  type Column,
} from "../../../components/tables/data-tables-component.vue";
import { useCashOut } from "../../../composables/useCashout";
import { useCommission } from "../../../composables/commission/useCommission";

const { items: cashOuts, fetchAll: fetchCashOuts } = useCashOut();
const { commissions, fetchAll: fetchCommissions } = useCommission();

onMounted(() => {
  fetchCashOuts();
  fetchCommissions();
});

// ===== Pemilih Periode =====
const now = new Date();
const selectedMonth = ref(now.getMonth() + 1); // 1-12
const selectedYear = ref(now.getFullYear());

const monthOptions = [
  { value: 1, label: "Januari" },
  { value: 2, label: "Februari" },
  { value: 3, label: "Maret" },
  { value: 4, label: "April" },
  { value: 5, label: "Mei" },
  { value: 6, label: "Juni" },
  { value: 7, label: "Juli" },
  { value: 8, label: "Agustus" },
  { value: 9, label: "September" },
  { value: 10, label: "Oktober" },
  { value: 11, label: "November" },
  { value: 12, label: "Desember" },
];

const yearOptions = computed(() => {
  const years = [];
  for (let y = now.getFullYear(); y >= now.getFullYear() - 3; y--)
    years.push(y);
  return years;
});

function isInPeriod(dateStr: string) {
  const d = new Date(dateStr);
  return (
    d.getMonth() + 1 === selectedMonth.value &&
    d.getFullYear() === selectedYear.value
  );
}

const incomeInPeriod = computed(() =>
  commissions.value.filter(
    (c) =>
      (c.status === "completed" || c.status === "delivered") &&
      isInPeriod(c.order_date)
  )
);

const expenseInPeriod = computed(() =>
  cashOuts.value.filter((e) => isInPeriod(e.date))
);

const totalIncome = computed(() =>
  incomeInPeriod.value.reduce((sum, c) => sum + (c.total_price || 0), 0)
);

const totalExpense = computed(() =>
  expenseInPeriod.value.reduce((sum, e) => sum + e.amount, 0)
);

const netProfit = computed(() => totalIncome.value - totalExpense.value);
const isProfit = computed(() => netProfit.value >= 0);

// ===== Rincian per Kategori Pengeluaran =====
const expenseByCategory = computed(() => {
  const groups: Record<string, number> = {};
  expenseInPeriod.value.forEach((e) => {
    groups[e.category] = (groups[e.category] || 0) + e.amount;
  });
  return Object.entries(groups)
    .map(([category, total]) => ({ category, total }))
    .sort((a, b) => b.total - a.total);
});

function categoryPercent(total: number) {
  if (totalExpense.value === 0) return 0;
  return Math.round((total / totalExpense.value) * 100);
}

// ===== Tabel Gabungan Transaksi =====
interface Transaction {
  date: string;
  type: "in" | "out";
  description: string;
  category: string;
  amount: number;
}

const allTransactions = computed<Transaction[]>(() => {
  const incomes: Transaction[] = incomeInPeriod.value.map((c) => ({
    date: c.order_date,
    type: "in",
    description: `Komisi #${c.id} — ${c.client?.name}`,
    category: "Pemasukan Komisi",
    amount: c.total_price,
  }));

  const expenses: Transaction[] = expenseInPeriod.value.map((e) => ({
    date: e.date,
    type: "out",
    description: e.description,
    category: e.category,
    amount: e.amount,
  }));

  return [...incomes, ...expenses].sort(
    (a, b) => new Date(a.date).getTime() - new Date(b.date).getTime()
  );
});

const columns: Column[] = [
  { key: "date", label: "Tanggal" },
  { key: "description", label: "Keterangan" },
  { key: "category", label: "Kategori" },
  { key: "amount", label: "Jumlah", align: "right" },
];

function formatDate(dateStr: string) {
  return new Date(dateStr).toLocaleDateString("id-ID", {
    day: "numeric",
    month: "short",
    year: "numeric",
  });
}

function formatRupiah(num: number) {
  return num.toLocaleString("id-ID");
}

function handlePrintReport() {
  window.print();
}
</script>

<template>
  <AppLayout>
    <div class="flex items-center justify-between mb-6 print:hidden">
      <div>
        <h1 class="text-2xl font-semibold text-[#3A2E1F] mb-1">
          Laporan Keuangan
        </h1>
        <p class="text-sm text-[#8A7A5C]">
          Ringkasan kas masuk, kas keluar, dan laba rugi bulanan.
        </p>
      </div>
      <button
        @click="handlePrintReport"
        class="flex items-center gap-2 px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition">
        <Printer :size="16" />
        Cetak Laporan
      </button>
    </div>

    <!-- Pemilih Periode -->
    <div
      class="bg-white border border-[#E5D9BF] rounded-xl p-4 mb-6 print:hidden">
      <div class="flex items-center gap-4">
        <Calendar :size="18" class="text-[#8A7A5C]" />
        <select
          v-model="selectedMonth"
          class="px-4 py-2 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40">
          <option v-for="m in monthOptions" :key="m.value" :value="m.value">
            {{ m.label }}
          </option>
        </select>
        <select
          v-model="selectedYear"
          class="px-4 py-2 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40">
          <option v-for="y in yearOptions" :key="y" :value="y">{{ y }}</option>
        </select>
      </div>
    </div>

    <!-- Header cetak (cuma keliatan pas print) -->
    <div class="hidden print:block mb-6">
      <h1 class="text-xl font-bold text-[#3A2E1F]">
        Dominic's Art — Laporan Keuangan
      </h1>
      <p class="text-sm text-[#6B5D45]">
        Periode:
        {{ monthOptions.find((m) => m.value === selectedMonth)?.label }}
        {{ selectedYear }}
      </p>
    </div>

    <!-- Ringkasan Laba Rugi -->
    <div class="grid grid-cols-3 gap-4 mb-6">
      <div
        class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between">
        <div>
          <p class="text-xs font-medium text-[#4C8C5B] tracking-wide">
            TOTAL PEMASUKAN
          </p>
          <p class="text-lg font-semibold text-[#3A2E1F] mt-1">
            Rp {{ formatRupiah(totalIncome) }}
          </p>
          <p class="text-xs text-[#8A7A5C] mt-0.5">
            {{ incomeInPeriod.length }} transaksi
          </p>
        </div>
        <div
          class="w-10 h-10 rounded-full bg-[#4C8C5B]/15 flex items-center justify-center shrink-0">
          <TrendingUp :size="20" class="text-[#4C8C5B]" />
        </div>
      </div>

      <div
        class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between">
        <div>
          <p class="text-xs font-medium text-[#B23A32] tracking-wide">
            TOTAL PENGELUARAN
          </p>
          <p class="text-lg font-semibold text-[#3A2E1F] mt-1">
            Rp {{ formatRupiah(totalExpense) }}
          </p>
          <p class="text-xs text-[#8A7A5C] mt-0.5">
            {{ expenseInPeriod.length }} transaksi
          </p>
        </div>
        <div
          class="w-10 h-10 rounded-full bg-[#B23A32]/15 flex items-center justify-center shrink-0">
          <TrendingDown :size="20" class="text-[#B23A32]" />
        </div>
      </div>

      <div
        class="border rounded-xl p-4 flex items-center justify-between"
        :class="
          isProfit
            ? 'bg-[#4C8C5B]/5 border-[#4C8C5B]/30'
            : 'bg-[#B23A32]/5 border-[#B23A32]/30'
        ">
        <div>
          <p
            class="text-xs font-medium tracking-wide"
            :class="isProfit ? 'text-[#4C8C5B]' : 'text-[#B23A32]'">
            {{ isProfit ? "LABA BERSIH" : "RUGI BERSIH" }}
          </p>
          <p class="text-lg font-semibold text-[#3A2E1F] mt-1">
            Rp {{ formatRupiah(Math.abs(netProfit)) }}
          </p>
        </div>
        <div
          class="w-10 h-10 rounded-full flex items-center justify-center shrink-0"
          :class="isProfit ? 'bg-[#4C8C5B]/15' : 'bg-[#B23A32]/15'">
          <Wallet
            :size="20"
            :class="isProfit ? 'text-[#4C8C5B]' : 'text-[#B23A32]'" />
        </div>
      </div>
    </div>

    <!-- Rincian per Kategori -->
    <div class="bg-white border border-[#E5D9BF] rounded-xl p-6 mb-6">
      <h2
        class="text-sm font-semibold text-[#3A2E1F] uppercase tracking-wide mb-4">
        Rincian Pengeluaran per Kategori
      </h2>
      <p v-if="expenseByCategory.length === 0" class="text-sm text-[#8A7A5C]">
        Tidak ada pengeluaran pada periode ini.
      </p>
      <div v-else class="space-y-3">
        <div v-for="cat in expenseByCategory" :key="cat.category">
          <div class="flex items-center justify-between text-sm mb-1">
            <span class="text-[#3A2E1F]">{{ cat.category }}</span>
            <span class="text-[#6B5D45]">
              Rp {{ formatRupiah(cat.total) }} ({{
                categoryPercent(cat.total)
              }}%)
            </span>
          </div>
          <div class="w-full h-2 bg-[#E5D9BF] rounded-full overflow-hidden">
            <div
              class="h-full bg-[#B23A32] rounded-full"
              :style="{ width: categoryPercent(cat.total) + '%' }"></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Daftar Transaksi -->
    <h2 class="text-lg font-semibold text-[#3A2E1F] mb-3">Rincian Transaksi</h2>
    <DataTablesComponent
      :columns="columns"
      :items="allTransactions"
      empty-message="Tidak ada transaksi pada periode ini.">
      <template #date="{ item }">
        <span class="text-[#6B5D45]">{{ formatDate(item.date) }}</span>
      </template>
      <template #description="{ item }">
        <span class="text-[#3A2E1F]">{{ item.description }}</span>
      </template>
      <template #category="{ item }">
        <span
          class="px-2 py-0.5 rounded-full text-xs font-medium"
          :class="
            item.type === 'in'
              ? 'bg-[#4C8C5B]/15 text-[#4C8C5B]'
              : 'bg-[#B23A32]/15 text-[#B23A32]'
          ">
          {{ item.category }}
        </span>
      </template>
      <template #amount="{ item }">
        <span
          :class="item.type === 'in' ? 'text-[#4C8C5B]' : 'text-[#B23A32]'"
          class="font-medium">
          {{ item.type === "in" ? "+" : "-" }} Rp
          {{ formatRupiah(item.amount) }}
        </span>
      </template>
    </DataTablesComponent>
  </AppLayout>
</template>

<style>
@media print {
  .print\:hidden {
    display: none !important;
  }
}
</style>
