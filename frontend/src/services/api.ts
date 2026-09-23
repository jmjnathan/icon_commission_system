import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:8080",
});

// Interceptor request — otomatis nambahin token JWT ke tiap request
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Interceptor response — handle status code tertentu
api.interceptors.response.use(
  (response) => response,
  (error) => {
    const status = error.response?.status;
    const requestUrl = error.config?.url ?? "";

    if (status === 401 && !requestUrl.endsWith("/login")) {
      localStorage.removeItem("token");
      window.location.href = "/login";
    }

    if (status === 403) {
      alert("Kamu tidak punya akses untuk aksi ini");
    }

    if (status === 500) {
      alert("Terjadi kesalahan pada server");
    }

    return Promise.reject(error);
  }
);
export default api;
