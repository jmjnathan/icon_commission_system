import { ref } from "vue";
import api from "../../services/api";

export interface MaterialComponentVariant {
  id: number;
  material_component_id: number;
  specification: string;
  length: number;
  width: number;
  thickness: number;
  unit: string;
  unit_price: number;
  remark: string;
  status: string;

  material_component?: {
    id: number;
    name: string;
    category: string;
  };
}

export interface CommissionItemMaterial {
  id: number;
  commission_item_id: number;
  material_component_variant_id: number;

  quantity: number;
  unit_price: number;
  total_cost: number;

  remark: string;
  status: string;

  material_component_variant: MaterialComponentVariant;
}

const items = ref<CommissionItemMaterial[]>([]);
const variants = ref<MaterialComponentVariant[]>([]);

const isLoading = ref(false);
const isVariantLoading = ref(false);

export function useCommissionItemMaterial() {
  async function fetchMaterials(
    commissionItemId: number
  ) {
    isLoading.value = true;

    try {
      const response = await api.get(
        `/commission-items/${commissionItemId}/materials`
      );

      items.value = response.data.data || [];
    } finally {
      isLoading.value = false;
    }
  }

  async function fetchVariants() {
    isVariantLoading.value = true;

    try {
      const response = await api.get(
        "/master/material-component/variants"
      );

      variants.value = response.data.data || [];
    } finally {
      isVariantLoading.value = false;
    }
  }

  async function createMaterial(
    commissionItemId: number,
    payload: {
      material_component_variant_id: number;
      quantity: number;
      remark: string;
    }
  ) {
    await api.post(
      `/commission-items/${commissionItemId}/materials`,
      payload
    );

    await fetchMaterials(commissionItemId);
  }

  async function updateMaterial(
    id: number,
    commissionItemId: number,
    payload: {
      material_component_variant_id: number;
      quantity: number;
      remark: string;
    }
  ) {
    await api.put(
      `/commission-item-materials/${id}`,
      payload
    );

    await fetchMaterials(commissionItemId);
  }

  async function deleteMaterial(
    id: number,
    commissionItemId: number
  ) {
    await api.delete(
      `/commission-item-materials/${id}`
    );

    await fetchMaterials(commissionItemId);
  }

  return {
    items,
    variants,

    isLoading,
    isVariantLoading,

    fetchMaterials,
    fetchVariants,

    createMaterial,
    updateMaterial,
    deleteMaterial,
  };
}