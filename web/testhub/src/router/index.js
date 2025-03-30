import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/components/Login.vue'
import TestItemEdit from '@/components/TestItemEdit.vue'
import TestItemList from '@/components/TestItemList.vue'
import Setting from '@/components/Setting.vue'

const routes = [
  { path: '/', component: Login },
  { path: '/testItemEdit', component: TestItemEdit },
  { path: '/testItemList', component: TestItemList },
  { path: '/setting', component: Setting },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
