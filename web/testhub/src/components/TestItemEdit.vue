<template>
  <TestItemForm :formTitle="'試験アイテム編集'" :modify="true" />
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

const selectedIds = []
let currentIndex = 0

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
  // 初期データを読み込む
  await loadTestItem()
})

// ページ遷移
const nextDispTestItem = () => {
  if (modify) {
    currentIndex += 1
    if (currentIndex < selectedIds.length) {
      // 次のデータへ遷移
      loadTestItem()
    } else {
      alert('全ての編集が完了しました。')
      // 最後のデータが終わったら一覧画面に戻る
      router.push({ path: '/testitem/list' })
    }
  } else {
    // 追加モードの場合は一覧画面に戻る
    router.push({ path: '/testitem/list' })
  }
}

// クエリパラメータからデータを取得する
const loadQueryParam = async () => {
  // クリアする
  selectedIds.length = 0
  currentIndex = 0
  // クエリパラメータから取得する

  // 試験ID
  if (route.query.selectedIds && Array.isArray(route.query.selectedIds)) {
    for (const id of route.query.selectedIds) {
      selectedIds.push(id)
    }
  }

  // カテゴリー
  if (route.query.category) {
    Number(route.query.category)
  }
}

// データの取得処理
const loadTestItem = async () => {
  if (selectedIds[currentIndex]) {
    // 現在のIDに基づいてデータを取得
    const params = { id: selectedIds[currentIndex] }
    testItem.value = await axios.get('/api/testitem', { params }).then((res) => res.data)
    console.log(testItem.value)
  } else {
    // 初期値を設定
    // testItem.value.category = categories.value.find((value) => value.id === selectedCategory)
  }
}

defineProps({
  nextDispTestItem: Function, // ここで受け取る
})
</script>
