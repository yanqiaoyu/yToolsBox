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
      <el-timeline-item timestamp="2023/01/13" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 为20个脆弱性添加swag接口文档,现在点击左上角的"POC工具盒"文字,即可跳转进入swag
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/12/28" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 适配了前5个安全事件,现在也支持一键触发了
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/12/12" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 风险模块部分风险改成了需要返回敏感数据,适配一下
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/12/08" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 根据雪涛总的反馈,修复了2条脆弱性的Bug
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/11/27" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 根据杰专的反馈,更改了泄密溯源POC的逻辑,为API代理网关的POC做准备
            <br />2. 修复了一个参数名传递错误引起的Bug
            <br />3. 后端新增好了修改大脑中的时间的功能,为触发需要学习基线的模型做好准备
            <br />
            <b>4. 现在已经支持触发胡总新增的5条脆弱性</b>
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/11/21" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 合入了军子与源总一键删除脆弱性与风险日志的功能
            <br />2. 修复了一个强制审计存在的Bug
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/11/04" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 根据军子和鑫总的反馈,内置风险策略字段发生变化,"一键调整/还原阈值"功能适配3.0.11,不再适配3.0.10
            <br />
            <b>2.换了一个猫猫图 --重要!</b>
            <br />3.根据鑫总的反馈,新增了一键下发数据源并扫描的功能
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/08/07" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 根据柳立涛的反馈,新增一键修改/恢复脆弱性与风险阈值的功能
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/08/05" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 根据张鑫宇反馈,修复一个测试ssh链接的Bug
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/07/23" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 根据吴淑媚的反馈,修复4个风险的危险级别
            <br />2. 根据吴淑媚的反馈,新增一个类别过滤的功能
            <br />3. 根据柳立涛的反馈,优化了强制审计按钮说明文案
            <br />4. 根据柳立涛的反馈,新增了分类分级的POC工具
            <br />5. 根据内部需求,新增了一个可以定制请求与响应的工具
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/07/09" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 优化DSC Agent安装工具
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/07/08" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 优化敏感数据类型
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/07/07" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 修复了账号提取的一个Bug
            <br />2. 修复了一个Agent没考虑到网卡名称的问题
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/07/03" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 适配了所有多IP多地址账号相关的风险
            <br />2. 引入了tcpreplay容器
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/07/01" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 适配了所有API遍历风险
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/06/30" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 适配了融合版本的所有风险模板,但并未完全实装
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/06/24" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 新增安装DSC-Agent与安装DAS-Agent的工具
            <br />2. 适配融合版本
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/06/22" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 支持触发部分风险(需要修改一下阈值)
            <br />2. 支持强制审计日志
            <br />3. 支持分开触发脆弱性与风险
          </p>
        </el-card>
      </el-timeline-item>

      <el-timeline-item timestamp="2022/06/19" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 支持查看触发脆弱性与风险的触发记录
            <br />2. 支持在更新配置之后,查看配置的更新结果
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/06/18" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 为一些可能会对一线造成困惑的选项添加注释
            <br />2. 添加测试ssh链接功能
            <br />3. 为文件上传以及测试链接等异步操作添加遮罩,避免一线失误操作
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/06/17" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            <br />1. 一键生成DSC-Agent配置并更新工具盒中的Agent
            <br />2. 一键生成可以触发脆弱性的账号提取配置,并上传至DSC中
            <br />3. 一键触发所有脆弱性(已在最新版本中验证)
          </p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="2022/06/10" placement="top">
        <el-card>
          <h4>更新内容</h4>
          <p>
            发布面向POC测试的定制版工具盒
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
