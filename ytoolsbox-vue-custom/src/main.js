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


// const noAuthAndCookieURL = [
//   '/api/auth/custom/mock/vulnerability/SensitiveAPINotSec',
//   '/api/auth/custom/mock/vulnerability/NoneSensitiveAPINotSec',
//   '/api/auth/custom/mock/vulnerability/ws/v1/cluster/apps/new-application',
//   '/api/auth/custom/mock/vulnerability/?reqType=ElasticSearch',
//   '/api/auth/custom/mock/vulnerability/baizhi/_search/',
//   '/api/auth/custom/mock/vulnerability/_nodes',
//   '/api/auth/custom/mock/vulnerability/v1.40/images/json',
//   '/api/auth/custom/mock/vulnerability/actuator/info',
//   '/api/auth/custom/mock/vulnerability/jolokia/list',
//   '/api/webservices/list',
//   '/api/auth/custom/mock/vulnerability/v2/api-docs',
//   '/api/auth/custom/mock/vulnerability/solr/admin/info/system',
//   '/api/auth/custom/mock/vulnerability/api/v1/namespaces/pods',
//   '/api/auth/custom/mock/vulnerability/pods/',
//   '/api/auth/custom/mock/vulnerability/api/v1/settings/pinner',
//   '/api/auth/custom/mock/vulnerability/graphql',
//   '/api/auth/custom/mock/vulnerability/services',
//   '/api/auth/custom/mock/vulnerability/ws_utc/resources/setting/keystore',
//   '/api/auth/custom/mock/vulnerability/apisix/admin/migrate/export',
//   '/api/auth/custom/mock/vulnerability/run',
//   '/api/auth/custom/mock/vulnerability/actuator/gateway/routes/hacktest',
//   '/api/auth/custom/mock/vulnerability/mailsms/s',
//   '/api/auth/custom/mock/vulnerability/druid/sql.json?orderBy=SQL&orderType=desc&page=1&perPageCount=1000000',
//   '/api/auth/custom/mock/vulnerability/v1/submissions/create',
//   '/api/auth/custom/mock/vulnerability/customers/1',
//   '/api/auth/custom/mock/vulnerability/apisix/batch-requests',
//   '/api/auth/custom/mock/vulnerability/api/v1/secret/kube-system/kubernetes-dashboard-certs',
//   '/api/auth/custom/mock/vulnerability/openapi.json',
//   '/api/auth/custom/mock/vulnerability/user/admin/infolog.log'
// ]

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

  // 所有请求内容不保存 或 缓存到 Internet的临时文件中
  config.headers['Cache-Control'] = 'no-store'

  if (config.url.includes('/api/auth/custom/mock/vulnerability/main/index.php')) {
    config.headers.Referer = config.headers.Referer + '?password=12345678&username=admin123 '
  }

  if (config.url.includes('/api/auth/custom/mock/vulnerability/get')) {
    config.headers.Cookie = 'jsessionid=1223212313'
    config.headers.Origin = 'http://www.aliyun.com:8009'
  }

  if (config.url.includes('/api/auth/custom/mock/vulnerability/?reqType=ElasticSearch')) {
    config.headers.Host = config.headers.Host + ':9200'
  }

  if (config.url.includes('/api/auth/custom/mock/vulnerability/pods/')) {
    config.headers.Host = config.headers.Host + ':10250'
  }


  // if (noAuthAndCookieURL.includes(config.url)) {
  //   // 指定API删除cookie
  //   document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
  //   // 指定API删除Authorization
  //   delete config.headers.Authorization
  // }

  

  // if (config.url.endsWith('reqType=ElasticSearch')) {
  //   config.headers.Host = 'localhost:9200'
  // }

  // if (config.url.includes('/api/auth/custom/mock/vulnerability/pods/')) {
  //   config.url = 'http://localhost:10250/api/auth/custom/mock/vulnerability/pods/'
  // }

  // if (config.url == "/api/auth/custom/mock/securityevents/RequestTraverseAndReturnTooMuchSensitiveDataAbnormalTime?user=") {
  //   var process = require('child_process')
  //   var cmd = 'date';
  //   process.exec(cmd, function(error, stdout, stderr){
  //     console.log("error:"+error);
  //     console.log("stdout:"+stdout);
  //     console.log("stderr:"+stderr);
  //   })
  // }

  // // 最后必须return这个config
  // if (config.url == "http://10.74.4.46/api/auth/custom/mock/vulnerability/main/index.php") {
  //   config.headers['Referer'] = 'http://10.74.4.46/?password=12345678&username=admin123'
  // }

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
