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
}

export interface MaterialComponent {
  id: number;
  name: string;
  category: string;
  remark: string;
  status: string;
  variants?: MaterialComponentVariant[];
}

const items = ref<MaterialComponent[]>([]);
const variants = ref<MaterialComponentVariant[]>([]);

const isLoading = ref(false);
const isVariantLoading = ref(false);

export function useMaterialComponent() {
  async function fetchAll() {
    isLoading.value = true;

    try {
      const response = await api.get(
        "/master/get-material-component"
      );

      items.value = response.data.data || [];
    } finally {
      isLoading.value = false;
    }
  }

  async function create(payload: {
    name: string;
    category: string;
    remark: string;
  }) {
    await api.post(
      "/master/create-material-component",
      payload
    );

    await fetchAll();
  }

  async function update(
    id: number,
    payload: {
      name: string;
      category: string;
      remark: string;
      status: string;
    }
  ) {
    await api.put(
      `/master/material-component/edit/${id}`,
      payload
    );

    await fetchAll();
  }

  async function remove(id: number) {
    await api.delete(
      `/master/material-component/delete/${id}`
    );

    await fetchAll();
  }

  async function fetchVariants(materialComponentId: number) {
    isVariantLoading.value = true;

    try {
      const response = await api.get(
        `/master/material-component/${materialComponentId}/variants`
      );

      variants.value = response.data.data || [];
    } finally {
      isVariantLoading.value = false;
    }
  }

  async function createVariant(
    materialComponentId: number,
    payload: {
      specification: string;
      length: number;
      width: number;
      thickness: number;
      unit: string;
      unit_price: number;
      remark: string;
    }
  ) {
    await api.post(
      `/master/material-component/${materialComponentId}/variants`,
      payload
    );

    await fetchVariants(materialComponentId);
  }

  async function updateVariant(
    id: number,
    materialComponentId: number,
    payload: {
      specification: string;
      length: number;
      width: number;
      thickness: number;
      unit: string;
      unit_price: number;
      remark: string;
      status: string;
    }
  ) {
    await api.put(
      `/master/material-component/variant/edit/${id}`,
      payload
    );

    await fetchVariants(materialComponentId);
  }

  async function removeVariant(
    id: number,
    materialComponentId: number
  ) {
    await api.delete(
      `/master/material-component/variant/delete/${id}`
    );

    await fetchVariants(materialComponentId);
  }

  return {
    items,
    variants,

    isLoading,
    isVariantLoading,

    fetchAll,
    create,
    update,
    remove,

    fetchVariants,
    createVariant,
    updateVariant,
    removeVariant,
  };
}