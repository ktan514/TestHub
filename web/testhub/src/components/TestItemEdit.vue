<template>
  <div class="max-w-3xl mx-auto p-4">
    <h2 class="text-xl font-semibold mb-4">試験アイテム登録</h2>
    <form @submit.prevent="submitTestItem">
      <div class="mb-4">
        <label class="block text-gray-700">試験分類</label>
        <select v-model="testItem.category" class="w-full border rounded px-3 py-2">
          <option value="" disabled>選択してください</option>
          <option v-for="category in categories" :key="category.id" :value="category">
            {{ category.name }}
          </option>
        </select>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Function Name</label>
        <input v-model="testItem.subject" type="text" class="w-full border rounded px-3 py-2">
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Keyword of Function</label>
        <textarea v-model="testItem.perspective" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Transport Unit</label>
        <textarea v-model="testItem.conditions" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Location</label>
        <textarea v-model="testItem.steps" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Situation</label>
        <textarea v-model="testItem.expectedResult" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Test ID</label>
        <textarea v-model="testItem.expectedResult" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Test ID(Old)</label>
        <textarea v-model="testItem.expectedResult" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Transport pattern & Procedure</label>
        <textarea v-model="testItem.expectedResult" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Expected result</label>
        <textarea v-model="testItem.expectedResult" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="mb-4">
        <label class="block text-gray-700">Scheduled Tester</label>
        <select v-model="testItem.scheduledTester" class="w-full border rounded px-3 py-2">
          <option value="" disabled>選択してください</option>
          <option v-for="tester in scheduledTesters" :key="tester.id" :value="tester">
            {{ tester.name }}
          </option>
        </select>
      </div>

      <div class="text-right">
        <button type="submit" class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded">
          登録
        </button>
        <button type="button" @click="handleCancel" class="bg-gray-500 hover:bg-gray-600 text-white px-4 py-2 rounded ml-2">
          戻る
        </button>
      </div>
    </form>
  </div>
</template>

<script setup>
import axios from 'axios';
import { ref, onMounted } from 'vue';

const testItem = ref({
  category: {
    id: 0,
    name: '',
  },
  funcName: '',
  keyword: '',
  transUnit: '',
  location: '',
  testId: '',
  testIdOld: '',
  transPtnProd: '',
  expResult: '',
  topic: '',
  plan: false,
  sim: '',
  scheduledTester: {
    id: 0,
    name: '',
  }
});

const scheduledTesters = ref([]);
const categories = ref([]);

// 画面表示時にテスター一覧を読み込む想定
onMounted(() => {
  // 仮データをセット（実際にはAPI経由で取得）
  // 予定担当者の選択アイテム
  scheduledTesters.value = [
    { id: 1, name: 'Tester1' },
    { id: 2, name: 'Tester2' },
    { id: 3, name: 'Tester3' },
    { id: 4, name: 'Tester4' },
    { id: 5, name: 'Tester5' },
  ];
  // 試験分類の選択アイテム
  categories.value = [
    { id: 1, name: 'New Items' },
    { id: 2, name: 'Normal Items' },
    { id: 3, name: 'Function A' },
    { id: 4, name: 'Function B' },
    { id: 5, name: 'GUI' },
  ];
});

// フォーム送信時の処理
const submitTestItem = async () => {
  console.log('登録する試験アイテム:', testItem.value);
  // ここでAPI呼び出しやバリデーション処理を実行
  try {
    const response = await axios.post('/api/testitems', testItem.value);
    alert('登録が完了しました。');
    console.log(response.data);
    window.history.back(); // 前の画面に戻る
  } catch (error) {
    alert('登録にエラーが発生しました。');
    console.error(error)
  }
};

// キャンセルボタンの処理
const handleCancel = () => {
  // 戻るボタンが押された時の処理
  // 例：フォームのクリアや、前のページへの遷移
  window.history.back(); // 前の画面に戻る
};
</script>

<style scoped>
textarea {
  resize: vertical;
}
</style>
