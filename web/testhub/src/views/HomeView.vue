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
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
const router = useRouter()
const store = useStore()

const onLogout = async () => {
  if (confirm('ログアウトしますか?')) {
    try {
      await store.dispatch('auth/logout')
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
