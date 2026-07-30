<script setup lang="ts">
import { ref } from "vue";
import { ImagePlus, X, Loader2 } from "lucide-vue-next";
import { useUpload } from "../../composables/useUpload";

const photoUrls = defineModel<string[]>({ default: () => [] });
const { uploadFile, isUploading } = useUpload();
const fileInput = ref<HTMLInputElement | null>(null);

async function handleFileChange(e: Event) {
  const files = (e.target as HTMLInputElement).files;
  if (!files) return;

  for (const file of Array.from(files)) {
    const url = await uploadFile(file);
    photoUrls.value = [...photoUrls.value, url];
  }
  if (fileInput.value) fileInput.value.value = "";
}

function removePhoto(index: number) {
  photoUrls.value = photoUrls.value.filter((_, i) => i !== index);
}
</script>

<template>
  <div>
    <label class="block text-xs font-medium text-[#6B5D45] mb-1.5">
      Foto Referensi
    </label>
    <div class="flex flex-wrap gap-3">
      <div
        v-for="(url, index) in photoUrls"
        :key="index"
        class="relative w-20 h-20 rounded-lg overflow-hidden border border-[#D9CBB0]">
        <img :src="url" class="w-full h-full object-cover" />
        <button
          @click="removePhoto(index)"
          type="button"
          class="absolute top-1 right-1 w-5 h-5 flex items-center justify-center rounded-full bg-black/60 text-white hover:bg-black/80">
          <X :size="12" />
        </button>
      </div>

      <label
        class="w-20 h-20 flex flex-col items-center justify-center gap-1 rounded-lg border-2 border-dashed border-[#D9CBB0] text-[#B0A588] hover:border-[#C9A24B] hover:text-[#8A6D1F] cursor-pointer transition">
        <Loader2 v-if="isUploading" :size="20" class="animate-spin" />
        <ImagePlus v-else :size="20" />
        <span class="text-[10px]">{{
          isUploading ? "Upload..." : "Tambah"
        }}</span>
        <input
          ref="fileInput"
          type="file"
          accept="image/*"
          multiple
          class="hidden"
          @change="handleFileChange" />
      </label>
    </div>
  </div>
</template>
