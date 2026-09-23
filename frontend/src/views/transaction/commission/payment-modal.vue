<script setup lang="ts">
import { ref, watch } from 'vue'
import ChildModalWrapper from '../../../components/modal/child-modal-wrapper.vue'
import { useCommission } from '../../../composables/commission/useCommission'
import { useToast } from '../../../composables/etc/useToast'

const props = defineProps<{
  visible: boolean
  commissionId: number | null
  remainingAmount: number
}>()

const emit = defineEmits<{ hide: [] }>()

const { addPayment } = useCommission()
const toast = useToast()

const amount = ref<number | null>(null)
const paymentType = ref('dp')
const method = ref('Transfer')
const notes = ref('')
const isSubmitting = ref(false)

watch(() => props.visible, (val) => {
  if (val) {
    amount.value = null
    paymentType.value = 'dp'
    method.value = 'Transfer'
    notes.value = ''
  }
})

async function handleSubmit() {
  if (!amount.value || !props.commissionId) return
  isSubmitting.value = true
  try {
    await addPayment(props.commissionId, {
      amount: amount.value,
      payment_type: paymentType.value,
      method: method.value,
      notes: notes.value,
    })
    toast.success('Pembayaran berhasil dicatat')
    emit('hide')
  } catch (err: any) {
    toast.error(err.response?.data?.error || 'Gagal mencatat pembayaran')
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <ChildModalWrapper :visible="visible" header-title="Catat Pembayaran" width-class="w-full max-w-md" @hide="emit('hide')">
    <form @submit.prevent="handleSubmit" class="space-y-4">
      <div class="bg-[#F5EFE3] rounded-lg px-4 py-2.5 flex items-center justify-between text-sm">
        <span class="text-[#8A7A5C]">Sisa Tagihan</span>
        <span class="font-semibold text-[#3A2E1F]">Rp {{ remainingAmount.toLocaleString('id-ID') }}</span>
      </div>

      <div>
        <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">Jumlah Bayar</label>
        <input v-model.number="amount" type="number" required
          class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
      </div>

      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">Jenis</label>
          <select v-model="paymentType" class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40">
            <option value="dp">DP (Tanda Jadi)</option>
            <option value="cicilan">Cicilan</option>
            <option value="pelunasan">Pelunasan</option>
          </select>
        </div>
        <div>
          <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">Metode</label>
          <select v-model="method" class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40">
            <option value="Transfer">Transfer</option>
            <option value="Cash">Cash</option>
            <option value="QRIS">QRIS</option>
          </select>
        </div>
      </div>

      <div>
        <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">Catatan</label>
        <input v-model="notes" type="text" placeholder="Opsional"
          class="w-full px-4 py-2.5 border border-[#D9CBB0] rounded-lg text-sm focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40" />
      </div>

      <button type="submit" :disabled="isSubmitting"
        class="w-full py-2.5 bg-[#7A1F2B] text-white rounded-lg font-medium hover:bg-[#5F1621] transition disabled:opacity-50">
        {{ isSubmitting ? "Menyimpan..." : "Simpan Pembayaran" }}
      </button>
    </form>
  </ChildModalWrapper>
</template>