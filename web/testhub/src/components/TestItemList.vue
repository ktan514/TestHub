<template>
  <p v-if="user">ログインユーザー: {{ user.name }}</p>
  <div class="max-w-4xl mx-auto p-4">
    <h2 class="text-xl font-semibold mb-4">試験項目一覧</h2>

    <!-- 分類選択メニュー -->
    <div class="mb-4">
      <label class="block text-gray-700">分類を選択</label>
      <select
        v-model="selectedCategory"
        @change="fetchTestItems"
        class="w-full border rounded px-3 py-2"
      >
        <option value="" disabled>選択してください</option>
        <option v-for="category in categories" :key="category.id" :value="category.id">
          {{ category.name }}
        </option>
      </select>
    </div>

    <button
      type="button"
      @click="handleModify"
      class="bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded ml-2"
    >
      編集
    </button>
    <button
      type="button"
      @click="handleRemove"
      class="bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded ml-2"
    >
      削除
    </button>

    <!-- 試験項目一覧表示 -->
    <table class="w-full border-collapse border-gray-400">
      <thead>
        <tr class="bg-gray-200">
          <th class="" hidden></th>
          <th class="border border-gray-400 px-4 py-2">Test ID</th>
          <th class="border border-gray-400 px-4 py-2">Situation</th>
          <th class="border border-gray-400 px-4 py-2">Topic</th>
          <th class="border border-gray-400 px-4 py-2">Scheduled Tester</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item in testItems"
          :key="item.id"
          :class="{ 'bg-blue-200': selectedRows.includes(item.id), 'cursor-pointer': true }"
          @click="toggleRowSelection(item.id)"
        >
          <td class="" hidden>
            <input type="checkbox" :value="item.id" v-model="selectedRows" />
          </td>
          <td class="border border-gray-400 px-4 py-2">{{ item.id }}</td>
          <td class="border border-gray-400 px-4 py-2">{{ item.situation }}</td>
          <td class="border border-gray-400 px-4 py-2">{{ item.topic }}</td>
          <td class="border border-gray-400 px-4 py-2">{{ item.scheduledTester.name }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import axios from 'axios'
import router from '@/router'
import { useAuthStore } from '@/store/auth'
import { useTestItemStore } from '@/store/testitem'

const authStore = useAuthStore()
const testItemStore = useTestItemStore()

const categories = ref([])
const testItems = ref([])
const selectedRows = ref([]) // 選択された行のIDを格納

const user = computed(() => authStore.user)

const selectedCategory = computed({
  get: () => testItemStore.categoryId,
  set: (value) => testItemStore.setCategoryId(value),
})

onMounted(async () => {
  if (!user.value) {
    // 未ログインなら拒否する
    alert('ログアウトされています。\nログインしてください。')
    // ログイン画面に飛ばす
    router.push({ path: '/login' })
    return
  }
  // Test Categoryの選択アイテム
  await loadTestCategorySelectBox()
  await fetchTestItems()
})

// 分類に応じた試験項目を取得する
const fetchTestItems = async () => {
  if (selectedCategory.value) {
    const params = { categoryId: selectedCategory.value }
    testItems.value = await axios.get('/api/testitems', { params }).then((res) => res.data)
  } else {
    testItems.value = []
  }
  // 選択をクリアする
  selectedRows.value.splice(0)
}

// 行くリック時に選択状態を切り替える
const toggleRowSelection = (id) => {
  const index = selectedRows.value.indexOf(id)
  if (index === -1) {
    selectedRows.value.push(id)
  } else {
    selectedRows.value.splice(index, 1)
  }
}

// 編集ボタン押下時の処理
const handleModify = () => {
  const category = selectedCategory.value
  const selectedIds = selectedRows.value
  router.push({ path: '/testitem/edit', query: { category, selectedIds } })
}

// キャンセルボタン押下時の処理
const handleRemove = async () => {
  const selectedIds = selectedRows.value

  if (selectedIds.length === 0) {
    alert('削除する項目を選択してください。')
    return
  }

  // 確認ダイアログを一度だけ表示
  if (!confirm(`選択した ${selectedIds.length} 件を削除しますか？`)) {
    return
  }

  try {
    // バルク削除
    const params = { ids: selectedIds }
    await axios.delete('/api/testitems', { params })

    alert('削除しました。')

    // 削除後、一覧を更新
    fetchTestItems() // ここで一覧を再取得する関数を呼び出す
  } catch (error) {
    console.error('削除エラー:', error)
    alert('削除に失敗しました。')
  }
}

// 試験カテゴリー選択ボックスの要素を取得する
const loadTestCategorySelectBox = async () => {
  categories.value = await axios.get('/api/categories').then((res) => res.data)
  console.log(categories.value)
}
</script>

<style scoped>
table {
  border-collapse: collapse;
}

table td {
  padding: 5px;
}

table th {
  padding: 5px;
}

.border-gray-400 {
  border: 1px solid gray;
  min-width: 100px;
}

.bg-blue-200 {
  background-color: #3dae7b80;
}
</style>
