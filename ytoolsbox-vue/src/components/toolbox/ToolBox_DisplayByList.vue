<template>
  <div>
    <!-- 任务列表展示区 -->
    <el-table :data="toolsList" stripe border style="width: 100%">
      <!-- 只要添加了type=index，就是序号列 -->
      <el-table-column type="index" label="序号" align="center" width="90"></el-table-column>
      <el-table-column prop="toolName" label="工具名称" width="270">
        <template slot-scope="scope">
          <el-link
            :underline="false"
            @click.native="toToolContent(scope.row)"
          >{{ scope.row.toolName }}</el-link>
        </template>
      </el-table-column>
      <el-table-column prop="toolDesc" label="工具介绍"></el-table-column>
      <el-table-column label="工具类型" align="center" width="200">
        <template slot-scope="scope">{{ scope.row.toolType === 'container' ? '容器' : '脚本' }}</template>
      </el-table-column>
      <el-table-column prop="toolAuthor" label="工具作者" align="center" width="200"></el-table-column>
      <el-table-column label="工具评分" align="center" width="150">
        <template slot-scope="scope">{{ scope.row.toolRate == '0' ? '暂无评分' : scope.row.toolRate }}</template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="180"></el-table-column>
    </el-table>

    <!-- 分页功能 -->
    <!-- <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="queryInfo.pagenum"
      :page-sizes="[10, 50, 100]"
      :page-size="queryInfo.pagesize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
    ></el-pagination>-->
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
}
</script>

<style>
</style>