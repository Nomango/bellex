import * as API from './'
export default {
  getControllers: params => {
    return API.GET('static/mock/controllers.json', params)
  },
  delControllers: params => {
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
