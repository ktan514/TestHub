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
import { ref } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

// 状態を定義
const username = ref('');
const password = ref('');
const loading = ref(false);
const errorMessage = ref('');

// Vue Router のインスタンスを取得
const router = useRouter();

// ログイン処理
const handleLogin = async () => {
  loading.value = true;
  errorMessage.value = '';

  try {
    const response = await axios.post('/api/login', {
      username: username.value,
      password: password.value
    }, {
      withCredentials: true, // クッキーを送信する
      headers: {
        'Content-Type': 'application/json'
      }
    });

    // ログイン成功時の処理
    console.log('ログイン成功:', response.data.message);
    // localStorageに保存
    const user = response.data.user;
    localStorage.setItem('user', JSON.stringify(user));

    // ログイン成功後、一覧画面へ遷移
    router.push('/testItemList');
  } catch (error) {
    if (error.response) {
      errorMessage.value = error.response.data.message || 'ログインに失敗しました';
    } else {
      errorMessage.value = 'ネットワークエラー';
    }
  } finally {
    loading.value = false;
  }
};
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
