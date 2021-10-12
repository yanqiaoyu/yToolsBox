<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/toolbox' }">工具盒</el-breadcrumb-item>
      <el-breadcrumb-item>{{ toolName }}</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs type="border-card">
      <el-tab-pane label="配置管理">
        <el-row :gutter="20">
          <!-- 这个列里面放的是搜索框 -->
          <el-col :span="6">
            <el-input
              placeholder="请输入配置名称"
              v-model="queryInfo.query"
              clearable
              @clear="GetToolConfig"
              @change="GetToolConfig"
            >
              <el-button
                slot="append"
                icon="el-icon-search"
                @click="GetToolConfig"
              ></el-button>
            </el-input>
          </el-col>

          <!-- 这个列里面放的是添加配置的框 -->
          <el-col :span="6">
            <el-button type="primary" @click="showAddConfigDialog"
              >添加配置</el-button
            >
          </el-col>
        </el-row>
        <!-- 分割线 -->
        <el-divider></el-divider>

        <!-- 工具配置展示区 -->
        <el-table :data="allConfigForm" stripe border style="width: 100%">
          <!-- 只要添加了type=index，就能序号列 -->
          <el-table-column
            type="index"
            label="序号"
            width="100"
          ></el-table-column>
          <el-table-column
            prop="toolConfigName"
            label="配置名称"
            width="450"
          ></el-table-column>
          <el-table-column
            prop="toolConfigDesc"
            label="配置描述"
            width="797"
          ></el-table-column>
          <el-table-column label="操作" width="300" align="center">
            <template slot-scope="scope">
              <!-- 编辑 -->
              <el-button
                type="primary"
                icon="el-icon-edit"
                circle
                @click="showEditToolConfig(scope.row)"
              ></el-button>
              <!-- 删除 -->
              <el-button
                :disabled="scope.row.toolConfigName === '默认配置'"
                type="danger"
                icon="el-icon-delete"
                circle
                @click="deleteToolConfig(scope.row)"
              ></el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="工具评价"> 工具评价 </el-tab-pane>
    </el-tabs>

    <!-- 查看配置详情的对话框 -->
    <el-dialog
      title="编辑配置"
      :visible.sync="editToolConfigVisible"
      width="30%"
      :close-on-click-modal="false"
      @close="closeEditDialog"
    >
      <!-- 查看配置的表单 -->
      <el-form
        :model="editConfigForm"
        :rules="editConfigFormRule"
        ref="editConfigForm"
        label-width="120px"
      >
        <!-- 编辑工具配置的弹窗界面中的表单 -->
        <el-form-item label="配置名称" prop="toolConfigName">
          <el-input
            v-model="editConfigForm.toolConfigName"
            :disabled="disableEditDefaultConfig(editConfigForm.toolConfigName)"
          ></el-input>
        </el-form-item>
        <el-form-item label="配置描述" prop="toolConfigDesc">
          <el-input
            v-model="editConfigForm.toolConfigDesc"
            :disabled="disableEditDefaultConfig(editConfigForm.toolConfigName)"
          ></el-input>
        </el-form-item>
        <el-form-item label="工具类型" prop="toolType">
          <!-- 选择工具类型下拉框 -->
          <el-select
            :disabled="disableEditDefaultConfig(editConfigForm.toolConfigName)"
            v-model="editConfigForm.toolType"
            placeholder="请选择工具类型"
          >
            <el-option label="容器化工具" value="container"></el-option>
            <el-option label="脚本工具" value="script"></el-option>
          </el-select>
        </el-form-item>

        <!-- 选择了脚本工具，出现这个 -->
        <div v-if="editConfigForm.toolType == 'script'">
          <el-form-item label="脚本名称" prop="toolScriptName">
            <el-input
              :disabled="true"
              v-model="editConfigForm.toolScriptName"
              placeholder="脚本名称"
            ></el-input>
          </el-form-item>
          <!-- python版本选择器 -->
          <el-form-item
            label="Python版本"
            v-if="editConfigForm.toolPythonVersion"
            prop="toolPythonVersion"
          >
            <!-- Python版本下拉框 -->
            <el-select
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              placeholder="请选择Python版本"
              v-model="editConfigForm.toolPythonVersion"
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
          <!-- shell版本选择器 -->
          <el-form-item
            label="Shell版本"
            v-if="editConfigForm.toolShellVersion"
            prop="toolShellVersion"
          >
            <!-- Shell版本下拉框 -->
            <el-select
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              placeholder="请选择Shell版本"
              v-model="editConfigForm.toolShellVersion"
            >
              <el-option label="sh" value="sh"></el-option>
              <el-option label="bash" value="bash"></el-option>
            </el-select>
          </el-form-item>
          <!-- 脚本位于远端服务器的位置 -->
          <el-form-item
            v-if="editConfigForm.toolExecuteLocation == 'remote'"
            label="脚本绝对路径"
            prop="toolScriptPath"
          >
            <el-input
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              v-model="editConfigForm.toolScriptPath"
              placeholder="脚本在远程环境的绝对路径,不含文件名 eg:/tmp/"
            ></el-input>
          </el-form-item>
          <!-- 运行参数 -->
          <el-form-item label="运行参数" prop="toolOptions">
            <el-input
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              type="textarea"
              v-model="editConfigForm.toolOptions"
              :autosize="{ minRows: 8, maxRows: 10 }"
              placeholder="工具运行时需要传入的参数"
            ></el-input>
          </el-form-item>
          <el-form-item label="最终执行语句" prop="toolRunCMD">
            <el-input
              type="textarea"
              :value="finalScriptCMD"
              :disabled="true"
              :autosize="{ minRows: 8, maxRows: 10 }"
            ></el-input>
          </el-form-item>
        </div>

        <!-- 选择了容器化工具，出现这个 -->
        <div v-else-if="editConfigForm.toolType == 'container'">
          <el-form-item label="Docker镜像名称" prop="toolDockerImageName">
            <el-input
              :disabled="true"
              v-model="editConfigForm.toolDockerImageName"
              placeholder="Docker镜像名称, eg:hello-world"
            ></el-input>
          </el-form-item>
          <el-form-item label="运行参数" prop="toolOptions">
            <el-input
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              type="textarea"
              v-model="editConfigForm.toolOptions"
              :autosize="{ minRows: 8, maxRows: 10 }"
              placeholder="工具运行时需要传入的参数"
            ></el-input>
          </el-form-item>

          <el-form-item label="最终执行语句" prop="toolRunCMD">
            <el-input
              type="textarea"
              :value="finalCMD"
              :disabled="true"
              :autosize="{ minRows: 8, maxRows: 10 }"
            ></el-input>
          </el-form-item>
        </div>
      </el-form>

      <!-- 底部的按钮 -->
      <span slot="footer" class="dialog-footer">
        <el-button @click="closeEditDialog">取 消</el-button>
        <el-button type="primary" @click="confirmEdit">确 定</el-button>
      </span>
    </el-dialog>

    <!-- 新增配置的对话框 -->
    <el-dialog
      title="新增配置"
      :visible.sync="addToolConfigVisible"
      width="30%"
      :close-on-click-modal="false"
    >
      <!-- 新增配置的表单 -->
      <el-form
        :model="editConfigForm"
        :rules="editConfigFormRule"
        ref="editConfigForm"
        label-width="120px"
      >
        <!-- 编辑工具配置的弹窗界面中的表单 -->
        <el-form-item label="配置名称" prop="toolConfigName">
          <el-input
            v-model="editConfigForm.toolConfigName"
            :disabled="disableEditDefaultConfig(editConfigForm.toolConfigName)"
          ></el-input>
        </el-form-item>
        <el-form-item label="配置描述" prop="toolConfigDesc">
          <el-input
            v-model="editConfigForm.toolConfigDesc"
            :disabled="disableEditDefaultConfig(editConfigForm.toolConfigName)"
          ></el-input>
        </el-form-item>
        <el-form-item label="工具类型" prop="toolType">
          <!-- 选择工具类型下拉框 -->
          <el-select
            :disabled="disableEditDefaultConfig(editConfigForm.toolConfigName)"
            v-model="editConfigForm.toolType"
            placeholder="请选择工具类型"
          >
            <el-option label="容器化工具" value="container"></el-option>
            <el-option label="脚本工具" value="script"></el-option>
          </el-select>
        </el-form-item>

        <!-- 选择了脚本工具，出现这个 -->
        <div v-if="editConfigForm.toolType == 'script'">
          <el-form-item label="脚本名称" prop="toolScriptName">
            <el-input
              :disabled="true"
              v-model="editConfigForm.toolScriptName"
              placeholder="脚本名称"
            ></el-input>
          </el-form-item>
          <!-- python版本选择器 -->
          <el-form-item
            label="Python版本"
            v-if="editConfigForm.toolPythonVersion"
            prop="toolPythonVersion"
          >
            <!-- Python版本下拉框 -->
            <el-select
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              placeholder="请选择Python版本"
              v-model="editConfigForm.toolPythonVersion"
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
          <!-- shell版本选择器 -->
          <el-form-item
            label="Shell版本"
            v-if="editConfigForm.toolShellVersion"
            prop="toolShellVersion"
          >
            <!-- Shell版本下拉框 -->
            <el-select
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              placeholder="请选择Shell版本"
              v-model="editConfigForm.toolShellVersion"
            >
              <el-option label="sh" value="sh"></el-option>
              <el-option label="bash" value="bash"></el-option>
            </el-select>
          </el-form-item>
          <!-- 脚本位于远端服务器的位置 -->
          <el-form-item
            v-if="editConfigForm.toolExecuteLocation == 'remote'"
            label="脚本绝对路径"
            prop="toolScriptPath"
          >
            <el-input
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              v-model="editConfigForm.toolScriptPath"
              placeholder="脚本在远程环境的绝对路径,不含文件名 eg:/tmp/"
            ></el-input>
          </el-form-item>
          <!-- 运行参数 -->
          <el-form-item label="运行参数" prop="toolOptions">
            <el-input
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              type="textarea"
              v-model="editConfigForm.toolOptions"
              :autosize="{ minRows: 8, maxRows: 10 }"
              placeholder="工具运行时需要传入的参数"
            ></el-input>
          </el-form-item>
          <el-form-item label="最终执行语句" prop="toolRunCMD">
            <el-input
              type="textarea"
              :value="finalScriptCMD"
              :disabled="true"
              :autosize="{ minRows: 8, maxRows: 10 }"
            ></el-input>
          </el-form-item>
        </div>

        <!-- 选择了容器化工具，出现这个 -->
        <div v-else-if="editConfigForm.toolType == 'container'">
          <el-form-item label="Docker镜像名称" prop="toolDockerImageName">
            <el-input
              :disabled="true"
              v-model="editConfigForm.toolDockerImageName"
              placeholder="Docker镜像名称, eg:hello-world"
            ></el-input>
          </el-form-item>
          <el-form-item label="运行参数" prop="toolOptions">
            <el-input
              :disabled="
                disableEditDefaultConfig(editConfigForm.toolConfigName)
              "
              type="textarea"
              v-model="editConfigForm.toolOptions"
              :autosize="{ minRows: 8, maxRows: 10 }"
              placeholder="工具运行时需要传入的参数"
            ></el-input>
          </el-form-item>

          <el-form-item label="最终执行语句" prop="toolRunCMD">
            <el-input
              type="textarea"
              :value="finalCMD"
              :disabled="true"
              :autosize="{ minRows: 8, maxRows: 10 }"
            ></el-input>
          </el-form-item>
        </div>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      toolID: 0,
      toolName: '',
      // created阶段拿到的这个工具的所有的配置信息
      allConfigForm: [],
      // 编辑配置需要的表单
      editConfigForm: {},
      // 编辑配置需要的表单验证规则
      editConfigFormRule: [],
      // 编辑配置的对话框可见度操控
      editToolConfigVisible: false,
      // 添加配置的对话框可见度操控
      addToolConfigVisible: false,
      // 新增配置的表单
      addConfigForm: [],
      // 查询配置信息的请求body
      queryInfo: { query: '' }
    }
  },
  computed: {
    finalCMD() {
      let { toolRunCMD, toolDockerImageName, toolOptions } = this.editConfigForm

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
    },
    finalScriptCMD() {
      let {
        toolRunCMD,
        toolScriptName,
        toolScriptPath,
        toolOptions,
        toolPythonVersion,
        toolShellVersion
      } = this.editConfigForm
      console.log(toolPythonVersion, toolShellVersion)
      if (toolPythonVersion) {
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
  created() {
    this.toolID = this.$route.query.toolID
    this.toolName = this.$route.query.toolName

    // 拿到这个工具的配置信息
    this.GetToolConfig()
  },
  methods: {
    // 发起请求拿到配置信息
    async GetToolConfig() {
      const { data: res } = await this.$http.get('tools/config/' + this.toolID)
      this.allConfigForm = res.data.toolsConfig

      if (res.meta.status_code != 200) {
        this.$message.error('获取工具配置信息失败')
      }
    },
    // 查看配置详情的对话框
    showEditToolConfig(toolInfo) {
      console.log(toolInfo)
      this.editToolConfigVisible = true
      // TBD 这里后续需要请求后端，返回这条配置的数据，而不是直接拿取前端的数据
      this.editConfigForm = toolInfo
    },
    // 关闭配置详情的对话框
    closeEditDialog() {
      this.editToolConfigVisible = false
    },
    confirmEdit() {
      this.editToolConfigVisible = false
      // TBD
      // 后续要发送请求确认修改
    },
    // 封装了一下针对默认配置禁止修改的条件
    disableEditDefaultConfig(configName) {
      if (configName === '默认配置') {
        return true
      } else return false
    },
    deleteToolConfig(id) {
      console.log(id)
    },
    // 展示添加配置的对话框
    showAddConfigDialog() {
      this.addToolConfigVisible = true
    }
  }
}
</script>

<style lang="less" scoped>
.el-form-item {
  margin-bottom: 22px;
  width: 450px;
}
</style>
