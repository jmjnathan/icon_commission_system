import { ref } from "vue";
import api from "../../../services/api";

export interface PurchaseOrderItem {
  id: number;
  purchase_order_id: number;
  product_id: number;
  product_name: string;
  brand: string | null;
  quantity: number;
  unit_price: number;
  subtotal: number;
  remark: string | null;
}

export interface PurchaseOrder {
  id: number;
  document_name: string;
  document_no: string;
  vendor_name: string;
  order_date: string;
  status: string;
  remark: string | null;
  total: number;
  created_at: string;
  updated_at: string;
}

export interface PurchaseOrderDetail extends PurchaseOrder {
  items: PurchaseOrderItem[];
}

export interface Product {
  id: number;
  name: string;
  unit: string;
  description: string | null;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreatePurchaseOrderItemRequest {
  product_id: number;
  product_name: string;
  quantity: number;
  unit_price: number;
  brand?: string | null;
  remark?: string | null;
}

export interface CreatePurchaseOrderRequest {
  document_no?: string;
  vendor_name: string;
  order_date: string;
  remark?: string | null;
  items: CreatePurchaseOrderItemRequest[];
}

export interface UpdatePurchaseOrderRequest {
  document_no?: string;
  vendor_name: string;
  order_date: string;
  remark?: string | null;
  items: CreatePurchaseOrderItemRequest[];
}

interface ApiResponse<T> {
  data: T;
  message?: string;
}

export function usePurchaseOrder() {
  const items = ref<PurchaseOrder[]>([]);
  const isLoading = ref(false);

  async function fetchAll() {
    isLoading.value = true;

    try {
      const response = await api.get<ApiResponse<PurchaseOrder[]>>(
        "/purchase-order"
      );

      items.value = response.data.data;
    } finally {
      isLoading.value = false;
    }
  }

  async function getById(id: number) {
    const response = await api.get<ApiResponse<PurchaseOrderDetail>>(
      `/purchase-order/${id}`
    );

    return response.data.data;
  }

  async function getItems(id: number) {
    const response = await api.get<ApiResponse<PurchaseOrderItem[]>>(
      `/purchase-order/${id}/items`
    );

    return response.data.data;
  }

  async function create(payload: CreatePurchaseOrderRequest) {
    const response = await api.post<ApiResponse<PurchaseOrder>>(
      "/purchase-order",
      payload
    );

    return response.data.data;
  }

  async function update(id: number, payload: UpdatePurchaseOrderRequest) {
    const response = await api.put<ApiResponse<PurchaseOrder>>(
      `/purchase-order/${id}`,
      payload
    );

    return response.data.data;
  }

  async function submit(id: number) {
    const response = await api.post(`/purchase-order/${id}/submit`);

    return response.data;
  }

  async function remove(id: number) {
    const response = await api.delete(`/purchase-order/${id}`);

    return response.data;
  }

  async function searchProduct(keyword: string) {
    const response = await api.get<ApiResponse<Product[]>>(
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
    getById,
    getItems,
    create,
    update,
    submit,
    remove,
    searchProduct,
  };
}
