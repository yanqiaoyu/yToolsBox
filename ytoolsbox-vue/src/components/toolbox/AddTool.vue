<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/toolbox' }">工具盒</el-breadcrumb-item>
      <el-breadcrumb-item>添加工具</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 底部背景卡片 -->
    <el-card>
      <!-- 步骤条 -->
      <el-steps :active="activeIndex" finish-status="success" align-center>
        <el-step
          class="stepClass"
          v-for="(step, index) in stepList"
          :key="index"
          :title="step"
        ></el-step>
      </el-steps>
    </el-card>

    <!-- 分割线 -->
    <el-divider></el-divider>

    <!-- 底部信息填充区域 -->
    <!-- 表单数据 -->

    <el-card style="width:70%;margin:0 auto">
      <!-- 引导信息 -->
      <el-alert
        :title="stepList[activeIndex]"
        type="success"
        center
        :closable="false"
      >
      </el-alert>

      <el-form
        :model="toolForm"
        :rules="toolRules"
        ref="addToolForm"
        label-width="130px"
      >
        <!-- 填写工具信息 -->
        <div
          v-if="activeIndex === 0 || activeIndex === stepList.length - 1"
          style="margin-top:10px;margin-left:20%"
        >
          <el-form-item label="工具类型" prop="toolType">
            <!-- 选择工具类型下拉框 -->
            <el-select
              :disabled="activeIndex == stepList.length - 1"
              v-model="toolForm.toolType"
              placeholder="请选择工具类型"
              style="width:450px"
            >
              <el-option label="容器化工具" value="container"></el-option>
              <el-option label="脚本工具" value="script"></el-option>
            </el-select>
          </el-form-item>

          <!-- 选择了容器化工具，出现这个 -->
          <div v-if="toolForm.toolType == 'container'">
            <el-form-item label="Docker镜像名称" prop="toolDockerImageName">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolDockerImageName"
                placeholder="Docker镜像名称, eg:hello-world"
                style="width:450px"
              ></el-input>
              <el-tooltip
                class="item"
                effect="dark"
                content="这里的名称与拉取镜像挂钩 例如:填写hello-world,后续拉取时会执行 docker pull hello-world"
                placement="top-start"
              >
                <i
                  class="header-icon el-icon-info"
                  style="margin-left:10px"
                ></i>
              </el-tooltip>
            </el-form-item>
            <el-form-item label="运行参数" prop="toolOptions">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                type="textarea"
                v-model="toolForm.toolOptions"
                style="width:450px"
                :autosize="{ minRows: 8, maxRows: 10 }"
                placeholder="工具运行时需要传入的参数"
              ></el-input>
            </el-form-item>

            <el-form-item label="最终执行语句" prop="toolRunCMD">
              <el-input
                type="textarea"
                :value="finalCMD"
                style="width:450px"
                :disabled="true"
                :autosize="{ minRows: 8, maxRows: 10 }"
              ></el-input>
            </el-form-item>
          </div>

          <!-- 工具名称的输入框 -->
          <el-form-item label="工具名称" prop="toolName">
            <el-input
              :disabled="activeIndex == stepList.length - 1"
              v-model="toolForm.toolName"
              placeholder="给工具起个名字"
              style="width:450px"
            ></el-input>
          </el-form-item>

          <!-- 选择执行位置的开关 -->
          <el-form-item>
            <el-switch
              :disabled="activeIndex == stepList.length - 1"
              v-model="toolForm.toolExecuteLocation"
              active-text="远程执行"
              active-value="remote"
              inactive-text="本地执行"
              inactive-value="local"
              @change="changeLocationSwitch($event)"
            ></el-switch>
            <!-- 针对执行位置的说明提示 -->
            <el-tooltip
              class="item"
              effect="dark"
              content="选择远程执行,如果是脚本,会把脚本上传到远程环境再执行;如果是容器,那么会先ssh进入远程环境后再拉取镜像"
              placement="top-start"
            >
              <i class="header-icon el-icon-info" style="margin-left:10px"></i>
            </el-tooltip>
          </el-form-item>

          <!-- 只有选择了远程执行，才会需要填写远程信息 -->
          <div v-if="toolForm.toolExecuteLocation == 'remote'">
            <el-form-item label="SSH IP地址" prop="toolRemoteIP">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolRemoteIP"
                placeholder="远程环境的IP"
                style="width:450px"
              ></el-input>
            </el-form-item>
            <el-form-item label="SSH端口" prop="toolRemoteSSH_Port">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolRemoteSSH_Port"
                placeholder="远程环境的端口"
                style="width:450px"
              ></el-input>
            </el-form-item>
            <el-form-item label="SSH账号" prop="toolRemoteSSH_Account">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolRemoteSSH_Account"
                placeholder="远程SSH账号"
                style="width:450px"
              ></el-input>
            </el-form-item>
            <el-form-item label="SSH密码" prop="toolRemoteSSH_Password">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolRemoteSSH_Password"
                placeholder="远程SSH密码"
                style="width:450px"
              ></el-input>
            </el-form-item>
          </div>

          <!-- 工具简介的输入框 -->
          <el-form-item label="工具简介" prop="toolDesc">
            <el-input
              :disabled="activeIndex == stepList.length - 1"
              type="textarea"
              v-model="toolForm.toolDesc"
              style="width:450px"
              :autosize="{ minRows: 8, maxRows: 20 }"
              placeholder="简要的介绍一下这个工具"
            ></el-input>
          </el-form-item>
        </div>

        <!-- 填写作者信息 -->
        <div
          v-if="activeIndex === 1 || activeIndex === stepList.length - 1"
          style="margin-top:10px;margin-left:20%"
        >
          <!-- 作者名称的输入框 -->
          <el-form-item label="作者名称" prop="toolAuthor">
            <el-input
              :disabled="activeIndex == stepList.length - 1"
              v-model="toolForm.toolAuthor"
              placeholder="留下作者的名称"
              style="width:450px"
            ></el-input>
          </el-form-item>

          <!-- 作者联系方式的输入框 -->
          <el-form-item label="作者手机号" prop="toolAuthorMobile">
            <el-input
              :disabled="activeIndex == stepList.length - 1"
              v-model="toolForm.toolAuthorMobile"
              placeholder="留下作者的联系方式"
              style="width:450px"
            ></el-input>
          </el-form-item>
        </div>
      </el-form>

      <!-- 上一步与下一步 -->
      <div style="text-align: right;">
        <el-button-group>
          <div v-if="activeIndex === 0">
            <el-button disabled type="primary" icon="el-icon-arrow-left"
              >上一步</el-button
            >
            <el-button type="primary" @click="nextStep"
              >下一步<i class="el-icon-arrow-right el-icon--right"></i
            ></el-button>
          </div>

          <div v-else-if="activeIndex === stepList.length - 1">
            <el-button type="primary" icon="el-icon-arrow-left" @click="preStep"
              >上一步</el-button
            >
            <el-button type="primary" @click="PostNewTool"
              >确认提交 <i class="el-icon-check"></i
            ></el-button>
          </div>

          <div v-else>
            <el-button type="primary" icon="el-icon-arrow-left"
              >上一步</el-button
            >
            <el-button type="primary" @click="nextStep"
              >下一步<i class="el-icon-arrow-right el-icon--right"></i
            ></el-button>
          </div>
        </el-button-group>
      </div>
    </el-card>
  </div>
</template>

<script>
import qs from 'qs'
export default {
  data() {
    return {
      activeIndex: 0,
      stepList: ['工具信息', '作者信息', '完成'],
      toolForm: {
        toolType: '',
        toolDockerImageName: '',
        toolName: '',
        toolAuthor: '',
        toolAuthorMobile: '',
        toolOptions: '',
        toolRunCMD: '',
        toolDesc: '',
        toolExecuteLocation: 'local',
        toolRemoteIP: '',
        toolRemoteSSH_Port: '',
        toolRemoteSSH_Account: '',
        toolRemoteSSH_Password: ''
      },
      toolRules: {
        toolType: [
          { required: true, message: '请选择工具类型', trigger: 'change' }
        ],
        toolDockerImageName: [
          { required: true, message: '请输入Docker镜像名称', trigger: 'blur' }
        ],
        toolName: [
          { required: true, message: '请填写工具名称', trigger: 'blur' },
          { min: 0, max: 10, message: '最好控制在10个字符内', trigger: 'blur' }
        ],
        toolDesc: [
          { min: 0, max: 100, message: '100个字符内', trigger: 'blur' }
        ],
        toolRemoteIP: [
          { required: true, message: '请填写远程环境的IP', trigger: 'blur' }
        ],
        toolRemoteSSH_Port: [
          {
            required: true,
            message: '请填写远程环境的SSH端口',
            trigger: 'blur'
          }
        ],
        toolRemoteSSH_Account: [
          {
            required: true,
            message: '请填写远程环境的SSH账号',
            trigger: 'blur'
          }
        ],
        toolRemoteSSH_Password: [
          {
            required: true,
            message: '请填写远程环境的SSH密码',
            trigger: 'blur'
          }
        ],
        toolAuthor: [
          {
            required: true,
            message: '请填写作者名称',
            trigger: 'blur'
          },
          { min: 0, max: 10, message: '最好控制在10个字符内', trigger: 'blur' }
        ]
      }
    }
  },
  computed: {
    finalCMD() {
      let { toolRunCMD, toolDockerImageName, toolOptions } = this.toolForm

      // 处理一下特殊字符
      let runName = toolDockerImageName.replace('/', '')
      runName = runName.replace(':', '')
      runName = runName.replace('.', '')

      toolRunCMD =
        'docker run' +
        ' ' +
        'yToolsBox-' +
        runName +
        ' ' +
        toolOptions +
        ' ' +
        toolDockerImageName

      return toolRunCMD
    }
  },
  methods: {
    clickStep(index) {
      this.activeIndex = index
    },

    preStep() {
      if (this.activeIndex == 0) {
        return
      }

      this.activeIndex -= 1
    },
    nextStep() {
      console.log(this.$refs)
      this.$refs.addToolForm.validate(valid => {
        if (valid) {
          if (this.activeIndex == this.stepList.length - 1) {
            return
          }

          this.activeIndex += 1
        } else {
          this.$message.error('参数校验失败')
          return
        }
      })
    },
    changeLocationSwitch(e) {
      this.toolExecuteLocation = e
      // console.log(this.toolExecuteLocation)
    },
    PostNewTool() {
      this.$confirm('参数确认无误, 确认新增此工具?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {
          this.$http.post('tools', qs.stringify(this.toolForm))

          this.$message({
            type: 'success',
            message: '添加工具成功!'
          })

          this.$router.push('/toolbox')
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '添加工具失败'
          })
        })
    }
  }
}
</script>

<style lang="less" scoped></style>
