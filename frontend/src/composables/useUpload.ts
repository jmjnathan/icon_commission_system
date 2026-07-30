import { ref } from "vue";
import api from "../services/api";

export function useUpload() {
  const isUploading = ref(false);

  async function uploadFile(file: File): Promise<string> {
    isUploading.value = true;
    try {
      const formData = new FormData();
      formData.append("file", file);
      const response = await api.post("/upload", formData, {
        headers: { "Content-Type": "multipart/form-data" },
      });
      return response.data.data.file_url;
    } finally {
      isUploading.value = false;
    }
  }

  return { uploadFile, isUploading };
}
