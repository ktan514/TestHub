import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/LoginView.vue'
import TestItem from '@/views/HomeView.vue'
import TestItemList from '@/components/TestItemList.vue'
import TestItemAdd from '@/components/TestItemAdd.vue'
import TestItemEdit from '@/components/TestItemEdit.vue'
// import Setting from '@/components/Setting.vue'
import { useAuthStore } from '@/store/auth'

const routes = [
  { path: '/', component: Login },
  { path: '/login', component: Login },
  {
    path: '/testitem',
    component: TestItem,
    meta: { requiresAuth: true },
    children: [
      {
        path: 'list',
        component: TestItemList,
        meta: { requiresAuth: true },
      },
      {
        path: 'edit',
        component: TestItemEdit,
        meta: { requiresAuth: true },
      },
      {
        path: 'add',
        component: TestItemAdd,
        meta: { requiresAuth: true },
      },
    ],
  },
  // { path: '/setting', component: Setting },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// グローバルナビゲーションガードでPiniaストアを都度呼び出す
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  const isAuthenticated = !!authStore.user

  if (to.matched.some((record) => record.meta.requiresAuth) && !isAuthenticated) {
    alert('ログアウトされています。ログインしてください')
    next({ path: '/login', query: { redirect: to.fullPath } })
  } else {
    next()
  }
})

export default router
