import axios from 'axios'

const state = {
  user: null,
}

const mutations = {
  setUser(state, user) {
    state.user = user
  },
  logout(state) {
    state.user = null
  },
}

const actions = {
  async login({ commit }, user) {
    let response = {}

    // ユーザー認証を行う
    try {
      response = await axios.post('/api/login', user, {
        withCredentials: true, // クッキーを送信する
        headers: {
          'Content-Type': 'application/json',
        },
      })

      console.log('ログイン成功')
    } catch (error) {
      console.log('ログイン失敗')

      if (error.response) {
        throw new Error('ログインに失敗しました')
      } else {
        throw new Error('ネットワークエラー')
      }
    }

    // localStorageに保存
    commit('setUser', { user: response.data.user })
  },
  logout({ commit }) {
    commit('logout')
  },
}

const getters = {
  isAuthenticated: (state) => !!state.user,
  userName: (state) => (state.user ? state.user.name : ''),
}

export default {
  namespaced: true, // モジュール名を有効にする
  state,
  mutations,
  actions,
  getters,
}
