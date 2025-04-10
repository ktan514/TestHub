<!-- components/TestItemForm.vue -->
<template>
  <div class="max-w-3xl mx-auto p-4 form-container">
    <h2 class="text-xl font-semibold form-group">{{ formTitle }}</h2>
    <form @submit.prevent="submitTestItem">
      <!-- Test ID -->
      <div class="form-group">
        <label class="block text-gray-700">Test ID</label>
        <input
          v-model="testItem.id"
          type="text"
          class="w-full border rounded px-3 py-2"
          required
          :readonly="modify"
        />
      </div>

      <!-- Test Category -->
      <div class="form-group">
        <label class="block text-gray-700">Test Category</label>
        <select v-model="testItem.category" class="w-full border rounded px-3 py-2" required>
          <option value="" disabled>選択してください</option>
          <option v-for="category in categories" :key="category.id" :value="category">
            {{ category.name }}
          </option>
        </select>
      </div>

      <!-- Scheduled Tester -->
      <div class="form-group">
        <label class="block text-gray-700">Scheduled Tester</label>
        <select v-model="testItem.scheduledTester" class="w-full border rounded px-3 py-2">
          <option value="" disabled>選択してください</option>
          <option v-for="tester in scheduledTesters" :key="tester.id" :value="tester">
            {{ tester.name }}
          </option>
        </select>
      </div>

      <!-- Other fields (Situation, Purpose, etc.) -->
      <div class="form-group" v-for="(field, index) in fields" :key="index">
        <label class="block text-gray-700">{{ field.label }}</label>
        <textarea
          v-model="testItem[field.model]"
          class="w-full border rounded px-3 py-2"
        ></textarea>
      </div>

      <div class="text-right">
        <button type="submit" class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded">
          決定
        </button>
        <button
          type="button"
          @click="handleCancel"
          class="bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded ml-2"
        >
          戻る
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const testItem = ref({
  id: '',
  situation: '',
  purpose: '',
  operation: '',
  expResult: '',
  topic: '',
  category: {
    id: 0,
    name: '',
  },
})

const categories = ref([])
const scheduledTesters = ref([])

const fields = [
  { label: 'Situation', model: 'situation' },
  { label: "What's Purpose", model: 'purpose' },
  { label: 'Operation', model: 'operation' },
  { label: 'Expected result', model: 'expResult' },
  { label: 'Topic', model: 'topic' },
]

const formTitle = '試験アイテム登録'

const modify = ref(false)

onMounted(async () => {
  // Test Categoryの選択アイテム
  await loadTestCategorySelectBox()
  // Scheduled Testerの選択アイテム
  await loadScheduledTesterSelectBox()
})

const handleCancel = () => {
  // 戻る処理をここに記述
}

const loadTestCategorySelectBox = async () => {
  categories.value = await axios.get('/api/categories').then((res) => res.data)
}

// 予定テスター選択ボックスの要素を取得する
const loadScheduledTesterSelectBox = async () => {
  scheduledTesters.value = await axios.get('/api/users').then((res) => res.data)
  console.log(scheduledTesters.value)
}

// フォーム送信時の処理
const submitTestItem = async () => {
  console.log('登録する試験アイテム:', testItem.value)
  // ここでAPI呼び出しやバリデーション処理を実行
  try {
    const result = await axios.post('/api/testitem', testItem.value).then((res) => res.data)
    alert('登録が完了しました。')
    console.log(result)

    // 正常時のコールバック関数
    if (props.nextDispTestItem) {
      props.nextDispTestItem()
    }
  } catch (error) {
    alert('登録にエラーが発生しました。')
    console.error(error)
  }
}

const props = defineProps(['onNext'])
</script>

<style scoped>
textarea {
  resize: vertical;
}
</style>
