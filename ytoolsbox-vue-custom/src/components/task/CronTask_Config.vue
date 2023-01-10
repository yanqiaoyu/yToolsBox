<template>
  <div>
    <!-- 搜索与添加 -->
    <el-row :gutter="20">
      <!-- 这个列里面放的是搜索框 -->
      <el-col :span="6">
        <el-input
          placeholder="请输入定时任务的名称"
          clearable
          v-model="queryInfo.query"
          @clear="getCronTasksList"
          @change="getCronTasksList"
        >
          <el-button slot="append" icon="el-icon-search"></el-button>
        </el-input>
      </el-col>

      <!-- 这个列里面放的是添加按钮 -->
      <el-col :span="6">
        <el-button type="primary" @click="openCreateCronTaskDialog">新建定时任务</el-button>
        <el-button type="danger" @click="openClearCronTaskDialog">清空所有定时任务</el-button>
      </el-col>
    </el-row>

    <!-- 分割线 -->
    <el-divider></el-divider>

    <!-- 定时任务列表展示区 -->
    <el-table :data="cronTasksList" stripe border style="width: 100%">
      <!-- 只要添加了type=index，就能序号列 -->
      <el-table-column type="index" label="序号" align="center" width="90"></el-table-column>
      <el-table-column prop="cronTaskName" label="定时任务名称"></el-table-column>
      <el-table-column prop="cronTaskDesc" label="定时任务描述"></el-table-column>
      <el-table-column prop="cronTaskTime" label="定时语句"></el-table-column>
      <el-table-column label="新建时间" align="center" width="200">
        <template slot-scope="scope">{{ $commonFun.FormatDate(scope.row.CreatedAt) }}</template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="180">
        <template slot-scope="scope">
          <!-- 删除任务 -->
          <el-tooltip
            class="item"
            effect="dark"
            content="删除定时任务"
            placement="top"
            :enterable="false"
          >
            <el-button
              type="danger"
              icon="el-icon-delete"
              circle
              @click="deleteCronTask(scope.row.ID, scope.row.cronTaskScheduleID)"
            ></el-button>
          </el-tooltip>
          <!-- 编辑任务 -->
          <el-tooltip
            class="item"
            effect="dark"
            content="编辑定时任务"
            placement="top"
            :enterable="false"
          >
            <el-button
              type="primary"
              icon="el-icon-edit"
              circle
              @click="showEditCronTaskDialog(scope.row.cronTaskScheduleID)"
            ></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <!-- 新增定时任务的对话框 -->
    <el-dialog
      title="新增定时任务"
      :visible.sync="addCronTaskDialogVisible"
      width="550px"
      :close-on-click-modal="false"
      @close="closeCronTaskDialog"
      v-loading="loading"
      element-loading-text="创建任务中"
      element-loading-spinner="el-icon-loading"
      element-loading-background="rgba(0, 0, 0, 0.8)"
      destroy-on-close
    >
      <!-- destroy-on-close -->
      <!-- 新增定时任务的表单 -->
      <el-form
        :model="addCronTaskForm"
        :rules="addCronTaskFormRule"
        ref="addCronTaskForm"
        label-width="120px"
      >
        <el-form-item label="定时任务名称" prop="cronTaskName">
          <el-input v-model="addCronTaskForm.cronTaskName" clearable></el-input>
        </el-form-item>

        <el-form-item label="定时任务描述" prop="cronTaskDesc">
          <el-input v-model="addCronTaskForm.cronTaskDesc" clearable></el-input>
        </el-form-item>

        <el-form-item label="工具配置选择" prop="cronTaskFinalList">
          <TaskCascader
            :avoid-clear-cascader-bug="avoidClearCascaderBug"
            :my-width="'335px'"
            :final-list.sync="finalList"
            :options="options"
            @deliverOptions="deliverOptions"
          ></TaskCascader>
        </el-form-item>

        <el-form-item label="定时语句" prop="cronTaskTime">
          <el-input v-model="addCronTaskForm.cronTaskTime" class="timeSpanClass" clearable></el-input>
          <el-tooltip class="item" effect="dark" placement="top-start">
            <div slot="content">
              *(秒) *(分) *(时) *(日) *(月) *(星期)
              <br />
              <br />按时间间隔执行
              <br />每隔5秒执行一次：*/5 * * * * ?
              <br />每隔1分钟执行一次：0 */1 * * * ?
              <br />
              <br />按时间点执行
              <br />每天23点执行一次：0 0 23 * * ?
              <br />每天凌晨1点执行一次：0 0 1 * * ?
              <br />每月1号凌晨1点执行一次：0 0 1 1 * ?
              <br />在26分、29分、33分执行一次：0 26,29,33 * * * ?
              <br />每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
            </div>
            <i class="header-icon el-icon-info" style="margin-left:10px"></i>
          </el-tooltip>
        </el-form-item>

        <el-form-item>
          <el-checkbox v-model="addCronTaskForm.cronRunAtOnce" checked>添加任务后立即执行一次</el-checkbox>
        </el-form-item>
      </el-form>

      <!-- 底部的按钮 -->
      <span slot="footer" class="dialog-footer">
        <el-button @click="closeCronTaskDialog">取 消</el-button>
        <el-button type="primary" @click="confirmAddCronTask">新增定时任务</el-button>
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

    <!-- 修改定时任务的提示框 -->
    <el-dialog
      title="修改定时任务"
      :visible.sync="editCronTaskDialogVisible"
      width="550px"
      :close-on-click-modal="false"
      @close="closeEditCronTaskDialog"
      destroy-on-close
    >
      <el-form
        :model="editCronTaskForm"
        :rules="editCronTaskFormRule"
        ref="editCronTaskForm"
        label-width="120px"
      >
        <el-form-item label="定时任务名称" prop="cronTaskName">
          <el-input v-model="editCronTaskForm.cronTaskName" clearable></el-input>
        </el-form-item>
        <el-form-item label="定时任务描述" prop="cronTaskDesc">
          <el-input v-model="editCronTaskForm.cronTaskDesc" clearable></el-input>
        </el-form-item>
        <el-form-item label="工具配置选择">
          <!-- 这里没有复用那个Cascader组件 是因为有一个Bug查了一天没解决...... -->
          <el-cascader
            v-model="editCronTaskForm.cronTaskCascaderSelectedInfo"
            :options="editCronTaskForm.cronTaskCascaderAllInfo"
            :props="{ multiple: true }"
            collapse-tags
            clearable
            size="medium"
            filterable
            placeholder="搜索工具名称或配置名称"
            style="width: 335px"
          ></el-cascader>
        </el-form-item>
        <el-form-item label="定时语句" prop="cronTaskTime">
          <el-input v-model="editCronTaskForm.cronTaskTime" class="timeSpanClass" clearable></el-input>
          <el-tooltip class="item" effect="dark" placement="top-start">
            <div slot="content">
              *(秒) *(分) *(时) *(日) *(月) *(星期)
              <br />
              <br />按时间间隔执行
              <br />每隔5秒执行一次：*/5 * * * * ?
              <br />每隔1分钟执行一次：0 */1 * * * ?
              <br />
              <br />按时间点执行
              <br />每天23点执行一次：0 0 23 * * ?
              <br />每天凌晨1点执行一次：0 0 1 * * ?
              <br />每月1号凌晨1点执行一次：0 0 1 1 * ?
              <br />在26分、29分、33分执行一次：0 26,29,33 * * * ?
              <br />每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
            </div>
            <i class="header-icon el-icon-info" style="margin-left:10px"></i>
          </el-tooltip>
        </el-form-item>
        <el-form-item>
          <el-checkbox v-model="addCronTaskForm.cronRunAtOnce" checked disabled>添加任务后立即执行一次</el-checkbox>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="closeEditCronTaskDialog">取 消</el-button>
        <el-button type="primary" @click="editCronTask">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import TaskCascader from './Task_Cascader.vue'

export default {
  components: {
    TaskCascader,
  },
  data() {
    return {
      // 与新增对话框中的级联选择器绑定的列表
      finalList: [],
      options: [],
      avoidClearCascaderBug: 0,

      // 与编辑对话框中的级联选择器绑定的列表
      editFinalList: [],
      editOptions: [],
      editAvoidClearCascaderBug: 0,

      // 存储查询结果
      cronTasksList: [],
      total: 0,

      // 查询定时任务需要的参数
      queryInfo: {
        // 查询字符
        query: '',
        // 这行属性其实就是当前在第几页
        pagenum: 1,
        // 这行属性其实就是当前每页展示多少条数据，这里最好与page-sizes里面的第一个元素值保持一致，否则在刷新的时候会出Bug
        pagesize: 10,
      },

      // 新增定时任务的对话框能见度
      addCronTaskDialogVisible: false,
      // 新增定时任务的表单
      addCronTaskForm: {
        cronTaskName: '',
        cronTaskDesc: '',
        cronTaskFinalList: [],
        cronTaskTime: '0 */1 * * * ?',
        cronRunAtOnce: true,
      },
      // 新增定时任务的表单验证规则
      addCronTaskFormRule: {
        cronTaskName: [
          { required: true, message: '请输入定时任务名称', trigger: 'blur' },
          {
            min: 1,
            max: 30,
            message: '长度在 1 到 30 个字符',
            trigger: 'blur',
          },
          //   { validator: checkConfigName, trigger: 'blur' },
        ],
        cronTaskFinalList: [
          { required: true, message: '请选择任务配置', trigger: 'blur' },
        ],
        cronTaskTime: [
          { required: true, message: '请选择填写定时语句', trigger: 'blur' },
        ],
      },

      // 编辑定时任务的对话框能见度
      editCronTaskDialogVisible: false,
      // 编辑定时任务的表单
      editCronTaskForm: {
        cronTaskName: '',
        cronTaskDesc: '',
        cronTaskTime: '0 */1 * * * ?',
        cronRunAtOnce: true,
        cronTaskCascaderAllInfo: '',
        cronTaskCascaderSelectedInfo: [],
      },
      // 编辑定时任务的表单的规则
      editCronTaskFormRule: {
        cronTaskName: [
          { required: true, message: '请输入定时任务名称', trigger: 'blur' },
          {
            min: 1,
            max: 30,
            message: '长度在 1 到 30 个字符',
            trigger: 'blur',
          },
          //   { validator: checkConfigName, trigger: 'blur' },
        ],
        cronTaskFinalList: [
          { required: true, message: '请选择任务配置', trigger: 'blur' },
        ],
        cronTaskTime: [
          { required: true, message: '请选择填写定时语句', trigger: 'blur' },
        ],
      },

      // 遮罩控制
      loading: false,
    }
  },
  watch: {
    finalList(newVal) {
      this.addCronTaskForm.cronTaskFinalList = newVal
    },
  },
  created() {
    this.getCronTasksList()
  },
  methods: {
    // 接收新增子组件传过来的options
    deliverOptions(data) {
      this.options = data
    },
    // 接收编辑子组件传过来的options
    editDeliverOptions(data) {
      this.editOptions = data
    },
    // 打开创建定时任务的对话框
    openCreateCronTaskDialog() {
      this.addCronTaskDialogVisible = true
    },
    // 关闭创建定时任务的对话框
    closeCronTaskDialog() {
      // 关闭的时候请求一次列表
      this.getCronTasksList()
      this.$refs.addCronTaskForm.resetFields()
      // https://blog.csdn.net/qq_34451048/article/details/106198550
      ++this.avoidClearCascaderBug
      this.finalList = []
      this.options = []
      this.addCronTaskDialogVisible = false
    },
    // 确认添加定时任务
    confirmAddCronTask() {
      this.$refs.addCronTaskForm.validate(async (valid) => {
        if (valid) {
          this.loading = true

          // 拿到所有配置ID，存放到一个List中
          // 这里需要清空一下
          this.addCronTaskForm.cronTaskFinalList = []
          for (var i in this.finalList) {
            this.addCronTaskForm.cronTaskFinalList.push(this.finalList[i]['1'])
          }

          const { data: res } = await this.$http.post('crontasks', {
            cronTaskName: this.addCronTaskForm.cronTaskName,
            cronTaskDesc: this.addCronTaskForm.cronTaskDesc,
            cronTaskFinalList: JSON.stringify(
              this.addCronTaskForm.cronTaskFinalList
            ),
            cronTaskTime: this.addCronTaskForm.cronTaskTime,
            cronRunAtOnce: this.addCronTaskForm.cronRunAtOnce,
            cronTaskCascaderAllInfo: JSON.stringify(this.options),
            cronTaskCascaderSelectedInfo: JSON.stringify(this.finalList),
          })
          if (res.meta.status_code !== 200) {
            this.loading = false
            return this.$message.error('创建定时任务失败')
          }

          this.loading = false
          // 成功了关闭对话框
          this.addCronTaskDialogVisible = false
          // 清空Cascader的选中条目
          // 将来看到这里，不要奇怪，这个++是为了解决一个清空级联选择器的Bug
          // https://blog.csdn.net/qq_34451048/article/details/106198550
          ++this.avoidClearCascaderBug
          this.finalList = []
          this.options = []

          this.$message.success('添加成功')
          this.dialogVisible = false
        } else {
          this.$message.error('添加信息验证失败')
        }
      })
    },
    // 清除所有定时任务
    openClearCronTaskDialog() {
      this.$confirm('此操作将删除所有定时任务, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          const { data: res } = await this.$http.delete('crontasks')

          if (res.meta.status_code == 200) this.getCronTasksList()
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

    // 打开 编辑定时任务详情的框
    async showEditCronTaskDialog(cronTaskScheduleID) {
      const { data: res } = await this.$http.get(
        'crontasks/' + cronTaskScheduleID
      )
      res.data.cronTaskCascaderAllInfo = JSON.parse(
        res.data.cronTaskCascaderAllInfo
      )

      res.data.cronTaskCascaderSelectedInfo = JSON.parse(
        res.data.cronTaskCascaderSelectedInfo
      )

      this.editCronTaskForm = res.data
      this.editCronTaskDialogVisible = true
      console.log('获取特定的定时任务的细节: ', this.editCronTaskForm)
    },
    // 关闭 编辑定时任务详情的框
    closeEditCronTaskDialog() {
      this.editCronTaskDialogVisible = false
    },
    // 确认 编辑定时任务
    editCronTask() {
      this.editCronTaskDialogVisible = false
    },

    // 查询所有定时任务
    async getCronTasksList() {
      const { data: res } = await this.$http.get('crontasks', {
        params: this.queryInfo,
      })
      if (res.meta.status_code !== 200)
        return this.$message.error('获取定时任务列表失败')

      console.log('查询定时任务结果: ', res.data)

      this.cronTasksList = res.data.cronTaskItemList
      this.total = res.data.total
    },

    // 删除定时任务
    deleteCronTask(cronTaskOriginID, cronTaskScheduleID) {
      console.log('删除定时任务:', cronTaskScheduleID)
      this.$confirm('确定删除此定时任务?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          // console.log(id)
          const { data: res } = await this.$http.delete(
            'crontasks/' + cronTaskOriginID + '/' + cronTaskScheduleID
          )

          if (res.meta.status_code == 200) {
            this.getCronTasksList()
            return this.$message({
              type: 'success',
              message: '删除成功!',
            })
          } else {
            this.getCronTasksList()
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

    // 翻页
    handleCurrentChange(val) {
      // console.log(`当前页: ${val}`)
      this.queryInfo.pagenum = val
      this.getCronTasksList()
    },
    // 改变每页大小
    handleSizeChange(val) {
      // console.log(`每页 ${val} 条`)
      this.queryInfo.pagesize = val
      this.getCronTasksList()
    },
  },
}
</script>

<style lang="less" scoped>
.el-form-item {
  margin-bottom: 22px;
  margin-right: 55px;
}

.timeSpanClass {
  width: 90%;
}
</style>