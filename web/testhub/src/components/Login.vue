<template>
  <div class="login-container">
    <h1>ログイン</h1>
    <form @submit.prevent="handleLogin">
      <div>
        <label for="username">ユーザー名</label>
        <input type="text" id="username" v-model="username" required />
      </div>
      <div>
        <label for="password">パスワード</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <button type="submit" :disabled="loading">ログイン</button>
      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import router from '@/router'
import { useRoute } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const authStore = useAuthStore()
const route = useRoute()

const username = ref('')
const password = ref('')
const loading = ref(false)
const errorMessage = ref('')

// ログイン処理
const handleLogin = async () => {
  errorMessage.value = ''

  try {
    loading.value = true
    const user = { username: username.value, password: password.value }

    await authStore.login(user) // Pinia store の login を呼び出す

    // ログイン成功後にリダイレクト先があればそこに、なければデフォルトページへ
    const redirectPath = route.query.redirect || '/testitem/list'
    router.push(redirectPath)
  } catch (error) {
    errorMessage.value = error.message
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  max-width: 300px;
  margin: 50px auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  text-align: center;
}

input {
  width: 100%;
  padding: 8px;
  margin: 5px 0;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

.error {
  color: red;
}
</style>
