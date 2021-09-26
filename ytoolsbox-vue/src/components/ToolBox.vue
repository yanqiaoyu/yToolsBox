<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>工具盒</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 卡片信息 -->
    <el-card class="box-card">
      <el-card
        v-for="o in 10"
        :key="o"
        :body-style="{ padding: '0px' }"
        class="toolbox"
      >
        <!-- <img
          src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
          class="image"
        /> -->
        <div style="padding: 14px;">
          <span>好吃的汉堡</span>
          <div class="bottom clearfix">
            <time class="time">{{ currentDate }}</time>
            <el-button type="text" class="button">操作按钮</el-button>
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
  }
}
</script>

<style>
.toolbox {
  width: 250px;
  display: inline-block;
  margin: 10px 10px;
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
