<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/toolbox' }">工具盒</el-breadcrumb-item>
      <el-breadcrumb-item>{{ "POC工具-泄密溯源" }}</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs type="border-card">
      <el-tab-pane label="配置">
        <POCToolConten_Config></POCToolConten_Config>
      </el-tab-pane>
      <el-tab-pane label="窃取数据">
        <el-empty description="点击下方按钮开始窃取数据" :image="require('../../../../assets/dataleakage.png')">
          <el-button type="primary" @click="stealData">窃取数据</el-button>
        </el-empty>
      </el-tab-pane>
    </el-tabs>

    <!-- 窃取数据详情的提示框 -->
    <el-dialog
      title="窃取数据详情"
      :visible.sync="dataLeakageDetailDialogVisible"
      width="60%"
      :close-on-click-modal="false"
      @close="dataLeakageDetailDialogVisible = false"
    >
      <div
        v-loading="isLoading"
        element-loading-text="窃取数据中..."
        element-loading-spinner="el-icon-loading"
        style="white-space: pre-line;"
      >
        <!-- pre可以展示制表符 -->
        <pre>{{ dataLeakageDetail }}</pre>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dataLeakageDetailDialogVisible = false">关闭</el-button>
      </span>
    </el-dialog>
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
      // 遮罩
      isLoading: false,
      loadingContent: '',
      dataLeakageDetailDialogVisible: false,
      dataLeakageDetail: '',

      data: {
        leakage: '',
      },
    }
  },
  methods: {
    async stealData() {
      this.dataLeakageDetail = ''
      this.dataLeakageDetailDialogVisible = true
      this.isLoading = true

      const { data: res } = await this.$http.get('custom/dataleakage')
      if (res.meta.status_code !== 200) {
        this.dataLeakageDetail = '窃取数据失败'
        this.$message.error(res.meta.message)
        this.isLoading = false
        return
      }

      // const { data: res2 } = await this.$http.get('custom/forceaudit')
      // if (res2.meta.status_code !== 200) {
      //   this.dataLeakageDetail = '窃取数据失败'
      //   this.$message.error(res2.meta.message)
      //   this.isLoading = false
      //   return
      // }

      this.dataLeakageDetail +=
        '以下为本次窃取的数据(特别备注:以下数据皆为代码生成,并非真实业务数据)\n-------------------------\n'
      for (let i = 0; i < res.data.Content.length; i++) {
        for (var key in res.data.Content[i]) {
          this.dataLeakageDetail += res.data.Content[i][key]
          this.dataLeakageDetail += '\n'
        }
      }

      this.dataLeakageDetail +=
        '\n-------------------------\n复制以上数据到泄密溯源功能中,进行溯源'
      this.isLoading = false
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