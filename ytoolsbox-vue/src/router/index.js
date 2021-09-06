import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Home from '../components/Home.vue'
Vue.use(VueRouter)

const routes = [
  // 将根路径 '/'，重定向到 '/login'
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/login',
    component: Login,
    meta: {
      // 页面标题title
      title: '登录'
    }
  },
  {
    path: '/home',
    component: Home,
    meta: {
      // 页面标题title
      title: '工具盒管理系统'
    }
  }
]

const router = new VueRouter({
  routes
})

export default router
