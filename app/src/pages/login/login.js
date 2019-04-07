import Vue from 'Vue'
import login from './login.vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import router from '../../router/login'
import 'styles/reset.css'
import 'styles/iconfont.css'
import VueRouter from 'vue-router'
Vue.use(VueRouter)
Vue.config.productionTip = false

Vue.use(ElementUI)
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render: h => h(login)
})
