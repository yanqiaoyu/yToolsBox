<template>
  <div>
    <!-- 搜索与清空 -->
    <el-row :gutter="20">
      <!-- 这个列里面放的是搜索框 -->
      <el-col :span="6">
        <el-input
          placeholder="请输入定时任务的名称"
          clearable
          v-model="queryInfo.query"
          @clear="getCronTasksResultsList"
          @change="getCronTasksResultsList"
        >
          <el-button slot="append" icon="el-icon-search"></el-button>
        </el-input>
      </el-col>

      <el-col :span="6">
        <el-button type="danger" @click="openClearCronTaskResultsDialog">清空所有定时任务执行结果</el-button>
      </el-col>
    </el-row>

    <!-- 分割线 -->
    <el-divider></el-divider>

    <!-- 定时任务执行结果列表展示区 -->
    <el-table :data="cronTasksResultsList" stripe border style="width: 100%">
      <!-- 只要添加了type=index，就能序号列 -->
      <el-table-column type="index" label="序号" align="center" width="90"></el-table-column>
      <el-table-column prop="cronTaskName" label="定时任务名称"></el-table-column>
      <el-table-column prop="cronTaskDesc" label="定时任务描述"></el-table-column>
      <el-table-column prop="toolName" label="选择的工具"></el-table-column>
      <el-table-column prop="toolConfigName" label="选择的配置"></el-table-column>
      <el-table-column label="新建时间" align="center" width="200">
        <template slot-scope="scope">{{ $commonFun.FormatDate(scope.row.CreatedAt) }}</template>
      </el-table-column>
      <el-table-column label="完成时间" align="center" width="200">
        <template slot-scope="scope">
          <!-- 未完成 -->
          <div v-if="scope.row.isDone == false">未完成</div>
          <!-- 已完成 -->
          <div v-else>{{ $commonFun.FormatDate(scope.row.UpdatedAt) }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="toolTaskProgress" label="任务进度" width="150">
        <template slot-scope="scope">
          <el-progress :text-inside="true" :stroke-width="26" :percentage="scope.row.progress"></el-progress>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="180">
        <template slot-scope="scope">
          <!-- 删除任务 -->
          <el-tooltip
            class="item"
            effect="dark"
            content="删除任务"
            placement="top"
            :enterable="false"
            v-if="scope.row.isDone"
          >
            <a v-if="scope.row.isDone" style="margin-right:10px; padding-top:10px">
              <el-button
                type="danger"
                icon="el-icon-delete"
                circle
                @click="DeleteTask(scope.row.ID)"
              ></el-button>
            </a>
          </el-tooltip>
          <!-- 取消任务 -->
          <el-tooltip
            class="item"
            effect="dark"
            content="取消任务"
            placement="top"
            :enterable="false"
            v-if="!scope.row.isDone"
          >
            <a v-if="!scope.row.isDone" style="margin-right:10px; padding-top:10px">
              <el-button
                type="danger"
                icon="el-icon-close"
                circle
                @click="CancelTask(scope.row.ID)"
                :disabled="scope.row.isDone"
              ></el-button>
            </a>
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
  </div>
</template>

<script>
export default {
  data() {
    return {
      // 查询定时任务需要的参数
      queryInfo: {
        // 查询字符
        query: '',
        // 这行属性其实就是当前在第几页
        pagenum: 1,
        // 这行属性其实就是当前每页展示多少条数据，这里最好与page-sizes里面的第一个元素值保持一致，否则在刷新的时候会出Bug
        pagesize: 10,
      },

      // 结果总数 用于翻页
      total: 0,

      // 定时任务执行结果
      cronTasksResultsList: [],

      // 定时任务执行结果的详情
      taskDetailDialogVisible: false,
      taskDetail: '',

      //定时器
      timer: '',
    }
  },
  created() {
    this.getCronTasksResultsList()
    // 每隔5秒更新定时任务详情
    this.timer = setInterval(this.getCronTasksResultsList, 5000)
  },
  // 销毁时记得要清理定时器
  beforeDestroy() {
    clearInterval(this.timer)
  },
  methods: {
    async getCronTasksResultsList() {
      const { data: res } = await this.$http.get('crontasksresult', {
        params: this.queryInfo,
      })
      if (res.meta.status_code !== 200)
        return this.$message.error('获取定时任务列表失败')

      this.cronTasksResultsList = res.data.cronTaskItemList
      this.total = res.data.total
    },
    openClearCronTaskResultsDialog() {
      this.$confirm('此操作将删除所有已完成任务, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          // console.log(id)
          const { data: res } = await this.$http.delete('crontasksresult')

          if (res.meta.status_code == 200) this.getCronTasksResultsList()
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
    // 任务详情
    opentaskDetailDialog(row) {
      this.taskDetailDialogVisible = true
      this.taskDetail = row.returnContent
    },

    // 翻页
    handleCurrentChange(val) {
      // console.log(`当前页: ${val}`)
      this.queryInfo.pagenum = val
      this.getCronTasksResultsList()
    },
    // 改变每页大小
    handleSizeChange(val) {
      // console.log(`每页 ${val} 条`)
      this.queryInfo.pagesize = val
      this.getCronTasksResultsList()
    },
    // 删除任务
    DeleteTask(cronTaskID) {
      console.log('删除任务:', cronTaskID)
      this.$confirm('确定删除此任务?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          // console.log(id)
          const { data: res } = await this.$http.delete(
            'crontasksresult/' + cronTaskID
          )

          if (res.meta.status_code == 200) {
            this.getCronTasksResultsList()
            return this.$message({
              type: 'success',
              message: '删除成功!',
            })
          } else {
            this.getCronTasksResultsList()
            return this.$message({
              type: 'fail',
              message: '删除失败!',
            })
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除',
          })
        })
    },
  },
}
</script>

<style>
</style>