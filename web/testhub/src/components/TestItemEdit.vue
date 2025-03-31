<template>
  <div class="max-w-3xl mx-auto p-4 form-container">
    <h2 class="text-xl font-semibold form-group">試験アイテム登録</h2>
    <form @submit.prevent="submitTestItem">

      <div>
        <div class="form-group">
          <label class="block text-gray-700">Test ID</label>
          <input v-model="testItem.id" type="text" class="w-full border rounded px-3 py-2" required :readonly="modify">
        </div>

        <div class="form-group">
          <label class="block text-gray-700">Test Category</label>
          <select v-model="testItem.category" class="w-full border rounded px-3 py-2" required>
            <option value="" disabled>選択してください</option>
            <option v-for="category in categories" :key="category.id" :value="category">
              {{ category.name }}
            </option>
          </select>
        </div>

        <div class="form-group">
          <label class="block text-gray-700">Situation</label>
          <textarea v-model="testItem.situation" class="w-full border rounded px-3 py-2"></textarea>
        </div>

        <div class="form-group">
          <label class="block text-gray-700">Scheduled Tester</label>
          <select v-model="testItem.scheduledTester" class="w-full border rounded px-3 py-2">
            <option value="" disabled>選択してください</option>
            <option v-for="tester in scheduledTesters" :key="tester.id" :value="tester">
              {{ tester.name }}
            </option>
          </select>
        </div>

        <div style="border-top: 1px black solid; padding-bottom: 10px;"></div>
      </div>

      <div class="form-group">
        <label class="block text-gray-700">What's Purpose</label>
        <textarea v-model="testItem.purpose" type="text" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="form-group">
        <label class="block text-gray-700">Operation</label>
        <textarea v-model="testItem.operation" type="text" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="form-group">
        <label class="block text-gray-700">Expected result</label>
        <textarea v-model="testItem.expResult" class="w-full border rounded px-3 py-2"></textarea>
      </div>

      <div class="form-group">
        <label class="block text-gray-700">Topic</label>
        <textarea v-model="testItem.topic" type="text" class="w-full border rounded px-3 py-2"></textarea>
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
import router from '@/router';
import { useRoute, onBeforeRouteUpdate } from 'vue-router';
import { ref, onMounted, computed } from 'vue';
import axios from 'axios';
import { useStore } from 'vuex';

const store = useStore();
const user = computed(() => store.state.auth.user);

const route = useRoute();
let modify = false;
const selectedIds = [];
let currentIndex = 0;
let selectedCategory = 0;

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
  scheduledTester: {
    id: 0,
    name: '',
    role: 0,
    admin: 0,
  },
});

const scheduledTesters = ref([]);
const categories = ref([]);

// 画面表示時にテスター一覧を読み込む想定
onMounted(async () => {
  if (!user.value) {
    // 未ログインなら拒否する
    alert('ログアウトされています。\nログインしてください。');
    // ログイン画面に飛ばす
    router.push({ path: '/login' });
    return;
  }

  // Test Categoryの選択アイテム
  await loadTestCategorySelectBox();
  // Scheduled Testerの選択アイテム
  await loadScheduledTesterSelectBox();
  // クエリパラメータを読み込む
  await loadQueryParam();
  // 初期データを読み込む
  await loadTestItem();
});

// フォーム送信時の処理
const submitTestItem = async () => {
  console.log('登録する試験アイテム:', testItem.value);
  // ここでAPI呼び出しやバリデーション処理を実行
  try {
    const result = await axios.post('/api/testitem', testItem.value)
      .then(res => res.data);
    alert('登録が完了しました。');
    console.log(result);

    // 保存処理が完了した後
    nextDispTestItem();
  } catch (error) {
    alert('登録にエラーが発生しました。');
    console.error(error)
  }
};

// キャンセルボタンの処理
const handleCancel = () => {
  // 戻るボタンが押された時の処理
  nextDispTestItem();
};

// データの取得処理
const loadTestItem = async () => {
  if (selectedIds[currentIndex]) {
    // 現在のIDに基づいてデータを取得
    const params = { id: selectedIds[currentIndex] };
    testItem.value = await axios.get('/api/testitem', { params })
      .then(res => res.data);
    console.log(testItem.value);
    // 編集モード
    modify = true;
  } else {
    // 初期値を設定
    testItem.value.category = categories.value.find(value => value.id === selectedCategory);
    // 追加モード
    modify = false;
  }
};

// 試験カテゴリー選択ボックスの要素を取得する
const loadTestCategorySelectBox = async () => {
  categories.value = await axios.get('/api/categories')
    .then(res => res.data);
  console.log(categories.value);
};

// 予定テスター選択ボックスの要素を取得する
const loadScheduledTesterSelectBox = async () => {
  scheduledTesters.value = await axios.get('/api/users')
    .then(res => res.data);
  console.log(scheduledTesters.value);
};

// クエリパラメータからデータを取得する
const loadQueryParam = async () => {
  // クリアする
  selectedIds.length = 0;
  currentIndex = 0;
  // クエリパラメータから取得する

  // 試験ID
  if (route.query.selectedIds && Array.isArray(route.query.selectedIds)) {
    for (const id of route.query.selectedIds) {
      selectedIds.push(id);
    }
  }

  // カテゴリー
  if (route.query.category) {
    selectedCategory = Number(route.query.category);
  }
};

// ページ遷移
const nextDispTestItem = () => {
  if (modify) {
    currentIndex += 1;
    if (currentIndex < selectedIds.length) {
      // 次のデータへ遷移
      loadTestItem();
    } else {
      alert("全ての編集が完了しました。");
      // 最後のデータが終わったら一覧画面に戻る
      router.push({ path: '/testitem/list' });
    }
  } else {
    // 追加モードの場合は一覧画面に戻る
    router.push({ path: '/testitem/list' });
  }
};
</script>

<style scoped>
textarea {
  resize: vertical;
}
</style>
