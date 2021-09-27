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
      >
        UP
      </div>
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
        v-for="o in 5"
        :key="o"
        :body-style="{ padding: '0px' }"
        class="toolbox"
        shadow="hover"
      >
        <div class="image_container">
          <el-image
            src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
            class="image"
            fit="scale-down"
          >
          </el-image>
        </div>

        <div class="text_container">
          <div class="rate_title">
            <span class="tool-title">工具名称</span>
            <el-rate
              v-model="value"
              disabled
              text-color="#ff9900"
              score-template="{value}"
            >
              >
            </el-rate>
          </div>

          <el-divider class="line"></el-divider>
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
          <el-image class="image"> </el-image>
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
      value: 3.7
    }
  },
  methods: {
    addTool() {
      console.log('新增工具')
      this.$router.push('/toolbox/add')
    }
  }
}
</script>

<style>
.toolbox {
  width: 255px;
  height: 350px;
  display: inline-block;
  margin: 6px 6px;
}

/* .text_container {
  display: flex;
} */

.image_container {
  height: 225px;
  width: 255px;
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
  margin: 0 auto;
  display: block;
}

.star-rate {
  display: inline-block;
}
</style>
