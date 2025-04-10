<template>
  <TestItemForm :formTitle="'試験アイテム登録'" :modify="false" />
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'
import TestItemForm from '@/components/TestItemForm.vue'
import { useAuthStore } from '@/store/auth'
import router from '@/router'
import { useRoute } from 'vue-router'

const authStore = useAuthStore()
const user = computed(() => authStore.user)
const route = useRoute()

onMounted(async () => {
  if (!user.value) {
    // 未ログインなら拒否する
    alert('ログアウトされています。\nログインしてください。')
    // ログイン画面に飛ばす
    router.push({ path: '/login' })
    return
  }

  // クエリパラメータを読み込む
  await loadQueryParam()
})

// クエリパラメータからデータを取得する
const loadQueryParam = async () => {
  // クエリパラメータから取得する

  // カテゴリー
  if (route.query.category) {
    Number(route.query.category)
  }
}
</script>
