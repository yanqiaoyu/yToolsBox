<template>
  <div>
    <!-- 回到顶部 -->
    <el-backtop>
      <div
        style="{
        height: 100%;
        width: 100%;
        background-color: #f2f5f6;
        box-shadow: 0 0 6px rgba(0,0,0, .12);
        text-align: center;
        line-height: 40px;
        color: #1989fa;
      }"
      >UP</div>
    </el-backtop>

    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>定时任务</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs type="border-card">
      <el-tab-pane label="定时任务配置中心">
        <Config :get-cron-tasks-list="getCronTasksList"></Config>
      </el-tab-pane>
      <el-tab-pane label="定时任务执行结果">
        <ExeResult></ExeResult>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import Config from './CronTask_Config.vue'
import ExeResult from './CronTask_ExeResult.vue'

export default {
  components: { Config, ExeResult },
  methods: {
    //获取定时任务列表
    async getCronTasksList() {
      console.log('查询定时任务列表')
    },
  },
}
</script>

<style lang="less" scoped>
</style>