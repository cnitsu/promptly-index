import { createMemoryHistory, createRouter, type RouteRecordRaw } from 'vue-router'
import IndexView from '../pages/IndexView.vue'

const routes: RouteRecordRaw[] = [
  { path: '/', component: IndexView }
]

const router = createRouter({
  history: createMemoryHistory(),
  routes,
})

export default router
