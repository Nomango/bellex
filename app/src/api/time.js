import * as API from './'
export default {
  getTimeList: params => {
    return API.GET('schedule/all', params)
  },
  putTimeList: params => {
    return API.PUT(`schedule/${params.id}`, params)
  },
  addTimeList: params => {
    return API.POST('schedule/new', params)
  },
  delTimeList: params => {
    return API.DELETE(`schedule/${params.id}`)
  }
}
