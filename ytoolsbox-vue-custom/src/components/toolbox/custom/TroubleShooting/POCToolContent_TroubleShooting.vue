<template>
  <div>
    <el-empty v-if="!isDone" description="暂无故障探测记录">
      <el-button type="primary" @click="TrobouleShooting">开始故障探测</el-button>
    </el-empty>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isDone: false,
    }
  },
}
</script>

<style>
</style>