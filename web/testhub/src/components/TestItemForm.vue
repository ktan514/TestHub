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
import { ref } from 'vue'

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

const fields = [
  { label: 'Situation', model: 'situation' },
  { label: "What's Purpose", model: 'purpose' },
  { label: 'Operation', model: 'operation' },
  { label: 'Expected result', model: 'expResult' },
  { label: 'Topic', model: 'topic' },
]

const formTitle = '試験アイテム登録'

const modify = ref(false)

const submitTestItem = () => {
  console.log('登録する試験アイテム:', testItem.value)
  // 送信処理をここに記述
}

const handleCancel = () => {
  // 戻る処理をここに記述
}
</script>

<style scoped>
textarea {
  resize: vertical;
}
</style>
