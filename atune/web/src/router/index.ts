import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import Home from "@/views/Home.vue";
import Atune from "@/views/atuneList.vue";
import Result from "@/views/ResultInfo.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "",
    redirect: "/task",
  },
  {
    path: "/task",
    component: Home,
    meta: { title: "任务列表" },
    children: [
      {
        path: "detail",
        name: "taskDetail",
        component: () => import("../views/taskDetail.vue"),
      },
    ],
  },
  {
    path: "/atune",
    component: Atune,
    meta: { title: "模板列表" },
    children: [
      {
        path: "detail",
        name: "atuneDetail",
        component: () => import("../views/atuneDetail.vue"),
      },
    ],
  },
  {
    path: "/result",
    component: Result,
    meta: { title: "结果展示" },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
