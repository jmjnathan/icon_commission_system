<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import {
  Plus,
  Wallet,
  TrendingDown,
  TrendingUp,
  Pencil,
  Trash2,
} from "lucide-vue-next";
import AppLayout from "../../../components/layout/sidebar-app-layout.vue";
import DataTablesComponent from "../../../components/tables/data-tables-component.vue";
import ChildModalWrapper from "../../../components/modal/child-modal-wrapper.vue";
import OverlayMenu from "../../../components/button/overlay-menu.vue";
import { useCashOut, type CashOut } from "../../../composables/useCashout.ts";
import { useCommission } from "../../../composables/commission/useCommission.ts";
import { useToast } from "../../../composables/etc/useToast";
import { useConfirm } from "../../../composables/etc/useConfirm";
import PaginationComponent from "../../../components/pagination/pagination-component.vue";

interface Column {
  key: string;
  label: string;
  align?: "left" | "center" | "right";
}

const { items, isLoading, fetchAll, create, update, remove } = useCashOut();
const { commissions, fetchAll: fetchCommissions } = useCommission();
const toast = useToast();
const { confirm } = useConfirm();

onMounted(async () => {
  isLoading.value = true;
  await new Promise((resolve) => setTimeout(resolve, 500));
  fetchAll();
  fetchCommissions();
});

// ===== Ringkasan Saldo =====
const totalIncome = computed(() =>
  commissions.value
    .filter((c) => {
      if (c.status !== "completed" && c.status !== "delivered") return false;
      const commissionDate = c.order_date.split("T")[0];
      const matchStart =
        !filterStartDate.value || commissionDate >= filterStartDate.value;
      const matchEnd =
        !filterEndDate.value || commissionDate <= filterEndDate.value;
      return matchStart && matchEnd;
    })
    .reduce((sum, c) => sum + (c.total_price || 0), 0)
);

const totalExpense = computed(() =>
  filteredItems.value.reduce((sum, i) => sum + i.amount, 0)
);

const balance = computed(() => totalIncome.value - totalExpense.value);

// Tabel
const columns: Column[] = [
  { key: "aksi", label: "Aksi", align: "center" },
  { key: "date", label: "Tanggal" },
  { key: "category", label: "Kategori" },
  { key: "description", label: "Keterangan" },
  { key: "qty", label: "Qty", align: "center" },
  { key: "unit_price", label: "Harga Satuan", align: "right" },
  { key: "amount", label: "Subtotal", align: "right" },
  { key: "notes", label: "Catatan", align: "left" },
];

function formatDate(dateStr: string) {
  return new Date(dateStr).toLocaleDateString("id-ID", {
    day: "numeric",
    month: "short",
    year: "numeric",
  });
}

function categoryBadge(category: string) {
  const map: Record<string, string> = {
    "Bahan Baku": "bg-[#C9A24B]/15 text-[#8A6D1F]",
    Operasional: "bg-[#3B6FA8]/15 text-[#3B6FA8]",
    Lainnya: "bg-[#8A7A5C]/15 text-[#8A7A5C]",
  };
  return map[category] || "bg-gray-100 text-gray-600";
}

// ===== Form Modal =====
const showModal = ref(false);
const editingId = ref<number | null>(null);
const category = ref("Bahan Baku");
const description = ref("");
const amount = ref<number | null>(null);
const date = ref("");
const isSubmitting = ref(false);
const formError = ref("");
const qty = ref<number | null>(1);
const unitPrice = ref<number | null>(null);
const vendor = ref("");
const paymentMethod = ref("Cash");
const receiptUrls = ref<string[]>([]);
const notes = ref("");
const paymentOptions = ["Cash", "Transfer", "QRIS"];

const subtotal = computed(() => (qty.value ?? 0) * (unitPrice.value ?? 0));

const isEditMode = computed(() => editingId.value !== null);
const modalTitle = computed(() =>
  isEditMode.value ? "Edit Pengeluaran" : "Tambah Pengeluaran"
);

const categoryOptions = ["Bahan Baku", "Operasional", "Lainnya"];

function resetForm() {
  editingId.value = null;
  category.value = "Bahan Baku";
  description.value = "";
  qty.value = 1;
  amount.value = 0;
  unitPrice.value = null;
  vendor.value = "";
  paymentMethod.value = "Cash";
  receiptUrls.value = [];
  notes.value = "";
  date.value = "";
  formError.value = "";
}

function openCreateModal() {
  resetForm();
  showModal.value = true;
}

function openEditModal(item: CashOut) {
  editingId.value = item.id;
  category.value = item.category;
  description.value = item.description;
  qty.value = item.qty;
  unitPrice.value = item.unit_price;
  vendor.value = item.vendor;
  paymentMethod.value = item.payment_method || "Cash";
  receiptUrls.value = item.receipt_url ? [item.receipt_url] : [];
  notes.value = item.notes;
  date.value = item.date.split("T")[0];
  formError.value = "";
  showModal.value = true; // hapus baris "amount = item.amount;"
}

function closeModal() {
  showModal.value = false;
  resetForm();
}

async function handleSubmit() {
  if (!description.value || !qty.value || !unitPrice.value || !date.value) {
    formError.value =
      "Field wajib (Keterangan, Qty, Harga Satuan, Tanggal) harus diisi";
    return;
  }

  isSubmitting.value = true;
  formError.value = "";
  try {
    const payload = {
      category: category.value,
      description: description.value,
      qty: qty.value,
      unit_price: unitPrice.value,
      vendor: vendor.value,
      payment_method: paymentMethod.value,
      receipt_url: receiptUrls.value[0] || "",
      notes: notes.value,
      date: date.value,
    };
    if (isEditMode.value && editingId.value !== null) {
      await update(editingId.value, payload);
      toast.success("Pengeluaran berhasil diperbarui");
    } else {
      await create(payload);
      toast.success("Pengeluaran baru berhasil dicatat");
    }
    closeModal();
  } catch (err: any) {
    formError.value = err.response?.data?.error || "Gagal menyimpan data";
  } finally {
    isSubmitting.value = false;
  }
}

async function handleDelete(item: CashOut) {
  const confirmed = await confirm({
    title: "Hapus Pengeluaran",
    message: `Yakin ingin menghapus catatan "${item.description}"?`,
    confirmLabel: "Ya, Hapus",
    danger: true,
  });
  if (!confirmed) return;

  try {
    await remove(item.id);
    toast.success("Pengeluaran berhasil dihapus");
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal menghapus data");
  }
}

function buildMenuItems(item: CashOut) {
  return [
    { label: "Edit", icon: Pencil, command: () => openEditModal(item) },
    {
      label: "Hapus",
      icon: Trash2,
      danger: true,
      command: () => handleDelete(item),
    },
  ];
}

const filterStartDate = ref("");
const filterEndDate = ref("");

const filteredItems = computed(() => {
  return items.value.filter((item) => {
    const itemDate = item.date.split("T")[0]; // ambil bagian YYYY-MM-DD aja
    const matchStart =
      !filterStartDate.value || itemDate >= filterStartDate.value;
    const matchEnd = !filterEndDate.value || itemDate <= filterEndDate.value;
    return matchStart && matchEnd;
  });
});

function resetDateFilter() {
  filterStartDate.value = "";
  filterEndDate.value = "";
}
function setThisMonth() {
  const now = new Date();
  const first = new Date(now.getFullYear(), now.getMonth(), 1);
  const last = new Date(now.getFullYear(), now.getMonth() + 1, 0);
  filterStartDate.value = first.toISOString().split("T")[0];
  filterEndDate.value = last.toISOString().split("T")[0];
}

const currentPage = ref(1);
const perPage = ref(50);
</script>

<template>
  <AppLayout>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-semibold text-[#3A2E1F] mb-1">
          Arus Kas Keluar
        </h1>
        <p class="text-sm text-[#8A7A5C]">
          Catat pengeluaran yang menggunakan pemasukan dari transaksi komisi.
        </p>
      </div>
      <button
        @click="openCreateModal"
        class="flex items-center gap-2 px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition">
        <Plus :size="16" />
        Catat Pengeluaran
      </button>
    </div>

    <!-- Ringkasan Saldo -->
    <div class="grid grid-cols-3 gap-4 mb-8">
      <!-- Skeleton -->
      <template v-if="isLoading">
        <div
          v-for="n in 3"
          :key="`skeleton-${n}`"
          class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between animate-pulse">
          <div class="flex-1">
            <div class="h-3 w-24 bg-[#E5D9BF] rounded mb-2"></div>
            <div class="h-5 w-28 bg-[#E5D9BF] rounded"></div>
          </div>
          <div class="w-10 h-10 rounded-full bg-[#E5D9BF] shrink-0"></div>
        </div>
      </template>

      <!-- Data asli -->
      <template v-else>
        <div
          class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between">
          <div>
            <p class="text-xs font-medium text-[#4C8C5B] tracking-wide">
              TOTAL PEMASUKAN
            </p>
            <p class="text-lg font-semibold text-[#3A2E1F] mt-1">
              Rp {{ totalIncome.toLocaleString("id-ID") }}
            </p>
          </div>
          <div
            class="w-10 h-10 rounded-full bg-[#4C8C5B]/15 flex items-center justify-center">
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
              Rp {{ totalExpense.toLocaleString("id-ID") }}
            </p>
          </div>
          <div
            class="w-10 h-10 rounded-full bg-[#B23A32]/15 flex items-center justify-center">
            <TrendingDown :size="20" class="text-[#B23A32]" />
          </div>
        </div>

        <div
          class="bg-white border border-[#E5D9BF] rounded-xl p-4 flex items-center justify-between">
          <div>
            <p class="text-xs font-medium text-[#C9A24B] tracking-wide">
              SALDO
            </p>
            <p class="text-lg font-semibold text-[#3A2E1F] mt-1">
              Rp {{ balance.toLocaleString("id-ID") }}
            </p>
          </div>
          <div
            class="w-10 h-10 rounded-full bg-[#C9A24B]/15 flex items-center justify-center">
            <Wallet :size="20" class="text-[#C9A24B]" />
          </div>
        </div>
      </template>
    </div>
    <div class="bg-white border border-[#E5D9BF] rounded-xl p-4 mb-6">
      <div
        class="flex flex-col lg:flex-row lg:items-end lg:justify-between gap-4">
        <!-- Filter -->
        <div class="flex flex-wrap items-end gap-4">
          <div>
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
              Dari Tanggal
            </label>
            <input
              v-model="filterStartDate"
              type="date"
              class="px-4 py-2 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
          </div>

          <div>
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
              Sampai Tanggal
            </label>
            <input
              v-model="filterEndDate"
              type="date"
              class="px-4 py-2 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
          </div>

          <button
            @click="setThisMonth"
            class="px-3 py-2 text-sm text-[#4C8C5B] border border-[#7A1F2B]/30 rounded-lg hover:bg-[#7A1F2B]/5 transition">
            Atur Ke Bulan Ini
          </button>

          <button
            v-if="filterStartDate || filterEndDate"
            @click="resetDateFilter"
            class="flex items-center gap-1.5 px-3 py-2 text-sm text-[#8A7A5C] hover:text-[#3A2E1F] transition">
            <XIcon :size="14" />
            Reset Filter
          </button>
        </div>

        <!-- Pagination -->
        <div class="flex justify-end lg:ml-auto">
          <PaginationComponent
            v-model:current-page="currentPage"
            v-model:per-page="perPage"
            :total-items="filteredItems.length" />
        </div>
      </div>
      <div class="mt-4">
        <DataTablesComponent
          :columns="columns"
          :items="filteredItems"
          :is-loading="isLoading"
          empty-message="Belum ada catatan pengeluaran."
          row-key="id">
          <template #aksi="{ item }">
            <OverlayMenu :items="buildMenuItems(item)" />
          </template>
          <template #date="{ item }">
            <span class="text-[#6B5D45]">{{ formatDate(item.date) }}</span>
          </template>
          <template #category="{ item }">
            <span
              class="px-3 py-1 rounded-full text-xs font-semibold"
              :class="categoryBadge(item.category)">
              {{ item.category }}
            </span>
          </template>
          <template #description="{ item }">
            <span class="text-[#3A2E1F]">{{ item.description }}</span>
          </template>
          <template #amount="{ item }">
            <span class="text-[#B23A32] font-medium"
              >- Rp {{ item.amount.toLocaleString("id-ID") }}</span
            >
          </template>
          <template #qty="{ item }">
            <span class="text-[#6B5D45]">{{ item.qty }}</span>
          </template>
          <template #unit_price="{ item }">
            <span class="text-[#6B5D45]"
              >Rp {{ item.unit_price.toLocaleString("id-ID") }}</span
            >
          </template>
        </DataTablesComponent>
      </div>
    </div>

    <ChildModalWrapper
      :visible="showModal"
      :header-title="modalTitle"
      width-class="w-full max-w-md"
      @hide="closeModal">
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
            >Kategori</label
          >
          <select
            v-model="category"
            class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40">
            <option v-for="opt in categoryOptions" :key="opt" :value="opt">
              {{ opt }}
            </option>
          </select>
        </div>

        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
            >Keterangan</label
          >
          <input
            v-model="description"
            type="text"
            required
            placeholder="Misal: Kanvas ukuran 6R"
            class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
              >Qty</label
            >
            <input
              v-model.number="qty"
              type="number"
              min="1"
              required
              class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
          </div>
          <div>
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
              >Harga Satuan</label
            >
            <input
              v-model.number="unitPrice"
              type="number"
              required
              class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
          </div>
        </div>

        <div
          class="bg-[#F5EFE3] rounded-lg px-4 py-2.5 flex items-center justify-between">
          <span class="text-xs text-[#8A7A5C]">Subtotal</span>
          <span class="text-sm font-semibold text-[#3A2E1F]"
            >Rp {{ subtotal.toLocaleString("id-ID") }}</span
          >
        </div>

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
              >Vendor / Toko</label
            >
            <input
              v-model="vendor"
              type="text"
              placeholder="Opsional"
              class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
          </div>
          <div>
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
              >Metode Bayar</label
            >
            <select
              v-model="paymentMethod"
              class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40">
              <option v-for="opt in paymentOptions" :key="opt" :value="opt">
                {{ opt }}
              </option>
            </select>
          </div>
        </div>

        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
            >Tanggal</label
          >
          <input
            v-model="date"
            type="date"
            required
            class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
        </div>

        <PhotoPicker v-model="receiptUrls" />

        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
            >Catatan</label
          >
          <input
            v-model="notes"
            type="text"
            placeholder="Opsional"
            class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
        </div>

        <p
          v-if="formError"
          class="text-sm text-[#B23A32] bg-[#B23A32]/10 px-3 py-2 rounded-lg">
          {{ formError }}
        </p>

        <button
          type="submit"
          :disabled="isSubmitting"
          class="w-full py-2.5 bg-[#7A1F2B] text-white rounded-lg font-medium hover:bg-[#5F1621] transition disabled:opacity-50">
          {{ isSubmitting ? "Menyimpan..." : "Simpan" }}
        </button>
      </form>
    </ChildModalWrapper>
  </AppLayout>
</template>
