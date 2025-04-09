import { defineStore } from 'pinia'

export const useTestItemStore = defineStore('testitem', {
  state: () => ({
    testId: null,
    categoryId: null,
  }),
  getters: {
    getTestId: (state) => state.testId ?? '',
    getCategoryId: (state) => state.categoryId ?? '',
  },
  actions: {
    setTestId(testId) {
      this.testId = testId
    },
    setCategoryId(categoryId) {
      this.categoryId = categoryId
    },
  },
})
