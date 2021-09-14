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
import './plugins/element.js'
// 导入全局样式
import './assets/css/global.css'
// 导入axios发送ajax请求
import axios from 'axios'

// 配置请求的路径
// axios.defaults.baseURL = 'http://103.44.241.227/api/auth/'
axios.defaults.baseURL = 'http://localhost/api/auth/'

axios.interceptors.request.use(config => {
  // console.log(config)
  config.headers.Authorization = window.sessionStorage.getItem('token')
  // 最后必须return这个config
  return config
})

// 在Vue的原型上挂载axios，让所有示例都能发送http请求
Vue.prototype.$http = axios

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
