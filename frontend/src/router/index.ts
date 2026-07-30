import { createRouter, createWebHistory } from "vue-router";
import LoginView from "../views/login-view.vue";
import DashboardView from "../views/dashboard/dashboard-view.vue";
import SizeView from "../views/master/size/size-view.vue";
import MaterialView from "../views/master/material/material-view.vue";
import StyleView from "../views/master/style/style-view.vue";
import ObjectView from "../views/master/object/object-view.vue";
import ClientView from "../views/transaction/client/client-view.vue";
import CommissionView from "../views/transaction/commission/commission-view.vue";

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
  {
    path: "/master/material",
    name: "Material",
    component: MaterialView,
    meta: { requiresAuth: true },
  },
  {
    path: "/master/styles",
    name: "Style",
    component: StyleView,
    meta: { requiresAuth: true },
  },
  {
    path: "/master/saints",
    name: "Saints",
    component: ObjectView,
    meta: { requiresAuth: true },
  },
  {
    path: "/transaction/clients",
    name: "Client",
    component: ClientView,
    meta: { requiresAuth: true },
  },
  {
    path: "/transaction/commissions",
    name: "CreateCommission",
    component: CommissionView,
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
