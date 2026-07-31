<script setup lang="ts">
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { Plus, Trash2, ArrowLeft } from "lucide-vue-next";
import AppLayout from "../../../components/layout/sidebar-app-layout.vue";
import SearchableSelect from "../../../components/input-text/serchable-select.vue";
import InputText from "../../../components/input-text/input-text-component.vue";
import InputNumber from "../../../components/input-text/input-number-component.vue";
import PhotoPicker from "../../../components/upload-file/upload-photo.vue";
import { useCommission } from "../../../composables/commission/useCommission";
import { useClientApi } from "../../../composables/commission/client_api.ts";
import { useMasterSize } from "../../../composables/master/m_size.ts";
import { useMasterMaterial } from "../../../composables/master/m_material.ts";
import { useMasterStyle } from "../../../composables/master/m_style.ts";
import { useMasterSaints } from "../../../composables/master/m_saints.ts";
import { useToast } from "../../../composables/etc/useToast.ts";

const router = useRouter();
const toast = useToast();

const { create } = useCommission();
const { items: clients, fetchAll: fetchClients } = useClientApi();
const { items: sizes, fetchAll: fetchSizes } = useMasterSize();
const { items: materials, fetchAll: fetchMaterials } = useMasterMaterial();
const { items: styles, fetchAll: fetchStyles } = useMasterStyle();
const { items: saints, fetchAll: fetchSaints } = useMasterSaints();

onMounted(() => {
  fetchClients();
  fetchSizes();
  fetchMaterials();
  fetchStyles();
  fetchSaints();
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

// ===== State form header =====
const clientId = ref<number | null>(null);
const deadline = ref("");
const notes = ref("");
const photoUrls = ref<string[]>([]);

// ===== State items (array baris) =====
interface ItemRow {
  saintId: number | null;
  sizeId: number | null;
  materialId: number | null;
  styleId: number | null;
  price: number | null;
  notes: string;
}

function emptyRow(): ItemRow {
  return {
    saintId: null,
    sizeId: null,
    materialId: null,
    styleId: null,
    price: null,
    notes: "",
  };
}

const itemRows = ref<ItemRow[]>([emptyRow()]);

function addRow() {
  itemRows.value.push(emptyRow());
}

function removeRow(index: number) {
  if (itemRows.value.length === 1) return; // minimal 1 item
  itemRows.value.splice(index, 1);
}

const totalPrice = computed(() =>
  itemRows.value.reduce((sum, row) => sum + (row.price ?? 0), 0)
);

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
      items: itemRows.value.map((row) => ({
        saint_id: row.saintId,
        size_id: row.sizeId,
        material_id: row.materialId,
        style_id: row.styleId,
        price: row.price,
        notes: row.notes,
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

function resetForm() {
  clientId.value = null;
  deadline.value = "";
  notes.value = "";
  photoUrls.value = [];
  itemRows.value = [emptyRow()];
  errors.value = {};
}
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
            </div>

            <p
              v-if="errors[`item_${index}`]"
              class="text-xs text-[#B23A32] mt-2">
              {{ errors[`item_${index}`] }}
            </p>
          </div>
        </div>

        <div class="flex justify-end mt-4 pt-4 border-t border-[#E5D9BF]">
          <div class="text-right">
            <p class="text-xs text-[#8A7A5C]">Total Harga</p>
            <p class="text-lg font-semibold text-[#3A2E1F]">
              Rp {{ totalPrice.toLocaleString("id-ID") }}
            </p>
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
</template>
