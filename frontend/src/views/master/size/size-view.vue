<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import SidebarAppLayout from "../../../components/layout/sidebar-app-layout.vue";
import DataTablesComponent from "../../../components/tables/data-tables-component.vue";
import PaginationComponent from "../../../components/pagination/pagination-component.vue";
import {
  useMasterSize,
  type MasterSize,
} from "../../../composables/master/m_size.ts";
import ChildModalWrapper from "../../../components/modal/child-modal-wrapper.vue";
import OverlayMenu from "../../../components/button/overlay-menu.vue";
import { Plus, Pencil, Trash2 } from "lucide-vue-next";
import { useToast } from "../../../composables/etc/useToast.ts";
import { useConfirm } from "../../../composables/etc/useConfirm.ts";

interface Column {
  key: string;
  label: string;
  align?: "left" | "center" | "right";
}

const { items, isLoading, fetchAll, create, update, remove } = useMasterSize();

onMounted(fetchAll);
const toast = useToast();
const { confirm } = useConfirm();
const search = ref("");
const statusFilter = ref<"all" | "active" | "inactive">("all");
const currentPage = ref(1);
const perPage = ref(50);

const filteredItems = computed(() => {
  return items.value.filter((item) => {
    const matchSearch =
      item.name.toLowerCase().includes(search.value.toLowerCase()) ||
      item.size.toLowerCase().includes(search.value.toLowerCase());
    const matchStatus =
      statusFilter.value === "all" ||
      (statusFilter.value === "active" &&
        item.status.toLowerCase() === "active") ||
      (statusFilter.value === "inactive" &&
        item.status.toLowerCase() !== "active");
    return matchSearch && matchStatus;
  });
});

const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * perPage.value;
  return filteredItems.value.slice(start, start + perPage.value);
});

const columns: Column[] = [
  { key: "aksi", label: "Aksi", align: "center" },
  { key: "name", label: "Kode" },
  { key: "size", label: "Ukuran" },
  { key: "status", label: "Status" },
];

function statusBadge(status: string) {
  return status.toLowerCase() === "active"
    ? "bg-[#4C8C5B]/15 text-[#4C8C5B]"
    : "bg-[#B23A32]/15 text-[#B23A32]";
}

// ===== State form modal =====
const showModal = ref(false);
const editingId = ref<number | null>(null); // null = mode Create, ada isi = mode Edit
const name = ref("");
const size = ref("");
const status = ref("Active");
const isSubmitting = ref(false);
const formError = ref("");

const isEditMode = computed(() => editingId.value !== null);
const modalTitle = computed(() =>
  isEditMode.value ? "Edit Ukuran" : "Tambah Ukuran Baru"
);

function resetForm() {
  editingId.value = null;
  name.value = "";
  size.value = "";
  status.value = "Active";
  formError.value = "";
}

function openCreateModal() {
  resetForm();
  showModal.value = true;
}

function openEditModal(item: MasterSize) {
  editingId.value = item.id;
  name.value = item.name;
  size.value = item.size;
  status.value = item.status;
  formError.value = "";
  showModal.value = true;
}

function closeModal() {
  showModal.value = false;
  resetForm();
}

async function handleAdd() {
  isSubmitting.value = true;
  formError.value = "";
  try {
    if (isEditMode.value && editingId.value !== null) {
      await update(editingId.value, {
        name: name.value,
        size: size.value,
        status: status.value,
      });
      toast.success("Ukuran berhasil diperbarui");
    } else {
      await create({ name: name.value, size: size.value });
      toast.success("Ukuran baru berhasil ditambahkan");
    }
    closeModal();
  } catch (err: any) {
    formError.value = err.response?.data?.error || "Gagal menyimpan data";
  } finally {
    isSubmitting.value = false;
  }
}
async function handleDelete(item: MasterSize) {
  const confirmed = await confirm({
    title: "Hapus Ukuran",
    message: `Yakin ingin menghapus ukuran "${item.name}"? Tindakan ini tidak dapat dibatalkan.`,
    confirmLabel: "Ya, Hapus",
    danger: true,
  });

  if (!confirmed) return;

  try {
    await remove(item.id);
    toast.success("Ukuran berhasil dihapus");
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal menghapus data");
  }
}

function buildMenuItems(item: MasterSize) {
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
      <div class="flex items-center justify-between mb-6">
        <h1
          class="text-xl font-semibold text-[#3A2E1F] flex items-center gap-2">
          Ukuran
        </h1>
        <button
          @click="openCreateModal"
          class="flex items-center gap-2 px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition">
          <Plus :size="16" />
          Tambah
        </button>
      </div>

      <div class="grid grid-cols-12 gap-6 mb-6">
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
      <div class="flex justify-end mb-2">
        <PaginationComponent
          v-model:current-page="currentPage"
          v-model:per-page="perPage"
          :total-items="filteredItems.length" />
      </div>

      <DataTablesComponent
        :columns="columns"
        :items="paginatedItems"
        :is-loading="isLoading"
        empty-message="Tidak ditemukan data."
        row-key="id">
        <template #aksi="{ item }">
          <OverlayMenu :items="buildMenuItems(item)" />
        </template>
        <template #status="{ item }">
          <span
            class="px-3 rounded-full text-xs font-semibold"
            :class="statusBadge(item.status)">
            {{ item.status.toUpperCase() }}
          </span>
        </template>
      </DataTablesComponent>

      <PaginationComponent
        v-model:current-page="currentPage"
        v-model:per-page="perPage"
        :total-items="filteredItems.length" />
    </div>

    <ChildModalWrapper
      :visible="showModal"
      :header-title="modalTitle"
      width-class="w-full max-w-md"
      @hide="closeModal">
      <form @submit.prevent="handleAdd" class="space-y-4">
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
            >Kode</label
          >
          <input
            v-model="name"
            type="text"
            required
            class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
        </div>
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
            >Ukuran</label
          >
          <input
            v-model="size"
            type="text"
            required
            class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
        </div>

        <!-- Field Status cuma muncul pas mode Edit, karena Create selalu default Active -->
        <div v-if="isEditMode">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
            >Status</label
          >
          <select
            v-model="status"
            class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40">
            <option value="Active">Active</option>
            <option value="Inactive">Inactive</option>
          </select>
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
  </SidebarAppLayout>
</template>
