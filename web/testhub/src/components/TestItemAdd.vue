<template>
  <TestItemForm
    :testItem="testItem"
    :categories="categories"
    :formTitle="'試験アイテム追加'"
    @submitTestItem="submitTestItem"
  />
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import TestItemForm from '@/components/TestItemForm.vue'

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

onMounted(async () => {
  await loadTestCategorySelectBox()
})

const loadTestCategorySelectBox = async () => {
  categories.value = await axios.get('/api/categories').then((res) => res.data)
}

const submitTestItem = async () => {
  try {
    await axios.post('/api/testitem', testItem.value)
    alert('登録が完了しました')
  } catch (error) {
    alert('登録に失敗しました')
  }
}
</script>
