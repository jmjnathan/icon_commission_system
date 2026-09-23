<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { Plus, Trash2, ArrowLeft } from "lucide-vue-next";

import AppLayout from "../../../components/layout/sidebar-app-layout.vue";
import SearchableSelect from "../../../components/input-text/serchable-select.vue";
import InputText from "../../../components/input-text/input-text-component.vue";
import InputNumber from "../../../components/input-text/input-number-component.vue";
import PhotoPicker from "../../../components/upload-file/upload-photo.vue";
// NOTE: sebelumnya dipake di template tapi belum pernah di-import.
// Sesuaikan path-nya sama lokasi asli komponen ini di project kamu.
import ChildModalWrapper from "../../../components/modal/child-modal-wrapper.vue";

import { useCommission } from "../../../composables/commission/useCommission";
import { useClientApi } from "../../../composables/commission/client_api.ts";
import { useMasterSize } from "../../../composables/master/m_size.ts";
import { useMasterMaterial } from "../../../composables/master/m_material.ts";
import { useMasterStyle } from "../../../composables/master/m_style.ts";
import { useMasterSaints } from "../../../composables/master/m_saints.ts";
import { useToast } from "../../../composables/etc/useToast.ts";
import {
  useCommissionItemMaterial,
  type MaterialComponentVariant,
} from "../../../composables/commission/useCommissionMaterial.ts";

const router = useRouter();
const toast = useToast();

// ===== Composables (data master & submit) =====
const { create } = useCommission();
const { items: clients, fetchAll: fetchClients } = useClientApi();
const { items: sizes, fetchAll: fetchSizes } = useMasterSize();
const { items: materials, fetchAll: fetchMaterials } = useMasterMaterial();
const { items: styles, fetchAll: fetchStyles } = useMasterStyle();
const { items: saints, fetchAll: fetchSaints } = useMasterSaints();

const {
  variants: materialVariants,
  isVariantLoading,
  fetchVariants,
  // Berikut ini didestructure tapi belum kepake di file ini.
  // Kelihatannya disiapin buat fitur edit/hapus material di kemudian hari.
  // items: commissionMaterials,
  // isLoading: isMaterialLoading,
  // createMaterial,
  // updateMaterial,
  // deleteMaterial,
} = useCommissionItemMaterial();

onMounted(async () => {
  await Promise.all([
    fetchClients(),
    fetchSizes(),
    fetchMaterials(),
    fetchVariants(),
    fetchStyles(),
    fetchSaints(),
  ]);
});

// ===== Opsi dropdown =====
const clientOptions = computed(() =>
  clients.value.map((c) => ({ label: c.name, value: c.id }))
);
const sizeOptions = computed(() =>
  sizes.value.map((s) => ({ label: `${s.name} (${s.size})`, value: s.id }))
);
const materialOptions = computed(() =>
  materials.value.map((m) => ({ label: m.name, value: m.id }))
);
const styleOptions = computed(() =>
  styles.value.map((s) => ({ label: s.name, value: s.id }))
);
const saintOptions = computed(() =>
  saints.value.map((s) => ({ label: s.name, value: s.id }))
);

// ===== Tipe data form =====
interface ItemMaterialRow {
  materialComponentVariantId: number | null;
  quantity: number | null;
  remark: string;
}

interface ItemRow {
  saintId: number | null;
  sizeId: number | null;
  materialId: number | null;
  styleId: number | null;
  price: number | null;
  notes: string;
  materials: ItemMaterialRow[];
}

function emptyRow(): ItemRow {
  return {
    saintId: null,
    sizeId: null,
    materialId: null,
    styleId: null,
    price: null,
    notes: "",
    materials: [],
  };
}

// ===== State form header =====
const clientId = ref<number | null>(null);
const deadline = ref("");
const notes = ref("");
const photoUrls = ref<string[]>([]);

// ===== State item-item pesanan =====
const itemRows = ref<ItemRow[]>([emptyRow()]);

function addRow() {
  itemRows.value.push(emptyRow());
}

function removeRow(index: number) {
  if (itemRows.value.length === 1) return; // minimal 1 item
  itemRows.value.splice(index, 1);
}

// ===== State modal tambah komponen bahan =====
const showMaterialModal = ref(false);
const activeItemIndex = ref<number | null>(null);
const selectedMaterialVariantId = ref<number | null>(null);
const materialQuantity = ref<number | null>(null);
const materialRemark = ref("");

const selectedVariant = computed(() =>
  materialVariants.value.find(
    (item) => item.id === selectedMaterialVariantId.value
  )
);

const estimatedMaterialCost = computed(() => {
  if (!selectedVariant.value || !materialQuantity.value) return 0;
  return selectedVariant.value.unit_price * materialQuantity.value;
});

function openMaterialModal(index: number) {
  activeItemIndex.value = index;
  selectedMaterialVariantId.value = null;
  materialQuantity.value = null;
  materialRemark.value = "";
  showMaterialModal.value = true;
}

function closeMaterialModal() {
  activeItemIndex.value = null;
  selectedMaterialVariantId.value = null;
  materialQuantity.value = null;
  materialRemark.value = "";
  showMaterialModal.value = false;
}

function addMaterialToItem() {
  if (activeItemIndex.value === null) return;

  if (
    !selectedMaterialVariantId.value ||
    !materialQuantity.value ||
    materialQuantity.value <= 0
  ) {
    toast.error("Komponen bahan dan quantity wajib diisi");
    return;
  }

  itemRows.value[activeItemIndex.value].materials.push({
    materialComponentVariantId: selectedMaterialVariantId.value,
    quantity: materialQuantity.value,
    remark: materialRemark.value,
  });

  closeMaterialModal();
}

// ===== Format helper =====
function formatRupiah(value: number) {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    maximumFractionDigits: 0,
  }).format(value);
}

function formatVariantLabel(variant: MaterialComponentVariant) {
  const componentName = variant.material_component?.name ?? "Unknown";

  const dimensions = [variant.length, variant.width, variant.thickness]
    .filter((value) => value > 0)
    .join(" × ");

  const specification =
    dimensions !== ""
      ? `${dimensions} cm`
      : variant.specification || variant.unit;

  return `${componentName} — ${specification} — ${formatRupiah(
    variant.unit_price
  )}/${variant.unit}`;
}

function formatMaterialLabel(variantId: number | null) {
  const variant = materialVariants.value.find((item) => item.id === variantId);
  if (!variant) return "Komponen tidak ditemukan";
  return formatVariantLabel(variant);
}

// ===== Validasi & Submit =====
const errors = ref<Record<string, string>>({});
const isSubmitting = ref(false);

function validate(): boolean {
  errors.value = {};
  if (!clientId.value) errors.value.client = "Client wajib dipilih";
  if (!deadline.value) errors.value.deadline = "Deadline wajib diisi";

  itemRows.value.forEach((row, i) => {
    if (
      !row.saintId ||
      !row.sizeId ||
      !row.materialId ||
      !row.styleId ||
      !row.price
    ) {
      errors.value[`item_${i}`] = "Semua field item wajib diisi";
    }
  });

  return Object.keys(errors.value).length === 0;
}

function resetForm() {
  clientId.value = null;
  deadline.value = "";
  notes.value = "";
  photoUrls.value = [];
  itemRows.value = [emptyRow()];
  errors.value = {};
}

async function handleSubmit() {
  if (!validate()) {
    toast.error("Mohon lengkapi semua field yang wajib diisi");
    return;
  }

  isSubmitting.value = true;
  try {
    await create({
      client_id: clientId.value,
      deadline: deadline.value,
      notes: notes.value,
      photo_urls: photoUrls.value,
      discount_type: discountType.value || undefined,
      discount_value: discountValue.value || 0,
      items: itemRows.value.map((row) => ({
        saint_id: row.saintId,
        size_id: row.sizeId,
        material_id: row.materialId,
        style_id: row.styleId,
        price: row.price,
        notes: row.notes,
        materials: row.materials.map((material) => ({
          material_component_variant_id: material.materialComponentVariantId,
          quantity: material.quantity,
          remark: material.remark,
        })),
      })),
    });

    toast.success("Pemesanan baru berhasil dibuat");
    resetForm();
    router.push("/transaction/commissions");
  } catch (err: any) {
    toast.error(err.response?.data?.error || "Gagal membuat komisi");
  } finally {
    isSubmitting.value = false;
  }
}

const discountType = ref<"" | "percent" | "fixed">("");
const discountValue = ref<number | null>(null);

const subtotal = computed(() =>
  itemRows.value.reduce((sum, row) => sum + (row.price ?? 0), 0)
);

const discountAmount = computed(() => {
  if (!discountType.value || !discountValue.value) return 0;
  if (discountType.value === "percent")
    return subtotal.value * (discountValue.value / 100);
  return discountValue.value;
});

const totalPrice = computed(() =>
  Math.max(subtotal.value - discountAmount.value, 0)
);
</script>

<template>
  <AppLayout>
    <div class="flex items-center gap-3 mb-6">
      <button
        @click="router.back()"
        class="w-9 h-9 flex items-center justify-center rounded-lg text-[#8A7A5C] hover:bg-[#E5D9BF]/50 transition">
        <ArrowLeft :size="18" />
      </button>
      <h1 class="text-2xl font-semibold text-[#3A2E1F]">Buat Pemesanan Baru</h1>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">
      <!-- Info Umum -->
      <div class="bg-white border border-[#E5D9BF] rounded-xl p-6">
        <h2
          class="text-sm font-semibold text-[#3A2E1F] mb-4 uppercase tracking-wide">
          Informasi Umum
        </h2>
        <div class="grid grid-cols-12 gap-6">
          <div class="col-span-6">
            <SearchableSelect
              required
              v-model="clientId"
              label="Client"
              placeholder="Pilih client"
              :options="clientOptions"
              :error-message="errors.client" />
          </div>
          <div class="col-span-6">
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
              Deadline <span class="text-[#B23A32]">*</span>
            </label>
            <input
              v-model="deadline"
              type="date"
              class="w-full px-4 py-2.5 bg-white border rounded-lg text-sm text-[#3A2E1F] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 transition"
              :class="
                errors.deadline ? 'border-[#B23A32]' : 'border-[#D9CBB0]'
              " />
            <p v-if="errors.deadline" class="text-xs text-[#B23A32] mt-1">
              {{ errors.deadline }}
            </p>
          </div>
          <div class="col-span-12">
            <InputText
              v-model="notes"
              label="Catatan Umum"
              placeholder="Catatan tambahan untuk komisi ini" />
          </div>
          <div class="col-span-12">
            <PhotoPicker v-model="photoUrls" />
          </div>
        </div>
      </div>

      <!-- Items -->
      <div class="bg-white border border-[#E5D9BF] rounded-xl p-6">
        <div class="flex items-center justify-between mb-4">
          <h2
            class="text-sm font-semibold text-[#3A2E1F] uppercase tracking-wide">
            Detail Ikon Pesanan
          </h2>
          <button
            type="button"
            @click="addRow"
            class="flex items-center gap-2 px-3 py-1.5 text-sm text-[#7A1F2B] border border-[#7A1F2B]/30 rounded-lg hover:bg-[#7A1F2B]/5 transition">
            <Plus :size="14" />
            Tambah Item
          </button>
        </div>

        <div class="space-y-4">
          <div
            v-for="(row, index) in itemRows"
            :key="index"
            class="border border-[#E5D9BF] rounded-lg p-4 relative">
            <button
              v-if="itemRows.length > 1"
              type="button"
              @click="removeRow(index)"
              class="absolute top-3 right-3 w-7 h-7 flex items-center justify-center rounded-lg text-[#B23A32] hover:bg-[#B23A32]/10 transition">
              <Trash2 :size="14" />
            </button>

            <p class="text-xs font-semibold text-[#8A7A5C] mb-3">
              Item #{{ index + 1 }}
            </p>

            <div class="grid grid-cols-12 gap-4">
              <div class="col-span-6">
                <SearchableSelect
                  v-model="row.saintId"
                  label="Santo / Santa"
                  placeholder="Pilih santo pelindung"
                  :options="saintOptions" />
              </div>
              <div class="col-span-6">
                <SearchableSelect
                  v-model="row.sizeId"
                  label="Ukuran"
                  placeholder="Pilih ukuran"
                  :options="sizeOptions" />
              </div>
              <div class="col-span-4">
                <SearchableSelect
                  v-model="row.materialId"
                  label="Material"
                  placeholder="Pilih material"
                  :options="materialOptions" />
              </div>
              <div class="col-span-4">
                <SearchableSelect
                  v-model="row.styleId"
                  label="Style"
                  placeholder="Pilih style"
                  :options="styleOptions" />
              </div>
              <div class="col-span-4">
                <InputNumber
                  v-model="row.price"
                  label="Harga"
                  prefix="Rp"
                  :max-fraction-digits="0" />
              </div>
              <div class="col-span-12">
                <InputText
                  v-model="row.notes"
                  label="Catatan Item"
                  placeholder="Detail khusus untuk item ini" />
              </div>

              <!-- Komponen Bahan -->
              <div class="col-span-12">
                <div
                  class="border border-[#E5D9BF] rounded-lg p-4 bg-[#FBF8F1]">
                  <div class="flex items-center justify-between mb-3">
                    <div>
                      <p class="text-sm font-semibold text-[#3A2E1F]">
                        Komponen Bahan
                      </p>
                      <p class="text-xs text-[#8A7A5C] mt-1">
                        Bahan yang digunakan untuk item ini
                      </p>
                    </div>
                    <button
                      type="button"
                      @click="openMaterialModal(index)"
                      class="flex items-center gap-2 px-3 py-1.5 text-sm text-[#7A1F2B] border border-[#7A1F2B]/30 rounded-lg hover:bg-[#7A1F2B]/5 transition">
                      <Plus :size="14" />
                      Tambah Bahan
                    </button>
                  </div>

                  <div
                    v-if="row.materials.length === 0"
                    class="text-xs text-[#8A7A5C]">
                    Belum ada komponen bahan.
                  </div>

                  <div v-else class="space-y-2">
                    <div
                      v-for="(material, materialIndex) in row.materials"
                      :key="materialIndex"
                      class="flex items-center justify-between bg-white border border-[#E5D9BF] rounded-lg px-3 py-2">
                      <div>
                        <p class="text-sm text-[#3A2E1F]">
                          {{
                            formatMaterialLabel(
                              material.materialComponentVariantId
                            )
                          }}
                        </p>
                        <p class="text-xs text-[#8A7A5C]">
                          Qty: {{ material.quantity }}
                        </p>
                      </div>
                      <button
                        type="button"
                        @click="row.materials.splice(materialIndex, 1)"
                        class="text-[#B23A32] hover:bg-[#B23A32]/10 w-7 h-7 rounded-lg flex items-center justify-center">
                        <Trash2 :size="14" />
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <p
              v-if="errors[`item_${index}`]"
              class="text-xs text-[#B23A32] mt-2">
              {{ errors[`item_${index}`] }}
            </p>
          </div>
        </div>

        <div class="flex justify-end mt-4 pt-4 border-t border-[#E5D9BF]">
          <div class="w-full max-w-xs space-y-3">
            <div class="grid grid-cols-2 gap-3">
              <div>
                <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
                  >Tipe Diskon</label
                >
                <select
                  v-model="discountType"
                  class="w-full px-3 py-2 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40">
                  <option value="">Tanpa Diskon</option>
                  <option value="percent">Persen (%)</option>
                  <option value="fixed">Nominal (Rp)</option>
                </select>
              </div>
              <div>
                <label class="block text-xs font-medium text-[#6B5D45] mb-1.5"
                  >Nilai Diskon</label
                >
                <input
                  v-model.number="discountValue"
                  type="number"
                  :disabled="!discountType"
                  class="w-full px-3 py-2 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 disabled:bg-[#F5EFE3]" />
              </div>
            </div>

            <div class="space-y-1 pt-2 border-t border-[#E5D9BF] text-sm">
              <div class="flex justify-between text-[#6B5D45]">
                <span>Subtotal</span>
                <span>Rp {{ subtotal.toLocaleString("id-ID") }}</span>
              </div>
              <div
                v-if="discountAmount > 0"
                class="flex justify-between text-[#B23A32]">
                <span>Diskon</span>
                <span>- Rp {{ discountAmount.toLocaleString("id-ID") }}</span>
              </div>
              <div
                class="flex justify-between text-base font-semibold text-[#3A2E1F] pt-1">
                <span>Total</span>
                <span>Rp {{ totalPrice.toLocaleString("id-ID") }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="flex justify-end gap-3">
        <button
          type="button"
          @click="router.back()"
          class="px-5 py-2.5 text-sm font-medium text-[#6B5D45] rounded-lg hover:bg-[#E5D9BF]/50 transition">
          Batal
        </button>
        <button
          type="submit"
          :disabled="isSubmitting"
          class="px-5 py-2.5 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition disabled:opacity-50">
          {{ isSubmitting ? "Menyimpan..." : "Simpan Komisi" }}
        </button>
      </div>
    </form>
  </AppLayout>

  <ChildModalWrapper
    :visible="showMaterialModal"
    header-title="Tambah Komponen Bahan"
    width-class="w-full max-w-md"
    @hide="closeMaterialModal">
    <div class="space-y-4">
      <SearchableSelect
        v-model="selectedMaterialVariantId"
        label="Komponen Bahan"
        placeholder="Pilih komponen bahan"
        :options="
          materialVariants.map((variant) => ({
            label: formatVariantLabel(variant),
            value: variant.id,
          }))
        " />

      <div v-if="selectedVariant">
        <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
          Harga Satuan
        </label>
        <div
          class="w-full px-4 py-2.5 bg-[#F5EFE3] border border-[#E5D9BF] rounded-lg text-sm text-[#3A2E1F]">
          {{ formatRupiah(selectedVariant.unit_price) }}
          / {{ selectedVariant.unit }}
        </div>
      </div>

      <InputNumber
        v-model="materialQuantity"
        label="Quantity"
        :max-fraction-digits="3" />

      <InputText
        v-model="materialRemark"
        label="Catatan"
        placeholder="Catatan penggunaan bahan" />

      <div
        v-if="selectedVariant && materialQuantity"
        class="bg-[#7A1F2B]/5 border border-[#7A1F2B]/10 rounded-lg p-4">
        <p class="text-xs text-[#8A7A5C]">Estimasi Material Cost</p>
        <p class="text-lg font-semibold text-[#7A1F2B] mt-1">
          Rp {{ estimatedMaterialCost.toLocaleString("id-ID") }}
        </p>
      </div>

      <button
        type="button"
        @click="addMaterialToItem"
        class="w-full py-2.5 bg-[#7A1F2B] text-white rounded-lg font-medium hover:bg-[#5F1621] transition">
        Tambahkan
      </button>
    </div>
  </ChildModalWrapper>
</template>
