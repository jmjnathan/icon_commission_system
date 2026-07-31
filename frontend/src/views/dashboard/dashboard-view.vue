<script setup lang="ts">
import {
  onMounted,
  computed,
  ref,
  watch,
  onBeforeUnmount,
  nextTick,
} from "vue";
import AppLayout from "../../components/layout/sidebar-app-layout.vue";
import { useCommission } from "../../composables/commission/useCommission.ts";
import NotificationComponent from "../../components/notification/notification-component.vue";
import DataTablesComponent from "../../components/tables/data-tables-component.vue";
import {
  Eye,
  Zap,
  Clock,
  CheckCircle2,
  PackageCheck,
  PenTool,
  TrendingUp,
  TrendingDown,
  Wallet,
} from "lucide-vue-next";
import { useToast } from "../../composables/etc/useToast.ts";
import { useConfirm } from "../../composables/etc/useConfirm.ts";
import CommissionDetailModal from "./modals/commission-detail-modal.vue";
import PaginationComponent from "../../components/pagination/pagination-component.vue";
import { useCashOut } from "../../composables/useCashout.ts";
import {
  Chart,
  BarController,
  BarElement,
  CategoryScale,
  LinearScale,
  Tooltip,
} from "chart.js";

Chart.register(BarController, BarElement, CategoryScale, LinearScale, Tooltip);

const { commissions, isLoading, fetchAll, updateStatus } = useCommission();
const {
  items: cashOutItems,
  isLoading: isLoadingCashOut,
  fetchAll: fetchCashOut,
} = useCashOut();
const username = localStorage.getItem("username") || "Admin";

onMounted(async () => {
  isLoading.value = true;
  await new Promise((resolve) => setTimeout(resolve, 500));
  await fetchAll();
  fetchCashOut();
  nextTick(() => renderChart());
});

// Ringkasan Keuangan
const isFinanceLoading = computed(
  () => isLoading.value || isLoadingCashOut.value
);

const totalIncome = computed(() =>
  commissions.value
    .filter((c) => c.status === "completed" || c.status === "delivered")
    .reduce((sum, c) => sum + (c.total_price || 0), 0)
);

const totalExpense = computed(() =>
  cashOutItems.value.reduce((sum, i) => sum + (i.amount || 0), 0)
);

const balance = computed(() => totalIncome.value - totalExpense.value);

const financeStats = computed(() => [
  {
    label: "TOTAL PEMASUKAN",
    value: `Rp ${totalIncome.value.toLocaleString("id-ID")}`,
    color: "#4C8C5B",
    bg: "#4C8C5B",
    icon: TrendingUp,
  },
  {
    label: "TOTAL PENGELUARAN",
    value: `Rp ${totalExpense.value.toLocaleString("id-ID")}`,
    color: "#B23A32",
    bg: "#B23A32",
    icon: TrendingDown,
  },
  {
    label: "SALDO SAAT INI",
    value: `Rp ${balance.value.toLocaleString("id-ID")}`,
    color: "#C9A24B",
    bg: "#C9A24B",
    icon: Wallet,
  },
]);

const statusCount = computed(() => {
  const counts = { pending: 0, in_progress: 0, completed: 0, delivered: 0 };
  commissions.value.forEach((c) => {
    if (c.status in counts) counts[c.status as keyof typeof counts]++;
  });
  return counts;
});

const stats = computed(() => [
  {
    label: "PENDING",
    value: `${statusCount.value.pending} Pesanan`,
    color: "#C9A24B",
    bg: "#C9A24B",
    icon: Clock,
  },
  {
    label: "IN PROGRESS",
    value: `${statusCount.value.in_progress} Pesanan`,
    color: "#3B6FA8",
    bg: "#3B6FA8",
    icon: PenTool,
  },
  {
    label: "COMPLETED",
    value: `${statusCount.value.completed} Selesai`,
    color: "#4C8C5B",
    bg: "#4C8C5B",
    icon: CheckCircle2,
  },
  {
    label: "DELIVERED",
    value: `${statusCount.value.delivered} Terkirim`,
    color: "#8A7A5C",
    bg: "#8A7A5C",
    icon: PackageCheck,
  },
]);
const toast = useToast();
const { confirm } = useConfirm();

const ongoingCommissions = computed(() =>
  commissions.value.filter((c) => c.status !== "delivered")
);

const statusFlow = ["pending", "in_progress", "completed", "delivered"];
const statusLabel: Record<string, string> = {
  pending: "Pending",
  in_progress: "In Progress",
  completed: "Completed",
  delivered: "Delivered",
};

function nextStatus(current: string) {
  const idx = statusFlow.indexOf(current);
  if (idx === -1 || idx === statusFlow.length - 1) return null;
  return statusFlow[idx + 1];
}

async function handleAdvanceStatus(item: any) {
  console.log("handleAdvanceStatus dipanggil", item);
  const next = nextStatus(item.status);
  console.log("next status:", next);
  if (!next) return;

  console.log("sebelum confirm() dipanggil");
  const confirmed = await confirm({
    title: "Ubah Status Komisi",
    message: `Ubah status komisi #${item.id} (${item.client?.name}) dari "${
      statusLabel[item.status]
    }" menjadi "${statusLabel[next]}"?`,
    confirmLabel: "Ya, Ubah",
  });
  console.log("setelah confirm(), hasilnya:", confirmed);

  if (!confirmed) return;

  try {
    await updateStatus(item.id, next);
    toast.success(`Status berhasil diubah menjadi "${statusLabel[next]}"`);
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal mengubah status");
  }
}

function progressPercent(status: string) {
  const map: Record<string, number> = {
    pending: 10,
    in_progress: 50,
    completed: 90,
    delivered: 100,
  };
  return map[status] ?? 0;
}

function progressColor(status: string) {
  const map: Record<string, string> = {
    pending: "#C9A24B",
    in_progress: "#3B6FA8",
    completed: "#4C8C5B",
    delivered: "#8A7A5C",
  };
  return map[status] ?? "#B0A588";
}

function statusBadge(status: string) {
  const map: Record<string, string> = {
    pending: "bg-[#C9A24B]/15 text-[#8A6D1F]",
    in_progress: "bg-[#3B6FA8]/15 text-[#3B6FA8]",
    completed: "bg-[#4C8C5B]/15 text-[#4C8C5B]",
    delivered: "bg-[#8A7A5C]/15 text-[#8A7A5C]",
  };
  return map[status] || "bg-gray-100 text-gray-600";
}

function formatDate(dateStr: string) {
  return new Date(dateStr).toLocaleDateString("id-ID", {
    day: "numeric",
    month: "short",
    year: "numeric",
  });
}

function subjectSummary(c: any) {
  if (!c.items || c.items.length === 0) return "-";
  const names = c.items.map((i: any) => i.saint?.name).filter(Boolean);
  if (names.length === 1) return names[0];
  return `${names[0]} +${names.length - 1} lainnya`;
}

function detailSummary(c: any) {
  if (!c.items || c.items.length === 0) return "";
  const first = c.items[0];
  const parts = [first.size?.size, first.material?.name].filter(Boolean);
  return parts.join(", ");
}

function daysUntil(dateStr: string) {
  const diff = new Date(dateStr).getTime() - new Date().getTime();
  return Math.ceil(diff / (1000 * 60 * 60 * 24));
}

const urgentItems = computed(() =>
  commissions.value
    .filter(
      (c) =>
        c.status !== "delivered" &&
        daysUntil(c.deadline) <= 14 &&
        daysUntil(c.deadline) >= 0
    )
    .sort((a, b) => daysUntil(a.deadline) - daysUntil(b.deadline))
    .map((c) => ({
      id: c.id,
      clientName: c.client?.name ?? "-",
      subject: subjectSummary(c),
      daysLeft: daysUntil(c.deadline),
      status: c.status,
    }))
);
// Header tabel
interface Column {
  key: string;
  label: string;
  align?: "left" | "center" | "right";
}

const columns: Column[] = [
  { key: "aksi", label: "Aksi", align: "center" },
  { key: "client", label: "Client" },
  { key: "subject", label: "Subjek & Detail Ikon" },
  { key: "deadline", label: "Deadline" },
  { key: "progress", label: "Progress", align: "left" },
];

const showDetailModal = ref(false);
const selectedCommission = ref<any>(null);

function openDetailModal(item: any) {
  console.log("Data commission yang dibuka:", item);
  selectedCommission.value = item;
  showDetailModal.value = true;
}

function closeDetailModal() {
  showDetailModal.value = false;
  selectedCommission.value = null;
}
const currentPage = ref(1);
const perPage = ref(50);

// ===== Grafik: Barang/Subjek Paling Banyak Dipesan =====
const topSubjects = computed(() => {
  const counts: Record<string, number> = {};

  commissions.value.forEach((c) => {
    c.items?.forEach((item: any) => {
      const name = item.saint?.name;
      if (!name) return;
      counts[name] = (counts[name] || 0) + 1;
    });
  });

  return Object.entries(counts)
    .sort((a, b) => b[1] - a[1])
    .slice(0, 6)
    .map(([name, count]) => ({ name, count }));
});

const chartCanvas = ref<HTMLCanvasElement | null>(null);
let chartInstance: Chart | null = null;

function renderChart() {
  if (!chartCanvas.value || topSubjects.value.length === 0) return;

  if (chartInstance) {
    chartInstance.destroy();
  }

  chartInstance = new Chart(chartCanvas.value, {
    type: "bar",
    data: {
      labels: topSubjects.value.map((d) => d.name),
      datasets: [
        {
          label: "Jumlah Dipesan",
          data: topSubjects.value.map((d) => d.count),
          backgroundColor: "#C9A24B",
          borderRadius: 6,
          maxBarThickness: 28,
        },
      ],
    },
    options: {
      indexAxis: "y",
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: { display: false },
        tooltip: {
          backgroundColor: "#3A2E1F",
          padding: 8,
          cornerRadius: 6,
        },
      },
      scales: {
        x: {
          beginAtZero: true,
          ticks: { stepSize: 1, color: "#8A7A5C" },
          grid: { color: "#F0E8D8" },
        },
        y: {
          ticks: { color: "#3A2E1F" },
          grid: { display: false },
        },
      },
    },
  });
}

watch(topSubjects, () => {
  nextTick(() => renderChart());
});

onBeforeUnmount(() => {
  chartInstance?.destroy();
});
</script>

<template>
  <AppLayout>
    <!-- Header: stack di mobile, row di desktop -->
    <div
      class="flex flex-col sm:flex-row sm:items-start sm:justify-between gap-4 mb-6">
      <div>
        <h1 class="text-xl sm:text-2xl font-semibold text-[#3A2E1F] mb-1">
          Dashboard Overview
        </h1>
        <p class="text-sm text-[#8A7A5C]">
          Selamat datang kembali. Berikut ringkasan aktivitas komisi ikonografi
          Anda.
        </p>
      </div>

      <div
        class="flex items-center gap-2 text-sm text-[#3A2E1F] font-medium self-start sm:self-auto">
        <div
          class="w-8 h-8 rounded-full bg-[#C9A24B] flex items-center justify-center text-white text-xs font-semibold shrink-0">
          {{ username.charAt(0).toUpperCase() }}
        </div>
        <span>{{ username }}</span>
      </div>
    </div>

    <!-- Perhatian Khusus -->
    <NotificationComponent :items="urgentItems" />

    <h2 class="text-md font-semibold text-[#3A2E1F] mt-4">
      Monitoring Pemasukan
    </h2>
    <div class="grid grid-cols-1 sm:grid-cols-3 gap-4 mt-2">
      <!-- Skeleton -->
      <template v-if="isFinanceLoading">
        <div
          v-for="n in 3"
          :key="`finance-skeleton-${n}`"
          class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between animate-pulse">
          <div class="flex-1">
            <div class="h-3 w-28 bg-[#E5D9BF] rounded mb-2"></div>
            <div class="h-6 w-32 bg-[#E5D9BF] rounded"></div>
          </div>
          <div class="w-10 h-10 rounded-full bg-[#E5D9BF] shrink-0"></div>
        </div>
      </template>

      <!-- Data asli -->
      <template v-else>
        <div
          v-for="stat in financeStats"
          :key="stat.label"
          class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between">
          <div class="min-w-0">
            <p
              class="text-xs font-medium tracking-wide truncate"
              :style="{ color: stat.color }">
              {{ stat.label }}
            </p>
            <p
              class="text-lg sm:text-xl font-bold text-[#3A2E1F] mt-1 truncate">
              {{ stat.value }}
            </p>
          </div>
          <div
            class="w-10 h-10 rounded-full flex items-center justify-center shrink-0 ml-2"
            :style="{ backgroundColor: stat.bg + '20' }">
            <component :is="stat.icon" :size="20" :style="{ color: stat.bg }" />
          </div>
        </div>
      </template>
    </div>

    <h2 class="text-md font-semibold text-[#3A2E1F] mt-4">
      Monitoring Pesanan
    </h2>
    <!-- Stats: 1 kolom mobile, 2 kolom tablet, 4 kolom desktop -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-4 mt-2">
      <!-- Skeleton saat loading -->
      <template v-if="isLoading">
        <div
          v-for="n in 4"
          :key="`skeleton-${n}`"
          class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between animate-pulse">
          <div class="flex-1">
            <div class="h-3 w-20 bg-[#E5D9BF] rounded mb-2"></div>
            <div class="h-5 w-24 bg-[#E5D9BF] rounded"></div>
          </div>
          <div class="w-10 h-10 rounded-full bg-[#E5D9BF] shrink-0"></div>
        </div>
      </template>

      <!-- Data asli -->
      <template v-else>
        <div
          v-for="stat in stats"
          :key="stat.label"
          class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between">
          <div class="min-w-0">
            <p
              class="text-xs font-medium tracking-wide truncate"
              :style="{ color: stat.color }">
              {{ stat.label }}
            </p>
            <p
              class="text-base sm:text-lg font-semibold text-[#3A2E1F] mt-1 truncate">
              {{ stat.value }}
            </p>
          </div>
          <div
            class="w-10 h-10 rounded-full flex items-center justify-center shrink-0 ml-2">
            <component :is="stat.icon" :size="20" :style="{ color: stat.bg }" />
          </div>
        </div>
      </template>
    </div>

    <h2 class="text-md font-semibold text-[#3A2E1F] mb-3">
      Pesanan Berlangsung
    </h2>

    <div class="bg-white border border-[#E5D9BF] rounded-xl p-2 mb-5">
      <div class="flex justify-end overflow-x-auto">
        <PaginationComponent
          v-model:current-page="currentPage"
          v-model:per-page="perPage"
          :total-items="ongoingCommissions.length" />
      </div>

      <!-- Bungkus tabel supaya bisa scroll horizontal di layar kecil -->
      <div class="overflow-x-auto">
        <DataTablesComponent
          :columns="columns"
          :items="ongoingCommissions"
          :is-loading="isLoading"
          empty-message="Tidak ada pesanan yang sedang berlangsung."
          row-key="id"
          class="mb-2 min-w-[720px]">
          <template #aksi="{ item }">
            <div class="flex gap-2">
              <button
                @click="openDetailModal(item)"
                class="w-8 h-8 flex items-center justify-center rounded-lg bg-[#C9A24B]/15 text-[#8A6D1F] hover:bg-[#C9A24B]/25 transition shrink-0"
                title="Lihat detail">
                <Eye :size="16" />
              </button>
              <button
                @click="handleAdvanceStatus(item)"
                class="w-8 h-8 flex items-center justify-center rounded-lg bg-[#3B6FA8]/15 text-[#3B6FA8] hover:bg-[#3B6FA8]/25 transition shrink-0"
                :title="`Ubah ke ${
                  statusLabel[nextStatus(item.status) ?? '']
                }`">
                <Zap :size="16" />
              </button>
            </div>
          </template>
          <template #client="{ item }">
            <span class="text-[#3A2E1F] font-medium">{{
              item.client?.name
            }}</span>
          </template>

          <template #subject="{ item }">
            <p class="text-[#3A2E1F]">{{ subjectSummary(item) }}</p>
            <p class="text-xs text-[#8A7A5C]">{{ detailSummary(item) }}</p>
          </template>

          <template #deadline="{ item }">
            <span class="text-[#6B5D45] whitespace-nowrap">{{
              formatDate(item.deadline)
            }}</span>
          </template>

          <template #progress="{ item }">
            <div class="flex items-center gap-2">
              <div class="w-24 h-2 bg-[#E5D9BF] rounded-full overflow-hidden">
                <div
                  class="h-full rounded-full"
                  :style="{
                    width: progressPercent(item.status) + '%',
                    backgroundColor: progressColor(item.status),
                  }"></div>
              </div>
              <span class="text-xs text-[#6B5D45]"
                >{{ progressPercent(item.status) }}%</span
              >
            </div>
            <span
              class="inline-block mt-1 px-2 py-0.5 rounded-full text-xs font-medium"
              :class="statusBadge(item.status)">
              {{ item.status }}
            </span>
          </template>
        </DataTablesComponent>
      </div>

      <div class="flex justify-end overflow-x-auto">
        <PaginationComponent
          v-model:current-page="currentPage"
          v-model:per-page="perPage"
          :total-items="ongoingCommissions.length" />
      </div>
    </div>

    <h2 class="text-md font-semibold text-[#3A2E1F] mb-3">
      Subjek Paling Banyak Dipesan
    </h2>

    <div class="bg-white border border-[#E5D9BF] rounded-xl p-4 mb-5">
      <!-- Skeleton -->
      <div v-if="isLoading" class="animate-pulse space-y-3">
        <div
          v-for="n in 5"
          :key="`chart-skeleton-${n}`"
          class="h-6 bg-[#E5D9BF] rounded"
          :style="{ width: `${90 - n * 10}%` }"></div>
      </div>

      <!-- Kosong -->
      <p
        v-else-if="topSubjects.length === 0"
        class="text-sm text-[#8A7A5C] text-center py-6">
        Belum ada data pesanan untuk ditampilkan.
      </p>

      <!-- Chart -->
      <div v-else class="h-56">
        <canvas ref="chartCanvas"></canvas>
      </div>
    </div>
  </AppLayout>
  <CommissionDetailModal
    :visible="showDetailModal"
    :commission="selectedCommission"
    @hide="closeDetailModal" />
</template>
