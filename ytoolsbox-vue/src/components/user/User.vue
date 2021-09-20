<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>全局配置</el-breadcrumb-item>
      <el-breadcrumb-item>用户管理</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 卡片信息 -->
    <el-card class="box-card">
      <el-row :gutter="20">
        <!-- 这个列里面放的是搜索框 -->
        <el-col :span="6">
          <el-input placeholder="请输入内容">
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>

        <!-- 这个列里面放的是添加用户的框 -->
        <el-col :span="6">
          <el-button type="primary" @click="dialogVisible = true"
            >添加用户</el-button
          >
        </el-col>
      </el-row>

      <!-- 用户列表展示区 -->
      <el-table :data="userList" stripe border style="width: 100%">
        <!-- 只要添加了type=index，就能序号列 -->
        <el-table-column type="index" label="序号" width="60">
        </el-table-column>
        <el-table-column prop="username" label="账户名" width="210">
        </el-table-column>
        <el-table-column prop="email" label="邮箱" width="430">
        </el-table-column>
        <el-table-column prop="mobile" label="电话" width="430">
        </el-table-column>
        <el-table-column prop="role" label="角色" width="140">
        </el-table-column>
        <!-- 根据实际的数据渲染出一个状态按钮 -->
        <el-table-column label="状态" width="100">
          <template slot-scope="scope">
            <!-- {{ scope.row.mgstate }} -->
            <el-switch v-model="scope.row.mgstate"> </el-switch>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template>
            <!-- 编辑 -->
            <el-button type="primary" icon="el-icon-edit" circle></el-button>
            <!-- 删除 -->
            <el-button type="danger" icon="el-icon-delete" circle></el-button>
            <!-- 分配角色 -->
            <el-tooltip
              class="item"
              effect="dark"
              content="分配角色"
              placement="top"
              :enterable="false"
            >
              <el-button type="warning" icon="el-icon-user" circle></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.pagenum"
        :page-sizes="[5, 10]"
        :page-size="queryInfo.pagesize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
      >
      </el-pagination>
    </el-card>

    <!-- 添加用户的提示框 -->
    <el-dialog
      title="添加用户"
      :visible.sync="dialogVisible"
      width="30%"
      :close-on-click-modal="false"
      @close="closeDialog"
    >
      <el-form
        :model="ruleForm"
        :rules="rules"
        ref="ruleForm"
        label-width="60px"
      >
        <el-form-item label="账户" prop="name">
          <el-input v-model="ruleForm.name"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="ruleForm.password"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="ruleForm.email"></el-input>
        </el-form-item>
        <el-form-item label="工号" prop="workNum">
          <el-input v-model="ruleForm.workNum"></el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="Phone">
          <el-input v-model="ruleForm.Phone"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addUser">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    var checkEmail = (rule, value, callback) => {
      const regEmail = /^([a-zA-Z]|[0-9])(\w|-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/
      if (regEmail.test(value)) {
        return callback()
      }
      callback(new Error('邮箱格式错误'))
    }

    var checkPhone = (rule, value, callback) => {
      const regPhone = /^1(3|4|5|7|8)\d{9}$/
      if (regPhone.test(value)) {
        return callback()
      }
      callback(new Error('手机号格式错误'))
    }

    return {
      // 获取用户的参数列表
      queryInfo: {
        query: '',
        // 这行属性其实就是当前在第几页
        pagenum: 1,
        // 这行属性其实就是当前每页展示多少条数据，这里最好与page-sizes里面的第一个元素值保持一致，否则在刷新的时候会出Bug
        pagesize: 5
      },
      userList: [],
      total: 0,
      // 决定添加用户框是否弹出
      dialogVisible: false,
      ruleForm: {
        name: '',
        password: '',
        email: '',
        workNum: ''
      },
      rules: {
        name: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 2, max: 10, message: '长度在 2 到 10 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 6, max: 10, message: '长度在 6 到 10 个字符', trigger: 'blur' }
        ],
        email: [
          { required: false, message: '请输入邮箱', trigger: 'blur' },
          {
            min: 1,
            max: 30,
            message: '长度在 1 到 30 个字符',
            trigger: 'blur'
          },
          { validator: checkEmail, trigger: 'blur' }
        ],
        workNum: [
          { required: false, message: '请输入工号', trigger: 'blur' },
          { min: 5, max: 10, message: '长度在 5 到 10 个字符', trigger: 'blur' }
        ],
        Phone: [
          { required: false, message: '请输入手机号', trigger: 'blur' },
          {
            min: 10,
            max: 18,
            message: '长度在 10 到 18 个字符',
            trigger: 'blur'
          },
          { validator: checkPhone, trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.GetUsersList()
  },
  methods: {
    async GetUsersList() {
      // console.log("获取用户列表被调用了")
      const { data: res } = await this.$http.get('users', {
        params: this.queryInfo
      })
      console.log(res)
      // 获取用户失败
      if (res.meta.status_code !== 200)
        return this.$message.error('获取用户信息失败')

      // 成功了就开始赋值
      this.userList = res.data.users
      this.total = res.data.total
    },
    handleCurrentChange(val) {
      // console.log(`当前页: ${val}`)
      this.queryInfo.pagenum = val
      this.GetUsersList()
    },
    handleSizeChange(val) {
      // console.log(`每页 ${val} 条`)
      this.queryInfo.pagesize = val
      this.GetUsersList()
    },
    closeDialog() {
      this.$refs.ruleForm.resetFields()
    },
    addUser() {
      // console.log('添加用户')
      this.$refs.ruleForm.validate(valid => {
        // console.log(valid)
        if (valid) {
          this.$message.success('添加成功')
          this.dialogVisible = false
        } else {
          this.$message.error('添加信息验证失败')
        }
      })
      // this.dialogVisible = false
    }
  }
}
</script>

<style lang="less" scoped></style>
