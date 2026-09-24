<script setup lang="ts">
import SidebarAppLayout from "../../../components/layout/sidebar-app-layout.vue";
import { ref, computed, onMounted } from "vue";
import {
  useMasterProduct,
  type MasterProduct,
} from "../../../composables/master/m_product.ts";
import { useToast } from "../../../composables/etc/useToast.ts";
import { useConfirm } from "../../../composables/etc/useConfirm.ts";
import { Plus, Pencil, Trash2 } from "lucide-vue-next";
import ChildModalWrapper from "../../../components/modal/child-modal-wrapper.vue";
import DataTablesComponent from "../../../components/tables/data-tables-component.vue";
import PaginationComponent from "../../../components/pagination/pagination-component.vue";
import InputTextComponent from "../../../components/input-text/input-text-component.vue";
import OverlayMenu from "../../../components/button/overlay-menu.vue";

interface Column {
  key: string;
  label: string;
  align?: "left" | "center" | "right";
}

const { items, isLoading, fetchAll, create, update, remove } =
  useMasterProduct();

const toast = useToast();
const { confirm } = useConfirm();

const search = ref("");
const statusFilter = ref<"all" | "active" | "inactive">("all");
const currentPage = ref(1);
const perPage = ref(50);

const showModal = ref(false);
const editingId = ref<number | null>(null);

const name = ref("");
const unit = ref("");
const description = ref("");
const is_active = ref(true);

const isSubmitting = ref(false);
const formError = ref("");

const isEditMode = computed(() => editingId.value !== null);

const modalTitle = computed(() =>
  isEditMode.value ? "Edit Product" : "Tambah Product Baru"
);

/* =========================
   FETCH DATA
========================= */

onMounted(async () => {
  isLoading.value = true;

  try {
    await new Promise((resolve) => setTimeout(resolve, 500));
    await fetchAll();
  } finally {
    isLoading.value = false;
  }
});

/* =========================
   FILTER
========================= */

const filteredItems = computed(() => {
  const keyword = search.value.trim().toLowerCase();

  return items.value.filter((item) => {
    const matchSearch =
      keyword === "" || item.name.toLowerCase().includes(keyword);

    const matchStatus =
      statusFilter.value === "all" ||
      (statusFilter.value === "active" && item.is_active) ||
      (statusFilter.value === "inactive" && !item.is_active);

    return matchSearch && matchStatus;
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
   TABLE
========================= */

const columns: Column[] = [
  {
    key: "aksi",
    label: "Aksi",
    align: "center",
  },
  {
    key: "name",
    label: "Nama Product",
  },
  {
    key: "unit",
    label: "Unit",
  },
  {
    key: "description",
    label: "Deskripsi",
  },
  {
    key: "is_active",
    label: "Status",
  },
];

/* =========================
   STATUS
========================= */

function statusBadge(status: boolean) {
  return status
    ? "bg-[#4C8C5B]/15 text-[#4C8C5B]"
    : "bg-[#B23A32]/15 text-[#B23A32]";
}

/* =========================
   FORM
========================= */

function resetForm() {
  editingId.value = null;

  name.value = "";
  unit.value = "";
  description.value = "";
  is_active.value = true;

  formError.value = "";
}

function openCreateModal() {
  resetForm();
  showModal.value = true;
}

function openEditModal(item: MasterProduct) {
  editingId.value = item.id;

  name.value = item.name;
  unit.value = item.unit;
  description.value = item.description ?? "";
  is_active.value = item.is_active;

  formError.value = "";
  showModal.value = true;
}

function closeModal() {
  showModal.value = false;
  resetForm();
}

/* =========================
   CREATE / UPDATE
========================= */

async function handleAdd() {
  isSubmitting.value = true;
  formError.value = "";

  try {
    if (!name.value.trim()) {
      formError.value = "Nama product wajib diisi.";
      return;
    }

    if (!unit.value.trim()) {
      formError.value = "Unit wajib diisi.";
      return;
    }

    if (isEditMode.value && editingId.value !== null) {
      await update(editingId.value, {
        name: name.value.trim(),
        unit: unit.value.trim(),
        description: description.value.trim(),
        is_active: is_active.value,
      });

      toast.success("Product berhasil diperbarui");
    } else {
      await create({
        name: name.value.trim(),
        unit: unit.value.trim(),
        description: description.value.trim(),
      });

      toast.success("Product berhasil ditambahkan");
    }

    await fetchAll();

    closeModal();
  } catch (err: any) {
    formError.value =
      err?.response?.data?.message ||
      err?.response?.data?.error ||
      "Gagal menyimpan data product.";
  } finally {
    isSubmitting.value = false;
  }
}

/* =========================
   DELETE
========================= */

async function handleDelete(item: MasterProduct) {
  const confirmed = await confirm({
    title: "Hapus Product",
    message: `Yakin ingin menghapus product "${item.name}"? Tindakan ini tidak dapat dibatalkan.`,
    confirmLabel: "Ya, Hapus",
    danger: true,
  });

  if (!confirmed) return;

  try {
    await remove(item.id);

    toast.success("Product berhasil dihapus");

    await fetchAll();
  } catch (err: any) {
    toast.error(
      err?.response?.data?.message ||
        err?.response?.data?.error ||
        "Gagal menghapus product."
    );
  }
}

/* =========================
   MENU
========================= */

function buildMenuItems(item: MasterProduct) {
  return [
    {
      label: "Edit",
      icon: Pencil,
      command: () => openEditModal(item),
    },
    {
      label: "Hapus",
      icon: Trash2,
      danger: true,
      command: () => handleDelete(item),
    },
  ];
}
</script>

<template>
  <SidebarAppLayout>
    <div class="bg-white border border-[#E5D9BF] rounded-xl p-6 mb-6">
      <!-- HEADER -->
      <div class="flex items-center justify-between mb-6">
        <h1
          class="text-xl font-semibold text-[#3A2E1F] flex items-center gap-2">
          Produk
        </h1>

        <button
          @click="openCreateModal"
          class="flex items-center gap-2 px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition">
          <Plus :size="16" />
          Tambah
        </button>
      </div>

      <!-- FILTER -->
      <div class="grid grid-cols-12 gap-6 mb-6">
        <!-- SEARCH -->
        <div class="col-span-6">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Search
          </label>

          <input
            v-model="search"
            type="text"
            placeholder="Ketik untuk mencari..."
            class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition" />
        </div>

        <!-- STATUS -->
        <div class="col-span-6">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Status
          </label>

          <div class="flex items-center gap-4 h-10.5">
            <label
              class="flex items-center gap-2 text-sm text-[#3A2E1F] cursor-pointer">
              <input
                type="radio"
                value="all"
                v-model="statusFilter"
                class="accent-[#C9A24B]" />
              Semua
            </label>

            <label
              class="flex items-center gap-2 text-sm text-[#3A2E1F] cursor-pointer">
              <input
                type="radio"
                value="active"
                v-model="statusFilter"
                class="accent-[#C9A24B]" />
              Aktif
            </label>

            <label
              class="flex items-center gap-2 text-sm text-[#3A2E1F] cursor-pointer">
              <input
                type="radio"
                value="inactive"
                v-model="statusFilter"
                class="accent-[#C9A24B]" />
              Non-Aktif
            </label>
          </div>
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
        empty-message="Tidak ditemukan data."
        row-key="id">
        <!-- ACTION -->
        <template #aksi="{ item }">
          <OverlayMenu :items="buildMenuItems(item)" />
        </template>

        <!-- STATUS -->
        <template #is_active="{ item }">
          <span
            class="px-3 py-1 rounded-full text-xs font-semibold"
            :class="statusBadge(item.is_active)">
            {{ item.is_active ? "ACTIVE" : "INACTIVE" }}
          </span>
        </template>
      </DataTablesComponent>

      <!-- PAGINATION BOTTOM -->
      <PaginationComponent
        v-model:current-page="currentPage"
        v-model:per-page="perPage"
        :total-items="filteredItems.length" />
    </div>

    <!-- MODAL -->
    <ChildModalWrapper
      v-model:visible="showModal"
      :header-title="modalTitle"
      @hide="closeModal">
      <div class="space-y-5">
        <!-- ERROR -->
        <div
          v-if="formError"
          class="px-4 py-3 bg-[#B23A32]/10 text-[#B23A32] rounded-lg text-sm">
          {{ formError }}
        </div>

        <!-- NAME -->
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Nama Product
          </label>

          <InputTextComponent
            v-model="name"
            placeholder="Masukkan nama product" />
        </div>

        <!-- UNIT -->
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Unit
          </label>

          <InputTextComponent
            v-model="unit"
            placeholder="Contoh: Pcs, Kg, Liter" />
        </div>

        <!-- DESCRIPTION -->
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Deskripsi
          </label>

          <textarea
            v-model="description"
            rows="4"
            placeholder="Masukkan deskripsi product"
            class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition resize-none" />
        </div>

        <!-- STATUS -->
        <div v-if="isEditMode">
          <label class="block text-xs font-medium text-[#6B5D45] mb-2">
            Status
          </label>

          <div class="flex items-center gap-5">
            <label
              class="flex items-center gap-2 text-sm text-[#3A2E1F] cursor-pointer">
              <input
                v-model="is_active"
                type="radio"
                :value="true"
                class="accent-[#C9A24B]" />
              Aktif
            </label>

            <label
              class="flex items-center gap-2 text-sm text-[#3A2E1F] cursor-pointer">
              <input
                v-model="is_active"
                type="radio"
                :value="false"
                class="accent-[#C9A24B]" />
              Non-Aktif
            </label>
          </div>
        </div>

        <!-- BUTTON -->
        <div class="flex justify-end gap-3 pt-4">
          <button
            type="button"
            @click="closeModal"
            class="px-4 py-2 text-sm font-medium text-[#6B5D45] border border-[#D9CBB0] rounded-lg hover:bg-[#F8F4EC] transition">
            Batal
          </button>

          <button
            type="button"
            @click="handleAdd"
            :disabled="isSubmitting"
            class="px-4 py-2 text-sm font-medium text-white bg-[#7A1F2B] rounded-lg hover:bg-[#5F1621] disabled:opacity-50 disabled:cursor-not-allowed transition">
            {{
              isSubmitting
                ? "Menyimpan..."
                : isEditMode
                ? "Simpan Perubahan"
                : "Simpan"
            }}
          </button>
        </div>
      </div>
    </ChildModalWrapper>
  </SidebarAppLayout>
</template>
