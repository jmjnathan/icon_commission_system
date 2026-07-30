import { ref } from "vue";

interface ConfirmOptions {
  title?: string;
  message: string;
  confirmLabel?: string;
  cancelLabel?: string;
  danger?: boolean;
}

const isVisible = ref(false);
const options = ref<ConfirmOptions>({ message: "" });
let resolvePromise: ((value: boolean) => void) | null = null;

export function useConfirm() {
  function confirm(opts: ConfirmOptions): Promise<boolean> {
    options.value = {
      title: "Konfirmasi",
      confirmLabel: "Ya, Lanjutkan",
      cancelLabel: "Batal",
      danger: false,
      ...opts,
    };
    isVisible.value = true;

    return new Promise((resolve) => {
      resolvePromise = resolve;
    });
  }

  function handleConfirm() {
    isVisible.value = false;
    resolvePromise?.(true);
    resolvePromise = null;
  }

  function handleCancel() {
    isVisible.value = false;
    resolvePromise?.(false);
    resolvePromise = null;
  }

  return { isVisible, options, confirm, handleConfirm, handleCancel };
}
