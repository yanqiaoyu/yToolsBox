import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'
// 导入全局样式
import './assets/css/global.css'
// 导入axios发送ajax请求
import axios from 'axios'

// 配置请求的路径
axios.defaults.baseURL = 'http://localhost/'
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
