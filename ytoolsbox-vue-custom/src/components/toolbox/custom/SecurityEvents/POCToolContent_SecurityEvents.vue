<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/toolbox' }">工具盒</el-breadcrumb-item>
      <el-breadcrumb-item>{{ "POC工具-安全事件" }}</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs type="border-card" @tab-click="handleClick">
      <el-tab-pane label="配置">
        <POCToolConten_Config></POCToolConten_Config>
      </el-tab-pane>
      <el-tab-pane label="触发安全事件">
        <POCToolContent_TriggerSecurityEvents></POCToolContent_TriggerSecurityEvents>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import POCToolContent_TriggerSecurityEvents from './POCToolContent_TriggerSecurityEvents.vue'
import POCToolConten_Config from '../Config/POCToolContent_Config.vue'

export default {
  components: {
    POCToolContent_TriggerSecurityEvents,
    POCToolConten_Config,
  },
  methods: {},
}
</script>

<style>
</style>