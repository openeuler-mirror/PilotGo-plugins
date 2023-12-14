import { createRouter, createWebHistory } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
import Home from '@/views/Home.vue';
import Result from '@/views/ResultInfo.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/plugin/atune',
    component: Home,
    meta: { title: '首页' },
  },
  {
    path: '/plugin/atune/result',
    component: Result,
    meta: { title: '结果展示' },
  },
];

const router = createRouter({
  history: createWebHistory('/'),
  routes,
});

export default router;
