<template>
  <div>
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

    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>任务</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 底部背景卡片 -->
    <el-card>
      <!-- 搜索与添加 -->
      <el-row :gutter="20">
        <!-- 这个列里面放的是搜索框 -->
        <el-col :span="6">
          <el-input
            placeholder="请输入任务对应的工具名称"
            clearable
            v-model="queryInfo.query"
            @clear="GetTasksList"
            @change="GetTasksList"
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>

        <!-- 这个列里面放的是添加按钮 -->
        <el-col :span="6">
          <el-button type="primary" @click="openDialog">新建任务</el-button>
          <el-button type="danger" @click="openClearTaskDialog">清空任务</el-button>
        </el-col>
      </el-row>

      <!-- 分割线 -->
      <el-divider></el-divider>

      <!-- 任务列表展示区 -->
      <el-table :data="tasksList" stripe border style="width: 100%">
        <!-- 只要添加了type=index，就能序号列 -->
        <el-table-column type="index" label="任务序号" align="center" width="90"></el-table-column>
        <el-table-column prop="toolName" label="选择的工具"></el-table-column>
        <el-table-column prop="toolConfigName" label="选择的配置"></el-table-column>
        <el-table-column label="新建时间" align="center" width="200">
          <!-- <template slot-scope="scope">{{ scope.row.addTime | formatDate }}</template> -->
          <template slot-scope="scope">{{ FomatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="完成时间" align="center" width="200">
          <template slot-scope="scope">
            <!-- 未完成 -->
            <div v-if="scope.row.isDone == false">未完成</div>
            <!-- 已完成 -->
            <div v-else>{{ FomatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="toolTaskProgress" label="任务进度" width="150">
          <template slot-scope="scope">
            <el-progress :text-inside="true" :stroke-width="26" :percentage="scope.row.progress"></el-progress>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center" width="180">
          <template slot-scope="scope">
            <!-- 取消任务 -->
            <el-tooltip
              class="item"
              effect="dark"
              content="取消任务"
              placement="top"
              :enterable="false"
            >
              <el-button
                type="danger"
                icon="el-icon-close"
                circle
                @click="cancelTask(scope.row.id)"
                :disabled="scope.row.isDone"
              ></el-button>
            </el-tooltip>
            <!-- 再次执行 -->
            <el-tooltip
              class="item"
              effect="dark"
              content="再次执行任务"
              placement="top"
              :enterable="false"
            >
              <el-button
                type="success"
                icon="el-icon-refresh-left"
                circle
                @click="restartTask(scope.row.id)"
                :disabled="!scope.row.isDone"
              ></el-button>
            </el-tooltip>
            <!-- 任务详情 -->
            <el-tooltip
              class="item"
              effect="dark"
              content="任务执行详情"
              placement="top"
              :enterable="false"
            >
              <el-button
                type="warning"
                icon="el-icon-question"
                @click="opentaskDetailDialog(scope.row)"
                circle
              ></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <!-- 添加任务的提示框 -->
      <el-dialog
        title="添加任务"
        :visible.sync="dialogVisible"
        width="550px"
        :close-on-click-modal="false"
        @close="closeDialog"
        v-loading="loading"
        element-loading-text="创建任务中"
        element-loading-spinner="el-icon-loading"
        element-loading-background="rgba(0, 0, 0, 0.8)"
      >
        <!-- 任务配置的级联选择器 -->
        <div>
          <span style="margin-right: 20px">选择配置</span>
          <el-cascader
            v-model="finalList"
            :options="options"
            :props="props"
            collapse-tags
            clearable
            size="medium"
            filterable
            placeholder="搜索工具名称或配置名称"
          ></el-cascader>
        </div>

        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="PostAddTask">确 定</el-button>
        </span>
      </el-dialog>

      <!-- 任务详情的提示框 -->
      <el-dialog
        title="任务详情"
        :visible.sync="taskDetailDialogVisible"
        width="50%"
        :close-on-click-modal="false"
        @close="taskDetailDialogVisible = false"
      >
        <div style="white-space: pre-line;">{{ taskDetail }}</div>
        <span slot="footer" class="dialog-footer">
          <el-button @click="taskDetailDialogVisible = false">关闭</el-button>
        </span>
      </el-dialog>

      <!-- 分页功能 -->
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.pagenum"
        :page-sizes="[10, 50, 100]"
        :page-size="queryInfo.pagesize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total"
      ></el-pagination>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 与级联选择器绑定的列表
      finalList: [],
      options: [],
      props: { multiple: true },
      queryInfo: {
        query: '',
        // 这行属性其实就是当前在第几页
        pagenum: 1,
        // 这行属性其实就是当前每页展示多少条数据，这里最好与page-sizes里面的第一个元素值保持一致，否则在刷新的时候会出Bug
        pagesize: 10,
      },
      tasksList: [],
      total: 0,
      dialogVisible: false,
      taskDetailDialogVisible: false,
      loading: false,
      configIDList: [],
      taskDetail: '',
      //定时器
      timer: '',
    }
  },
  created() {
    // 要获取任务的列表
    this.GetTasksList()
    // 每隔10秒定时更新任务详情
    this.timer = setInterval(this.GetTasksList, 10000)
  },
  // 销毁时记得要清理定时器
  beforeDestroy() {
    clearInterval(this.timer)
  },
  methods: {
    // 获取任务列表的请求
    async GetTasksList() {
      // console.log('Get Task')
      const { data: res } = await this.$http.get('tasks', {
        params: this.queryInfo,
      })
      if (res.meta.status_code !== 200)
        return this.$message.error('获取任务列表失败')

      this.tasksList = res.data.TaskItemList
      this.total = res.data.Total
    },
    // 获取级联选择器中的信息的请求
    async GetCascaderList() {
      console.log('打开了工具配置窗口')
      const { data: res } = await this.$http.get('tasks/cascader')
      console.log('所有工具的配置如下', res)
      if (res.meta.status_code !== 200)
        return this.$message.error('获取配置信息失败')

      this.options = res.data.CascaderList
    },
    // 确认新增一个任务
    async PostAddTask() {
      this.loading = true
      // 拿到所有配置ID，存放到一个List中
      for (var i in this.finalList) {
        this.configIDList.push(this.finalList[i]['1'])
      }

      console.log('选中的配置如下', this.configIDList)

      if (this.finalList.length == 0) {
        this.loading = false
        return this.$message.error('未选择配置信息,无法创建任务')
      }
      const { data: res } = await this.$http.post('tasks', {
        ConfigList: JSON.stringify(this.configIDList),
      })
      if (res.meta.status_code !== 200) {
        this.loading = false
        return this.$message.error('获取配置信息失败')
      }
      this.loading = false
      this.$message.success('新建任务成功')
      // 成功了关闭对话框
      this.dialogVisible = false
      // 清空Cascader的选中条目
      this.finalList = []
      // 清空configIDList
      this.configIDList = []
    },
    // 取消任务
    CancelTask() {},
    // 打开新增任务的对话框
    openDialog() {
      // 清空Cascader的选中条目
      this.finalList = []
      this.dialogVisible = true
      // 打开的时候，请求工具以及对应的工具配置
      this.GetCascaderList()
    },
    openClearTaskDialog() {
      this.$confirm('此操作将删除所有已完成任务, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          // console.log(id)
          const { data: res } = await this.$http.delete('tasks')

          if (res.meta.status_code == 200) this.GetTasksList()
          return this.$message({
            type: 'success',
            message: '删除成功!',
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除',
          })
        })
    },
    // 关闭新增任务的对话框
    closeDialog() {
      // 关闭的时候，请求一次任务列表
      this.GetTasksList()
      this.dialogVisible = false
    },
    // 转换时间戳
    FomatDate(unixtime) {
      return this.$moment.unix(unixtime).format('YYYY-MM-DD HH:mm:ss')
    },
    // 翻页
    handleCurrentChange(val) {
      // console.log(`当前页: ${val}`)
      this.queryInfo.pagenum = val
      this.GetTasksList()
    },
    // 改变每页大小
    handleSizeChange(val) {
      // console.log(`每页 ${val} 条`)
      this.queryInfo.pagesize = val
      this.GetTasksList()
    },
    // 任务详情
    opentaskDetailDialog(row) {
      this.taskDetailDialogVisible = true
      this.taskDetail = row.returnContent
    },
  },
}
</script>

<style lang="less" scoped>
</style>
