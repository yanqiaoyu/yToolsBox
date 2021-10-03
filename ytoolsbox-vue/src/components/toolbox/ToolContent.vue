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
        <el-table :data="configForm" stripe border style="width: 100%">
          <!-- 只要添加了type=index，就能序号列 -->
          <el-table-column type="index" label="序号" width="100"></el-table-column>
          <el-table-column prop="toolConfigName" label="配置名称" width="450"></el-table-column>
          <el-table-column prop="toolConfigDesc" label="配置描述" width="797"></el-table-column>
          <el-table-column label="操作" width="300">
            <template slot-scope="scope">
              <!-- 编辑 -->
              <el-button
                type="primary"
                icon="el-icon-edit"
                circle
                @click="showEditToolConfig(scope.row.id)"
              ></el-button>
              <!-- 删除 -->
              <el-button
                type="danger"
                icon="el-icon-delete"
                circle
                @click="delteToolConfig(scope.row.id)"
              ></el-button>
            </template>
          </el-table-column>
        </el-table>
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
    // 发起请求拿到配置信息
    async GetToolConfig() {
      const { data: res } = await this.$http.get(
        'tools/config/' + this.toolInfo.id
      )

      this.configForm = res.data.toolsConfig
      console.log(this.configForm)

      if (res.meta.status_code != 200) {
        this.$message.error('获取工具配置信息失败')
      }
    },
    showEditToolConfig() {},
    delteToolConfig() {},
  },
}
</script>

<style>
</style>