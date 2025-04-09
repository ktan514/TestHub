// stores/auth.js
import { defineStore } from 'pinia'
import axios from 'axios'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.user,
    userName: (state) => (state.user ? state.user.name : ''),
  },
  actions: {
    async login(user) {
      try {
        const response = await axios.post('/api/login', user, {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
          },
        })
        this.user = response.data.user
      } catch (error) {
        throw new Error(error.response ? 'ログインに失敗しました' : 'ネットワークエラー')
      }
    },
    logout() {
      this.user = null
    },
  },
})
