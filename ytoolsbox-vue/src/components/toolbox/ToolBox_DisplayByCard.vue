<template>
  <div>
    <!-- 已有的工具卡片 -->
    <el-card
      v-for="tool in toolsList"
      :key="tool.id"
      :body-style="{ padding: '0px' }"
      class="toolbox"
      shadow="hover"
      @click.native="toToolContent(tool)"
    >
      <div v-if="tool.toolType === 'container'" class="image_container">
        <el-image :src="require('../../assets/container.png')" class="image" fit="scale-down"></el-image>
      </div>
      <div v-if="tool.toolType === 'script'" class="image_container">
        <el-image :src="require('../../assets/script.png')" class="image" fit="scale-down"></el-image>
      </div>

      <div class="text_container">
        <div class="rate_title">
          <span class="tool-title" :title="tool.toolName">{{ tool.toolName }}</span>
          <!-- 只有被评分过的工具，才会显示评分 -->
          <div v-if="tool.toolRate != 0 || tool.toolRateCount != 0">
            <el-tooltip
              class="item"
              effect="dark"
              :content="
                '平均分:' +
                  tool.toolRate +
                  '分 ' +
                  '评分人数:' +
                  tool.toolRateCount +
                  '人'
              "
              placement="top"
            >
              <el-rate
                v-model="tool.toolRate"
                disabled
                text-color="#ff9900"
                score-template="{value}"
                class="el-rate"
              ></el-rate>
            </el-tooltip>
          </div>
          <!-- 否则显示暂无评分 -->
          <div v-else>
            <span>暂无评分</span>
          </div>
        </div>

        <el-divider class="line"></el-divider>
        <div class="toolsDesc">
          <el-tooltip class="item" effect="dark" :content="tool.toolDesc" placement="bottom">
            <span>{{ tool.toolDesc }}</span>
          </el-tooltip>
        </div>
        <div class="toolTags">
          <el-tag type="info">{{ tool.toolType === 'container' ? '容器' : '脚本' }}</el-tag>
          <el-tag type="success">作者:{{ tool.toolAuthor }}</el-tag>
        </div>
      </div>
    </el-card>

    <!-- 新增工具的卡片 -->
    <el-card
      :body-style="{ padding: '0px' }"
      class="toolbox"
      shadow="hover"
      @click.native="addTool()"
    >
      <div class="image_container">
        <el-image fit="scale-down" class="image" :src="require('../../assets/add.png')"></el-image>
      </div>

      <div style="padding: 14px;">
        <span class="add-title">添加工具</span>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  props: {
    // 父组件传过来的所有工具信息
    toolsList: {
      type: Array,
      default: function () {
        return []
      },
    },
    // 父组件传过来的新增工具功能
    addTool: {
      type: Function,
      default: null,
    },
    // 父组件传过来的跳转至工具详情的功能
    toToolContent: {
      type: Function,
      default: null,
    },
  },
  methods: {},
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