<template>
  <div>
    <el-input
      type="textarea"
      placeholder="工具的使用说明"
      v-model="toolTutorial"
      :autosize="{ minRows: 10, maxRows: 100}"
      :readonly="isReadOnly"
    ></el-input>
    <!-- 分割线 -->
    <el-divider></el-divider>
    <el-button type="success" :disabled="!isReadOnly" @click="enableEdit">启用编辑</el-button>
    <el-button type="primary" :disabled="isReadOnly" @click="saveEdit">保存修改</el-button>
  </div>
</template>

<script>
export default {
  props: {
    toolTutorial: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      isReadOnly: true,
      textarea: '',
    }
  },
  created() {
    // console.log('toolTutorial:', this.toolTutorial)
  },
  methods: {
    // 点击启用编辑后 禁用启用编辑按钮 同时启用文本编辑框
    enableEdit() {
      this.isReadOnly = false
    },
    // 点击保存编辑后 启用编辑按钮 同时禁用文本编辑框
    saveEdit() {
      this.isReadOnly = true
    },
  },
}
</script>

<style lang="less" scoped>
</style>