<template>
  <div id="app">
    <main>
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, provide } from 'vue'

const user = ref(null)
const isAdmin = ref(false)

onMounted(() => {
  // 子コンポーネントに提供
  provide('user', user)
  provide('setUser', setUser)
})

// ユーザー情報を設定する関数
const setUser = (newUser) => {
  user.value = newUser
  isAdmin.value = newUser?.role === 1 // role=1なら管理者
}
</script>

<style scoped>
#app {
  font-family: 'Segoe UI', sans-serif;
  padding: 1rem;
  margin: 0 auto;
}
</style>
