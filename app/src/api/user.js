import * as API from './'
export default {
  getUserList: params => {
    return API.GET('user/all', params)
  },
  getUserStatus: params => {
    return API.GET('user/status', params)
  },
  addUser: params => {
    return API.POST('user/new', params)
  },
  putUser: params => {
    return API.PUT(`user/${params.institution_id}`, params)
  },
  getInstitution: params => {
    return API.GET('institution/all', params)
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
