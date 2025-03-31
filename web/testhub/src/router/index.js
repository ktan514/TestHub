import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/LoginView.vue'
import TestItem from '@/views/HomeView.vue'
import TestItemEdit from '@/components/TestItemEdit.vue'
import TestItemList from '@/components/TestItemList.vue'
// import Setting from '@/components/Setting.vue'

const routes = [
  { path: '/', component: Login },
  { path: '/login', component: Login },
  // { path: '/home', component: Home },
  {
    path: '/testitem',
    component: TestItem,
    children: [
      {
        path: 'list',
        component: TestItemList,
      },
      {
        path: 'edit',
        component: TestItemEdit,
      },
    ],
  },
  // { path: '/setting', component: Setting },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
