<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>全局配置</el-breadcrumb-item>
      <el-breadcrumb-item>权限管理</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 卡片 -->
    <el-card class="box-card">
      <el-row :gutter="20"> </el-row>

      <!-- 权限列表展示区 -->
      <el-table :data="rightsList" stripe border style="width: 100%">
        <!-- 只要添加了type=index，就能序号列 -->
        <el-table-column type="index" label="序号"> </el-table-column>
        <el-table-column prop="authname" label="模块名称"> </el-table-column>
        <el-table-column prop="path" label="路径"> </el-table-column>
        <el-table-column prop="level" label="权限等级">
          <template slot-scope="scope">
            <el-tag effect="dark" :type="scope.row.type">
              {{ scope.row.label }}
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      rightsList: []
    }
  },
  created() {
    this.GetRightsList()
  },
  methods: {
    async GetRightsList() {
      const { data: res } = await this.$http.get('rights')
      this.rightsList = res.data.rightslist

      // 为权限数组添加了额外的两组属性，这样渲染时可以进行区分
      this.rightsList.forEach(item => {
        if (item.level === 0) {
          item.type = 'success'
          item.label = '级别一'
        } else if (item.level === 1) {
          item.type = 'warning'
          item.label = '级别二'
        } else if (item.level === 2) {
          item.type = 'danger'
          item.label = '级别三'
        }
      })
    }
  }
}
</script>

<style></style>
