import * as API from './'
export default {
  getLogin: params => {
    return API.POST('user/login', params)
  },
  logout: params => {
    return API.POST('user/logout', params)
  },
  resetPsd: params => {
    return API.REPOST('user/password', params)
  },
  getProfileInfo: params => {
    return API.GET('user/status', params)
  },
  updateProfileInfo: params => {
    return API.POST('user/profile', params)
  }
}
