<template>
  <div>
    <el-input
      type="textarea"
      placeholder="工具的使用说明"
      v-model="sonToolTutorial"
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
  // 这里的toolTutorial来自toolbox.vue的toToolContent
  props: {
    fatherToolTutorial: {
      type: String,
      default: '',
    },
  },

  data() {
    return {
      isReadOnly: true,
      textarea: '',
      sonToolTutorial: '',
    }
  },
  watch: {
    fatherToolTutorial(newVal) {
      this.sonToolTutorial = newVal
    },
  },
  created() {
    // console.log('toolTutorial:', this.fatherToolTutorial)
    this.sonToolTutorial = this.fatherToolTutorial
  },
  methods: {
    // 点击启用编辑后 禁用启用编辑按钮 同时启用文本编辑框
    enableEdit() {
      this.$confirm('此操作将启用工具使用说明的编辑, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(() => {
          this.isReadOnly = false
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '取消启用编辑',
          })
        })
    },
    // 点击保存编辑后 启用编辑按钮 同时禁用文本编辑框
    saveEdit() {
      this.$confirm('确认保存修改?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          const { data: res } = await this.$http.put('tools/tutorial', {
            toolID: 1,
            toolTutorial: this.sonToolTutorial,
          })
          if (res.meta.status_code == 200) this.isReadOnly = true
          return this.$message({
            type: 'success',
            message: '保存成功!',
          })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '取消保存更改',
          })
        })
    },
  },
}
</script>

<style lang="less" scoped>
</style>