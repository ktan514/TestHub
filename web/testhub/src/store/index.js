import { createStore } from 'vuex'
import auth from './auth'
import setting from './setting'

const store = createStore({
  modules: {
    auth,
    setting,
  },
})

export default store
