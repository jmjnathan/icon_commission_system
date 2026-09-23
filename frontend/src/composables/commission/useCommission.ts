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

export interface CommissionPayment {
  id: number;
  amount: number;
  payment_type: string;
  method: string;
  paid_at: string;
  notes: string;
}

export interface Commission {
  id: number;
  client: {
    name: string;
    phone: string;
    address: string;
    city: string;
    province: string;
    postal_code: string;
  };
  order_date: string;
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
      console.log("COMMISSION DATA:", commissions.value);
    } catch (err: any) {
      errorMessage.value =
        err.response?.data?.error || "Gagal memuat data komisi";
    } finally {
      isLoading.value = false;
    }
  }

  async function create(payload: any) {
    await api.post("/commissions/create", payload);
    await fetchAll();
  }

  async function updateStatus(id: number, status: string) {
    await api.patch(`/commissions/edit/${id}/status`, { status });
    await fetchAll();
  }

  async function addPayment(
    commissionId: number,
    payload: {
      amount: number;
      payment_type: string;
      method: string;
      notes?: string;
    }
  ) {
    await api.post(`/commissions/${commissionId}/payments`, payload);
    await fetchAll();
  }

  return {
    commissions,
    isLoading,
    errorMessage,
    fetchAll,
    create,
    updateStatus,
    addPayment,
  };
}
