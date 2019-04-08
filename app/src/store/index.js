import Vue from 'vue'
import Vuex from 'vuex'
import userAjax from '@/api/user'
Vue.use(Vuex)
const state = {
  isShowLoading: false,
  isCollapse: false,
  ajaxLoading: false,
  roles: 0
}
const mutations = {
  changLoading: (state, isShow) => {
    state.isShowLoading = isShow
  },
  changeCollapse: (state) => {
    state.isCollapse = !state.isCollapse
  },
  SET_ROLE: (state, role) => {
    state.roles = role
  }
}
const actions = {
  userStatus ({
    commit
  }) {
    return new Promise((resolve, reject) => {
      userAjax.getUserStatus()
        .then(response => {
          const {
            data
          } = response
          commit('SET_ROLE', data.user.role)
          resolve(data.user.role)
        }).catch(error => {
          reject(error)
        })
    })
  }
}

export default new Vuex.Store({
  actions,
  state,
  mutations
})
