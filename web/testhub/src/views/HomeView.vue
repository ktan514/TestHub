<template>
  <header>
    <h1>🧪 TestHub</h1>
    <button @click="onLogout">ログアウト</button>
  </header>

  <main>
    <router-view />
  </main>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const authStore = useAuthStore()

const onLogout = () => {
  if (confirm('ログアウトしますか?')) {
    try {
      authStore.logout()
      router.push('/login') // ログアウト後にログイン画面へ遷移
    } catch (error) {
      console.error('ログアウトに失敗しました:', error)
    }
  }
}
</script>

<style scoped>
header {
  background-color: #3eaf7c;
  color: white;
  padding: 1rem;
  text-align: center;
}

footer {
  margin-top: 2rem;
  text-align: center;
  color: #888;
  font-size: 0.9rem;
}
</style>
