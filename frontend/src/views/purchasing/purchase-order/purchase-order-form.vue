<script setup lang="ts">
import { computed, onMounted, ref } from "vue";

import { useRoute, useRouter } from "vue-router";

import { ArrowLeft, Plus, Pencil, Trash2, Search } from "lucide-vue-next";

import SidebarAppLayout from "../../../components/layout/sidebar-app-layout.vue";
import ChildModalWrapper from "../../../components/modal/child-modal-wrapper.vue";
import InputTextComponent from "../../../components/input-text/input-text-component.vue";
import OverlayMenu from "../../../components/button/overlay-menu.vue";

import {
  usePurchaseOrder,
  type PurchaseOrderItem,
  type Product,
} from "../../../composables/purchasing/purchase-order/usePurchaseOrder.ts";

import { useToast } from "../../../composables/etc/useToast.ts";
import { useConfirm } from "../../../composables/etc/useConfirm.ts";

const route = useRoute();
const router = useRouter();

const { getById, getItems, create, update, submit, remove, searchProduct } =
  usePurchaseOrder();

const toast = useToast();
const { confirm } = useConfirm();

/* =========================
   MODE
========================= */

const id = computed(() => {
  const value = route.params.id;

  if (!value) return null;

  return Number(value);
});

const isEditMode = computed(() => id.value !== null);

const isComplete = ref(false);

const pageLoading = ref(false);
const isSubmitting = ref(false);

/* =========================
   HEADER
========================= */

const documentNo = ref("");
const orderDate = ref("");
const vendorName = ref("");
const remark = ref("");

/* =========================
   ITEMS
========================= */

const items = ref<PurchaseOrderItem[]>([]);

/* =========================
   ITEM MODAL
========================= */

const showItemModal = ref(false);

function closeModal() {
  showItemModal.value = false;
  //   resetForm();
}

const editingItemId = ref<number | null>(null);

const itemProductId = ref<number | null>(null);
const itemProductName = ref("");
const itemBrand = ref("");
const itemQuantity = ref<number | null>(null);
const itemUnitPrice = ref<number | null>(null);
const itemRemark = ref("");

const itemError = ref("");

const productKeyword = ref("");
const productOptions = ref<Product[]>([]);
const productLoading = ref(false);

const itemModalTitle = computed(() =>
  editingItemId.value !== null ? "Edit Item" : "Tambah Item"
);

/* =========================
   TOTAL
========================= */

const total = computed(() => {
  return items.value.reduce((sum, item) => sum + Number(item.subtotal || 0), 0);
});

/* =========================
   LOAD DETAIL
========================= */

onMounted(async () => {
  if (!id.value) {
    orderDate.value = new Date().toISOString().substring(0, 10);

    return;
  }

  pageLoading.value = true;

  try {
    const data = await getById(id.value);
    const orderItems = await getItems(id.value);

    documentNo.value = data.document_no;
    orderDate.value = data.order_date.substring(0, 10);
    vendorName.value = data.vendor_name;
    remark.value = data.remark ?? "";

    items.value = orderItems;

    isComplete.value = data.status === "ordered";
  } catch (err: any) {
    toast.error(
      err?.response?.data?.message ||
        err?.response?.data?.error ||
        "Gagal mengambil Pemesanan Bahan Baku."
    );

    router.push("/purchasing/purchase-order");
  } finally {
    pageLoading.value = false;
  }
});

/* =========================
   PRODUCT SEARCH
========================= */

async function handleProductSearch() {
  const keyword = productKeyword.value.trim();

  // User mulai mengubah value setelah memilih product
  if (keyword !== itemProductName.value) {
    itemProductId.value = null;
    itemProductName.value = "";
  }

  if (!keyword) {
    productOptions.value = [];
    return;
  }

  productLoading.value = true;

  try {
    productOptions.value = await searchProduct(keyword);
  } catch (err) {
    productOptions.value = [];
  } finally {
    productLoading.value = false;
  }
}

/* =========================
   SELECT PRODUCT
========================= */

function selectProduct(product: Product) {
  itemProductId.value = product.id;
  itemProductName.value = product.name;

  productKeyword.value = product.name;
  productOptions.value = [];
}

/* =========================
   ITEM FORM
========================= */

function resetItemForm() {
  editingItemId.value = null;

  itemProductId.value = null;
  itemProductName.value = "";

  itemBrand.value = "";

  itemQuantity.value = null;
  itemUnitPrice.value = null;

  itemRemark.value = "";

  itemError.value = "";

  productKeyword.value = "";
  productOptions.value = [];
}

function openAddItemModal() {
  resetItemForm();

  showItemModal.value = true;
}

function openEditItemModal(item: PurchaseOrderItem) {
  editingItemId.value = item.id;

  itemProductId.value = item.product_id;
  itemProductName.value = item.product_name;

  itemBrand.value = item.brand ?? "";

  itemQuantity.value = item.quantity;
  itemUnitPrice.value = item.unit_price;

  itemRemark.value = item.remark ?? "";

  itemError.value = "";

  productKeyword.value = item.product_name;
  productOptions.value = [];

  showItemModal.value = true;
}

function closeItemModal() {
  showItemModal.value = false;

  resetItemForm();
}

/* =========================
   SAVE ITEM
========================= */

function handleSaveItem() {
  itemError.value = "";

  if (!itemProductId.value) {
    itemError.value = "Product wajib dipilih.";

    return;
  }

  if (!itemQuantity.value || itemQuantity.value <= 0) {
    itemError.value = "Quantity harus lebih dari 0.";

    return;
  }

  if (itemUnitPrice.value === null || itemUnitPrice.value < 0) {
    itemError.value = "Harga satuan tidak valid.";

    return;
  }

  const subtotal = Number(itemQuantity.value) * Number(itemUnitPrice.value);

  if (editingItemId.value !== null) {
    const index = items.value.findIndex(
      (item) => item.id === editingItemId.value
    );

    if (index !== -1) {
      items.value[index] = {
        ...items.value[index],
        product_id: itemProductId.value,
        product_name: itemProductName.value,
        brand: itemBrand.value.trim() || null,
        quantity: Number(itemQuantity.value),
        unit_price: Number(itemUnitPrice.value),
        subtotal,
        remark: itemRemark.value.trim() || null,
      };
    }
  } else {
    items.value.push({
      id: -Date.now(),
      purchase_order_id: id.value ?? 0,
      product_id: itemProductId.value,
      product_name: itemProductName.value,
      brand: itemBrand.value.trim() || null,
      quantity: Number(itemQuantity.value),
      unit_price: Number(itemUnitPrice.value),
      subtotal,
      remark: itemRemark.value.trim() || null,
    });
  }

  closeItemModal();
}

/* =========================
   DELETE ITEM
========================= */

async function handleDeleteItem(item: PurchaseOrderItem) {
  const confirmed = await confirm({
    title: "Hapus Item",
    message: `Yakin ingin menghapus item "${item.product_name}"?`,
    confirmLabel: "Ya, Hapus",
    danger: true,
  });

  if (!confirmed) return;

  items.value = items.value.filter((value) => value.id !== item.id);
}

/* =========================
   ITEM MENU
========================= */

function buildItemMenu(item: PurchaseOrderItem) {
  return [
    {
      label: "Edit",
      icon: Pencil,
      command: () => openEditItemModal(item),
    },
    {
      label: "Delete",
      icon: Trash2,
      danger: true,
      command: () => handleDeleteItem(item),
    },
  ];
}

/* =========================
   VALIDATION
========================= */

function validateHeader() {
  if (!vendorName.value.trim()) {
    toast.error("Vendor wajib diisi.");

    return false;
  }

  if (!orderDate.value) {
    toast.error("Tanggal dokumen wajib diisi.");

    return false;
  }

  if (items.value.length === 0) {
    toast.error("Minimal harus ada 1 item.");

    return false;
  }

  return true;
}

/* =========================
   CREATE / UPDATE
========================= */

async function handleCreate() {
  if (!validateHeader()) return;

  isSubmitting.value = true;

  try {
    const payload = {
      document_no: documentNo.value.trim(),
      vendor_name: vendorName.value.trim(),
      order_date: orderDate.value,
      remark: remark.value.trim() || null,
      items: items.value.map((item) => ({
        product_id: item.product_id,
        product_name: item.product_name,
        quantity: item.quantity,
        unit_price: item.unit_price,
        brand: item.brand,
        remark: item.remark,
      })),
    };

    const result = await create(payload);

    toast.success("Pemesanan Bahan Baku berhasil dibuat.");

    router.replace(`/purchasing/purchase-order/form/${result.id}`);
  } catch (err: any) {
    toast.error(
      err?.response?.data?.message ||
        err?.response?.data?.error ||
        "Gagal membuat Pemesanan Bahan Baku."
    );
  } finally {
    isSubmitting.value = false;
  }
}

async function handleUpdate() {
  if (!id.value) return;

  if (!validateHeader()) return;

  isSubmitting.value = true;

  try {
    const payload = {
      document_no: documentNo.value.trim(),
      vendor_name: vendorName.value.trim(),
      order_date: orderDate.value,
      remark: remark.value.trim() || null,
      items: items.value.map((item) => ({
        product_id: item.product_id,
        product_name: item.product_name,
        quantity: item.quantity,
        unit_price: item.unit_price,
        brand: item.brand,
        remark: item.remark,
      })),
    };

    await update(id.value, payload);

    toast.success("Pemesanan Bahan Baku berhasil diperbarui.");
  } catch (err: any) {
    toast.error(
      err?.response?.data?.message ||
        err?.response?.data?.error ||
        "Gagal memperbarui Pemesanan Bahan Baku."
    );
  } finally {
    isSubmitting.value = false;
  }
}

/* =========================
   SUBMIT
========================= */

async function handleSubmit() {
  if (!id.value) return;

  if (!validateHeader()) return;

  const confirmed = await confirm({
    title: "Submit Pemesanan Bahan Baku",
    message:
      "Setelah di-submit, Pemesanan Bahan Baku tidak dapat diedit lagi. Lanjutkan?",
    confirmLabel: "Ya, Submit",
  });

  if (!confirmed) return;

  isSubmitting.value = true;

  try {
    await submit(id.value);

    isComplete.value = true;

    toast.success("Pemesanan Bahan Baku berhasil di-submit.");
  } catch (err: any) {
    toast.error(
      err?.response?.data?.message ||
        err?.response?.data?.error ||
        "Gagal submit Pemesanan Bahan Baku."
    );
  } finally {
    isSubmitting.value = false;
  }
}

/* =========================
   DELETE DOCUMENT
========================= */

async function handleDeleteDocument() {
  if (!id.value) return;

  const confirmed = await confirm({
    title: "Hapus Pemesanan Bahan Baku",
    message: "Yakin ingin menghapus Pemesanan Bahan Baku ini?",
    confirmLabel: "Ya, Hapus",
    danger: true,
  });

  if (!confirmed) return;

  try {
    await remove(id.value);

    toast.success("Pemesanan Bahan Baku berhasil dihapus.");

    router.push("/purchasing/purchase-order");
  } catch (err: any) {
    toast.error(
      err?.response?.data?.message ||
        err?.response?.data?.error ||
        "Gagal menghapus Pemesanan Bahan Baku."
    );
  }
}

/* =========================
   BACK
========================= */

function handleBack() {
  router.push("/purchasing/purchase-order");
}

/* =========================
   FORMAT
========================= */

function formatCurrency(value: number) {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    maximumFractionDigits: 0,
  }).format(value);
}
</script>

<template>
  <SidebarAppLayout>
    <div class="bg-white border border-[#E5D9BF] rounded-xl p-6 mb-6">
      <!-- PAGE HEADER -->

      <div class="flex items-center justify-between mb-6">
        <div class="flex items-center gap-3">
          <button
            type="button"
            @click="handleBack"
            class="p-2 rounded-lg text-[#6B5D45] hover:bg-[#F8F4EC] transition">
            <ArrowLeft :size="18" />
          </button>

          <div>
            <h1 class="text-xl font-semibold text-[#3A2E1F]">
              Pemesanan Bahan Baku
            </h1>

            <p v-if="isEditMode" class="text-xs text-[#8A7A60] mt-1">
              {{ documentNo }}
            </p>
          </div>
        </div>

        <!-- ACTION -->

        <div class="flex items-center gap-2">
          <!-- NEW -->

          <template v-if="!isEditMode">
            <button
              type="button"
              @click="handleCreate"
              :disabled="isSubmitting"
              class="px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] disabled:opacity-50 disabled:cursor-not-allowed transition">
              {{ isSubmitting ? "Membuat Dokumen..." : "Buat Dokumen" }}
            </button>
          </template>

          <!-- DRAFT -->

          <template v-else-if="!isComplete">
            <button
              type="button"
              @click="handleDeleteDocument"
              :disabled="isSubmitting"
              class="px-4 py-2 text-sm font-medium text-[#B23A32] border border-[#D9CBB0] rounded-lg hover:bg-[#B23A32]/5 disabled:opacity-50 transition">
              Hapus
            </button>

            <button
              type="button"
              @click="handleUpdate"
              :disabled="isSubmitting"
              class="px-4 py-2 text-sm font-medium text-[#6B5D45] border border-[#D9CBB0] rounded-lg hover:bg-[#F8F4EC] disabled:opacity-50 transition">
              Simpan
            </button>

            <button
              type="button"
              @click="handleSubmit"
              :disabled="isSubmitting"
              class="px-4 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] disabled:opacity-50 transition">
              Submit
            </button>
          </template>

          <!-- COMPLETE -->

          <template v-else>
            <span
              class="px-3 py-2 rounded-lg bg-[#4C8C5B]/10 text-[#4C8C5B] text-sm font-semibold">
              COMPLETE
            </span>
          </template>
        </div>
      </div>

      <div v-if="pageLoading" class="py-20 text-center text-sm text-[#8A7A60]">
        Loading...
      </div>

      <template v-else>
        <!-- HEADER SECTION -->

        <div class="p-5 mb-6">
          <div class="mb-5">
            <h2 class="text-base font-semibold text-[#3A2E1F]">Detail</h2>

            <p class="text-xs text-[#8A7A60] mt-1">
              Informasi utama untuk pemesanan Barang
            </p>
          </div>

          <div class="grid grid-cols-12 gap-6">
            <!-- DOCUMENT NO -->

            <div class="col-span-4">
              <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
                No. Dokumen
              </label>

              <InputTextComponent
                v-model="documentNo"
                :disabled="isComplete"
                placeholder="Kosongkan untuk generate otomatis" />
            </div>

            <!-- DATE -->

            <div class="col-span-4">
              <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
                Tanggal Dokumen
              </label>

              <input
                v-model="orderDate"
                type="date"
                :disabled="isComplete"
                class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] disabled:bg-[#F8F4EC] disabled:text-[#9A8B73] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition" />
            </div>

            <!-- VENDOR -->

            <div class="col-span-4">
              <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
                Vendor
              </label>

              <InputTextComponent
                v-model="vendorName"
                :disabled="isComplete"
                placeholder="Masukkan nama vendor" />
            </div>

            <!-- REMARK -->

            <div class="col-span-12">
              <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
                Remark
              </label>

              <textarea
                v-model="remark"
                :disabled="isComplete"
                rows="3"
                placeholder="Catatan"
                class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] placeholder-[#B0A588] disabled:bg-[#F8F4EC] disabled:text-[#9A8B73] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition resize-none"></textarea>
            </div>
          </div>
        </div>

        <!-- ITEM SECTION -->

        <div class="p-5">
          <div class="flex items-center justify-between mb-5">
            <div>
              <h2 class="text-base font-semibold text-[#3A2E1F]">Items</h2>

              <p class="text-xs text-[#8A7A60] mt-1">
                Daftar barang yang dibeli
              </p>
            </div>

            <button
              v-if="!isComplete"
              type="button"
              @click="openAddItemModal"
              class="flex items-center gap-2 px-3 py-2 bg-[#7A1F2B] text-white text-sm font-medium rounded-lg hover:bg-[#5F1621] transition">
              <Plus :size="16" />

              Tambah Item
            </button>
          </div>

          <!-- TABLE -->

          <div class="overflow-x-auto border border-[#E5D9BF] rounded-lg">
            <table class="w-full text-sm">
              <thead class="bg-[#F8F4EC] border-b border-[#E5D9BF]">
                <tr>
                  <th
                    v-if="!isComplete"
                    class="px-4 py-3 text-center font-semibold text-[#6B5D45]">
                    Aksi
                  </th>
                  <th class="px-4 py-3 text-left font-semibold text-[#6B5D45]">
                    Product
                  </th>

                  <th class="px-4 py-3 text-left font-semibold text-[#6B5D45]">
                    Brand
                  </th>

                  <th class="px-4 py-3 text-right font-semibold text-[#6B5D45]">
                    Qty
                  </th>

                  <th class="px-4 py-3 text-right font-semibold text-[#6B5D45]">
                    Unit Price
                  </th>

                  <th class="px-4 py-3 text-right font-semibold text-[#6B5D45]">
                    Subtotal
                  </th>
                </tr>
              </thead>

              <tbody>
                <tr
                  v-for="item in items"
                  :key="item.id"
                  class="border-b border-[#EEE5D5] last:border-b-0">
                  <td v-if="!isComplete" class="px-4 py-3 text-center">
                    <OverlayMenu :items="buildItemMenu(item)" />
                  </td>
                  <td class="px-4 py-3 text-[#3A2E1F]">
                    {{ item.product_name }}
                  </td>

                  <td class="px-4 py-3 text-[#6B5D45]">
                    {{ item.brand || "-" }}
                  </td>

                  <td class="px-4 py-3 text-right text-[#3A2E1F]">
                    {{ item.quantity }}
                  </td>

                  <td class="px-4 py-3 text-right text-[#3A2E1F]">
                    {{ formatCurrency(item.unit_price) }}
                  </td>

                  <td class="px-4 py-3 text-right font-medium text-[#3A2E1F]">
                    {{ formatCurrency(item.subtotal) }}
                  </td>
                </tr>

                <!-- EMPTY -->

                <tr v-if="items.length === 0">
                  <td
                    :colspan="isComplete ? 5 : 6"
                    class="px-2 py-6 text-center text-sm text-[#9A8B73]">
                    Belum ada item.
                  </td>
                </tr>
              </tbody>

              <!-- TOTAL -->

              <tfoot>
                <tr class="bg-[#F8F4EC]">
                  <td
                    colspan="5"
                    class="px-4 py-4 text-right font-semibold text-[#6B5D45]">
                    Total
                  </td>

                  <td class="px-4 py-4 text-right font-bold text-[#3A2E1F]">
                    {{ formatCurrency(total) }}
                  </td>
                </tr>
              </tfoot>
            </table>
          </div>
        </div>
      </template>
    </div>

    <!-- ITEM MODAL -->

    <ChildModalWrapper
      v-model:visible="showItemModal"
      :headerTitle="itemModalTitle"
      @hide="closeModal">
      <div class="space-y-5">
        <!-- ERROR -->

        <div
          v-if="itemError"
          class="px-4 py-3 bg-[#B23A32]/10 text-[#B23A32] rounded-lg text-sm">
          {{ itemError }}
        </div>

        <!-- PRODUCT -->

        <div class="relative">
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Product
          </label>

          <div class="relative">
            <Search
              :size="16"
              class="absolute left-3 top-1/2 -translate-y-1/2 text-[#9A8B73]" />

            <input
              v-model="productKeyword"
              @input="handleProductSearch"
              type="text"
              placeholder="Cari product..."
              class="w-full pl-9 pr-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition" />
          </div>

          <!-- PRODUCT OPTIONS -->

          <div
            v-if="productOptions.length > 0"
            class="absolute z-50 left-0 right-0 mt-1 bg-white border border-[#D9CBB0] rounded-lg shadow-lg overflow-hidden">
            <button
              v-for="product in productOptions"
              :key="product.id"
              type="button"
              @click="selectProduct(product)"
              class="w-full px-4 py-3 text-left hover:bg-[#F8F4EC] transition">
              <div class="text-sm font-medium text-[#3A2E1F]">
                {{ product.name }}
              </div>

              <div class="text-xs text-[#8A7A60] mt-0.5">
                Unit: {{ product.unit }}
              </div>
            </button>
          </div>

          <div v-if="productLoading" class="text-xs text-[#8A7A60] mt-1">
            Mencari product...
          </div>
        </div>

        <!-- BRAND -->

        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Brand
          </label>

          <InputTextComponent
            v-model="itemBrand"
            placeholder="Masukkan brand" />
        </div>

        <!-- QUANTITY -->

        <div class="grid grid-cols-2 gap-4">
          <div>
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
              Quantity
            </label>

            <input
              v-model.number="itemQuantity"
              type="number"
              min="0"
              step="0.01"
              placeholder="0"
              class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition" />
          </div>

          <div>
            <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
              Unit Price
            </label>

            <input
              v-model.number="itemUnitPrice"
              type="number"
              min="0"
              step="1"
              placeholder="0"
              class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition" />
          </div>
        </div>

        <!-- REMARK -->

        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
            Remark
          </label>

          <textarea
            v-model="itemRemark"
            rows="3"
            placeholder="Catatan item"
            class="w-full px-4 py-2.5 bg-white border border-[#D9CBB0] rounded-lg text-sm text-[#3A2E1F] placeholder-[#B0A588] focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40 focus:border-[#C9A24B] transition resize-none"></textarea>
        </div>

        <!-- BUTTON -->

        <div class="flex justify-end gap-3 pt-4">
          <button
            type="button"
            @click="closeItemModal"
            class="px-4 py-2 text-sm font-medium text-[#6B5D45] border border-[#D9CBB0] rounded-lg hover:bg-[#F8F4EC] transition">
            Batal
          </button>

          <button
            type="button"
            @click="handleSaveItem"
            class="px-4 py-2 text-sm font-medium text-white bg-[#7A1F2B] rounded-lg hover:bg-[#5F1621] transition">
            {{ editingItemId !== null ? "Edit Item" : "Tambah Item" }}
          </button>
        </div>
      </div>
    </ChildModalWrapper>
  </SidebarAppLayout>
</template>
