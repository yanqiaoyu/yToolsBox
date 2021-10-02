<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/toolbox'}">工具盒</el-breadcrumb-item>
      <el-breadcrumb-item>{{ toolInfo.toolName }}</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs type="border-card">
      <el-tab-pane label="配置管理">
        <!-- 用户列表展示区 -->
        <!-- <el-table :data="configForm" stripe border style="width: 100%"> -->
        <!-- 只要添加了type=index，就能序号列 -->
        <!-- <el-table-column type="index" label="序号"></el-table-column>
        <el-table-column prop="toolType" label="配置名称"></el-table-column>-->
        <!-- </el-table>     -->
      </el-tab-pane>
      <el-tab-pane label="工具详情"></el-tab-pane>
      <el-tab-pane label="执行"></el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
export default {
  data() {
    return {
      toolInfo: {},
      configForm: {},
    }
  },
  created() {
    // 拿到传进来的这个工具的信息
    this.toolInfo = this.$route.query.toolInfo
    // 拿到这个工具的配置信息
    this.GetToolConfig()
  },
  methods: {
    async GetToolConfig() {
      console.log(this.$route.query.toolInfo)
      const { data: res } = await this.$http.get(
        'tools/config/' + this.toolInfo.id
      )
      console.log(res)

      if (res.meta.status_code != 200) {
        this.$message.error('获取工具配置信息失败')
      }
    },
  },
}
</script>

<style>
</style>