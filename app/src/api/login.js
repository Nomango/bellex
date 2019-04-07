import * as API from './'
export default {
  getLogin: params => {
    return API.POST('user/login', params)
  },
  logout: params => {
    return API.POST('user/logout', params)
  }
}
