<script setup lang="ts">
import ChildModalWrapper from "../../../components/modal/child-modal-wrapper.vue";
import { Calendar, User, FileText } from "lucide-vue-next";
import { computed } from "vue";

interface CommissionItem {
  id: number;
  saint?: { name: string };
  size?: { name: string; size: string };
  material?: { name: string };
  style?: { name: string };
  price: number;
  notes: string;
}

interface CommissionPhoto {
  id: number;
  file_url: string;
}

interface CommissionPayment {
  id: number;
  amount: number;
  payment_type: string;
  method: string;
  paid_at: string;
  notes: string;
}

interface CommissionDetail {
  id: number;
  client?: { name: string };
  deadline: string;
  notes: string;
  status: string;
  subtotal: number;
  discount_type: string;
  discount_value: number;
  total_price: number;
  payment_status: string;
  items: CommissionItem[];
  photos?: CommissionPhoto[];
  payments?: CommissionPayment[];
}

const props = defineProps<{
  visible: boolean;
  commission: CommissionDetail | null;
}>();

function formatDate(dateStr: string) {
  return new Date(dateStr).toLocaleDateString("id-ID", {
    day: "numeric",
    month: "long",
    year: "numeric",
  });
}

function statusBadge(status: string) {
  const map: Record<string, string> = {
    pending: "bg-[#C9A24B]/15 text-[#8A6D1F]",
    in_progress: "bg-[#3B6FA8]/15 text-[#3B6FA8]",
    completed: "bg-[#4C8C5B]/15 text-[#4C8C5B]",
    delivered: "bg-[#8A7A5C]/15 text-[#8A7A5C]",
  };
  return map[status] || "bg-gray-100 text-gray-600";
}

const emit = defineEmits<{
  hide: [];
  "add-payment": [commission: CommissionDetail];
}>();

const remainingAmount = computed(() => {
  if (!props.commission) return 0;
  const paid = (props.commission.payments || []).reduce(
    (sum, p) => sum + p.amount,
    0
  );
  return Math.max(props.commission.total_price - paid, 0);
});

function paymentBadge(status: string) {
  const map: Record<string, string> = {
    unpaid: "bg-[#B23A32]/15 text-[#B23A32]",
    partial: "bg-[#C9A24B]/15 text-[#8A6D1F]",
    paid: "bg-[#4C8C5B]/15 text-[#4C8C5B]",
  };
  return map[status] || "bg-gray-100 text-gray-600";
}

function paymentStatusLabel(status: string) {
  const map: Record<string, string> = {
    unpaid: "Belum Bayar",
    partial: "DP / Sebagian",
    paid: "Lunas",
  };
  return map[status] || status;
}

function paymentTypeLabel(type: string) {
  const map: Record<string, string> = {
    dp: "DP",
    cicilan: "Cicilan",
    pelunasan: "Pelunasan",
  };
  return map[type] || type;
}
</script>

<template>
  <ChildModalWrapper
    :visible="visible"
    header-title="Detail Komisi"
    width-class="w-full max-w-2xl"
    @hide="emit('hide')">
    <div v-if="commission" class="space-y-6">
      <!-- Info Umum -->
      <div class="flex items-start justify-between">
        <div class="space-y-2">
          <div class="flex items-center gap-2 text-sm text-[#3A2E1F]">
            <User :size="16" class="text-[#8A7A5C]" />
            <span class="font-medium">{{ commission.client?.name }}</span>
          </div>
          <div class="flex items-center gap-2 text-sm text-[#6B5D45]">
            <Calendar :size="16" class="text-[#8A7A5C]" />
            <span>Deadline: {{ formatDate(commission.deadline) }}</span>
          </div>
          <div
            v-if="commission.notes"
            class="flex items-start gap-2 text-sm text-[#6B5D45]">
            <FileText :size="16" class="text-[#8A7A5C] mt-0.5" />
            <span>{{ commission.notes }}</span>
          </div>
        </div>
        <span
          class="px-3 py-1 rounded-full text-xs font-semibold shrink-0"
          :class="statusBadge(commission.status)">
          {{ commission.status.toUpperCase() }}
        </span>
      </div>

      <!-- Foto Referensi -->
      <div v-if="commission.photos && commission.photos.length > 0">
        <p
          class="text-xs font-semibold text-[#8A7A5C] uppercase tracking-wide mb-2">
          Foto Referensi
        </p>
        <div class="flex flex-wrap gap-3">
          <img
            v-for="photo in commission.photos"
            :key="photo.id"
            :src="photo.file_url"
            class="w-20 h-20 rounded-lg object-cover border border-[#E5D9BF]" />
        </div>
      </div>

      <!-- Daftar Item -->
      <div>
        <p
          class="text-xs font-semibold text-[#8A7A5C] uppercase tracking-wide mb-2">
          Detail Ikon Pesanan ({{ commission.items?.length || 0 }} item)
        </p>
        <div class="space-y-3">
          <div
            v-for="(item, index) in commission.items"
            :key="item.id"
            class="border border-[#E5D9BF] rounded-lg p-4">
            <div class="flex items-start justify-between mb-2">
              <p class="text-sm font-medium text-[#3A2E1F]">
                #{{ index + 1 }} — {{ item.saint?.name }}
              </p>
              <p class="text-sm font-semibold text-[#3A2E1F]">
                Rp {{ item.price?.toLocaleString("id-ID") }}
              </p>
            </div>
            <div class="grid grid-cols-2 gap-2 text-xs text-[#6B5D45]">
              <p>
                Ukuran:
                <span class="text-[#3A2E1F]"
                  >{{ item.size?.name }} ({{ item.size?.size }})</span
                >
              </p>
              <p>
                Material:
                <span class="text-[#3A2E1F]">{{ item.material?.name }}</span>
              </p>
              <p>
                Style:
                <span class="text-[#3A2E1F]">{{ item.style?.name }}</span>
              </p>
            </div>
            <p v-if="item.notes" class="text-xs text-[#8A7A5C] mt-2 italic">
              Catatan: {{ item.notes }}
            </p>
          </div>
        </div>
      </div>

      <!-- Total -->
      <div class="flex justify-end pt-4 border-t border-[#E5D9BF]">
        <div class="text-right">
          <p class="text-xs text-[#8A7A5C]">Total Harga</p>
          <p class="text-lg font-semibold text-[#3A2E1F]">
            Rp {{ commission.total_price?.toLocaleString("id-ID") }}
          </p>
        </div>
      </div>
      <!-- Ringkasan Harga dengan Diskon -->
      <div class="pt-4 border-t border-[#E5D9BF] space-y-1 text-sm">
        <div class="flex justify-between text-[#6B5D45]">
          <span>Subtotal</span>
          <span>Rp {{ commission.subtotal?.toLocaleString("id-ID") }}</span>
        </div>
        <div
          v-if="commission.discount_value > 0"
          class="flex justify-between text-[#B23A32]">
          <span
            >Diskon ({{
              commission.discount_type === "percent"
                ? commission.discount_value + "%"
                : "Nominal"
            }})</span
          >
          <span
            >- Rp
            {{
              (commission.subtotal - commission.total_price).toLocaleString(
                "id-ID"
              )
            }}</span
          >
        </div>
        <div
          class="flex justify-between text-base font-semibold text-[#3A2E1F]">
          <span>Total Tagihan</span>
          <span>Rp {{ commission.total_price?.toLocaleString("id-ID") }}</span>
        </div>
      </div>

      <!-- Riwayat Pembayaran -->
      <div class="pt-4 border-t border-[#E5D9BF]">
        <div class="flex items-center justify-between mb-2">
          <p
            class="text-xs font-semibold text-[#8A7A5C] uppercase tracking-wide">
            Riwayat Pembayaran
          </p>
          <span
            class="px-2 py-0.5 rounded-full text-xs font-medium"
            :class="paymentBadge(commission.payment_status)">
            {{ paymentStatusLabel(commission.payment_status) }}
          </span>
        </div>

        <p
          v-if="!commission.payments || commission.payments.length === 0"
          class="text-sm text-[#8A7A5C]">
          Belum ada pembayaran.
        </p>
        <div v-else class="space-y-2">
          <div
            v-for="p in commission.payments"
            :key="p.id"
            class="flex items-center justify-between text-sm border border-[#E5D9BF] rounded-lg px-3 py-2">
            <div>
              <p class="text-[#3A2E1F] font-medium">
                {{ paymentTypeLabel(p.payment_type) }} — {{ p.method }}
              </p>
              <p class="text-xs text-[#8A7A5C]">{{ formatDate(p.paid_at) }}</p>
            </div>
            <span class="text-[#4C8C5B] font-semibold"
              >+ Rp {{ p.amount.toLocaleString("id-ID") }}</span
            >
          </div>
        </div>

        <div
          class="flex items-center justify-between mt-3 pt-2 border-t border-dashed border-[#E5D9BF] text-sm">
          <span class="text-[#8A7A5C]">Sisa Tagihan</span>
          <span class="font-semibold text-[#3A2E1F]"
            >Rp {{ remainingAmount.toLocaleString("id-ID") }}</span
          >
        </div>

        <button
          v-if="commission.payment_status !== 'paid'"
          @click="emit('add-payment', commission)"
          class="w-full mt-3 py-2 text-sm font-medium text-[#7A1F2B] border border-[#7A1F2B]/40 rounded-lg hover:bg-[#7A1F2B]/5 transition">
          + Catat Pembayaran
        </button>
      </div>
    </div>
  </ChildModalWrapper>
</template>
