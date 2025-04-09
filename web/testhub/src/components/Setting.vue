<template>
  <p v-if="user">ログインユーザー: {{ user.name }}</p>
  <div class="settings-container">
    <h1>設定</h1>
    <ul>
      <li>
        <router-link to="/user-registration">ユーザー(テスター)登録</router-link>
      </li>
      <li>
        <router-link to="/test-category">試験カテゴリー登録</router-link>
      </li>
      <li>
        <router-link to="/user-group">ユーザーのグループ登録</router-link>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted, inject } from 'vue'
import router from '@/router'

const user = inject('user')

onMounted(async () => {
  if (!user.value) {
    // 未ログインなら拒否する
    alert('ログアウトされています。\nログインしてください。')
    // ログイン画面に飛ばす
    router.push({ path: '/login' })
    return
  }
})
</script>

<style scoped>
.settings-container {
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  text-align: center;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  margin: 10px 0;
}

router-link {
  text-decoration: none;
  color: #007bff;
  font-weight: bold;
}

router-link:hover {
  text-decoration: underline;
}
</style>
