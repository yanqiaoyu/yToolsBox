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
      <el-breadcrumb-item>{{ "POC工具-探针部署" }}</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs type="border-card">
      <el-tab-pane label="部署">
        <el-collapse v-model="activeNames">
          <el-collapse-item class="titleStyle" title="安装DSC-Agent" name="dscAgentCollapse">
            <el-form ref="dscAgentForm" :model="dscAgentForm" align-center>
              <el-form-item label-width="80px" label="IP地址" style="width:450px">
                <el-tooltip class="item" effect="dark" placement="top-start">
                  <div slot="content">想要把DSC-Agent安装在哪里,这里就填哪里的IP</div>
                  <el-input v-model="dscAgentForm.ip" placeholder="请输入IP地址"></el-input>
                </el-tooltip>
              </el-form-item>
              <el-form-item label-width="80px" label="SSH端口" style="width:450px">
                <el-input v-model="dscAgentForm.port" placeholder="请输入SSH端口"></el-input>
              </el-form-item>
              <el-form-item label-width="80px" label="SSH账号" style="width:450px">
                <el-input v-model="dscAgentForm.username" placeholder="请输入SSH账号"></el-input>
              </el-form-item>
              <el-form-item label-width="80px" label="SSH密码" style="width:450px">
                <el-input v-model="dscAgentForm.password" placeholder="请输入SSH密码" show-password></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="warning" @click="TestConnection(dscAgentForm)">测试SSH连通性</el-button>
                <el-button type="success" @click="Save(dscAgentForm)">保存配置</el-button>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="DoDeploy(dscAgentForm)">开始执行部署</el-button>
                <el-button type="primary">查看部署情况</el-button>
              </el-form-item>
            </el-form>
          </el-collapse-item>

          <el-collapse-item class="titleStyle" title="安装DAS-Agent" name="dasAgentCollapse">
            <el-form ref="dasAgentForm" :model="dasAgentForm" align-center>
              <el-form-item label-width="80px" label="IP地址" style="width:450px">
                <el-tooltip class="item" effect="dark" placement="top-start">
                  <div slot="content">想要把DAS-Agent安装在哪里,这里就填哪里的IP</div>
                  <el-input v-model="dasAgentForm.ip" placeholder="请输入IP地址"></el-input>
                </el-tooltip>
              </el-form-item>
              <el-form-item label-width="80px" label="SSH端口" style="width:450px">
                <el-input v-model="dasAgentForm.port" placeholder="请输入SSH端口"></el-input>
              </el-form-item>
              <el-form-item label-width="80px" label="SSH账号" style="width:450px">
                <el-input v-model="dasAgentForm.username" placeholder="请输入SSH账号"></el-input>
              </el-form-item>
              <el-form-item label-width="80px" label="SSH密码" style="width:450px">
                <el-input v-model="dasAgentForm.password" placeholder="请输入SSH密码" show-password></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="warning" @click="TestConnection(dasAgentForm)">测试SSH连通性</el-button>
                <el-button type="success" @click="Save(dasAgentForm)">保存配置</el-button>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="DoDeploy(dasAgentForm)">开始执行部署</el-button>
                <el-button type="primary">查看部署情况</el-button>
              </el-form-item>
            </el-form>
          </el-collapse-item>
        </el-collapse>
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import qs from 'qs'
export default {
  data() {
    return {
      activeNames: ['dscAgentCollapse', 'dasAgentCollapse'],
      dscAgentForm: {
        type: 'dsc',
        ip: '',
        port: '',
        username: '',
        password: '',
      },
      dasAgentForm: {
        type: 'das',
        ip: '',
        port: '',
        username: '',
        password: '',
      },

      // 遮罩
      isLoading: false,
      loadingContent: '',
    }
  },
  created() {
    this.GetAgentInstallConfig()
  },
  methods: {
    // 保存配置
    async Save(form) {
      const { data: res } = await this.$http.post(
        'custom/installconfig',
        qs.stringify(form)
      )
      if (res.meta.status_code !== 200) {
        this.$message.error('保存失败')
        return
      }
      this.$message({
        message: res.meta.message,
        type: 'success',
      })
      // 保存后更新一次
      this.GetAgentInstallConfig()
    },
    // 查询配置
    async GetAgentInstallConfig() {
      const { data: res } = await this.$http.get('custom/installconfig')

      for (var i in res.data.AgentInstallConfigList) {
        console.log('Agent安装配置: ', res.data.AgentInstallConfigList[i])
        if (res.data.AgentInstallConfigList[i]['type'] == 'das') {
          this.dasAgentForm.ip = res.data.AgentInstallConfigList[i]['ip']
          this.dasAgentForm.port = res.data.AgentInstallConfigList[i]['port']
          this.dasAgentForm.username =
            res.data.AgentInstallConfigList[i]['username']
          this.dasAgentForm.password =
            res.data.AgentInstallConfigList[i]['password']
        } else if (res.data.AgentInstallConfigList[i]['type'] == 'dsc') {
          this.dscAgentForm.ip = res.data.AgentInstallConfigList[i]['ip']
          this.dscAgentForm.port = res.data.AgentInstallConfigList[i]['port']
          this.dscAgentForm.username =
            res.data.AgentInstallConfigList[i]['username']
          this.dscAgentForm.password =
            res.data.AgentInstallConfigList[i]['password']
        }
      }
    },
    // 测试SSH链接
    async TestConnection(form) {
      this.loadingContent = '测试SSH链接中'
      this.isLoading = true

      // 为了防止有人修改ssh配置后忘记保存直接测试连接 这里保存一下
      this.Save(form)
      //   console.log(form)

      const { data: res } = await this.$http.post(
        'custom/testssh',
        qs.stringify(form)
      )
      this.isLoading = false
      if (res.meta.status_code !== 200) {
        this.$message.error('SSH连接失败: ' + res.meta.message)
        return
      }
      this.$message({
        message: res.meta.message,
        type: 'success',
      })
    },
    // 执行部署
    async DoDeploy(form) {
      this.$confirm('确定开始部署?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          this.loadingContent = '部署agent中'
          this.isLoading = true

          const { data: res } = await this.$http.post(
            'custom/installagent',
            qs.stringify(form)
          )

          this.isLoading = false
          if (res.meta.status_code == 200) {
            return this.$message({
              type: 'success',
              message: '部署成功!',
            })
          } else {
            return this.$message.error({
              message: res.meta.message,
            })
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消部署',
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