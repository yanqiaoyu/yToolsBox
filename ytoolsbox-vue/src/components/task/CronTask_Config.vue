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

    <!-- 新增配置的对话框 -->
    <el-dialog
      title="新增定时任务"
      :visible.sync="addCronTaskDialogVisible"
      width="550px"
      :close-on-click-modal="false"
      @close="closeCronTaskDialog"
      element-loading-text="创建任务中"
      element-loading-spinner="el-icon-loading"
      element-loading-background="rgba(0, 0, 0, 0.8)"
    >
      <!-- 新增配置的表单 -->
      <el-form
        :model="addCronTaskForm"
        :rules="addCronTaskFormRule"
        ref="addCronTaskForm"
        label-width="120px"
      >
        <el-form-item label="定时任务名称" prop="cronTaskName">
          <el-input v-model="addCronTaskForm.cronTaskName"></el-input>
        </el-form-item>

        <el-form-item label="定时任务描述" prop="cronTaskDesc">
          <el-input v-model="addCronTaskForm.cronTaskDesc"></el-input>
        </el-form-item>

        <el-form-item label="配置选择" prop="cronTaskFinalList">
          <TaskCascader
            :my-width="'335px'"
            :final-list.sync="finalList"
            :options="options"
            @deliverOptions="deliverOptions"
          ></TaskCascader>
        </el-form-item>
      </el-form>

      <!-- 底部的按钮 -->
      <span slot="footer" class="dialog-footer">
        <el-button @click="closeCronTaskDialog">取 消</el-button>
        <el-button type="primary" @click="confirmAddCronTask">新增定时任务</el-button>
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
  props: {
    // 父组件传过来的查询所有定时任务的功能
    getCronTasksList: {
      type: Function,
      default: null,
    },
  },
  data() {
    return {
      // 与级联选择器绑定的列表
      finalList: [],
      options: [],

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
        cronTaskFinalList: '',
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
      },
    }
  },
  watch: {
    finalList(newVal) {
      this.addCronTaskForm.cronTaskFinalList = newVal
    },
  },
  methods: {
    // 接收子组件传过来的options
    deliverOptions(data) {
      this.options = data
    },

    // 打开创建定时任务的对话框
    openCreateCronTaskDialog() {
      this.addCronTaskDialogVisible = true
    },
    // 关闭创建定时任务的对话框
    closeCronTaskDialog() {
      this.addCronTaskDialogVisible = false
    },
    // 确认添加定时任务
    confirmAddCronTask() {
      this.$refs.addCronTaskForm.validate(async (valid) => {
        if (valid) {
          this.loading = true
          // console.log(this.addCronTaskForm, this.finalList, this.options)

          this.loading = false
          // 成功了关闭对话框
          this.addCronTaskDialogVisible = false
          // 清空Cascader的选中条目
          this.finalList = []

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
  },
}
</script>

<style lang="less" scoped>
.el-form-item {
  margin-bottom: 22px;
  margin-right: 55px;
}
</style>