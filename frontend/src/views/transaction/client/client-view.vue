<script setup lang="ts">
import SidebarAppLayout from "../../../components/layout/sidebar-app-layout.vue";
import { ref, computed, onMounted } from "vue";
import {
  useClientApi,
  type ClientApi,
} from "../../../composables/commission/client_api.ts";
import { useToast } from "../../../composables/etc/useToast.ts";
import { useConfirm } from "../../../composables/etc/useConfirm.ts";
import { Pencil, Trash2, Plus } from "lucide-vue-next";
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
const { items, isLoading, fetchAll, create, update, remove } = useClientApi();

onMounted(async () => {
  isLoading.value = true;
  await new Promise((resolve) => setTimeout(resolve, 500));
  fetchAll();
});

const toast = useToast();
const { confirm } = useConfirm();
const search = ref("");
const currentPage = ref(1);
const perPage = ref(50);

const filteredItems = computed(() => {
  return items.value.filter((item) => {
    const matchSearch = item.name
      .toLowerCase()
      .includes(search.value.toLowerCase());
    return matchSearch;
  });
});

const paginatedItems = computed(() => {
  const start = (currentPage.value - 1) * perPage.value;
  return filteredItems.value.slice(start, start + perPage.value);
});

const columns: Column[] = [
  { key: "aksi", label: "Aksi", align: "center" },
  { key: "name", label: "Nama" },
  { key: "email", label: "Email" },
  { key: "phone", label: "Telepon" },
  { key: "address", label: "Alamat" },
  { key: "remark", label: "Catatan" },
];

// ===== State form modal =====
const showModal = ref(false);
const editingId = ref<number | null>(null); // null = mode Create, ada isi = mode Edit

const name = ref("");
const nickname = ref("");
const phone = ref("");
const email = ref("");
const instagram_handle = ref("");
const address = ref("");
const city = ref("");
const province = ref("");
const postal_code = ref("");
const denomination = ref("");
const patron_saint_preference = ref("");
const preferred_payment_method = ref("");

const isSubmitting = ref(false);
const formError = ref("");

const isEditMode = computed(() => editingId.value !== null);
const modalTitle = computed(() =>
  isEditMode.value ? "Edit Klien" : "Tambah Klien Baru"
);

function resetForm() {
  editingId.value = null;
  name.value = "";
  nickname.value = "";
  phone.value = "";
  email.value = "";
  instagram_handle.value = "";
  address.value = "";
  city.value = "";
  province.value = "";
  postal_code.value = "";
  denomination.value = "";
  patron_saint_preference.value = "";
  preferred_payment_method.value = "";
  formError.value = "";
}

function openCreateModal() {
  resetForm();
  showModal.value = true;
}

function openEditModal(item: ClientApi) {
  editingId.value = item.id;
  name.value = item.name;
  nickname.value = item.nickname;
  phone.value = item.phone;
  email.value = item.email;
  instagram_handle.value = item.instagram_handle;
  address.value = item.address;
  city.value = item.city;
  province.value = item.province;
  postal_code.value = item.postal_code;
  denomination.value = item.denomination;
  patron_saint_preference.value = item.patron_saint_preference;
  preferred_payment_method.value = item.preferred_payment_method;
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
    const payload = {
      name: name.value,
      nickname: nickname.value,
      phone: phone.value,
      email: email.value,
      instagram_handle: instagram_handle.value,
      address: address.value,
      city: city.value,
      province: province.value,
      postal_code: postal_code.value,
      denomination: denomination.value,
      patron_saint_preference: patron_saint_preference.value,
      preferred_payment_method: preferred_payment_method.value,
    };

    if (isEditMode.value && editingId.value !== null) {
      await update(editingId.value, {
        ...payload,
      });
      toast.success("Client berhasil diperbarui");
    } else {
      await create(payload);
      toast.success("Client baru berhasil ditambahkan");
    }
    closeModal();
  } catch (err: any) {
    formError.value = err.response?.data?.error || "Gagal menyimpan data";
  } finally {
    isSubmitting.value = false;
  }
}
async function handleDelete(item: ClientApi) {
  const confirmed = await confirm({
    title: "Hapus Klien",
    message: `Yakin ingin menghapus klien "${item.name}"? Tindakan ini tidak dapat dibatalkan.`,
    confirmLabel: "Ya, Hapus",
    danger: true,
  });

  if (!confirmed) return;

  try {
    await remove(item.id);
    toast.success("Klien berhasil dihapus");
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal menghapus data");
  }
}

function buildMenuItems(item: ClientApi) {
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
          Klien
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

        <div class="col-span-6"></div>
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
      </DataTablesComponent>

      <PaginationComponent
        v-model:current-page="currentPage"
        v-model:per-page="perPage"
        :total-items="filteredItems.length" />
    </div>

    <ChildModalWrapper
      :visible="showModal"
      :header-title="modalTitle"
      width-class="w-full max-w-6xl"
      @hide="closeModal">
      <form @submit.prevent="handleAdd" class="space-y-4">
        <div class="grid grid-cols-3 gap-4">
          <InputTextComponent
            v-model="name"
            label="Nama Lengkap"
            placeholder="Masukkan nama"
            required />

          <InputTextComponent
            v-model="nickname"
            label="Nickname"
            placeholder="Masukkan nickname" />

          <InputTextComponent
            v-model="phone"
            label="Telepon"
            placeholder="Masukkan nomor telepon" />

          <InputTextComponent
            v-model="email"
            label="Email"
            placeholder="Masukkan email" />

          <InputTextComponent
            v-model="instagram_handle"
            label="Instagram"
            placeholder="Masukkan handle instagram" />

          <InputTextComponent
            v-model="address"
            label="Alamat"
            placeholder="Masukkan alamat" />

          <InputTextComponent
            v-model="city"
            label="Kota"
            placeholder="Masukkan kota" />

          <InputTextComponent
            v-model="province"
            label="Provinsi"
            placeholder="Masukkan provinsi" />

          <InputTextComponent
            v-model="postal_code"
            label="Kode Pos"
            placeholder="Masukkan kode pos" />

          <InputTextComponent
            v-model="denomination"
            label="Denominasi"
            placeholder="Masukkan denominasi" />

          <InputTextComponent
            v-model="patron_saint_preference"
            label="Santo Pelindung"
            placeholder="Masukkan santo pelindung pilihan" />

          <InputTextComponent
            v-model="preferred_payment_method"
            label="Metode Pembayaran"
            placeholder="Masukkan metode pembayaran pilihan" />

          <div v-if="isEditMode"></div>
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
