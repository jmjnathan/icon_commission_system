import { ref } from "vue";
import api from "../../services/api";

export interface MasterSaints {
  id: number;
  name: string;
  size: string;
  status: string;
  created_at: string;
  created_username: string;
}

const items = ref<MasterSaints[]>([]);
const isLoading = ref(false);
const errorMessage = ref("");

export function useMasterSaints() {
  async function fetchAll() {
    isLoading.value = true;
    errorMessage.value = "";
    try {
      const response = await api.get("/master/get-saints");
      items.value = response.data.data || [];
    } catch (err: any) {
      errorMessage.value = err.response?.data?.error || "Gagal memuat data";
    } finally {
      isLoading.value = false;
    }
  }

  async function create(payload: { name: string; size: string }) {
    await api.post("/master/create-saints", payload);
    await fetchAll();
  }

  async function update(
    id: number,
    payload: { name: string; size: string; status: string }
  ) {
    await api.put(`/master/saints/edit/${id}`, payload);
    await fetchAll();
  }

  async function remove(id: number) {
    await api.delete(`/master/saints/delete/${id}`);
    await fetchAll();
  }

  return { items, isLoading, errorMessage, fetchAll, create, update, remove };
}
