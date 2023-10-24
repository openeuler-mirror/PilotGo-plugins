import { createRouter, createWebHashHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import Home from '@/views/Home.vue'

const routes:Array<RouteRecordRaw> = [
    {
        path: '/',
        redirect: '/index',   //使用redirect重定向，默认系统显示的第一页
    },
    {
        path: '/index',
        component: Home,
        meta: { title: '首页'}
    },
]

const router =  createRouter({
    history: createWebHashHistory(),
    routes
})

export default router
