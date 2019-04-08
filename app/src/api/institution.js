import * as API from './'
export default {
  getInstitutionList: params => {
    return API.GET('institution/all', params)
  },
  putInstitutionList: params => {
    return API.PUT(`institution/${params.id}`, params)
  },
  addInstitutionList: params => {
    return API.POST('institution/new', params)
  },
  delInstitutionList: params => {
    return API.DELETE(`institution/${params.id}`)
  }
}
