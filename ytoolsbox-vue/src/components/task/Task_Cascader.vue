<template>
  <div class="myCasDiv">
    <el-cascader
      :key="avoidClearCascaderBug"
      v-model="tmpFinalList"
      :options="options"
      :props="props"
      collapse-tags
      clearable
      size="medium"
      filterable
      placeholder="搜索工具名称或配置名称"
      :style="{width: myWidth}"
      :ref="myCascaderRef"
    ></el-cascader>
  </div>
</template>

<script>
export default {
  props: {
    // 父组件传过来的,与级联选择器相关的内容
    finalList: {
      type: Array,
      default: function () {
        return []
      },
    },
    options: {
      type: Array,
      default: function () {
        return []
      },
    },
    myWidth: {
      type: String,
      default: function () {
        return ''
      },
    },
    avoidClearCascaderBug: {
      type: Number,
      default: function () {
        return 0
      },
    },
    myCascaderRef: {
      type: String,
      default: function () {
        return ''
      },
    },
  },
  data() {
    return {
      props: { multiple: true },
      tmpFinalList: [],
    }
  },
  watch: {
    tmpFinalList(newVal) {
      this.$emit('update:final-list', newVal)
    },
    immediate: true,
  },
  created() {
    this.tmpFinalList = this.finalList
    console.log('级联选择器已经被创建')
    this.GetCascaderList()
  },
  beforeDestroy() {
    console.log('级联选择器即将被销毁')
  },

  methods: {
    // 获取级联选择器中的信息的请求
    async GetCascaderList() {
      console.log('打开了工具配置窗口')
      const { data: res } = await this.$http.get('tasks/cascader')
      console.log('工具及其配置的级联选择器信息如下', res)
      if (res.meta.status_code !== 200)
        return this.$message.error('获取配置信息失败')

      //   this.options = res.data.CascaderList
      this.$emit('deliverOptions', res.data.CascaderList)
    },
  },
}
</script>

<style lang="less" scoped>
.myCasDiv {
  display: inline-block;
  .el-cascader-menu__list {
    width: 30px !important;
  }
}
</style>