import { defineStore } from 'pinia'

export const useSettingStore = defineStore('setting', {
  state: () => ({
    someSetting: null, // 必要に応じて初期状態を定義
  }),
  getters: {
    getSetting: (state) => state.someSetting,
  },
  actions: {
    setSetting(settingValue) {
      this.someSetting = settingValue
    },
  },
})
