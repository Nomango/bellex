import * as API from './'
export default {
  getControllers: params => {
    return API.GET('static/mock/controllers.json', params)
  }
}
