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
    if (error.response?.status === 401) {
      localStorage.removeItem("token");
      window.location.href = "/login";
    }
    if (error.response?.status === 403) {
      alert("Kamu tidak punya akses untuk aksi ini");
    }
    if (error.response?.status === 500) {
      alert("Terjadi kesalahan pada server");
    }
    return Promise.reject(error);
  }
);

export default api;
