<template>
  <div
    v-loading.fullscreen.lock="isLoading"
    :element-loading-text="loadingContent"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/toolbox' }">工具盒</el-breadcrumb-item>
      <el-breadcrumb-item>{{ "POC工具-分类分级" }}</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs v-model="activeTabName" type="border-card">
      <el-tab-pane label="配置" name="config">
        <POCToolConten_Config></POCToolConten_Config>
      </el-tab-pane>
      <el-tab-pane label="使用" name="use">
        <el-collapse v-model="activeCollapseName">
          <el-collapse-item class="titleStyle" title="下发扫描" name="scan">
            <div>
              <el-button type="success" @click="addDataClassifyTask">自动新增数据源并下发扫描任务</el-button>
              <el-tooltip class="item" effect="dark" placement="top-start">
                <div slot="content">
                  自动在大脑上新增一个Postgres数据源
                  <br />
                  IP为{{ toolboxIP }},端口为5432,账号为postgres,密码为test123456
                  <br />并自动下发扫描任务
                  <br />该数据源内置可以匹配所有敏感规则的数据
                </div>
                <i class="header-icon el-icon-info" style="margin-left:10px"></i>
              </el-tooltip>
            </div>
          </el-collapse-item>
        </el-collapse>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import POCToolConten_Config from '../Config/POCToolContent_Config.vue'
export default {
  components: {
    POCToolConten_Config,
  },
  data() {
    return {
      activeCollapseName: 'scan',
      activeTabName: 'use',

      toolboxIP: '',

      // 遮罩
      isLoading: false,
      loadingContent: '',
    }
  },
  created() {
    this.toolboxIP = window.location.host
  },
  methods: {
    // 下发数据源,并新增AI任务
    async addDataClassifyTask() {
      this.$confirm('确定立即添加数据源并下发扫描任务?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          // const { data: res } = await this.$http.delete('tasks/')
          this.loadingContent = '下发数据源中'
          this.isLoading = true

          const { data: res } = await this.$http.post('custom/adddataclassify')

          this.isLoading = false

          if (res.meta.status_code == 200) {
            return this.$message({
              type: 'success',
              message: '下发数据源成功!',
            })
          } else {
            return this.$message.error(res.meta.message)
          }
        })
        .catch(() => {
          this.isLoading = false
          this.$message({
            type: 'info',
            message: '已取下发数据源',
          })
        })
    },
  },
}
</script>

<style lang="less" scoped>
.titleStyle {
  /deep/ .el-collapse-item__header {
    font-weight: bold;
    font-size: 18px;
  }
}
</style>