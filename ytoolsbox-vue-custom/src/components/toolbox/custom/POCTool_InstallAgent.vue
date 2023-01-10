<template>
  <!-- 探针部署工具 -->
  <el-card
    :body-style="{ padding: '0px' }"
    class="toolbox"
    shadow="hover"
    @click.native="toRiskandVulnerability()"
  >
    <div class="image_container">
      <el-image fit="scale-down" class="image" :src="require('../../../assets/POC.png')"></el-image>
    </div>

    <div class="text_container">
      <div class="rate_title">
        <span class="tool-title" title="POC工具-探针部署">{{ "POC工具-探针部署" }}</span>
        <!-- 否则显示暂无评分 -->
        <div>
          <span>暂无评分</span>
        </div>
      </div>

      <el-divider class="line"></el-divider>

      <div class="toolsDesc">
        <el-tooltip
          class="item"
          effect="dark"
          content="协助完成DSC-Agent与DAS-Agent的部署"
          placement="bottom"
        >
          <span>{{ "协助完成DSC-Agent与DAS-Agent的部署" }}</span>
        </el-tooltip>
      </div>

      <div class="toolTags">
        <el-tag type="info">{{ "POC" }}</el-tag>
        <el-tag type="success">作者:{{ "颜桥宇" }}</el-tag>
      </div>
    </div>
  </el-card>
</template>

<script>
export default {
  methods: {
    toRiskandVulnerability() {
      this.$router.push({
        name: 'poc_tool_install_agent',
      })
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
  width: 200px; /*一定要设置宽度，或者元素内含的百分比*/
  white-space: nowrap; /*文本不换行*/
  text-overflow: ellipsis; /*ellipsis:文本溢出显示省略号（...）；clip：不显示省略标记（...），而是简单的裁切*/
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

.line.el-divider--horizontal {
  width: 90%;
  margin: 10px auto;
}

.toolTags {
  height: 50px;
  width: 95%;
  position: relative;
  text-align: right;
  margin-top: 10px;
}
</style>