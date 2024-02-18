import {createRouter, createWebHistory} from 'vue-router'
import ViewContainer from '@/views/layout/container.vue'
import ViewList from '@/views/list/index.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/404',
            component: () => import('@/views/404.vue')
        },
        // 404 page must be placed at the end !!!
        { path: '/:pathMatch(.*)*', redirect: '/404', hidden: true },
        {
            path: '/login',
            component: () => import('@/views/login/index.vue'),
            hidden: true
        },
        {
            path: '/register',
            component: () => import('@/views/register/index.vue'),
            hidden: true
        },
        {
            path: '/',
            name: 'index',
            component: ViewContainer,

            children: [
                {
                    path: '',
                    name: 'list',
                    component: ViewList
                },
                {
                    path: 'create',
                    component: () => import('@/views/create/index.vue')
                },
                {
                    path: 'detail',
                    component: () => import('@/views/detail/index.vue')
                },
                {
                    path: 'edit',
                    name: 'edit',
                    component: () => import('@/views/edit/index.vue')
                }
            ]
        }
    ]
})

export default router
