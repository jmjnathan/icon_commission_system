import { createRouter, createWebHistory } from "vue-router";
import LoginView from "../views/login-view.vue";
import DashboardView from "../views/dashboard/dashboard-view.vue";
import SizeView from "../views/master/size/size-view.vue";
import MaterialView from "../views/master/material/material-view.vue";
import StyleView from "../views/master/style/style-view.vue";
import ObjectView from "../views/master/object/object-view.vue";
import ClientView from "../views/transaction/client/client-view.vue";
import CommissionView from "../views/transaction/commission/commission-view.vue";
import HistoryTransactionView from "../views/history-transaction/history-transaction-view.vue";
import CashoutView from "../views/cashflow/cashout/cashout-view.vue";
import FinancialReportView from "../views/cashflow/financial-report/financial-report-view.vue";
import MaterialComponent from "../views/master/materialComponent/materialComponent.vue";
import ProductView from "../views/master/product/product-view.vue";

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
    path: "/master/product",
    name: "Product",
    component: ProductView,
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
  {
    path: "/transaction/history-transaction",
    name: "HistoryTransaction",
    component: HistoryTransactionView,
    meta: { requiresAuth: true },
  },
  {
    path: "/cashflow/cash-out",
    name: "CashOut",
    component: CashoutView,
    meta: { requiresAuth: true },
  },
  {
    path: "/cashflow/report",
    name: "FinancialReport",
    component: FinancialReportView,
    meta: { requiresAuth: true },
  },
  {
    path: "/master/material-components",
    name: "MaterialComponent",
    component: MaterialComponent,
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
