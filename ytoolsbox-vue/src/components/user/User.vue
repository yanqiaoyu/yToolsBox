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
          <el-button type="primary">添加用户</el-button>
        </el-col>
      </el-row>

      <!-- 用户列表展示区 -->
      <el-table :data="userList" stripe border style="width: 100%">
        <!-- 只要添加了type=index，就能序号列 -->
        <el-table-column type="index" label="序号" width="60">
        </el-table-column>
        <el-table-column prop="username" label="姓名" width="180">
        </el-table-column>
        <el-table-column prop="email" label="邮箱" width="180">
        </el-table-column>
        <el-table-column prop="mobile" label="电话"> </el-table-column>
        <el-table-column prop="RoleName" label="角色"> </el-table-column>
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
  </div>
</template>

<script>
export default {
  data() {
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
      total: 0
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
      // console.log(res)
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
    }
  }
}
</script>

<style lang="less" scoped></style>
