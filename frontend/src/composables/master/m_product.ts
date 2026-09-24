import { ref } from "vue";
import api from "../../services/api";

export interface MasterProduct {
  id: number;
  name: string;
  unit: string;
  description: string | null;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

interface CreateProductRequest {
  name: string;
  unit: string;
  description: string | null;
}

interface UpdateProductRequest {
  name: string;
  unit: string;
  description: string | null;
  is_active: boolean;
}

interface ApiResponse<T> {
  data: T;
}

export function useMasterProduct() {
  const items = ref<MasterProduct[]>([]);
  const isLoading = ref(false);

  async function fetchAll() {
    isLoading.value = true;

    try {
      const response = await api.get<ApiResponse<MasterProduct[]>>(
        "/master/products"
      );

      items.value = response.data.data;
    } finally {
      isLoading.value = false;
    }
  }

  async function create(payload: CreateProductRequest) {
    const response = await api.post<ApiResponse<MasterProduct>>(
      "/master/products",
      payload
    );

    return response.data.data;
  }

  async function update(id: number, payload: UpdateProductRequest) {
    const response = await api.put<ApiResponse<MasterProduct>>(
      `/master/products/${id}`,
      payload
    );

    return response.data.data;
  }

  async function remove(id: number) {
    const response = await api.delete(`/master/products/${id}`);

    return response.data;
  }

  async function search(keyword: string) {
    const response = await api.get<ApiResponse<MasterProduct[]>>(
      "/master/products/search",
      {
        params: {
          keyword,
        },
      }
    );

    return response.data.data;
  }

  return {
    items,
    isLoading,
    fetchAll,
    create,
    update,
    remove,
    search,
  };
}
