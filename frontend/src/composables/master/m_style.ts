import { ref } from "vue";
import api from "../../services/api";

export interface MasterStyle {
  id: number;
  name: string;
  remark: string;
  status: string;
  created_at: string;
  created_username: string;
}

const items = ref<MasterStyle[]>([]);
const isLoading = ref(false);
const errorMessage = ref("");

export function useMasterStyle() {
  async function fetchAll() {
    isLoading.value = true;
    errorMessage.value = "";
    try {
      const response = await api.get("/master/get-style");
      items.value = response.data.data || [];
    } catch (err: any) {
      errorMessage.value = err.response?.data?.error || "Gagal memuat data";
    } finally {
      isLoading.value = false;
    }
  }

  async function create(payload: { name: string; remark: string }) {
    await api.post("/master/create-style", payload);
    await fetchAll();
  }

  async function update(
    id: number,
    payload: { name: string; remark: string; status: string }
  ) {
    await api.put(`/master/style/edit/${id}`, payload);
    await fetchAll();
  }

  async function remove(id: number) {
    await api.delete(`/master/style/delete/${id}`);
    await fetchAll();
  }

  return { items, isLoading, errorMessage, fetchAll, create, update, remove };
}
