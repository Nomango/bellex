import * as API from '.'
export default {
  getLogin: params => {
    return API.POST('user/login', params)
  }
}
