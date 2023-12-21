import { createRouter, createWebHistory } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
import Home from '@/views/Home.vue';
import Result from '@/views/ResultInfo.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/atune'
  },
  {
    path: '/atune',
    component: Home,
    meta: { title: '任务列表' },
    children: [
      {
        path: 'detail',
        component: () => import("../components/atuneTemplete.vue")
      }
    ]
  },
  {
    path: '/result',
    component: Result,
    meta: { title: '结果展示' },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
