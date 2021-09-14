<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>全局配置</el-breadcrumb-item>
      <el-breadcrumb-item>用户管理</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 卡片信息 -->
    <el-card class="box-card">
      <el-row :gutter="20">
        <!-- 这个列里面放的是搜索框 -->
        <el-col :span="6">
          <el-input
            placeholder="请输入内容"
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>

        <!-- 这个列里面放的是添加用户的框 -->
        <el-col :span="6">
          <el-button type="primary">添加用户</el-button>
        </el-col>
      </el-row>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
        // 获取用户的参数列表
        queryInfo: {
            query: "",
            pagenum: 1,
            pagesize: 0
        }
    }
  },
  created() {
    this.GetUsersList()
  },
  methods: {
    GetUsersList() {
      const { data: res } = this.$http.get('users', {param: this.queryInfo})
      console.log(res)
    }
  }
}
</script>

<style lang="less" scoped></style>
