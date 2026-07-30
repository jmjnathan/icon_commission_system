<script setup lang="ts">
import ChildModalWrapper from "../../../components/modal/child-modal-wrapper.vue";
import { Calendar, User, FileText } from "lucide-vue-next";

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

interface CommissionDetail {
  id: number;
  client?: { name: string };
  deadline: string;
  notes: string;
  status: string;
  total_price: number;
  items: CommissionItem[];
  photos?: CommissionPhoto[];
}

const props = defineProps<{
  visible: boolean;
  commission: CommissionDetail | null;
}>();

const emit = defineEmits<{
  hide: [];
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
    </div>
  </ChildModalWrapper>
</template>
