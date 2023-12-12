import { createRouter, createWebHashHistory } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
import Home from '@/views/Home.vue';
import Result from '@/views/ResultInfo.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/plugin/atune', //使用redirect重定向，默认系统显示的第一页
  },
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
  history: createWebHashHistory(),
  routes,
});

export default router;
