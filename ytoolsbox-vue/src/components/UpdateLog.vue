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
      >UP</div>
    </el-backtop>

    <el-timeline>
      <el-timeline-item timestamp="2022/04/02" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <el-link
              href="https://github.com/yanqiaoyu/yToolsBox/issues/38"
              target="_blank"
            >1. 完成【需求】当想在某个工具下新增配置时, 应该允许从已有的配置中duplicate一下</el-link>
            <br />现在在工具的配置界面,可以直接从已有的配置复制一条新的配置出来了
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/04/01" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成【需求】对于已完成的任务，应该把'取消任务'的按钮，改成'再执行一次' #26
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/03/31" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 解决了Bug: 【前端】新增配置界面的表单，无法随着浏览器百分比变化而变化 #25
            <br />2. 开始开发: 【需求】对于已完成的任务，应该把‘取消任务’的按钮，改成‘再执行一次’ #26
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/03/29" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 解决了Bug: 【前端】下发任务的时候,偶然会出现一次性下发2次任务的情况 #27
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/26" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 解决了一个SSH建立连接时Ciphers异常引起的panic
            <br />2. 现在已完成的任务可以被清空了
            <br />3. 现在可以清空所有工具了
            <br />4. 为任务界面添加了异步的定时请求,现在任务执行情况可以实时更新了
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/25" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 统一侧边栏的色调
            <br />2. 隐藏未实现的权限管理模块
            <br />3. 让熊少帮忙解决了一个计算属性引起的Bug
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/21" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 现在脚本也支持本地执行了
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/20" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 现在可以修改配置中的执行位置与位置信息了
            <br />2. 现在docker支持远程和本地执行了
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/19" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成了前端，后端代码的打包，编译，docker化
            <br />2. 解决了一个nginx配置错误导致请求跨域的问题
            <br />3. 解决了el-upload上传不经过axios从而导致跨域的问题
            <br />4. 完善了中文readme文档，发布0.1版本
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/14" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>1. 完成配置界面的修改，删除，过滤搜索</p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/11" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 修复了一个脚本上传失败引起的后端崩溃的Bug
            <br />2.完成了容器化工具的默认配置展示功能
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/10" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 修复了一个容器执行失败引起的后端崩溃的Bug
            <br />2. 利用goroutine+channel实现了同步完成任务执行与进度更新
            <br />3. 完成了任务详情的展示
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/7" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 修复了解析配置ID异常的Bug
            <br />2. 修复了上传脚本但是执行语句是容器的Bug
            <br />3.
            完成了文件上传的功能
            <br />4.
            完成了工具执行的功能，至此关键功能基本OK，后续开始进行完善和改进
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/6" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 复用用户管理模块, 完成了任务模块相关的增删改查
            <br />2. 参考gorm的官方文档，解决了一个时间显示的问题
            <br />3.
            使用v-loading规避了一个重复请求的Bug，提升用户体验
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/5" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 尝试用goroutine并发新增任务
            <br />2. copy网上关于go-ssh的代码并融合使用
            <br />3.
            完成了前端任务界面的样式展示
            <br />4.
            选用级联选择器作为发起任务的前端样式，并解决了高度不一致的问题
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/4" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 解决一个主键重复的Bug
            <br />2.工具详情->工具配置界面已经可以正常显示
            <br />3.初步完成配置的编辑样式,暂时搁置编辑配置的功能,开发最为重要的后端执行工具逻辑
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/3" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 修复了一个错误处理的Bug
            <br />2.
            优化工具的库表结构，现在分为工具基本信息与工具配置信息两张表
            <br />3. 配合后端的优化，优化前端的新增工具的步骤
            <br />4.
            改用内联查询优化后端的查询逻辑
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/10/2" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成了工具详情展示的可携带参数跳转
            <br />2. 完成了工具搜索的功能
            <br />3. 优化了后端查询工具的逻辑
            <br />4. 优化了新增工具的一些逻辑与样式
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/30" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成了前端添加工具的内容
            <br />2. 完成了后端添加工具的逻辑
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/27" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 开始进行工具盒核心功能的编写
            <br />2. 在熊少的帮助下，改善了工具盒的布局与样式
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/26" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成了所有功能的模型绑定重写
            <br />2. 后端新增权限管理
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/23" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 参考gin的官方文档，使用模型绑定复写了参数解析的业务
            <br />2. 重写登录的逻辑
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/22" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成了删除用户的功能
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/21" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成了用户状态更新
            <br />2. 解决了一个更新状态后数据库顺序紊乱的问题
            <br />3.
            完成了修改用户信息的功能
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/20" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 完成添加新用户的表单
            <br />2. 配合dto和dao使用, 后端的数据库逻辑已完成
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/15" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 编写好了用户相关的后端service层代码
            <br />2. 准备引入PostgreSQL，并着手解耦后端代码
            <br />3.
            加入了一个更新日志界面
            <br />4.
            用ele的avatar组件重新导入了首页的猫猫
            <br />
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
          <p>
            1. 优化后端的架构，复写了前端需要用到的接口
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/6" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. 添加gin的后端代码
            <br />2. 前端的主页部分绘制完成
            <br />
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2021/9/3" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            1. init工具盒项目
            <br />2. 搭建好vue相关环境
            <br />
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
