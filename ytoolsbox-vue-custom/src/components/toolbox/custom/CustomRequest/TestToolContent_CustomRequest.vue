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
      <el-breadcrumb-item>{{ '测试工具-自定义请求器' }}</el-breadcrumb-item>
    </el-breadcrumb>

    <el-tabs type="border-card">
      <el-tab-pane label="配置">
        <POCToolConten_Config></POCToolConten_Config>
      </el-tab-pane>
      <el-tab-pane label="发送请求">
        <el-collapse v-model="activeNames">
          <el-collapse-item class="titleStyle" title="自定义请求方向" name="request">
            <el-form ref="requestForm" :model="requestForm" align-center>
              <el-form-item label-width="80px" label="请求URL" style="width:1500px">
                <el-input
                  placeholder="自定义的路径"
                  v-model="requestForm.customRequestPath"
                  clearable
                  style="width:500px"
                >
                  <template slot="prepend">{{ requestForm.customRequestURL }}</template>
                </el-input>
                <el-input
                  v-model="requestForm.customRequestQuery"
                  placeholder="自定义的参数"
                  style="width:500px"
                ></el-input>
              </el-form-item>
              <el-form-item label-width="80px" label="请求方式" style="width:450px">
                <el-select v-model="requestForm.customRequestMethod" placeholder="请选择">
                  <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                    :disabled="item.disabled"
                  ></el-option>
                </el-select>
              </el-form-item>

              <el-form-item>
                <el-checkbox v-model="requestForm.customRequestWhetherWithCookie">是否携带cookie</el-checkbox>
              </el-form-item>
              <el-form-item
                label-width="80px"
                label="Cookie"
                style="width:450px"
                v-show="requestForm.customRequestWhetherWithCookie"
              >
                <el-input v-model="requestForm.customRequestCookie" placeholder="请输入Cookie内容"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="success" @click="DoRequest()">发起请求</el-button>
                <el-button type="danger" @click="ForceAudit()">强制审计</el-button>
              </el-form-item>
            </el-form>
          </el-collapse-item>

          <el-collapse-item class="titleStyle" title="自定义响应方向" name="response">
            <el-form ref="responseForm" :model="responseForm" align-center>
              <el-form-item label-width="80px" label="响应内容" style="width:450px">
                <el-input
                  type="textarea"
                  :autosize="{ minRows: 20}"
                  placeholder="请输入你希望响应的内容"
                  v-model="responseForm.customResponseContent"
                ></el-input>
              </el-form-item>
            </el-form>
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

      activeNames: ['request', 'response'],

      requestForm: {
        customRequestURL: '',
        customRequestPath: 'default',
        customRequestQuery: 'key=value',
        customRequestMethod: 'Get',
        customRequestWhetherWithCookie: false,
        customRequestCookie: '',
      },
      responseForm: {
        customResponseContent: '',
      },

      options: [
        {
          value: 'Get',
          label: 'Get',
        },
        {
          value: 'Post',
          label: 'Post',
        },
        {
          value: 'Delete',
          label: 'Delete',
          disabled: true,
        },
        {
          value: 'Put',
          label: 'Put',
          disabled: true,
        },
        {
          value: 'Option',
          label: 'Option',
          disabled: true,
        },
      ],
      value: '',
    }
  },
  created() {
    // 拿取浏览器中的IP地址
    this.requestForm.customRequestURL =
      'http://' +
      window.location.host.split(':', 1)[0] +
      '/api/auth/custom/request/'
  },
  methods: {
    DoRequest() {
      this.loadingContent = '发送请求中'
      this.isLoading = true
      const xhr = new XMLHttpRequest()

      var encodeContent = ''

      if (this.requestForm.customRequestWhetherWithCookie == true) {
        document.cookie = this.requestForm.customRequestCookie
      }

      if (this.responseForm.customResponseContent != '') {
        encodeContent = Buffer.from(
          encodeURIComponent(this.responseForm.customResponseContent).replace(
            /%([0-9A-F]{2})/g,
            function toSolidBytes(match, p1) {
              return String.fromCharCode('0x' + p1)
            }
          ),
          'binary'
        ).toString('base64')
      }

      if (process.env.NODE_ENV == 'production') {
        xhr.open(
          this.requestForm.customRequestMethod,
          this.requestForm.customRequestURL +
            this.requestForm.customRequestPath +
            '?' +
            this.requestForm.customRequestQuery +
            '&needResponseData=' +
            encodeContent
        )
      } else {
        xhr.open(
          this.requestForm.customRequestMethod,
          'http://localhost:8081/api/auth/custom/request/' +
            this.requestForm.customRequestPath +
            '?' +
            this.requestForm.customRequestQuery +
            '&needResponseData=' +
            encodeContent
        )
      }
      xhr.send()

      this.isLoading = false

      xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
          //注意这里的responseText就是响应回来的数据
          console.log(xhr.responseText)
        }
      }

      document.cookie = 'expires=' + new Date(0).toUTCString()

      return this.$message({
        type: 'success',
        message: '发送请求成功!',
      })
    },
    ForceAudit() {
      this.$confirm('确定强制立刻进行审计并生成日志?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          // const { data: res } = await this.$http.delete('tasks/')
          this.loadingContent = '强制审计中'
          this.isLoading = true

          const { data: res } = await this.$http.get('custom/forceaudit')

          this.isLoading = false

          if (res.meta.status_code == 200) {
            return this.$message({
              type: 'success',
              message: '强制审计成功!',
            })
          } else {
            return this.$message.error(res.meta.message)
          }
        })
        .catch(() => {
          this.isLoading = false
          this.$message({
            type: 'info',
            message: '已取消强制审计',
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
