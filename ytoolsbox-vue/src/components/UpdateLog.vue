<template>
  <div class="block">
    <!-- 回到顶部 -->
    <el-backtop>
      <div
        style="{
        height: 100%;
        width: 100%;
        background-color: #f2f5f6;
        box-shadow: 0 0 6px rgba(0,0,0, .12);
        text-align: center;
        line-height: 40px;
        color: #1989fa;
      }"
      >
        UP
      </div>
    </el-backtop>

    <el-timeline>
      <el-timeline-item timestamp="2021/9/22" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>1. 完成了删除用户的功能 <br /></p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/21" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成了用户状态更新 <br />
            2. 解决了一个更新状态后数据库顺序紊乱的问题<br />
            3. 完成了修改用户信息的功能 <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/20" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成添加新用户的表单 <br />
            2. 配合dto和dao使用, 后端的数据库逻辑已完成<br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/15" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 编写好了用户相关的后端service层代码 <br />
            2. 准备引入PostgreSQL，并着手解耦后端代码 <br />
            3. 加入了一个更新日志界面 <br />
            4. 用ele的avatar组件重新导入了首页的猫猫 <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/14" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 广州一哥解决了一个elementUI引起的Menu收起导致图标显示不全的问题
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/10" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>1. 优化后端的架构，复写了前端需要用到的接口 <br /></p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/6" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 添加gin的后端代码 <br />
            2. 前端的主页部分绘制完成 <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/3" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. init工具盒项目 <br />
            2. 搭建好vue相关环境 <br />
          </p>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </div>
</template>

<script>
export default {}
</script>

<style></style>
