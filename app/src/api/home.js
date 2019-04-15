import * as API from './'
export default {
  getControllers: params => {
    return API.GET('mechine/all', params)
  },
  addControllers: params => {
    return API.POST('mechine/new', params)
  },
  getTimeList: params => {
    return API.GET('schedule/all', params)
  },
  delControllers: params => {
    return API.DELETE(`mechine/${params.id}`, params)
  },
  putControllers: params => {
    return API.PUT(`mechine/${params.id}`, params)
  },
  startControllers: params => {
    return API.POST(`mechine/${params.id}/start/current`, params)
  },
  timingControllers: params => {
    return API.REPOST(`mechine/${params.id}/start`, params)
  },
  closeControllers: params => {
    return API.POST(`mechine/${params.id}/close`)
  }
}
