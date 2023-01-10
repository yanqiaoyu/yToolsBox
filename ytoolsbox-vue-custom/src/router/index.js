import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'
import Home from '../components/Home.vue'
import User from '../components/user/User.vue'
import UpdateLog from '../components/UpdateLog.vue'
import Rights from '../components/power/Rights.vue'
import ToolBox from '../components/toolbox/ToolBox.vue'
import About from '../components/About.vue'
import Task from '../components/task/Task.vue'
import CronTask from '../components/task/CronTask.vue'
import AddTool from '../components/toolbox/AddTool.vue'
import ToolContent from '../components/toolbox/ToolContent.vue'
import POCToolContent_RiskandVulnerability from '../components/toolbox/custom/RiskAndVulnerability/POCToolContent_RiskandVulnerability.vue'
import POCToolContent_InstallAgent from '../components/toolbox/custom/InstallAgent/POCToolContent_InstallAgent.vue'
import POCToolContent_DSM from '../components/toolbox/custom/DSM/POCToolContent_DSM.vue'
import POCToolContent_DataLeakage from '../components/toolbox/custom/DataLeakage/POCToolContent_DataLeakage.vue'
import POCToolContent_SecurityEvents from '../components/toolbox/custom/SecurityEvents/POCToolContent_SecurityEvents.vue'
import TestToolContent_CustomRequest from '../components/toolbox/custom/CustomRequest/TestToolContent_CustomRequest.vue'
import TestToolContent_DeleteDSPInfo from '../components/toolbox/custom/DeleteDSPInfo/TestToolContent_DeleteDSPInfo.vue'
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
      { path: '/users', component: User, meta: { title: '用户管理' } },
      { path: '/updatelog', component: UpdateLog, meta: { title: '更新日志' } },
      { path: '/rights', component: Rights, meta: { title: '权限管理' } },
      { path: '/toolbox', component: ToolBox, meta: { title: '工具盒' } },
      { path: '/about', component: About, meta: { title: '关于' } },
      { path: '/task', component: Task, meta: { title: '普通任务' } },
      { path: '/crontask', component: CronTask, meta: { title: '定时任务' } },
      { path: '/toolbox/add', component: AddTool, meta: { title: '添加工具' } },
      {
        path: '/toolbox/tool',
        name: 'toolbox_tool',
        component: ToolContent,
        meta: { title: '工具详情' }
      },
      {
        path: '/toolbox/poc_tool_risk_and_vulnerability',
        name: 'poc_tool_risk_and_vulnerability',
        component: POCToolContent_RiskandVulnerability,
        meta: { title: 'POC工具-脆弱性与风险' }
      },
      {
        path: '/toolbox/poc_tool_security_events',
        name: 'poc_tool_security_events',
        component: POCToolContent_SecurityEvents,
        meta: { title: 'POC工具-安全事件' }
      },
      {
        path: '/toolbox/poc_tool_install_agent',
        name: 'poc_tool_install_agent',
        component: POCToolContent_InstallAgent,
        meta: { title: 'POC工具-探针部署' }
      },
      {
        path: '/toolbox/poc_tool_dsm',
        name: 'poc_tool_dsm',
        component: POCToolContent_DSM,
        meta: { title: 'POC工具-分类分级' }
      },
      {
        path: '/toolbox/poc_tool_dataleakage',
        name: 'poc_tool_dataleakage',
        component: POCToolContent_DataLeakage,
        meta: { title: 'POC工具-泄密溯源' }
      },
      {
        path: '/toolbox/test_tool_custom_request',
        name: 'test_tool_custom_request',
        component: TestToolContent_CustomRequest,
        meta: { title: '测试工具-自定义请求' }
      },
      {
        path: '/toolbox/test_tool_delete_dsp_info',
        name: 'test_tool_delete_dsp_info',
        component: TestToolContent_DeleteDSPInfo,
        meta: { title: '测试工具-DSP清除器' }
      }
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
