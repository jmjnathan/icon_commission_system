import { ref } from "vue";
import api from "../services/api";
import router from "../router";

const isLoading = ref(false);
const errorMessage = ref("");

export function useAuth() {
  async function login(username: string, password: string) {
    isLoading.value = true;
    errorMessage.value = "";

    try {
      console.log("🚀 LOGIN REQUEST", {
        username,
        password,
      });

      const response = await api.post("/login", { username, password });

      console.log("✅ LOGIN RESPONSE", response.status, response.data);

      const token = response.data.data.token;

      localStorage.setItem("token", token);
      localStorage.setItem("username", response.data.data.username);

      router.push("/");
    } catch (err: any) {
      console.error("❌ LOGIN ERROR", err);
      console.error("Status:", err.response?.status);
      console.error("Response:", err.response?.data);
      console.error("Request:", err.config);

      errorMessage.value =
        err.response?.data?.error || "Login gagal, coba lagi";
    } finally {
      isLoading.value = false;
    }
  }

  function logout() {
    localStorage.removeItem("token");
    localStorage.removeItem("username");
    router.push("/login");
  }

  return { login, logout, isLoading, errorMessage };
}
