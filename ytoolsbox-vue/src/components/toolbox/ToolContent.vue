<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/toolbox'}">工具盒</el-breadcrumb-item>
      <el-breadcrumb-item>{{ toolForm.toolName }}</el-breadcrumb-item>
    </el-breadcrumb>
  </div>
</template>

<script>
export default {
  data() {
    return {
      toolForm: {},
    }
  },
  created() {
    this.GetToolInfo()
  },
  methods: {
    GetToolInfo() {
      this.toolForm = this.$route.query.toolInfo
    },
  },
}
</script>

<style>
</style>