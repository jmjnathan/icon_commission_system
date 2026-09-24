<script setup lang="ts">
import { computed, onMounted, ref, watch } from "vue";
import { useRouter } from "vue-router";
import { Plus, Eye, Trash2 } from "lucide-vue-next";

import SidebarAppLayout from "../../../components/layout/sidebar-app-layout.vue";
import DataTablesComponent from "../../../components/tables/data-tables-component.vue";
import PaginationComponent from "../../../components/pagination/pagination-component.vue";
import OverlayMenu from "../../../components/button/overlay-menu.vue";
import InputTextComponent from "../../../components/input-text/input-text-component.vue";

import { usePurchaseOrder } from "../../../composables/purchasing/purchase-order/usePurchaseOrder.ts";
import { useToast } from "../../../composables/etc/useToast.ts";
import { useConfirm } from "../../../composables/etc/useConfirm.ts";

interface Column {
  key: string;
  label: string;
  align?: "left" | "center" | "right";
}

const router = useRouter();

const { items, isLoading, fetchAll, remove } = usePurchaseOrder();

const toast = useToast();
const { confirm } = useConfirm();

const activeTab = ref<"draft" | "complete">("draft");

const documentNoFilter = ref("");
const dateFrom = ref("");
const dateTo = ref("");

const currentPage = ref(1);
const perPage = ref(50);

/* =========================
   FETCH
========================= */

onMounted(async () => {
  await fetchAll();
});

/* =========================
   FILTER
========================= */

const filteredItems = computed(() => {
  const documentNo = documentNoFilter.value.trim().toLowerCase();

  return items.value.filter((item) => {
    const statusMatch =
      activeTab.value === "draft"
        ? item.status === "draft"
        : item.status === "ordered";

    const documentNoMatch =
      documentNo === "" || item.document_no.toLowerCase().includes(documentNo);

    const orderDate = item.order_date.substring(0, 10);

    const dateFromMatch = dateFrom.value === "" || orderDate >= dateFrom.value;

    const dateToMatch = dateTo.value === "" || orderDate <= dateTo.value;

    return statusMatch && documentNoMatch && dateFromMatch && dateToMatch;
  });
});

/* =========================
   PAGINATION
========================= */

const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * perPage.value;

  return filteredItems.value.slice(start, start + perPage.value);
});

/* =========================
   WATCH FILTER
========================= */

watch([activeTab, documentNoFilter, dateFrom, dateTo], () => {
  currentPage.value = 1;
});

/* =========================
   TABLE
========================= */

const columns: Column[] = [
  {
    key: "aksi",
    label: "Aksi",
    align: "center",
  },
  {
    key: "status",
    label: "Status",
    align: "center",
  },
  {
    key: "document_no",
    label: "No. Dokumen",
  },
  {
    key: "order_date",
    label: "Tanggal",
  },
  {
    key: "vendor_name",
    label: "Vendor",
  },
  {
    key: "total",
    label: "Total",
    align: "right",
  },
];

/* =========================
   FORMAT
========================= */

function formatDate(value: string) {
  if (!value) return "-";

  const date = new Date(value);

  return date.toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
}

function formatCurrency(value: number) {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    maximumFractionDigits: 0,
  }).format(value);
}

function statusBadge(status: string) {
  if (status === "draft") {
    return "bg-[#C9A24B]/15 text-[#9A761F]";
  }

  if (status === "ordered") {
    return "bg-[#4C8C5B]/15 text-[#4C8C5B]";
  }

  return "bg-[#8A7A60]/15 text-[#6B5D45]";
}

function statusLabel(status: string) {
  if (status === "draft") {
    return "DRAFT";
  }

  if (status === "ordered") {
    return "COMPLETE";
  }

  return status.toUpperCase();
}

/* =========================
   NAVIGATION
========================= */

function handleCreate() {
  router.push("/purchasing/purchase-order/form");
}

function handleDetail(item: any) {
  router.push(`/purchasing/purchase-order/form/${item.id}`);
}

/* =========================
   DELETE
========================= */

async function handleDelete(item: any) {
  const confirmed = await confirm({
    title: "Hapus Purchase Order",
    message: `Yakin ingin menghapus dokumen "${item.document_no}"? Tindakan ini tidak dapat dibatalkan.`,
    confirmLabel: "Ya, Hapus",
    danger: true,
  });

  if (!confirmed) return;

  try {
    await remove(item.id);

    toast.success("Purchase Order berhasil dihapus");

    await fetchAll();
  } catch (err: any) {
    toast.error(
      err?.response?.data?.message ||
        err?.response?.data?.error ||
        "Gagal menghapus Purchase Order."
    );
  }
}

/* =========================
   MENU
========================= */

function buildMenuItems(item: any) {
  const menu = [
    {
      label: "Detail",
      icon: Eye,
      command: () => handleDetail(item),
    },
  ];

  if (item.status === "draft") {
    menu.push({
      label: "Delete",
      icon: Trash2,
      command: () => handleDelete(item),
    });
  }

  return menu;
}
</script>

<template>
  <SidebarAppLayout>
    <div class="bg-white border border-[#E5D9BF] rounded-xl p-6 mb-6">
      <!-- HEADER -->

      <div class="flex items-center justify-between mb-6">
        <h1 class="text-xl font-semibold text-[#3A2E1F]">
          Pemesanan Barang Baku
        </h1>

        <button
          type="button"
          @click="handleCreate"
          class="flex items-center gap-2 px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition">
          <Plus :size="16" />

          Buat Pemesanan
        </button>
      </div>

      <!-- TAB -->

      <div class="flex items-center gap-1 border-b border-[#E5D9BF] mb-6">
        <button
          type="button"
          @click="activeTab = 'draft'"
          class="px-5 py-3 text-sm font-medium border-b-2 transition"
          :class="
            activeTab === 'draft'
              ? 'text-[#7A1F2B] border-[#7A1F2B]'
              : 'text-[#8A7A60] border-transparent hover:text-[#3A2E1F]'
          ">
          Draft
        </button>

        <button
          type="button"
          @click="activeTab = 'complete'"
          class="px-5 py-3 text-sm font-medium border-b-2 transition"
          :class="
            activeTab === 'complete'
              ? 'text-[#7A1F2B] border-[#7A1F2B]'
              : 'text-[#8A7A60] border-transparent hover:text-[#3A2E1F]'
          ">
          Complete
        </button>
      </div>

      <!-- FILTER -->

      <div class="grid grid-cols-12 gap-6 mb-6">
        <div class="col-span-5">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            No. Dokumen
          </label>

          <InputTextComponent
            v-model="documentNoFilter"
            placeholder="Cari nomor dokumen..." />
        </div>

        <div class="col-span-3">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Dari Tanggal
          </label>

          <input
            v-model="dateFrom"
            type="date"
            class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition" />
        </div>

        <div class="col-span-3">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Sampai Tanggal
          </label>

          <input
            v-model="dateTo"
            type="date"
            class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition" />
        </div>

        <div class="col-span-1 flex items-end">
          <button
            type="button"
            @click="
              documentNoFilter = '';
              dateFrom = '';
              dateTo = '';
            "
            class="w-full px-3 py-2.5 text-sm font-medium text-[#6B5D45] border border-[#D9CBB0] rounded-lg hover:bg-[#F8F4EC] transition">
            Reset
          </button>
        </div>
      </div>

      <!-- PAGINATION TOP -->

      <div class="flex justify-end mb-2">
        <PaginationComponent
          v-model:current-page="currentPage"
          v-model:per-page="perPage"
          :total-items="filteredItems.length" />
      </div>

      <!-- TABLE -->

      <DataTablesComponent
        :columns="columns"
        :items="paginatedItems"
        :is-loading="isLoading"
        empty-message="Tidak ditemukan dokumen."
        row-key="id">
        <!-- ACTION -->

        <template #aksi="{ item }">
          <OverlayMenu :items="buildMenuItems(item)" />
        </template>

        <!-- STATUS -->

        <template #status="{ item }">
          <span
            class="inline-flex items-center px-2 py-1 rounded-full text-xs font-semibold"
            :class="statusBadge(item.status)">
            {{ statusLabel(item.status) }}
          </span>
        </template>

        <!-- DATE -->

        <template #order_date="{ item }">
          {{ formatDate(item.order_date) }}
        </template>

        <!-- TOTAL -->

        <template #total="{ item }">
          <span class="font-medium text-[#3A2E1F]">
            {{ formatCurrency(item.total) }}
          </span>
        </template>
      </DataTablesComponent>

      <!-- PAGINATION BOTTOM -->

      <PaginationComponent
        v-model:current-page="currentPage"
        v-model:per-page="perPage"
        :total-items="filteredItems.length" />
    </div>
  </SidebarAppLayout>
</template>
