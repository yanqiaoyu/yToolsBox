<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>工具盒</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 底部背景 -->
    <el-card>
      <!-- 已有的工具卡片 -->
      <el-card
        v-for="o in 10"
        :key="o"
        :body-style="{ padding: '0px' }"
        class="toolbox"
        shadow="hover"
      >
        <!-- <img
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          class="image"
        /> -->

        <el-image
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          class="image"
        >
        </el-image>

        <div style="padding: 10px;">
          <span>工具名称</span>
          <div class="bottom clearfix">
            <!-- <time class="time">{{ currentDate }}</time> -->
          </div>
        </div>
      </el-card>

      <!-- 新增工具的卡片 -->
      <el-card
        :body-style="{ padding: '0px' }"
        class="toolbox"
        shadow="hover"
        @click.native="addNewTool"
      >
        <!-- <el-image
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          class="image"
        >
        </el-image> -->

        <div style="padding: 14px;">
          <span>文本</span>
          <div class="bottom clearfix">
            <!-- <time class="time">{{ currentDate }}</time> -->
          </div>
        </div>
      </el-card>
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentDate: new Date()
    }
  },
  methods: {
    addNewTool() {
      console.log('新增工具')
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

.addtoolbox {
}

.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 0;
  float: right;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: '';
}

.clearfix:after {
  clear: both;
}
</style>
