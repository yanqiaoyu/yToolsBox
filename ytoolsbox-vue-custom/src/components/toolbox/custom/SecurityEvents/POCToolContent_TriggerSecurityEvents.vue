<template>
  <div
    v-loading.fullscreen.lock="isLoading"
    :element-loading-text="loadingContent"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <el-row :gutter="20">
      <!-- 这个列里面放的是搜索框 -->
      <el-col :span="6">
        <el-input
          placeholder="请输入安全事件名称"
          clearable
          v-model="queryInfo.query"
          @clear="GetSecurityEventsList"
          @change="GetSecurityEventsList"
        >
          <el-button slot="append" icon="el-icon-search"></el-button>
        </el-input>
      </el-col>

      <!-- 这个列里面放的是触发全部脆弱性与风险的框 -->
      <el-col :span="8">
        <el-button type="primary" @click="TriggerAll">全部触发</el-button>

        <el-tooltip class="item" effect="dark" placement="top-start">
          <div slot="content">强制容器进行审计,实现日志落盘</div>
          <el-button type="danger" @click="ForceAudit">立即审计</el-button>
        </el-tooltip>
      </el-col>
    </el-row>
    <!-- 分割线 -->
    <el-divider></el-divider>

    <el-table
      :data="securityEventsList"
      stripe
      border
      style="width: 100%"
      @filter-change="filterChangeInit"
    >
      <el-table-column type="index" label="序号" width="50" align="center"></el-table-column>
      <el-table-column
        label="类别"
        width="100"
        align="center"
        prop="type"
        column-key="type"
        :filters="[{ text: '安全事件', value: '安全事件' }]"
        :filter-method="filterTag"
        :filter-multiple="false"
      >
        <template slot-scope="scope">
          <el-tag v-if="scope.row.type == '安全事件'">{{ scope.row.type }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="name" label="名称" width="200"></el-table-column>
      <el-table-column prop="desc" label="描述"></el-table-column>
      <el-table-column label="危险级别" width="80" align="center">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.level=='高危'" type="danger" effect="dark">{{ scope.row.level }}</el-tag>
          <el-tag v-if="scope.row.level=='中危'" type="warning" effect="dark">{{ scope.row.level }}</el-tag>
          <el-tag v-if="scope.row.level=='低危'" effect="dark">{{ scope.row.level }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150" align="center">
        <template slot-scope="scope">
          <!-- 触发安全事件 -->
          <el-tooltip
            class="item"
            effect="dark"
            content="触发该安全事件"
            placement="top"
            :enterable="false"
          >
            <a style="margin-right:10px; padding-top:10px">
              <el-button
                type="success"
                icon="el-icon-aim"
                circle
                @click="TriggerThis(scope.row.name)"
              ></el-button>
            </a>
          </el-tooltip>
          <!-- 触发方式介绍 -->
          <el-tooltip
            class="item"
            effect="dark"
            content="触发方式介绍"
            placement="top"
            :enterable="false"
          >
            <el-button
              type="warning"
              icon="el-icon-question"
              @click="OpenTriggerMethodDialog(scope.row.triggermethod)"
              circle
            ></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页功能 -->
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="queryInfo.pagenum"
      :page-sizes="[10, 50, 100]"
      :page-size="queryInfo.pagesize"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
    ></el-pagination>

    <el-dialog
      title="触发方法介绍"
      :visible.sync="triggerMethodDialogVisible"
      width="30%"
      :close-on-click-modal="false"
      @close="triggerMethodDialogVisible = false"
    >
      <div style="white-space: pre-line;">{{ triggerMethod }}</div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="triggerMethodDialogVisible = false">关闭</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
// import qs from 'qs'
import { faker } from '@faker-js/faker'
// faker.setLocale('zh_CN')
// import restoreservertime from '@util/restoreservertime'

export default {
  data() {
    return {
      queryInfo: {
        // 查询字符
        query: '',
        // 这行属性其实就是当前在第几页
        pagenum: 1,
        // 这行属性其实就是当前每页展示多少条数据，这里最好与page-sizes里面的第一个元素值保持一致，否则在刷新的时候会出Bug
        pagesize: 10,
        // 这行是类别,确定是风险还是脆弱性
        type: '',
      },

      passwordInfo: {
        password: '123456',
        // account: 'user',
      },

      // 结果总数 用于翻页
      total: 0,

      // 获取到的脆弱性风险数据
      securityEventsList: [],

      // 脆弱性与风险触发方法的对话框
      triggerMethodDialogVisible: false,

      // 触发方法
      triggerMethod: '',

      // 遮罩
      isLoading: false,
      loadingContent: '',
    }
  },
  created() {
    this.GetSecurityEventsList()
  },
  methods: {
    // 强制立即审计
    async ForceAudit() {
      this.$confirm(
        '确定强制立刻进行审计并生成日志?该功能仅建议测试使用',
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          // const { data: res } = await this.$http.delete('tasks/')
          this.loadingContent = '强制审计中'
          this.isLoading = true

          const { data: res } = await this.$http.get('custom/forceaudit')

          this.isLoading = false

          if (res.meta.status_code == 200) {
            return this.$message({
              type: 'success',
              message: '强制审计成功!',
            })
          } else {
            return this.$message.error(res.meta.message)
          }
        })
        .catch(() => {
          this.isLoading = false
          this.$message({
            type: 'info',
            message: '已取消强制审计',
          })
        })
    },

    // 获取有哪些安全事件
    async GetSecurityEventsList() {
      const { data: res } = await this.$http.get('custom/securityevents', {
        params: this.queryInfo,
      })
      if (res.meta.status_code !== 200)
        return this.$message.error('获取安全事件失败')

      // 成功了就开始赋值
      this.securityEventsList = res.data.SecurityEventsList
      this.total = res.data.Total
      console.log('脆弱性与风险: ', this.securityEventsList)
    },
    // 翻页
    handleCurrentChange(val) {
      // console.log(`当前页: ${val}`)
      this.queryInfo.pagenum = val
      this.GetSecurityEventsList()
    },
    // 改变每页大小
    handleSizeChange(val) {
      // console.log(`每页 ${val} 条`)
      this.queryInfo.pagesize = val
      this.GetSecurityEventsList()
    },
    // 进行请求
    async DoRequest(name) {
      switch (name) {
        // 1. 参数遍历获取大量敏感数据 安全事件包含的风险:
        //    A:单个IP在一段时间内进行请求参数值遍历； B：单个IP在一段时间内返回大量敏感数据
        case '参数遍历获取大量敏感数据': {
          for (let index = 0; index < 100; index++) {
            await this.$http.get(
              'custom/mock/securityevents/RequestTraverseAndReturnTooMuchSensitiveData?userName=' +
                faker.name.lastName() + faker.name.firstName()
            )
          }
          return 200
        }

        case '频繁访问获取大量敏感数据': {
          // for (let index = 0; index < 100; index++) {
          //   await this.$http.get(
          //     'custom/mock/securityevents/RequestTraverseAndReturnTooMuchSensitiveData?userName=' +
          //       faker.name.lastName()
          //   )
          // }
          return 200
        }

        case '异常时间段频繁访问获取大量敏感数据': {
          // for (let index = 0; index < 100; index++) {
          //   await this.$http.get(
          //     'custom/mock/securityevents/RequestTraverseAndReturnTooMuchSensitiveData?userName=' +
          //       faker.name.lastName()
          //   )
          // }
          return 200
        }

        case '发生探测攻击并通过参数遍历获取过量敏感数据': {
          for (let index = 0; index < 100; index++) {
            await this.$http.get(
              'custom/mock/risk/SingleIPReturnTooMuch4XXPeriod'
            )
          }
          for (let index = 0; index < 100; index++) {
            await this.$http.get(
              'custom/mock/risk/SingleIPRequestTraversePeriod?userName=' +
                faker.name.lastName() + faker.name.firstName()
            )
          }
          return 200
        }

        case '发生探测攻击并通过频繁访问获取过量敏感数据': {
          // for (let index = 0; index < 100; index++) {
          //   await this.$http.get(
          //     'custom/mock/risk/SingleIPReturnTooMuch4XXPeriod'
          //   )
          // }
          // for (let index = 0; index < 100; index++) {
          //   await this.$http.get(
          //     'custom/mock/risk/SingleIPRequestTraversePeriod?userName=' +
          //       faker.name.lastName() + faker.name.firstName()
          //   )
          // }
          return 200
        }

        // 本次需要实现的安全事件-mxq-2023-3-7
        // 6. 发生探测攻击并在异常时间段频繁访问获取非预期敏感数据 安全事件包含的风险：
        case '发生探测攻击并在异常时间段频繁访问获取非预期敏感数据': {
          for (let index = 0; index < 100; index++) {
            await this.$http.get(
              'custom/mock/risk/SingleIPReturnTooMuch4XXPeriod'
            )
          }

          for (let index = 0; index < 100; index++) {
            await this.$http.get(
              'custom/mock/risk/SingleIPVisitSameAPIAbnormalPeriod'
            )
          }
          return 200
        }

        // 7. 通过恶意构造请求窃取额外敏感数据 安全事件包含的风险：
        //      A:请求方法异常 or 请求参数名缺失 or 请求参数出现非预期的参数名 or 请求参数值出现新类型
        //      B:单个IP单次返回新类型的敏感数据 or 单个IP单次返回过多敏感数据类型
        //    安全事件触发条件：
        //      A和B告警需要对应到相同的请求和响应包
        case '通过恶意构造请求窃取额外敏感数据': {
          const { data: res } = await this.$http.get(
            'custom/mock/risk/LackOfVarName?reqType=risk'
          )
          return res.meta.status_code
        }

        // 8. API接口遭遇渗透攻击 安全事件包含的风险：
        //      A:请求方法异常 or 请求参数名缺失 or 请求参数出现非预期的参数名 or 请求参数值出现新类型 or 单个IP在一段时间内进行路径遍历
        //      B:单个IP在一段时间内大量4XX
        //    安全事件触发条件：
        //    （1）A中出现的事件和B中出现的事件需要同源IP/账号；
        //    （2）A和B出现事件（窗）相差不超过6小时
        case 'API接口遭遇渗透攻击': {
          // const { data: res } = await this.$http.get(
          //   'custom/mock/risk/LackOfVarName?reqType=risk'
          // )
          return 200
        }

        // 9. 通过恶意构造请求获取大量敏感数据 安全事件包含的风险：
        //      A：请求参数值出现新类型 or 请求参数名缺失 or 请求参数出现非预期的参数名;
        //      B: 单个IP在一段时间内频繁访问同一API or 单个IP在异常时间内频繁访问同一API;
        //      C: 单个IP在一段时间内返回大量敏感数据
        case '通过恶意构造请求获取大量敏感数据': {
          // const { data: res } = await this.$http.get(
          //   'custom/mock/risk/LackOfVarName?reqType=risk'
          // )
          return 200
        }

        // 10.账号失陷并下载了大量敏感数据 安全事件包含的风险：
        //    A: 账号多地访问
        //    B：单个账号在一段时间内返回大量敏感数据
        case '账号失陷并下载了大量敏感数据': {
          //2.2
          // const { data: res } = await this.$http.get(
          //   'custom/mock/securityevents/OneAccountGetSensitiveData'
          // )
          // return res.meta.status_code
          // 2.1
          const { data: res } = await this.$http.get(
            'custom/mock/risk/SingleAccountReturnNewTypeSensiDataOnce?reqType=securityevent'
          )
          return res.meta.status_code
        }
      }
    },
    // 触发风险事件
    TriggerThis(name) {
      this.$confirm('确定触发此风险事件?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          const status_code = await this.DoRequest(name)
          console.log('响应码是', status_code)

          if (status_code == 200) {
            return this.$message({
              type: 'success',
              message: '触发成功!',
            })
          } else {
            return this.$message.error({
              message: '触发失败!',
            })
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '',
          })
        })
    },
    async TriggerAll() {
      this.$confirm('确定触发全部安全事件?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          var flag = 0
          const { data: res } = await this.$http.get('custom/securityevents', {
            params: {
              pagenum: 1,
              pagesize: 100,
            },
          })
          // 显示遮罩
          this.isLoading = true

          for (var i = 0; i < res.data.Total; i++) {
            this.loadingContent =
              '触发' + res.data.SecurityEventsList[i]['name'] + '中'
            var status_code = await this.DoRequest(
              res.data.SecurityEventsList[i]['name']
            )
            if (status_code != 200) {
              this.isLoading = false
              console.log(
                res.data.SecurityEventsList[i]['name'],
                '触发失败, 返回值为:',
                status_code
              )
              flag = 1
            }
          }

          this.isLoading = false

          if (flag == 0) {
            return this.$message({
              type: 'success',
              message: '全部风险事件触发成功!',
            })
          } else {
            return this.$message({
              type: 'fail',
              message: '存在风险事件触发失败!',
            })
          }
        })
        .catch(() => {
          this.isLoading = false
          this.$message({
            type: 'info',
            message: '已取消触发',
          })
        })
    },
    // 展示触发的方法
    OpenTriggerMethodDialog(tiggermethod) {
      this.triggerMethodDialogVisible = true
      this.triggerMethod = tiggermethod
    },
    // 过滤
    filterTag() {
      return true
    },
    //
    filterChangeInit(filters) {
      // console.log('filters: ', filters.type[0])
      if (filters.type) {
        this.queryInfo.type = filters.type[0]
        console.log('this.queryInfo.type', this.queryInfo.type)
        this.GetSecurityEventsList()
      }
    },
  },
}
</script>

<style>
</style>