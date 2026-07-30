import { ref } from "vue";
import api from "../../services/api";

export interface ClientApi {
  id: number;
  name: string;
  nickname: string;
  phone: string;
  email: string;
  instagram_handle: string;
  address: string;
  city: string;
  province: string;
  postal_code: string;
  denomination: string;
  patron_saint_preference: string;
  preferred_payment_method: string;
  created_at: string;
  created_username: string;
}

const items = ref<ClientApi[]>([]);
const isLoading = ref(false);
const errorMessage = ref("");

export function useClientApi() {
  async function fetchAll() {
    isLoading.value = true;
    errorMessage.value = "";
    try {
      const response = await api.get("/master/get-clients");
      items.value = response.data.data || [];
    } catch (err: any) {
      errorMessage.value = err.response?.data?.error || "Gagal memuat data";
    } finally {
      isLoading.value = false;
    }
  }

  async function create(payload: {
    name: string;
    nickname: string;
    phone: string;
    email: string;
    instagram_handle: string;
    address: string;
    city: string;
    province: string;
    postal_code: string;
    denomination: string;
    patron_saint_preference: string;
    preferred_payment_method: string;
  }) {
    await api.post("/master/create-clients", payload);
    await fetchAll();
  }

  async function update(
    id: number,
    payload: {
      name: string;
      nickname: string;
      phone: string;
      email: string;
      instagram_handle: string;
      address: string;
      city: string;
      province: string;
      postal_code: string;
      denomination: string;
      patron_saint_preference: string;
      preferred_payment_method: string;
    }
  ) {
    await api.put(`/master/clients/edit/${id}`, payload);
    await fetchAll();
  }

  async function remove(id: number) {
    await api.delete(`/master/clients/delete/${id}`);
    await fetchAll();
  }

  return { items, isLoading, errorMessage, fetchAll, create, update, remove };
}
