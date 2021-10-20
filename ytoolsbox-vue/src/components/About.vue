<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>关于</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>
      <about />
    </el-card>
  </div>
</template>

<script>
import about from '../assets/doc/about.md'
// 代码高亮
import 'highlight.js/styles/github.css'
// 其他元素使用 github 的样式
import 'github-markdown-css'

export default {
  components: {
    about
  }
}
</script>

<style></style>
