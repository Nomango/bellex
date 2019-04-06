import * as API from './'
export default {
  getControllers: params => {
    return API.GET('static/mock/controllers.json', params)
  },
  getLogin: params => {
    return API.GET('user/status', params)
  }
}
