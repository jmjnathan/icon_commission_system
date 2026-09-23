<script setup lang="ts">
import { computed, onMounted, ref } from "vue";
import SidebarAppLayout from "../../../components/layout/sidebar-app-layout.vue";
import DataTablesComponent from "../../../components/tables/data-tables-component.vue";
import PaginationComponent from "../../../components/pagination/pagination-component.vue";
import ChildModalWrapper from "../../../components/modal/child-modal-wrapper.vue";
import InputTextComponent from "../../../components/input-text/input-text-component.vue";
import InputNumberComponent from "../../../components/input-text/input-number-component.vue";
import OverlayMenu from "../../../components/button/overlay-menu.vue";

import { Plus, Pencil, Trash2, Settings2 } from "lucide-vue-next";

import {
  useMaterialComponent,
  type MaterialComponent,
  type MaterialComponentVariant,
} from "../../../composables/master/m_material_component";

import { useToast } from "../../../composables/etc/useToast";
import { useConfirm } from "../../../composables/etc/useConfirm";

const {
  items,
  variants,
  isLoading,
  isVariantLoading,

  fetchAll,
  create,
  update,
  remove,

  fetchVariants,
  createVariant,
  updateVariant,
  removeVariant,
} = useMaterialComponent();

const toast = useToast();
const { confirm } = useConfirm();

onMounted(fetchAll);

const search = ref("");
const statusFilter = ref<"all" | "active" | "inactive">("all");

const currentPage = ref(1);
const perPage = ref(50);

const filteredItems = computed(() => {
  return items.value.filter((item) => {
    const keyword = search.value.toLowerCase();

    const matchSearch =
      item.name.toLowerCase().includes(keyword) ||
      item.category.toLowerCase().includes(keyword) ||
      item.remark.toLowerCase().includes(keyword);

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

function formatRupiah(value: number) {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    maximumFractionDigits: 0,
  }).format(value);
}

function formatSpecification(variant: MaterialComponentVariant) {
  if (variant.length || variant.width || variant.thickness) {
    const dimensions = [variant.length, variant.width, variant.thickness]
      .filter((value) => value > 0)
      .join(" × ");

    return `${dimensions} cm`;
  }

  return variant.specification || "-";
}

const showComponentModal = ref(false);
const editingComponentId = ref<number | null>(null);

const componentName = ref("");
const componentCategory = ref("");
const componentRemark = ref("");
const componentStatus = ref("Active");

const showVariantModal = ref(false);
const editingVariantId = ref<number | null>(null);

const selectedComponent = ref<MaterialComponent | null>(null);

const variantSpecification = ref("");
const variantLength = ref<number | null>(null);
const variantWidth = ref<number | null>(null);
const variantThickness = ref<number | null>(null);
const variantUnit = ref("pcs");
const variantUnitPrice = ref<number | null>(null);
const variantRemark = ref("");
const variantStatus = ref("Active");

const isComponentEditMode = computed(() => editingComponentId.value !== null);

const isVariantEditMode = computed(() => editingVariantId.value !== null);

function resetComponentForm() {
  editingComponentId.value = null;
  componentName.value = "";
  componentCategory.value = "";
  componentRemark.value = "";
  componentStatus.value = "Active";
}

function openCreateComponent() {
  resetComponentForm();
  showComponentModal.value = true;
}

function openEditComponent(item: MaterialComponent) {
  editingComponentId.value = item.id;
  componentName.value = item.name;
  componentCategory.value = item.category;
  componentRemark.value = item.remark;
  componentStatus.value = item.status;

  showComponentModal.value = true;
}

async function submitComponent() {
  try {
    if (isComponentEditMode.value && editingComponentId.value !== null) {
      await update(editingComponentId.value, {
        name: componentName.value,
        category: componentCategory.value,
        remark: componentRemark.value,
        status: componentStatus.value,
      });

      toast.success("Komponen bahan berhasil diperbarui");
    } else {
      await create({
        name: componentName.value,
        category: componentCategory.value,
        remark: componentRemark.value,
      });

      toast.success("Komponen bahan berhasil ditambahkan");
    }

    showComponentModal.value = false;
    resetComponentForm();
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal menyimpan komponen");
  }
}

async function openVariants(item: MaterialComponent) {
  selectedComponent.value = item;

  await fetchVariants(item.id);

  showVariantModal.value = true;
}

function resetVariantForm() {
  editingVariantId.value = null;

  variantSpecification.value = "";
  variantLength.value = null;
  variantWidth.value = null;
  variantThickness.value = null;
  variantUnit.value = "pcs";
  variantUnitPrice.value = null;
  variantRemark.value = "";
  variantStatus.value = "Active";
}

function openCreateVariant() {
  resetVariantForm();
}

function openEditVariant(variant: MaterialComponentVariant) {
  editingVariantId.value = variant.id;

  variantSpecification.value = variant.specification;

  variantLength.value = variant.length || null;

  variantWidth.value = variant.width || null;

  variantThickness.value = variant.thickness || null;

  variantUnit.value = variant.unit;
  variantUnitPrice.value = variant.unit_price;
  variantRemark.value = variant.remark;
  variantStatus.value = variant.status;
}

async function submitVariant() {
  if (!selectedComponent.value) return;

  if (variantUnitPrice.value === null) {
    toast.error("Harga wajib diisi");
    return;
  }

  try {
    const payload = {
      specification: variantSpecification.value,

      length: variantLength.value || 0,
      width: variantWidth.value || 0,
      thickness: variantThickness.value || 0,

      unit: variantUnit.value,
      unit_price: variantUnitPrice.value,
      remark: variantRemark.value,
    };

    if (isVariantEditMode.value && editingVariantId.value !== null) {
      await updateVariant(editingVariantId.value, selectedComponent.value.id, {
        ...payload,
        status: variantStatus.value,
      });

      toast.success("Spesifikasi berhasil diperbarui");
    } else {
      await createVariant(selectedComponent.value.id, payload);

      toast.success("Spesifikasi berhasil ditambahkan");
    }

    resetVariantForm();
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal menyimpan spesifikasi");
  }
}

async function handleDeleteVariant(variant: MaterialComponentVariant) {
  if (!selectedComponent.value) return;

  const confirmed = await confirm({
    title: "Hapus Spesifikasi",
    message: "Yakin ingin menghapus spesifikasi ini?",
    confirmLabel: "Ya, Hapus",
    danger: true,
  });

  if (!confirmed) return;

  try {
    await removeVariant(variant.id, selectedComponent.value.id);

    toast.success("Spesifikasi berhasil dihapus");
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal menghapus spesifikasi");
  }
}

async function handleDeleteComponent(item: MaterialComponent) {
  const confirmed = await confirm({
    title: "Hapus Komponen",
    message: `Yakin ingin menghapus "${item.name}"?`,
    confirmLabel: "Ya, Hapus",
    danger: true,
  });

  if (!confirmed) return;

  try {
    await remove(item.id);

    toast.success("Komponen berhasil dihapus");
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal menghapus komponen");
  }
}
</script>

<template>
  <SidebarAppLayout>
    <div class="bg-white border border-[#E5D9BF] rounded-xl p-6 mb-6">
      <div class="flex items-center justify-between mb-6">
        <h1 class="text-xl font-semibold text-[#3A2E1F]">Komponen Bahan</h1>

        <button
          @click="openCreateComponent"
          class="flex items-center gap-2 px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition">
          <Plus :size="16" />
          Tambah Komponen
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
            class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F]" />
        </div>

        <div class="col-span-6">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Status
          </label>

          <div class="flex items-center gap-4 h-10.5">
            <label
              v-for="status in [
                ['all', 'Semua'],
                ['active', 'Aktif'],
                ['inactive', 'Non-Aktif'],
              ]"
              :key="status[0]"
              class="flex items-center gap-2 text-sm cursor-pointer">
              <input
                v-model="statusFilter"
                type="radio"
                :value="status[0]"
                class="accent-[#C9A24B]" />

              {{ status[1] }}
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
        :columns="[
          { key: 'aksi', label: 'Aksi', align: 'center' },
          { key: 'name', label: 'Komponen' },
          { key: 'category', label: 'Kategori' },
          { key: 'remark', label: 'Catatan' },
          { key: 'status', label: 'Status', align: 'center' },
        ]"
        :items="paginatedItems"
        :is-loading="isLoading"
        empty-message="Tidak ditemukan data."
        row-key="id">
        <template #aksi="{ item }">
          <OverlayMenu
            :items="[
              {
                label: 'Spesifikasi & Harga',
                icon: Settings2,
                command: () => openVariants(item),
              },
              {
                label: 'Edit',
                icon: Pencil,
                command: () => openEditComponent(item),
              },
              {
                label: 'Hapus',
                icon: Trash2,
                danger: true,
                command: () => handleDeleteComponent(item),
              },
            ]" />
        </template>

        <template #status="{ item }">
          <span
            class="px-3 rounded-full text-xs font-semibold"
            :class="
              item.status.toLowerCase() === 'active'
                ? 'bg-[#4C8C5B]/15 text-[#4C8C5B]'
                : 'bg-[#B23A32]/15 text-[#B23A32]'
            ">
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
      :visible="showComponentModal"
      :header-title="
        isComponentEditMode ? 'Edit Komponen Bahan' : 'Tambah Komponen Bahan'
      "
      width-class="w-full max-w-md"
      @hide="showComponentModal = false">
      <form @submit.prevent="submitComponent" class="space-y-4">
        <InputTextComponent
          v-model="componentName"
          label="Nama Komponen"
          placeholder="Contoh: Kayu Pinus"
          required />

        <InputTextComponent
          v-model="componentCategory"
          label="Kategori"
          placeholder="Contoh: Wood"
          required />

        <InputTextComponent
          v-model="componentRemark"
          label="Catatan"
          placeholder="Catatan tambahan" />

        <div v-if="isComponentEditMode">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Status
          </label>

          <select
            v-model="componentStatus"
            class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm">
            <option value="Active">Active</option>
            <option value="Inactive">Inactive</option>
          </select>
        </div>

        <button
          type="submit"
          class="w-full py-2.5 bg-[#7A1F2B] text-white rounded-lg font-medium">
          Simpan
        </button>
      </form>
    </ChildModalWrapper>

    <ChildModalWrapper
      :visible="showVariantModal"
      :header-title="
        selectedComponent
          ? `${selectedComponent.name} — Spesifikasi & Harga`
          : 'Spesifikasi & Harga'
      "
      width-class="w-full max-w-2xl"
      @hide="showVariantModal = false">
      <div class="space-y-5">
        <!-- LIST -->
        <div class="border border-[#E5D9BF] rounded-lg overflow-hidden">
          <div class="px-4 py-3 bg-[#F5EFE3] flex items-center justify-between">
            <span class="font-medium text-sm text-[#3A2E1F]">
              Spesifikasi
            </span>

            <span class="text-xs text-[#8A7A5C]">
              {{ variants.length }} data
            </span>
          </div>

          <div
            v-if="isVariantLoading"
            class="p-6 text-center text-sm text-[#8A7A5C]">
            Memuat data...
          </div>

          <div
            v-else-if="variants.length === 0"
            class="p-6 text-center text-sm text-[#8A7A5C]">
            Belum ada spesifikasi.
          </div>

          <div
            v-for="variant in variants"
            :key="variant.id"
            class="flex items-center justify-between px-4 py-3 border-t border-[#E5D9BF]">
            <div>
              <p class="text-sm font-medium text-[#3A2E1F]">
                {{ formatSpecification(variant) }}
              </p>

              <p
                v-if="variant.specification"
                class="text-xs text-[#8A7A5C] mt-0.5">
                {{ variant.specification }}
              </p>

              <p class="text-xs text-[#8A7A5C] mt-1">
                {{ variant.unit }}
              </p>
            </div>

            <div class="flex items-center gap-4">
              <span class="font-semibold text-sm text-[#7A1F2B]">
                {{ formatRupiah(variant.unit_price) }}
              </span>

              <OverlayMenu
                :items="[
                  {
                    label: 'Edit',
                    icon: Pencil,
                    command: () => openEditVariant(variant),
                  },
                  {
                    label: 'Hapus',
                    icon: Trash2,
                    danger: true,
                    command: () => handleDeleteVariant(variant),
                  },
                ]" />
            </div>
          </div>
        </div>

        <!-- FORM -->
        <div class="border border-[#E5D9BF] rounded-lg p-4">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-sm font-semibold text-[#3A2E1F]">
              {{
                isVariantEditMode ? "Edit Spesifikasi" : "Tambah Spesifikasi"
              }}
            </h3>

            <button
              v-if="isVariantEditMode"
              type="button"
              @click="resetVariantForm"
              class="text-xs text-[#7A1F2B]">
              Batal Edit
            </button>
          </div>

          <div class="grid grid-cols-3 gap-3">
            <InputNumberComponent
              label="Panjang (cm)"
              v-model.number="variantLength"
              type="number"
              step="0.01"
              class="w-full px-4 bg-white rounded-lg text-sm" />

            <InputNumberComponent
              v-model.number="variantWidth"
              type="number"
              label="Lebar (cm)"
              step="0.01"
              class="w-full px-4 bg-white rounded-lg text-sm" />

            <InputNumberComponent
              v-model.number="variantThickness"
              label="Tebal (cm)"
              class="w-full px-4 bg-white text-sm"
              placeholder="2" />
          </div>

          <div class="mt-4 space-y-4">
            <InputTextComponent
              v-model="variantSpecification"
              label="Spesifikasi"
              placeholder="Contoh: Kayu grade A" />

            <div class="grid grid-cols-2 gap-4">
              <InputTextComponent
                v-model="variantUnit"
                label="Satuan"
                placeholder="pcs"
                required />

              <InputNumberComponent
                v-model.number="variantUnitPrice"
                label="Harga Satuan"
                placeholder="6000"
                required />
            </div>

            <InputTextComponent
              v-model="variantRemark"
              label="Catatan"
              placeholder="Catatan tambahan" />

            <button
              type="button"
              @click="submitVariant"
              class="w-full py-2.5 bg-[#7A1F2B] text-white rounded-lg font-medium">
              {{
                isVariantEditMode ? "Simpan Perubahan" : "Tambah Spesifikasi"
              }}
            </button>
          </div>
        </div>
      </div>
    </ChildModalWrapper>
  </SidebarAppLayout>
</template>
