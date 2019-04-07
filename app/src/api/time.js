import * as API from './'
export default {
  getTimeList: params => {
    return API.GET('static/mock/time.json', params)
  },
  delTimeList: params => {
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
