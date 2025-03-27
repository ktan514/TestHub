import { createRouter, createWebHistory } from 'vue-router'
import Home from '../components/HelloWorld.vue'
import TestItemEdit from '../components/TestItemEdit.vue'
import TestItemList from '../components/TestItemList.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/testItemEdit', component: TestItemEdit },
  { path: '/testItemList', component: TestItemList },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
