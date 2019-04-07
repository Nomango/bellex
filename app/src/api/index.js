import axios from 'axios'
import store from '../store/index'
import {
  Message
} from 'element-ui'
axios.defaults.withCredentials = true
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8' || 'application/json' // 配置请求头
// 添加一个请求拦截器
axios.interceptors.request.use(
  config => {
    if (window.sessionStorage.getItem('access_token')) {
      // 每次请求都为http头增加Authorization字段，其内容为token
      // config.headers.Authorization = window.sessionStorage.getItem('access_token');
    }
    if (config.method === 'get') {
      config.params = {
        // _t: Date.parse(new Date()) / 1000,
        // access_token: window.sessionStorage.getItem('access_token'),
        ...config.params
      }
    } else if (config.method === 'post') {
      // config.data = {
      //     _t: Date.parse(new Date()) / 1000,
      //     ...config.data,
      // }
    }
    store.state.ajaxLoading = true
    return config
  },
  error => {
    Message.error({
      message: '请求超时!'
    })
    return Promise.reject(error)
  }
)
// 响应拦截
axios.interceptors.response.use(function (response) {
  store.state.ajaxLoading = false
  return response
},
function (err) {
  store.state.ajaxLoading = false
  if (err.response) {
    Message.error({
      message: err.response.data.message
    })
    return Promise.resolve(null)
  } else {
    Message.error({
      message: '未知异常！'
    })
    return Promise.reject(err)
  }
})
// 基地址 http://localhost:8080/static/mock/index.json
export let base = '/api/v1/'
// 通用方法
export const POST = (url, params) => {
  console.info('POST请求路径' + `${base}${url}`);
  return axios.post(`${base}${url}`, params)
}
export const GET = (url, params) => {
  console.info('GET请求路径' + `${base}${url}`)
  return axios.get(`./${url}`, params)
}
