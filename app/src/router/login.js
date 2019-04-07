import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/pages/login/home/index'
Vue.use(Router)

const router = new Router({
  routes: [{
    path: '/',
    name: 'login',
    component: Home
  }]
})

// 路由守卫
// const toPath = ['/login/code', '/login/index', '/login', '/login/forget', '/login/newPsd']
// router.beforeEach ((to, from, next) => {
//   const isLogin = getSession('ele_login')
//   if (toPath.indexOf(to.path) >= 0) {
//     next()
//   } else { // 是否在登录状态下
//     if (isLogin) {
//       next()
//     } else {
//       next('/login')
//     }
//   }
// })
export default router
