<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>工具盒</el-breadcrumb-item>
    </el-breadcrumb>

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

    <!-- 底部背景卡片 -->
    <el-card>
      <!-- 搜索与添加 -->
      <el-row :gutter="20">
        <!-- 这个列里面放的是搜索框 -->
        <el-col :span="6">
          <el-input
            placeholder="请输入工具名称"
            clearable
            v-model="queryInfo.query"
            @clear="GetToolsList"
            @change="GetToolsList"
          >
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>

        <!-- 这个列里面放的是添加按钮 -->
        <el-col :span="6">
          <el-button type="primary" @click="addTool">添加工具</el-button>
          <el-button type="danger" @click="clearAllTool">清除所有工具</el-button>
        </el-col>

        <!-- 这个列里面放的是切换展示模式的按钮 -->
        <el-col :span="3.5" :offset="8" style="width: 260px; float: right; text-align: right;">
          <el-radio-group v-model="tabDefault" @change="handleRadioClick">
            <el-radio-button label="卡片模式">卡片模式</el-radio-button>
            <el-radio-button label="列表模式">列表模式</el-radio-button>
          </el-radio-group>
        </el-col>
      </el-row>

      <!-- 分割线 -->
      <el-divider></el-divider>

      <!-- 卡片模式 -->
      <DisplayByCard
        v-if="tabDefault == '卡片模式'"
        :tools-list="toolsList"
        :add-tool="addTool"
        :to-tool-content="toToolContent"
      ></DisplayByCard>

      <!-- 列表模式 -->
      <DisplayByList
        v-else-if="tabDefault == '列表模式'"
        :tools-list="toolsList"
        :add-tool="addTool"
        :to-tool-content="toToolContent"
      ></DisplayByList>
    </el-card>
  </div>
</template>

<script>
import DisplayByCard from './ToolBox_DisplayByCard'
import DisplayByList from './ToolBox_DisplayByList'

export default {
  components: { DisplayByCard, DisplayByList },
  data() {
    return {
      currentDate: new Date(),
      value: 0,
      toolsList: [],
      totalTools: 0,
      queryInfo: {
        query: '',
      },
      tabDefault: '卡片模式',
    }
  },
  created() {
    this.GetToolsList()
  },
  methods: {
    // 添加工具,跳转页面
    addTool() {
      this.$router.push('/toolbox/add')
    },
    // 请求所有的工具信息
    async GetToolsList() {
      const { data: res } = await this.$http.get('tools', {
        params: this.queryInfo,
      })
      // console.log(res)
      if (res.meta.status_code !== 200)
        return this.$message.error('获取工具列表失败')

      // 成功了就开始赋值
      this.toolsList = res.data.tools
      this.totalTools = res.data.total
    },
    // 删除所有工具
    clearAllTool() {
      this.$confirm(
        '此操作将删除所有工具,所有配置,以及所有已完成的任务, 是否继续?',
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          // console.log(id)
          const { data: res } = await this.$http.delete('tools')

          if (res.meta.status_code == 200) this.GetToolsList()
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
    // 工具详情跳转
    toToolContent(tool) {
      // console.log(tool.toolTutorial)
      this.$router.push({
        // path: 'toolbox/tool',
        name: 'toolbox_tool',
        query: {
          toolID: tool.id,
          toolName: tool.toolName,
        },
        // props: { toolTutorial: tool.toolTutorial },
        // 新增了工具使用说明的功能 这里要传过去
        // 如果提供了 path，params 会被忽略
        // 所以把path换成了name 当然router那里也要配置一下
        params: { toolTutorial: tool.toolTutorial },
      })
    },
    // 处理tab点击事件
    handleRadioClick(tab) {
      this.tabDefault = tab
      console.log('切换到', tab)
      console.log('现在默认模式为', this.tabDefault)
      // 在切换的时候, 进行一次请求
      this.GetToolsList()
    },
  },
}
</script>

<style lang="less" scoped>
</style>
