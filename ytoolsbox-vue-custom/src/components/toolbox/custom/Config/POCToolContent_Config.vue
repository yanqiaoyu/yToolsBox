<template>
  <div
    v-loading.fullscreen.lock="isLoading"
    :element-loading-text="loadingContent"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
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

    <el-collapse v-model="activeNames">
      <el-collapse-item class="titleStyle" title="工具盒配置" name="toolBoxCollapse">
        <el-form ref="toolBoxForm" :model="toolBoxForm" align-center>
          <el-form-item label-width="130px" label="工具盒IP地址" style="width:450px">
            <el-tooltip class="item" effect="dark" placement="top-start">
              <div slot="content">工具盒安装在哪里,这里就填哪里的IP,默认提取浏览器中的IP地址</div>
              <el-input v-model="toolBoxForm.toolBoxAddress" placeholder="请输入工具盒IP地址"></el-input>
            </el-tooltip>
          </el-form-item>
          <el-form-item label-width="130px" label="工具盒SSH端口" style="width:450px">
            <el-input v-model="toolBoxForm.toolBoxSSHPort" placeholder="请输入工具盒SSH端口"></el-input>
          </el-form-item>
          <el-form-item label-width="130px" label="工具盒SSH账号" style="width:450px">
            <el-input v-model="toolBoxForm.toolBoxSSHUserName" placeholder="请输入工具盒SSH账号"></el-input>
          </el-form-item>
          <el-form-item label-width="130px" label="工具盒SSH密码" style="width:450px">
            <el-input
              v-model="toolBoxForm.toolBoxSSHPassword"
              placeholder="请输入工具盒SSH密码"
              show-password
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="warning" @click="testConnection(toolBoxForm, '工具盒')">测试工具盒SSH连通性</el-button>
            <el-button type="success" @click="save(toolBoxForm)">保存</el-button>
          </el-form-item>
        </el-form>
      </el-collapse-item>

      <el-collapse-item class="titleStyle" title="安全大脑配置" name="DSCCollapse">
        <el-form ref="dscForm" :model="dscForm" align-center>
          <el-form-item label-width="160px" label="安全大脑IP地址" style="width:450px">
            <el-input v-model="dscForm.dscAddress" placeholder="请输入数据安全大脑地址"></el-input>
          </el-form-item>
          <el-form-item label-width="160px" label="安全大脑SSH端口" style="width:450px">
            <el-input v-model="dscForm.dscSSHPort" placeholder="请输入数据安全大脑SSH端口"></el-input>
          </el-form-item>
          <el-form-item label-width="160px" label="安全大脑SSH账号" style="width:450px">
            <el-input v-model="dscForm.dscSSHUserName" placeholder="请输入数据安全大脑SSH账号"></el-input>
          </el-form-item>
          <el-form-item label-width="160px" label="安全大脑SSH密码" style="width:450px">
            <el-input v-model="dscForm.dscPassword" placeholder="请输入数据安全大脑SSH密码" show-password></el-input>
          </el-form-item>
          <el-form-item label-width="160px" label="DSC-Agent认证token" style="width:450px">
            <el-tooltip class="item" effect="dark" placement="top-start">
              <div slot="content">进入数据安全大脑-系统管理-设备管理,找到设备类型为DP的设备(没有就新建),复制那个认证token到此处</div>
              <el-input v-model="dscForm.dp_Token" placeholder="请输入DSC-Agent认证token"></el-input>
            </el-tooltip>
          </el-form-item>

          <el-form-item label-width="160px" label="安全大脑前端账号" style="width:450px">
            <el-input v-model="dscForm.dscWebUserName" placeholder="请输入数据安全大脑前端账号"></el-input>
          </el-form-item>
          <el-form-item label-width="160px" label="安全大脑前端密码" style="width:450px">
            <el-input v-model="dscForm.dscWebPassword" placeholder="请输入数据安全大脑前端密码" show-password></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="warning" @click="testConnection(dscForm, '安全大脑')">测试安全大脑SSH连通性</el-button>

            <el-button type="success" @click="save(dscForm)">保存</el-button>
          </el-form-item>
        </el-form>
      </el-collapse-item>

      <el-collapse-item class="titleStyle" title="配置DSC-Agent" name="ConfigureAgentCollapse">
        <el-form align-center>
          <el-form-item>
            <el-button type="primary" @click="UpdateDSCAgent()">更新工具盒中Agent的配置</el-button>
            <el-tooltip class="item" effect="dark" placement="top-start">
              <div slot="content">
                这里更新的agent,更新的是工具盒所在虚拟机中的agent
                <br />点击按钮之后会执行以下步骤:
                <br />1. 根据你在上面填写的安全大脑IP地址, 更新agent中指向的大脑地址
                <br />2. 根据你在上面填写的DSC-Agent认证token, 更新agent中对接大脑的token
                <br />3. 重启工具盒中的agent, 使新配置生效
              </div>
              <i class="header-icon el-icon-info" style="margin-left:10px"></i>
            </el-tooltip>
          </el-form-item>
          <el-form-item>
            <el-button type="success" @click="GetDSCAgentConfig()">查看工具盒中Agent的配置</el-button>
            <el-tooltip class="item" effect="dark" placement="top-start">
              <div slot="content">更新完配置后,可以查看一下配置看看是否更新成功</div>
              <i class="header-icon el-icon-info" style="margin-left:10px"></i>
            </el-tooltip>
          </el-form-item>
        </el-form>
      </el-collapse-item>

      <el-collapse-item class="titleStyle" title="配置账号提取" name="ConfigureAccountExtract">
        <el-form align-center>
          <el-form-item>
            <el-button type="primary" @click="UpdateAccountExtract()">更新大脑中账号提取的配置</el-button>
            <el-tooltip class="item" effect="dark" placement="top-start">
              <div slot="content">
                这里更新的账号提取的配置,更新的是数据安全大脑所在虚拟机中的账号提取配置
                <br />点击按钮之后会执行以下步骤:
                <br />1. 在本地生成一份可以触发脆弱性与风险的账号提取配置
                <br />2. 根据你在上面填写的安全大脑IP地址, 上传这份配置到安全大脑所在的虚拟机中
                <br />3. 重启安全大脑的storage容器, 使配置生效
              </div>
              <i class="header-icon el-icon-info" style="margin-left:10px"></i>
            </el-tooltip>
          </el-form-item>
          <el-form-item>
            <el-button type="success" @click="GetAccountExtractConfig()">查看大脑中账号提取的配置</el-button>
            <el-tooltip class="item" effect="dark" placement="top-start">
              <div slot="content">更新完配置后,可以查看一下配置看看是否更新成功</div>
              <i class="header-icon el-icon-info" style="margin-left:10px"></i>
            </el-tooltip>
          </el-form-item>
        </el-form>
      </el-collapse-item>

      <el-collapse-item class="titleStyle" title="配置大脑风险与脆弱性阈值" name="ConfigureDSCThreshold">
        <el-form align-center>
          <el-form-item>
            <el-button type="primary" @click="ModifyDSCThreshold(mode='poc')">一键调整风险与脆弱性阈值</el-button>
            <el-tooltip class="item" effect="dark" placement="top-start">
              <div slot="content">这里会将大脑关于风险与脆弱性的阈值调整到适合POC工具触发的值</div>
              <i class="header-icon el-icon-info" style="margin-left:10px"></i>
            </el-tooltip>
          </el-form-item>
          <el-form-item>
            <el-button type="success" @click="ModifyDSCThreshold(mode='default')">一键还原风险与脆弱性阈值</el-button>
            <el-tooltip class="item" effect="dark" placement="top-start">
              <div slot="content">这里会将大脑关于风险与脆弱性的阈值还原为默认值</div>
              <i class="header-icon el-icon-info" style="margin-left:10px"></i>
            </el-tooltip>
          </el-form-item>
        </el-form>
      </el-collapse-item>
    </el-collapse>

    <!-- 配置详情的提示框 -->
    <el-dialog
      title="配置详情"
      :visible.sync="configDetailDialogVisible"
      width="60%"
      :close-on-click-modal="false"
      @close="configDetailDialogVisible = false"
    >
      <div
        v-loading="isConfigLoading"
        element-loading-text="拼命加载中"
        element-loading-spinner="el-icon-loading"
        style="white-space: pre-line;"
      >
        <!-- pre可以展示制表符 -->
        <pre>{{ configDetail }}</pre>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="configDetailDialogVisible = false">关闭</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import qs from 'qs'
export default {
  data() {
    return {
      activeNames: [
        'toolBoxCollapse',
        'DSCCollapse',
        'ConfigureAgentCollapse',
        'ConfigureAccountExtract',
        'ConfigureDSCThreshold',
      ],
      toolBoxForm: {
        toolBoxAddress: '',
        toolBoxSSHPort: '',
        toolBoxSSHUserName: '',
        toolBoxSSHPassword: '',
      },
      dscForm: {
        dscAddress: '',
        dscSSHPort: '',
        dscSSHUserName: '',
        dscPassword: '',
        dp_Token: '',
        dscWebUserName: '',
        dscWebPassword: '',
      },

      // 遮罩
      isLoading: false,
      isConfigLoading: false,
      loadingContent: '',

      // 配置内容
      configDetail: '',
      configDetailDialogVisible: false,
    }
  },
  created() {
    // 拿取浏览器中的IP地址
    this.toolBoxForm.toolBoxAddress = window.location.host.split(':', 1)[0]
    // 拿取配置
    this.GetPOCConfig()
  },
  methods: {
    // 查看工具盒中Agent的配置
    async GetDSCAgentConfig() {
      this.configDetail = ''
      this.configDetailDialogVisible = true
      this.isConfigLoading = true
      // this.taskDetail = row.returnContent
      const { data: res } = await this.$http.get('custom/dscagentconfig')
      if (res.meta.status_code !== 200) {
        this.$message.error(res.meta.message)
        this.isConfigLoading = false
        return
      }
      this.configDetail = res.data.Content
      this.isConfigLoading = false
    },
    // 查看大脑中账号提取的配置
    async GetAccountExtractConfig() {
      this.configDetail = ''
      this.configDetailDialogVisible = true
      this.isConfigLoading = true
      const { data: res } = await this.$http.get(
        'custom/getaccountextractconfig'
      )
      if (res.meta.status_code !== 200) {
        this.$message.error(res.meta.message)
        this.isConfigLoading = false
        return
      }
      this.configDetail = res.data.Content
      this.isConfigLoading = false
    },
    // 测试SSH链接
    async testConnection(form, flag) {
      this.loadingContent = '测试SSH链接中'
      this.isLoading = true

      // 为了防止有人修改ssh配置后忘记保存直接测试连接 这里保存一下
      this.save(form)

      var sendForm = {}
      if (flag == '工具盒') {
        sendForm['ip'] = form.toolBoxAddress
        sendForm['port'] = form.toolBoxSSHPort
        sendForm['username'] = form.toolBoxSSHUserName
        sendForm['password'] = form.toolBoxSSHPassword
      } else if (flag == '安全大脑') {
        sendForm['ip'] = form.dscAddress
        sendForm['port'] = form.dscSSHPort
        sendForm['username'] = form.dscSSHUserName
        sendForm['password'] = form.dscPassword
      }

      const { data: res } = await this.$http.post(
        'custom/testssh',
        qs.stringify(sendForm)
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
    handleChange() {},

    // 更新工具盒中Agent的配置
    async UpdateDSCAgent() {
      this.loadingContent = '正在更新工具盒中Agent的配置'
      this.isLoading = true

      // 遍历一次两个关键的form,确认关键信息都填写了
      console.log(this.toolBoxForm, this.dscForm)
      for (const i in this.toolBoxForm) {
        if (this.toolBoxForm[i] == '') {
          this.$message.error('请确保上面的配置都填写了')
          this.isLoading = false
          return
        }
      }
      for (const i in this.dscForm) {
        if (this.dscForm[i] == '') {
          this.$message.error('请确保上面的配置都填写了')
          this.isLoading = false
          return
        }
      }

      const { data: res } = await this.$http.get('custom/updatedscagent')
      this.isLoading = false

      if (res.meta.status_code !== 200) {
        this.$message.error('更新工具盒中Agent的配置失败')
        return
      }
      this.$message({
        message: res.meta.message,
        type: 'success',
      })
    },

    // 更新大脑中账号提取的配置
    async UpdateAccountExtract() {
      this.loadingContent = '正在更新大脑中账号提取的配置'
      this.isLoading = true

      // 遍历一次两个关键的form,确认关键信息都填写了
      console.log(this.toolBoxForm, this.dscForm)
      for (const i in this.toolBoxForm) {
        if (this.toolBoxForm[i] == '') {
          this.$message.error('请确保上面的配置都填写了')
          this.isLoading = false
          return
        }
      }
      for (const i in this.dscForm) {
        if (this.dscForm[i] == '') {
          this.$message.error('请确保上面的配置都填写了')
          this.isLoading = false
          return
        }
      }

      const { data: res } = await this.$http.get(
        'custom/updatedscaccountextract'
      )
      this.isLoading = false

      if (res.meta.status_code !== 200) {
        this.$message.error('更新大脑中账号提取的配置失败')
        return
      }
      this.$message({
        message: res.meta.message,
        type: 'success',
      })
    },

    // 拿取配置
    async GetPOCConfig() {
      const { data: res } = await this.$http.get('custom/pocconfig')
      // console.log(res.data)
      if (res.meta.status_code == 200) {
        this.toolBoxForm.toolBoxAddress = res.data.toolBoxAddress
        this.toolBoxForm.toolBoxSSHPassword = res.data.toolBoxSSHPassword
        this.toolBoxForm.toolBoxSSHPort = res.data.toolBoxSSHPort
        this.toolBoxForm.toolBoxSSHUserName = res.data.toolBoxSSHUserName

        this.dscForm.dscAddress = res.data.dscAddress
        this.dscForm.dscSSHPort = res.data.dscSSHPort
        this.dscForm.dscSSHUserName = res.data.dscSSHUserName
        this.dscForm.dscPassword = res.data.dscPassword
        this.dscForm.dp_Token = res.data.dp_Token

        this.dscForm.dscWebUserName = res.data.dscWebUserName
        this.dscForm.dscWebPassword = res.data.dscWebPassword

        // console.log(this.toolBoxForm)
      }
    },

    // 保存配置
    async save(form) {
      const { data: res } = await this.$http.post(
        'custom/pocconfig',
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
      this.GetPOCConfig()
    },

    // 调整大脑的阈值
    async ModifyDSCThreshold(mode) {
      console.log('调整阈值策略: ', mode)
      this.loadingContent = '正在调整/还原大脑中风险与脆弱性的阈值'
      this.isLoading = true
      const { data: res } = await this.$http.post(
        'custom/modifydscthreshold',
        qs.stringify({ mode: mode })
      )

      if (res.meta.status_code !== 200) {
        this.$message.error('调整/还原大脑中风险与脆弱性的阈值失败')
        this.isLoading = false
        return
      }

      this.$message({
        message: res.meta.message,
        type: 'success',
      })
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