<template>
  <div>
    <!-- 面包屑路径 -->
    <el-breadcrumb>
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item :to="{ path: '/toolbox' }">工具盒</el-breadcrumb-item>
      <el-breadcrumb-item>添加工具</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 底部背景卡片 -->
    <el-card>
      <!-- 步骤条 -->
      <el-steps :active="activeIndex" finish-status="success" align-center>
        <el-step
          class="stepClass"
          v-for="(step, index) in stepList"
          :key="index"
          @click.native="clickStep(index)"
          :title="step"
        ></el-step>
      </el-steps>
    </el-card>

    <!-- 分割线 -->
    <el-divider></el-divider>

    <!-- 底部信息填充区域 -->
    <!-- 表单数据 -->
    <el-form
      :model="toolForm"
      :rules="toolRules"
      ref="addToolForm"
      label-width="100px"
    >
      <el-card style="width:70%;margin:0 auto">
        <!-- 引导信息 -->
        <el-alert
          :title="stepList[activeIndex]"
          type="success"
          center
          :closable="false"
        >
        </el-alert>

        <!-- 填写工具信息 -->
        <div v-if="activeIndex == 0" style="margin-top:10px;margin-left:10%">
          <el-form-item label="工具类型" prop="tooltype">
            <!-- 选择工具类型下拉框 -->
            <el-select
              v-model="toolForm.tooltype"
              placeholder="请选择工具类型"
              style="width:250px"
            >
              <el-option label="容器化工具" value="container"></el-option>
              <el-option label="脚本工具" value="script"></el-option>
            </el-select>

            <!-- <el-tooltip
              class="item"
              effect="dark"
              content="Top Left 提示文字"
              placement="top-start"
            >
              <i class="header-icon el-icon-info" style="margin-left:10px"></i>
            </el-tooltip> -->
          </el-form-item>
          <el-form-item label="工具名称" prop="toolname">
            <el-input
              v-model="toolForm.toolname"
              placeholder="给工具起个名字"
              style="width:250px"
            ></el-input>
          </el-form-item>
        </div>

        <!--  -->
        <div v-else-if="activeIndex == 1">1</div>

        <!-- 上一步与下一步 -->
        <div style="text-align: right;">
          <el-button-group>
            <el-button type="primary" icon="el-icon-arrow-left"
              >上一步</el-button
            >
            <el-button type="primary"
              >下一步<i class="el-icon-arrow-right el-icon--right"></i
            ></el-button>
          </el-button-group>
        </div>
      </el-card>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeIndex: 0,
      stepList: ['填写工具信息', '填写作者信息', '选择上传图片'],
      toolForm: {
        tooltype: '',
        toolname: ''
      },
      toolRules: {
        tooltype: [
          { required: true, message: '请选择工具类型', trigger: 'change' }
        ],
        toolname: [
          { required: true, message: '请填写工具名称', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    clickStep(index) {
      this.activeIndex = index
    }
  }
}
</script>

<style lang="less" scoped>
.stepClass :hover {
  cursor: pointer;
}
</style>
