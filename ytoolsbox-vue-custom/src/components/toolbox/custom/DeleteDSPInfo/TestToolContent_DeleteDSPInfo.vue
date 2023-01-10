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
      <el-breadcrumb-item>{{ '测试工具-DSP清除器' }}</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs type="border-card">
      <el-tab-pane label="配置">
        <POCToolConten_Config></POCToolConten_Config>
      </el-tab-pane>
      <el-tab-pane label="清除器">
        <el-collapse v-model="activeNames">
          <el-collapse-item
            class="titleStyle"
            title="删除脆弱性与风险记录--作者:@周浩军62594(军子),@刘钦源13056(源总)"
            name="deleteRiskAndVunl"
          >
            <el-button type="danger" @click="deleteRiskAndVunl">开始删除</el-button>
          </el-collapse-item>
        </el-collapse>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
// import qs from 'qs'
import POCToolConten_Config from '../Config/POCToolContent_Config.vue'

export default {
  components: {
    POCToolConten_Config,
  },
  data() {
    return {
      // 遮罩
      isLoading: false,
      loadingContent: '',

      activeNames: ['deleteRiskAndVunl'],
    }
  },
  created() {},
  methods: {
    deleteRiskAndVunl() {
      this.$confirm('确定删除目标机器中的脆弱性与风险日志?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          this.loadingContent = '删除日志中...预计需要20秒的时间'
          this.isLoading = true

          const { data: res } = await this.$http.delete(
            'custom/cleaner/dsp/deleteRiskAndVunl'
          )

          this.isLoading = false

          if (res.meta.status_code == 200) {
            return this.$message({
              type: 'success',
              message: '删除成功!',
            })
          } else {
            return this.$message.error(res.meta.message)
          }
        })
        .catch(() => {
          this.isLoading = false
          this.$message({
            type: 'info',
            message: '已取消删除',
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
