import { createRouter, createWebHistory } from 'vue-router'
import Main from '@/components/Main.vue'
import XZkaknazvat from '@/components/XZkaknazvat.vue'
import RoomDetails from '@/components/RoomDetails.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Main,
    },
    {
      path: '/auth/yandex/callback',
      name: 'xz',
      component: XZkaknazvat
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
    },
    { path: '/rooms/:id',
      name: 'rooms',
       component: RoomDetails
    },
  ],
})

export default router
