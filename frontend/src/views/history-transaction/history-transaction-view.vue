<script setup lang="ts">
import { onMounted, computed, ref } from "vue";
import AppLayout from "../../components/layout/sidebar-app-layout.vue";
import DataTablesComponent from "../../components/tables/data-tables-component.vue";
import OverlayMenu from "../../components/button/overlay-menu.vue";
import CommissionDetailModal from ".././dashboard/modals/commission-detail-modal.vue";
import { useCommission } from "../../composables/commission/useCommission";
import { useToast } from "../../composables/etc/useToast";
import { printShippingLabel } from "../../utils/print-label.ts";
import { Eye, Printer, CheckCircle2, PackageCheck } from "lucide-vue-next";
import InputTextComponent from "../../components/input-text/input-text-component.vue";
import PaginationComponent from "../../components/pagination/pagination-component.vue";

interface Column {
  key: string;
  label: string;
  align?: "left" | "center" | "right";
}

const { commissions, isLoading, fetchAll } = useCommission();
const toast = useToast();

onMounted(async () => {
  isLoading.value = true;
  await new Promise((resolve) => setTimeout(resolve, 500));
  fetchAll();
});

const startDate = ref("");
const endDate = ref("");
// const dateRange = ref<any>([]);
const search = ref("");
const selectedStatus = ref("all");

const historyCommissions = computed(() => {
  return commissions.value.filter((c) => {
    // hanya history
    if (!["completed", "delivered"].includes(c.status)) return false;

    // status
    if (selectedStatus.value !== "all" && c.status !== selectedStatus.value)
      return false;

    // search
    if (search.value) {
      const keyword = search.value.toLowerCase();

      const found =
        c.client?.name?.toLowerCase().includes(keyword) ||
        c.items?.some((i: any) =>
          i.saint?.name?.toLowerCase().includes(keyword)
        );

      if (!found) return false;
    }

    // tanggal
    const created = new Date(c.order_date);

    if (startDate.value) {
      const start = new Date(startDate.value);
      start.setHours(0, 0, 0, 0);
      if (created < start) return false;
    }

    if (endDate.value) {
      const end = new Date(endDate.value);
      end.setHours(23, 59, 59, 999);
      if (created > end) return false;
    }

    return true;
  });
});
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

function statusBadge(status: string) {
  const map: Record<string, string> = {
    completed: "bg-[#4C8C5B]/15 text-[#4C8C5B]",
    delivered: "bg-[#8A7A5C]/15 text-[#8A7A5C]",
  };
  return map[status] || "bg-gray-100 text-gray-600";
}

function statusIcon(status: string) {
  return status === "completed" ? CheckCircle2 : PackageCheck;
}

const columns: Column[] = [
  { key: "aksi", label: "Aksi", align: "center" },
  { key: "client", label: "Client" },
  { key: "subject", label: "Subjek & Detail Ikon" },
  { key: "deadline", label: "Deadline" },
  { key: "total_price", label: "Total Harga" },
  { key: "status", label: "Status", align: "center" },
];

// ===== Modal Detail =====
const showDetailModal = ref(false);
const selectedCommission = ref<any>(null);

function openDetailModal(item: any) {
  selectedCommission.value = item;
  showDetailModal.value = true;
}

function closeDetailModal() {
  showDetailModal.value = false;
  selectedCommission.value = null;
}

// ===== Cetak Label =====
function handlePrintLabel(item: any) {
  if (!item.client) {
    toast.error("Data client tidak ditemukan");
    return;
  }
  printShippingLabel({
    id: item.id,
    clientName: item.client.name,
    phone: item.client.phone,
    address: item.client.address,
    city: item.client.city,
    province: item.client.province,
    postalCode: item.client.postal_code,
  });
}

function buildMenuItems(item: any) {
  const menu = [
    {
      label: "Lihat Detail",
      icon: Eye,
      command: () => openDetailModal(item),
    },
  ];

  if (item.status === "completed") {
    menu.push({
      label: "Cetak Label",
      icon: Printer,
      command: () => handlePrintLabel(item),
    });
  }

  return menu;
}

const grossIncome = computed(() =>
  historyCommissions.value.reduce(
    (total, item) => total + Number(item.total_price || 0),
    0
  )
);

const totalIcons = computed(() =>
  historyCommissions.value.reduce((sum, c) => sum + (c.items?.length || 0), 0)
);

const currentPage = ref(1);
const perPage = ref(50);
const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * perPage.value;

  return historyCommissions.value.slice(start, start + perPage.value);
});
</script>

<template>
  <AppLayout>
    <div class="mb-6">
      <h1 class="text-2xl font-semibold text-[#3A2E1F] mb-1">
        Riwayat Transaksi
      </h1>
      <p class="text-sm text-[#8A7A5C]">
        Daftar komisi yang sudah selesai dikerjakan atau sudah dikirim ke
        client.
      </p>
    </div>

    <div
      class="bg-white rounded-2xl border border-[#E8DCC7] p-6 mb-6 space-y-6">
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 items-start">
        <div>
          <label class="block text-sm font-semibold text-[#6B5D45] mb-2">
            Status
          </label>
          <InputTextComponent
            label=""
            v-model="search"
            type="text"
            placeholder="Ketik untuk memfilter berdasarkan klient" />
        </div>

        <div>
          <div class="relative">
            <label class="block text-sm font-semibold text-[#6B5D45] mb-2">
              Rentang Tanggal
            </label>

            <div
              class="mt-2 h-11.5 rounded-xl border border-[#E8DCC7] flex items-center px-4">
              <input class="flex-1 outline-none bg-transparent" type="date" />

              <span class="mx-3 text-[#8A7A5C]">—</span>

              <input class="flex-1 outline-none bg-transparent" type="date" />
            </div>
          </div>
        </div>
        <div>
          <label class="block text-sm font-semibold text-[#6B5D45] mb-2">
            Status
          </label>

          <div class="flex gap-6">
            <label class="flex items-center gap-2 cursor-pointer">
              <input v-model="selectedStatus" type="radio" value="all" />
              Semua
            </label>

            <label class="flex items-center gap-2 cursor-pointer">
              <input v-model="selectedStatus" type="radio" value="completed" />
              Selesai
            </label>

            <label class="flex items-center gap-2 cursor-pointer">
              <input v-model="selectedStatus" type="radio" value="delivered" />
              Dikirim
            </label>
          </div>
        </div>
      </div>
      <div
        class="rounded-2xl border border-[#E8DCC7] bg-[#FBF8F2] px-6 py-5 flex flex-col lg:flex-row lg:items-center lg:justify-between gap-4">
        <!-- Skeleton -->
        <template v-if="isLoading">
          <div class="animate-pulse">
            <div class="h-4 w-40 bg-[#E8DCC7] rounded mb-2"></div>
            <div class="h-9 w-56 bg-[#E8DCC7] rounded mb-2"></div>
            <div class="h-3 w-32 bg-[#E8DCC7] rounded"></div>
          </div>

          <div
            class="flex flex-wrap items-center gap-6 lg:justify-end animate-pulse">
            <div class="text-center">
              <div class="h-3 w-16 bg-[#E8DCC7] rounded mb-2 mx-auto"></div>
              <div class="h-7 w-10 bg-[#E8DCC7] rounded mx-auto"></div>
            </div>
            <div class="h-10 w-px bg-[#E8DCC7] hidden sm:block"></div>
            <div class="text-center">
              <div class="h-3 w-16 bg-[#E8DCC7] rounded mb-2 mx-auto"></div>
              <div class="h-7 w-10 bg-[#E8DCC7] rounded mx-auto"></div>
            </div>
          </div>
        </template>

        <!-- Data asli -->
        <template v-else>
          <!-- Kiri -->
          <div>
            <p class="text-sm font-medium text-[#8A7A5C]">
              Total Penghasilan Kotor
            </p>
            <h2 class="text-4xl font-bold text-[#7A1F2B] mt-1">
              Rp {{ grossIncome.toLocaleString("id-ID") }}
            </h2>
            <p class="text-sm text-[#8A7A5C] mt-1">
              Berdasarkan filter yang dipilih
            </p>
          </div>

          <!-- Kanan -->
          <div class="flex flex-wrap items-center gap-6 lg:justify-end text-sm">
            <div class="text-center">
              <p class="text-[#8A7A5C]">Transaksi</p>
              <p class="text-2xl font-bold text-[#3A2E1F]">
                {{ historyCommissions.length }}
              </p>
            </div>

            <div class="h-10 w-px bg-[#E8DCC7] hidden sm:block"></div>

            <div class="text-center">
              <p class="text-[#8A7A5C]">Total Icon</p>
              <p class="text-2xl font-bold text-[#3A2E1F]">
                {{ totalIcons }}
              </p>
            </div>
          </div>
        </template>
      </div>
      <div class="flex justify-end">
        <PaginationComponent
          v-model:current-page="currentPage"
          v-model:per-page="perPage"
          :total-items="paginatedItems.length" />
      </div>

      <DataTablesComponent
        :columns="columns"
        :items="historyCommissions"
        :is-loading="isLoading"
        empty-message="Belum ada riwayat transaksi."
        row-key="id">
        <template #aksi="{ item }">
          <OverlayMenu :items="buildMenuItems(item)" />
        </template>

        <template #client="{ item }">
          <span class="text-[#3A2E1F] font-medium">{{
            item.client?.name
          }}</span>
        </template>

        <template #subject="{ item }">
          <span class="text-[#3A2E1F]">{{ subjectSummary(item) }}</span>
        </template>

        <template #deadline="{ item }">
          <span class="text-[#6B5D45]">{{ formatDate(item.deadline) }}</span>
        </template>

        <template #total_price="{ item }">
          <span class="text-[#3A2E1F] font-medium">
            Rp {{ item.total_price?.toLocaleString("id-ID") }}
          </span>
        </template>

        <template #status="{ item }">
          <span
            class="inline-flex items-center gap-1.5 px-3 py-1 rounded-full text-xs font-semibold"
            :class="statusBadge(item.status)">
            <component :is="statusIcon(item.status)" :size="12" />
            {{ item.status.toUpperCase() }}
          </span>
        </template>
      </DataTablesComponent>
      <div class="flex justify-end">
        <PaginationComponent
          v-model:current-page="currentPage"
          v-model:per-page="perPage"
          :total-items="paginatedItems.length" />
      </div>
    </div>

    <CommissionDetailModal
      :visible="showDetailModal"
      :commission="selectedCommission"
      @hide="closeDetailModal" />
  </AppLayout>
</template>
