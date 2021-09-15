import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Home from '../components/Home.vue'
import User from '../components/user/User.vue'
import UpdateLog from '../components/UpdateLog.vue'
Vue.use(VueRouter)

const routes = [
  // 将根路径 '/'，重定向到 '/login'
  {
    path: '/',
    redirect: '/login'
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
    },
    redirect: '/updatelog',
    children: [
      {path: '/userconfig', component: User, meta: {title: "用户管理与配置"} },
      {path: '/updatelog', component: UpdateLog, meta:{title: "更新日志"}}
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  // 如果将要访问的路径是/login，那么放行
  if (to.path === '/login') return next()

  // 拿取本地的token
  const tokenStr = window.sessionStorage.getItem('token')
  // 如果本地没有token，那么返回登录界面
  if (!tokenStr) return next('/login')
  // 如果本地有token，那么放行
  next()
})

export default router
