<template>
  <div class="signup-container">
    <h1>サインアップ</h1>
    <form @submit.prevent="handleSignUp">
      <div>
        <label for="username">ユーザー名</label>
        <input type="text" id="username" v-model="username" required />
      </div>
      <div>
        <label for="password">パスワード</label>
        <input type="password" id="password" v-model="password" required />
      </div>
      <div>
        <label for="confirmPassword">パスワード確認</label>
        <input type="password" id="confirmPassword" v-model="confirmPassword" required />
      </div>
      <button type="submit" :disabled="loading">サインアップ</button>
      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

const username = ref('');
const password = ref('');
const confirmPassword = ref('');
const loading = ref(false);
const errorMessage = ref('');
const router = useRouter();

const handleSignUp = async () => {
  if (password.value !== confirmPassword.value) {
    errorMessage.value = 'パスワードが一致しません';
    return;
  }

  loading.value = true;
  errorMessage.value = '';

  try {
    await axios.post('/api/login', {
      username: username.value,
      password: password.value
    }, {
      withCredentials: true,
      headers: {
        'Content-Type': 'application/json'
      }
    });

    router.push('/login'); // 登録後にログイン画面へ遷移
  } catch (error) {
    errorMessage.value = error.response?.data?.message || 'サインアップに失敗しました';
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.signup-container {
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
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #218838;
}

.error {
  color: red;
}
</style>