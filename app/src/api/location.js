import * as API from './'
export default {
  getCity: params => {
    return API.GET('static/mock/city.json', params)
  }
}
