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
      ></el-alert>

      <el-form
        :model="toolForm"
        :rules="toolRules"
        ref="addToolForm"
        label-width="130px"
      >
        <!-- 填写工具基本信息 -->
        <div
          v-if="activeIndex === 0 || activeIndex === stepList.length - 1"
          style="margin-top:10px;margin-left:20%"
        >
          <!-- 工具名称的输入框 -->
          <el-form-item label="工具名称" prop="toolName">
            <el-input
              :disabled="activeIndex == stepList.length - 1"
              v-model="toolForm.toolName"
              placeholder="给工具起个名字,名字唯一"
              style="width:450px"
            ></el-input>
          </el-form-item>

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
          <el-form-item label="联系方式" prop="toolAuthorMobile">
            <el-input
              :disabled="activeIndex == stepList.length - 1"
              v-model="toolForm.toolAuthorMobile"
              placeholder="留下作者的联系方式"
              style="width:450px"
            ></el-input>
          </el-form-item>
        </div>

        <!-- 填写工具配置信息 -->
        <div
          v-if="activeIndex === 1 || activeIndex === stepList.length - 1"
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

          <!-- 选择了脚本工具，出现这个 -->
          <div v-if="toolForm.toolType == 'script'">
            <el-form-item>
              <el-upload
                style="width:450px"
                class="upload-demo"
                ref="upload"
                :action="uploadPath"
                :on-preview="handlePreview"
                :on-remove="handleRemove"
                :file-list="fileList"
                :auto-upload="false"
                :limit="1"
                :on-change="handleBeforeUpload"
                accept=".py, .sh"
                :data="uploadDataObj"
                :headers="uploadHeaderObj"
                :on-success="handleUploadSuc"
                :on-error="handleUploadErr"
              >
                <el-button slot="trigger" size="small" type="primary"
                  >选取文件</el-button
                >
                <!-- <el-button
                  style="margin-left: 10px;"
                  size="small"
                  type="success"
                  @click="submitUpload"
                >上传到服务器</el-button>-->
                <div slot="tip" class="el-upload__tip">
                  只能上传一个py或者sh文件，且不超过10MB
                </div>
              </el-upload>
            </el-form-item>

            <el-form-item label="脚本名称" prop="toolScriptName">
              <el-input
                :disabled="true"
                v-model="toolForm.toolScriptName"
                placeholder="脚本名称"
                style="width:450px"
              ></el-input>
            </el-form-item>

            <el-form-item
              label="Python版本"
              v-if="isPythonScript"
              prop="toolPythonVersion"
            >
              <!-- Python版本下拉框 -->
              <el-select
                :disabled="activeIndex == stepList.length - 1"
                placeholder="请选择Python版本"
                v-model="toolForm.toolPythonVersion"
                style="width:450px"
              >
                <el-option label="python2" value="python2"></el-option>
                <el-option label="python2.7" value="python2.7"></el-option>
                <el-option label="python3" value="python3"></el-option>
                <el-option label="python3.6" value="python3.6"></el-option>
                <el-option label="python3.7" value="python3.7"></el-option>
                <el-option label="python3.8" value="python3.8"></el-option>
                <el-option label="python3.9" value="python3.9"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item
              label="Shell版本"
              v-if="isShellScript"
              prop="toolShellVersion"
            >
              <!-- Shell版本下拉框 -->
              <el-select
                :disabled="activeIndex == stepList.length - 1"
                placeholder="请选择Shell版本"
                v-model="toolForm.toolShellVersion"
                style="width:450px"
              >
                <el-option label="sh" value="sh"></el-option>
                <el-option label="bash" value="bash"></el-option>
              </el-select>
            </el-form-item>

            <el-form-item
              v-if="toolForm.toolExecuteLocation == 'remote'"
              label="脚本绝对路径"
              prop="toolScriptPath"
            >
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolScriptPath"
                placeholder="脚本在远程环境的绝对路径,不含文件名 eg:/tmp/"
                style="width:450px"
              ></el-input>
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
                :value="finalScriptCMD"
                style="width:450px"
                :disabled="true"
                :autosize="{ minRows: 8, maxRows: 10 }"
              ></el-input>
            </el-form-item>
          </div>

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
          <div v-else-if="toolForm.toolExecuteLocation == 'local'">
            <el-form-item label="本机IP地址" prop="toolRemoteIP">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolRemoteIP"
                placeholder="本机的IP地址"
                style="width:450px"
              ></el-input>
            </el-form-item>
            <el-form-item label="本机SSH端口" prop="toolRemoteSSH_Port">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolRemoteSSH_Port"
                placeholder="本机的端口"
                style="width:450px"
              ></el-input>
            </el-form-item>
            <el-form-item label="本机SSH账号" prop="toolRemoteSSH_Account">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolRemoteSSH_Account"
                placeholder="本机的SSH账号"
                style="width:450px"
              ></el-input>
            </el-form-item>
            <el-form-item label="本机SSH密码" prop="toolRemoteSSH_Password">
              <el-input
                :disabled="activeIndex == stepList.length - 1"
                v-model="toolForm.toolRemoteSSH_Password"
                placeholder="本机的SSH密码"
                style="width:450px"
              ></el-input>
            </el-form-item>
          </div>
        </div>
      </el-form>

      <!-- 上一步与下一步 -->
      <div style="text-align: right;">
        <el-button-group>
          <!-- 在第一步 -->
          <div v-if="activeIndex === 0">
            <el-button
              type="primary"
              icon="el-icon-arrow-left"
              @click="back2ToolBoxPage"
              >取消</el-button
            >
            <el-button type="primary" @click="nextStep">
              下一步
              <i class="el-icon-arrow-right el-icon--right"></i>
            </el-button>
          </div>

          <!-- 在最后一步 -->
          <div v-else-if="activeIndex === stepList.length - 1">
            <el-button type="primary" icon="el-icon-arrow-left" @click="preStep"
              >上一步</el-button
            >
            <el-button type="primary" @click="PostNewTool">
              确认提交
              <i class="el-icon-check"></i>
            </el-button>
          </div>

          <!-- 在中间 -->
          <div v-else>
            <el-button type="primary" icon="el-icon-arrow-left" @click="preStep"
              >上一步</el-button
            >
            <el-button type="primary" @click="nextStep">
              下一步
              <i class="el-icon-arrow-right el-icon--right"></i>
            </el-button>
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
      stepList: ['工具基础信息', '工具配置信息', '完成'],
      toolForm: {
        // 工具基本信息,放在tools表里面
        toolName: '',
        toolDesc: '',
        toolAuthor: '',
        toolAuthorMobile: '',

        // 工具配置信息,放在toolsconfig表里面
        toolType: '',
        toolDockerImageName: '',
        toolScriptName: '',
        toolScriptPath: '',
        toolOptions: '',
        toolRunCMD: '',
        toolExecuteLocation: 'local',
        toolRemoteIP: '',
        toolRemoteSSH_Port: '',
        toolRemoteSSH_Account: '',
        toolRemoteSSH_Password: '',
        toolPythonVersion: '',
        toolShellVersion: ''
      },
      toolRules: {
        toolType: [
          { required: true, message: '请选择工具类型', trigger: 'change' }
        ],
        toolDockerImageName: [
          { required: true, message: '请输入Docker镜像名称', trigger: 'blur' }
        ],
        toolScriptName: [
          { required: true, message: '请输入脚本名称', trigger: 'input' }
        ],
        toolScriptPath: [
          { required: true, message: '请输入脚本绝对路径', trigger: 'blur' }
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
        ],
        toolPythonVersion: [
          {
            required: true,
            message: '请选择Python解释器版本',
            trigger: 'blur'
          }
        ],
        toolShellVersion: [
          {
            required: true,
            message: '请选择Shell解释器版本',
            trigger: 'blur'
          }
        ]
      },

      fileList: [],
      isPythonScript: false,
      isShellScript: false,
      uploadDataObj: {
        toolName: ''
      },
      uploadHeaderObj: {
        Authorization: ''
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
        'docker run --name' +
        ' ' +
        'yToolsBox-' +
        runName +
        ' ' +
        toolOptions +
        ' ' +
        toolDockerImageName

      return toolRunCMD
    },
    finalScriptCMD() {
      let {
        toolRunCMD,
        toolScriptName,
        toolScriptPath,
        toolOptions,
        toolPythonVersion,
        toolShellVersion
      } = this.toolForm

      if (this.isPythonScript) {
        toolRunCMD =
          toolPythonVersion +
          ' ' +
          toolScriptPath +
          toolScriptName +
          ' ' +
          toolOptions
      } else {
        toolRunCMD =
          toolShellVersion +
          ' ' +
          toolScriptPath +
          toolScriptName +
          ' ' +
          toolOptions
      }

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
        .then(async () => {
          // 新建工具的请求
          this.toolForm.toolType == 'script'
            ? (this.toolForm.toolRunCMD = this.finalScriptCMD)
            : (this.toolForm.toolRunCMD = this.finalCMD)
          const { data: res } = await this.$http.post(
            'tools',
            qs.stringify({ ...this.toolForm })
          )
          // 上传文件的请求
          this.submitUpload()
          // console.log(res, uploadResult)
          if (res.meta.status_code == 200) {
            this.$message({
              type: 'success',
              message: '添加工具成功!'
            })
            this.$router.push('/toolbox')
          } else if (res.meta.message == '不允许重复的工具名称') {
            this.$message.error('工具名称重复')
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '添加工具失败'
          })
        })
    },
    back2ToolBoxPage() {
      this.$router.push('/toolbox')
    },
    async submitUpload() {
      // 上传前为请求头添加授权
      this.uploadHeaderObj.Authorization = window.sessionStorage.getItem(
        'token'
      )
      // 把工具名也传过去，方便建立唯一的文件夹存储脚本
      this.uploadDataObj.toolName = this.toolForm.toolName
      console.log(window.location.host)
      this.$refs.upload.submit()
    },
    handleRemove(file, fileList) {
      this.isPythonScript = false
      this.isShellScript = false
      this.toolForm.toolScriptName = ''
      console.log(this.isPythonScript)
      console.log(file, fileList)
    },
    handlePreview(file) {
      console.log(file)
    },
    // 过滤文件格式，文件大小
    handleBeforeUpload(file) {
      let suffix = ''
      try {
        var fileArr = file.name.split('.')
        suffix = fileArr[fileArr.length - 1]
        // console.log(suffix)
        if (suffix != 'py' && suffix != 'sh') {
          this.$message.error('上传文件类型错误')
          this.$refs.upload.clearFiles()
          return
        }
      } catch (err) {
        // console.log(err)
        this.$message.error('文件异常')
        this.$refs.upload.clearFiles()
        return
      }

      // 是py脚本，就要允许选择Python的版本
      if (suffix == 'py') {
        this.isPythonScript = true
        this.isShellScript = false
      } else if (suffix == 'sh') {
        this.isPythonScript = false
        this.isShellScript = true
      }

      // Math.ceil 向上取整
      // console.log(Math.ceil(file.size / 1024 / 1024))
      if (Math.ceil(file.size / 1024 / 1024) > 10) {
        this.$message.error('上传文件必须小于10MB')
        this.$refs.upload.clearFiles()
        return
      }

      this.toolForm.toolScriptName = file.name
    },
    // 上传失败的钩子函数
    handleUploadErr(err) {
      console.log(err)
      this.$message.error({
        message: '上传文件失败!'
      })
    },
    // 上传成功的钩子函数
    handleUploadSuc(res) {
      if (res.meta.status_code == 200) {
        this.$message({
          type: 'success',
          message: '上传文件成成功!'
        })
      }
    },
    // 返回上传地址
    uploadPath() {
      // 测试，生产环境，不同的请求的路径
      if (process.env.NODE_ENV == 'production') {
        let uploadPath =
          'http://' + window.location.hostname + '/api/auth/upload'
        return uploadPath
      } else {
        let uploadPath =
          'http://' + window.location.hostname + ':8081/api/auth/upload'
        return uploadPath
      }
    }
  }
}
</script>

<style lang="less" scoped></style>
