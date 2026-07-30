import { createRouter, createWebHistory } from "vue-router";
import LoginView from "../views/login-view.vue";
import DashboardView from "../views/dashboard/dashboard-view.vue";
import SizeView from "../views/master/size/size-view.vue";

const routes = [
  {
    path: "/login",
    name: "Login",
    component: LoginView,
  },
  {
    path: "/",
    name: "Dashboard",
    component: DashboardView,
    meta: { requiresAuth: true },
  },
  {
    path: "/master/size",
    name: "Size",
    component: SizeView,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, _from) => {
  const token = localStorage.getItem("token");

  if (to.meta.requiresAuth && !token) {
    return "/login";
  } else if (to.name === "Login" && token) {
    return "/";
  }
});

export default router;
