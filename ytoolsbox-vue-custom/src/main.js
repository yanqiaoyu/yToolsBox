/*
 * @Author: YanQiaoYu
 * @Github: https://github.com/yanqiaoyu?tab=repositories
 * @Date: 2021-09-11 23:48:37
 * @LastEditors: YanQiaoYu
 * @LastEditTime: 2021-09-12 12:27:09
 * @FilePath: /ytoolsbox-vue/src/main.js
 */
import Vue from 'vue'
import App from './App.vue'
import router from './router'
import Moment from 'moment'
import './plugins/element.js'
// 导入全局样式
import './assets/css/global.css'
// 导入axios发送ajax请求
import axios from 'axios'

import nProgress from 'nprogress'
import 'nprogress/nprogress.css'

import commonFun from './assets/js/common.js'

// 允许携带cookie
axios.defaults.withCredentials = true
// 测试，生产环境，不同的请求的路径
if (process.env.NODE_ENV == 'production') {
  let host = window.location.host //主机
  axios.defaults.baseURL = 'http://' + host + '/api/auth/'
} else {
  axios.defaults.baseURL = 'http://localhost:8081/api/auth/'
}

// 请求拦截器
axios.interceptors.request.use(config => {
  // console.log(config)
  nProgress.start()
  config.headers['Authorization'] =
    'Bearer ' + window.sessionStorage.getItem('token')

  // 最后必须return这个config
  return config
})

// 响应拦截器
axios.interceptors.response.use(
  config => {
    nProgress.done()
    return config
  },
  error => {
    console.log('进入响应拦截器')
    // console.log(error.response.data.meta.message)
    console.log(error.response)
    // console.log(
    //   error.response.data.meta.message === '单个账号在一段时间内返回大量4XX' ||
    //     error.response.data.meta.message === '单个IP在一段时间内返回大量4XX'
    // )

    // if (
    //   error.response.data.meta.message === '单个账号在一段时间内返回大量4XX' ||
    //   error.response.data.meta.message === '单个IP在一段时间内返回大量4XX'
    // ) {
    //   console.log('退出响应拦截器')
    return Promise.resolve(error.response)
    // }
    // console.log('退出响应拦截器')
    // return Promise.reject.apply(error)
  }
)

// 在Vue的原型上挂载axios，让所有实例都能发送http请求
Vue.prototype.$http = axios
// 在Vue的原型上挂载Moment，让所有实例都能调用时间戳转换
Vue.prototype.$moment = Moment
// 在Vue的原型上挂载common，让所有实例都能调用公共方法
Vue.prototype.$commonFun = commonFun

Vue.config.productionTip = false
router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
