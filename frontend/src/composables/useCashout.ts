import { ref } from "vue";
import api from "../services/api";

export interface CashOut {
  id: number;
  category: string;
  description: string;
  qty: number;
  unit_price: number;
  amount: number;
  vendor: string;
  payment_method: string;
  receipt_url: string;
  notes: string;
  date: string;
}

interface CashOutPayload {
  category: string;
  description: string;
  qty: number;
  unit_price: number;
  vendor?: string;
  payment_method?: string;
  receipt_url?: string;
  notes?: string;
  date: string;
}

const items = ref<CashOut[]>([]);
const isLoading = ref(false);

export function useCashOut() {
  async function fetchAll() {
    isLoading.value = true;
    try {
      const response = await api.get("/cashflow/cashout/get-list");
      items.value = response.data.data || [];
    } finally {
      isLoading.value = false;
    }
  }

  async function create(payload: CashOutPayload) {
    await api.post("/cashflow/cashout/create", payload);
    await fetchAll();
  }

  async function update(id: number, payload: CashOutPayload) {
    await api.put(`/cashflow/cashout/edit/${id}`, payload);
    await fetchAll();
  }

  async function remove(id: number) {
    await api.delete(`/cashflow/cashout/delete/${id}`);
    await fetchAll();
  }

  return { items, isLoading, fetchAll, create, update, remove };
}
