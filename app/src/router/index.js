import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/pages/admin/Home/home'
Vue.use(Router)

const router = new Router({
  routes: [{
    path: '/',
    name: 'Home',
    component: Home,
    redirect: '/home/mainControl',
    children: [{
      path: '/home/building',
      component: () => import('@/pages/admin/Building/index.vue')
    }, {
      path: '/home/mainControl',
      component: () => import('@/pages/admin/MainControl/index.vue')
    }, {
      path: '/home/set',
      component: () => import('@/pages/admin/set/set.vue')
    }, {
      path: '/home/tableDetail',
      name: 'tableDetail',
      component: () => import('@/pages/admin/tableDetail/index.vue')
    }, {
      path: '/home/userManage',
      name: 'userManage',
      component: () => import('@/pages/admin/userManage/index.vue')
    }]
  }, {
    path: '/login',
    name: 'login',
    component: () => import('@/pages/admin/login/index.vue')
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