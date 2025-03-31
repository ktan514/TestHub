const state = {
  theme: 'light',
}

const mutations = {
  setTheme(state, theme) {
    state.theme = theme
  },
}

const actions = {
  updateTheme({ commit }, theme) {
    commit('setTheme', theme)
  },
}

const getters = {
  theme: (state) => state.theme,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}
