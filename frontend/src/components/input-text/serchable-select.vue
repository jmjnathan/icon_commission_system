<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { ChevronDown, X, Search, Loader2 } from 'lucide-vue-next'

interface Option {
  label: string
  value: any
}

const props = withDefaults(defineProps<{
  label?: string
  placeholder?: string
  modelValue: any
  options: Option[]
  loading?: boolean
  disabled?: boolean
  readOnly?: boolean
  errorMessage?: string
  clearable?: boolean
}>(), {
  label: '',
  placeholder: 'Pilih atau cari...',
  loading: false,
  disabled: false,
  readOnly: false,
  clearable: true,
})

const emit = defineEmits<{
  'update:modelValue': [value: any]
}>()

const isOpen = ref(false)
const search = ref('')
const wrapperRef = ref<HTMLDivElement | null>(null)

const selectedLabel = computed(() => {
  const found = props.options.find((o) => o.value === props.modelValue)
  return found?.label ?? ''
})

const filteredOptions = computed(() => {
  if (!search.value) return props.options
  return props.options.filter((o) =>
    o.label.toLowerCase().includes(search.value.toLowerCase())
  )
})

function toggleOpen() {
  if (props.disabled || props.readOnly) return
  isOpen.value = !isOpen.value
  if (isOpen.value) search.value = ''
}

function selectOption(option: Option) {
  emit('update:modelValue', option.value)
  isOpen.value = false
  search.value = ''
}

function clearSelection(e: Event) {
  e.stopPropagation()
  emit('update:modelValue', null)
}

function handleClickOutside(e: MouseEvent) {
  if (wrapperRef.value && !wrapperRef.value.contains(e.target as Node)) {
    isOpen.value = false
  }
}

onMounted(() => document.addEventListener('click', handleClickOutside))
onUnmounted(() => document.removeEventListener('click', handleClickOutside))
</script>

<template>
  <div ref="wrapperRef" class="relative">
    <label v-if="label" class="block text-xs font-medium mb-1.5" :class="errorMessage ? 'text-[#B23A32]' : 'text-[#6B5D45]'">
      {{ label }}
    </label>

    <button
      type="button"
      @click="toggleOpen"
      :disabled="disabled"
      class="w-full flex items-center justify-between px-4 py-2.5 bg-white border rounded-lg text-sm text-left transition disabled:bg-[#F5EFE3] disabled:cursor-not-allowed"
      :class="errorMessage ? 'border-[#B23A32]' : 'border-[#D9CBB0]'"
    >
      <span :class="selectedLabel ? 'text-[#3A2E1F]' : 'text-[#B0A588]'">
        {{ selectedLabel || placeholder }}
      </span>
      <div class="flex items-center gap-1 shrink-0">
        <Loader2 v-if="loading" :size="14" class="animate-spin text-[#8A7A5C]" />
        <button
          v-if="clearable && modelValue && !loading"
          @click="clearSelection"
          type="button"
          class="text-[#B0A588] hover:text-[#3A2E1F]"
        >
          <X :size="14" />
        </button>
        <ChevronDown :size="16" class="text-[#8A7A5C] transition-transform" :class="isOpen ? 'rotate-180' : ''" />
      </div>
    </button>

    <div
      v-if="isOpen"
      class="absolute z-20 mt-1 w-full bg-white border border-[#E5D9BF] rounded-lg shadow-lg overflow-hidden"
    >
      <div class="p-2 border-b border-[#E5D9BF]">
        <div class="relative">
          <Search :size="14" class="absolute left-3 top-1/2 -translate-y-1/2 text-[#B0A588]" />
          <input
            v-model="search"
            type="text"
            placeholder="Cari..."
            autofocus
            class="w-full pl-8 pr-3 py-1.5 text-sm border border-[#D9CBB0] rounded-md focus:outline-none focus:ring-2 focus:ring-[#C9A24B]/40"
          />
        </div>
      </div>

      <div class="max-h-56 overflow-y-auto">
        <p v-if="filteredOptions.length === 0" class="px-4 py-3 text-sm text-[#8A7A5C]">
          Tidak ditemukan.
        </p>
        <button
          v-for="option in filteredOptions"
          :key="option.value"
          type="button"
          @click="selectOption(option)"
          class="w-full text-left px-4 py-2 text-sm text-[#3A2E1F] hover:bg-[#F5EFE3] transition"
          :class="option.value === modelValue ? 'bg-[#C9A24B]/10 font-medium' : ''"
        >
          {{ option.label }}
        </button>
      </div>
    </div>

    <p v-if="errorMessage" class="text-xs text-[#B23A32] mt-1">{{ errorMessage }}</p>
  </div>
</template>