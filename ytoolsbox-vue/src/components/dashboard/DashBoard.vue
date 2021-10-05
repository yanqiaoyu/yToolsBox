<template>
  <div>
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
          <el-button type="primary" @click="dialogVisible = true">新建任务</el-button>
        </el-col>
      </el-row>

      <!-- 分割线 -->
      <el-divider></el-divider>

      <!-- 任务列表展示区 -->
      <el-table :data="tasksList" stripe border style="width: 100%">
        <!-- 只要添加了type=index，就能序号列 -->
        <el-table-column type="index" label="序号" width="70"></el-table-column>
        <el-table-column prop="toolName" label="工具名称"></el-table-column>
        <el-table-column prop="toolConfigName" label="配置名称"></el-table-column>
        <el-table-column prop="toolConfigDesc" label="配置描述"></el-table-column>
        <el-table-column prop="toolTaskProgress" label="任务进度">
          <el-progress :text-inside="true" :stroke-width="26" :percentage="0"></el-progress>
        </el-table-column>
        <el-table-column label="操作" align="center" width="150">
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
              <el-button type="warning" icon="el-icon-question" circle></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <!-- 添加任务的提示框 -->
      <el-dialog
        title="添加任务"
        :visible.sync="dialogVisible"
        width="30%"
        :close-on-click-modal="false"
        @close="closeDialog"
      >
        <div>
          <span class="demonstration" style="margin-right: 10px">选择配置</span>
          <el-cascader :options="options" :props="props" collapse-tags clearable size="medium"></el-cascader>
        </div>

        <span slot="footer" class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="PostAddTask">确 定</el-button>
        </span>
      </el-dialog>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      options: [
        {
          value: 1,
          label: '东南',
          children: [
            { value: 1, label: '普陀' },
            { value: 1, label: '黄埔' },
            { value: 1, label: '徐汇' },
          ],
        },
      ],
      props: { multiple: true },
      queryInfo: {
        query: '',
        params: '',
      },
      tasksList: [{ toolName: '1' }],
      dialogVisible: false,
    }
  },
  methods: {
    GetTasksList() {
      console.log('Get Task')
    },
    PostAddTask() {},
    CancelTask() {},
    closeDialog() {},
  },
}
</script>

<style lang="less" scoped>
</style>
