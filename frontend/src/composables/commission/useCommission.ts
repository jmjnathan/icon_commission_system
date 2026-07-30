import { ref } from "vue";
import api from "../../services/api";

export interface CommissionItem {
  id: number;
  saint: { name: string };
  size: { name: string; size: string };
  material: { name: string };
  style: { name: string };
  price: number;
  notes: string;
}

export interface CommissionPhoto {
  id: number;
  file_url: string;
}

export interface Commission {
  id: number;
  client: { name: string };
  order_date: string;
  deadline: string;
  notes: string;
  status: string;
  total_price: number;
  items: CommissionItem[];
  photos?: CommissionPhoto[];
}

const commissions = ref<Commission[]>([]);
const isLoading = ref(false);
const errorMessage = ref("");

export function useCommission() {
  async function fetchAll() {
    isLoading.value = true;
    errorMessage.value = "";
    try {
      const response = await api.get("/commissions/get-list");
      commissions.value = response.data.data || [];
    } catch (err: any) {
      errorMessage.value =
        err.response?.data?.error || "Gagal memuat data komisi";
    } finally {
      isLoading.value = false;
    }
  }

  async function updateStatus(id: number, status: string) {
    await api.patch(`/commissions/edit/${id}/status`, { status });
    await fetchAll();
  }

  async function create(payload: any) {
    await api.post("/commissions/create", payload);
  }

  return {
    commissions,
    isLoading,
    errorMessage,
    fetchAll,
    updateStatus,
    create,
  };
}
