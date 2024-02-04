import {createRouter, createWebHistory} from 'vue-router'
import ViewContainer from '@/views/layout/container.vue'
import ViewList from '@/views/list/index.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/login',
            component: () => import('@/views/login/index.vue'),
            hidden: true
        },
        {
            path: '/',
            name: 'index',
            component: ViewContainer,

            children: [
                {
                    path: '/',
                    component: ViewList
                },
                {
                    path: '/create',
                    component: () => import('@/views/create/index.vue')
                }
            ]
        },
        {
            path: '/404',
            component: () => import('@/views/404.vue')
        }
    ]
})

export default router
