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
          <el-input placeholder="请输入工具名称" clearable>
            <el-button slot="append" icon="el-icon-search"></el-button>
          </el-input>
        </el-col>

        <!-- 这个列里面放的是添加按钮 -->
        <el-col :span="6">
          <el-button type="primary" @click="addTool">添加工具</el-button>
        </el-col>
      </el-row>

      <!-- 分割线 -->
      <el-divider></el-divider>

      <!-- 已有的工具卡片 -->
      <el-card
        v-for="tool in toolsList"
        :key="tool.id"
        :body-style="{ padding: '0px' }"
        class="toolbox"
        shadow="hover"
      >
        <div class="image_container">
          <el-image :src="require('../../assets/users.png')" class="image" fit="scale-down"></el-image>
        </div>

        <div class="text_container">
          <div class="rate_title">
            <span class="tool-title" :title="tool.toolName">{{ tool.toolName }}</span>
            <el-rate
              v-model="value"
              disabled
              text-color="#ff9900"
              score-template="{value}"
              class="el-rate"
            ></el-rate>
          </div>

          <el-divider class="line"></el-divider>
          <div class="toolsDesc">
            <el-tooltip class="item" effect="dark" :content="tool.toolDesc" placement="bottom">
              <span>{{ tool.toolDesc }}</span>
            </el-tooltip>
          </div>
          <div class="toolTags">
            <el-tag type="info">{{ tool.toolType }}</el-tag>
            <el-tag type="success">作者:{{ tool.toolAuthor }}</el-tag>
          </div>
        </div>
      </el-card>

      <!-- 新增工具的卡片 -->
      <el-card
        :body-style="{ padding: '0px' }"
        class="toolbox"
        shadow="hover"
        @click.native="addTool"
      >
        <div class="image_container">
          <el-image fit="scale-down" class="image" :src="require('../../assets/add.png')"></el-image>
        </div>

        <div style="padding: 14px;">
          <span class="add-title">添加工具</span>
        </div>
      </el-card>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentDate: new Date(),
      value: 3.7,
      toolsList: [],
      totalTools: 0,
    }
  },
  created() {
    this.GetToolsList()
  },
  methods: {
    addTool() {
      this.$router.push('/toolbox/add')
    },
    async GetToolsList() {
      const { data: res } = await this.$http.get('tools')
      console.log(res)
      if (res.meta.status_code !== 200)
        return this.$message.error('获取工具列表失败')

      // 成功了就开始赋值
      this.toolsList = res.data.tools
      this.totalTools = res.data.total
    },
  },
}
</script>

<style lang="less" scoped>
.toolbox {
  width: 310px;
  height: 350px;
  display: inline-block;
  margin: 6px 6px;
}

.image_container {
  height: 225px;
  width: 310px;
}

.image {
  width: 100%;
  height: 100%;
  display: block;
}

.tool-title {
  font-weight: bold;
  height: 21px;
  display: inline-block;
  overflow: hidden;
  font-size: 18px;
  width: 150px; /*一定要设置宽度，或者元素内含的百分比*/
  white-space: nowrap; /*文本不换行*/
  text-overflow: ellipsis; /*ellipsis:文本溢出显示省略号（...）；clip：不显示省略标记（...），而是简单的裁切*/
}

.line.el-divider--horizontal {
  width: 90%;
  margin: 10px auto;
}

.rate_title {
  width: 90%;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
}

.add-title {
  font-weight: bold;
  height: 21px;
  width: 30%;
  margin: 0 108px;
  display: block;
}

.star-rate {
  display: inline-block;
}

.el-rate {
  padding-top: 2px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: '';
}

.clearfix:after {
  clear: both;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.toolsDesc {
  font-size: 12px;
  margin-left: 16px;
  margin-right: 8px;
  height: 20px;
  width: 277px;
  white-space: nowrap;
  display: inline-block;
  overflow: hidden;
  text-overflow: ellipsis;
}

// 官方文档给出的tag之间设置空隙的方法
.el-tag + .el-tag {
  margin-left: 10px;
}

.toolTags {
  height: 50px;
  width: 95%;
  position: relative;
  text-align: right;
  margin-top: 10px;
}
</style>
