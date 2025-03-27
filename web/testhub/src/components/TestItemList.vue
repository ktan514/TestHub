<template>
  <div class="max-w-4xl mx-auto p-4">
    <h2 class="text-xl font-semibold mb-4">試験項目一覧</h2>

    <!-- 分類選択メニュー -->
    <div class="mb-4">
      <label class="block text-gray-700">分類を選択</label>
      <select v-model="selectedCategory" @change="fetchTestItems" class="w-full border rounded px-3 py-2">
        <option value="" disabled>選択してください</option>
        <option v-for="category in categories" :key="category.id" :value="category.id">
          {{ category.name }}
        </option>
      </select>
    </div>

    <!-- 試験項目一覧表示 -->
    <table class="w-full border-collapse">
      <thead>
        <tr class="bg-gray-200">
          <th class="border px-4 py-2">試験対象</th>
          <th class="border px-4 py-2">試験観点</th>
          <th class="border px-4 py-2">実施予定者</th>
          <th class="border px-4 py-2">期待値</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in testItems" :key="item.id">
          <td class="border px-4 py-2">{{ item.subject }}</td>
          <td class="border px-4 py-2">{{ item.perspective }}</td>
          <td class="border px-4 py-2">{{ item.scheduledTester }}</td>
          <td class="border px-4 py-2">{{ item.expectedResult }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const categories = ref([]);
const selectedCategory = ref('');
const testItems = ref([]);

// カテゴリー一覧を取得
onMounted(async () => {
  categories.value = await axios.get('/api/categories').then(res => res.data);
});

// 分類に応じた試験項目を取得する関数
const fetchTestItems = async () => {
  if (selectedCategory.value) {
    testItems.value = await axios.get(`/api/testitems?category=${selectedCategory.value}`)
      .then(res => res.data);
  } else {
    testItems.value = [];
  }
};
</script>

<style scoped>
table th,
table td {
  text-align: left;
}
</style>