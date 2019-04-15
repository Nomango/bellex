import Vue from 'vue'
import Vuex from 'vuex'
import userAjax from '@/api/user'
Vue.use(Vuex)
const state = {
  isShowLoading: false,
  isCollapse: false,
  ajaxLoading: false,
  roles: 0,
  nickName: '',
  userInfo: null
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
  },
  SET_NICKNAME: (state, name) => {
    state.nickName = name
  },
  SET_USERINFO: (state, user) => {
    state.userInfo = user
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
          commit('SET_NICKNAME', data.user.nickname)
          commit('SET_USERINFO', data.user)
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
