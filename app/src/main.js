// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App'
import router from './router'
// import VueLazyLoad from 'vue-lazyload' // 图片懒加载
import store from './store'
import 'styles/reset.css'
import 'styles/iconfont.css'

Vue.config.productionTip = false

/* eslint-disable no-new */
// Vue.use(VueLazyLoad, {
//   loading: require('./assets/imgs/loading.gif')
// })
Vue.use(ElementUI)
new Vue({
  el: '#app',
  router,
  store,
  components: {
    App
  },
  template: '<App/>'
})
