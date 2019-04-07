import * as API from './'
export default {
  getUserList: params => {
    return API.GET('static/mock/user.json', params)
  },
  delUser: params => {
    return new Promise(function (resolve, reject) {
      if (params !== undefined) {
        resolve({
          code: 0
        })
      }
    })
    // return API.GET('static/mock/controllers.json', params)
  }
}
