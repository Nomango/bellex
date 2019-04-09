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
    return API.PUT(`user/${params.id}`, params)
  },
  getInstitution: params => {
    return API.GET('institution/all', params)
  },
  delUser: params => {
    return API.DELETE(`user/${params.id}`)
  }
}
