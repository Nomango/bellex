import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
const state = {
  isShowLoading: false,
  isCollapse: false,
  ajaxLoading: false
}
const mutations = {
  changLoading: function (state, isShow) {
    state.isShowLoading = isShow
  },
  changeCollapse: function (state) {
    state.isCollapse = !state.isCollapse
  }
}
export default new Vuex.Store({
  state,
  mutations
})
