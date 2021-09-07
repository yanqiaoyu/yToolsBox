<template>
  <div class="login_container">
    <!-- 登录界面 -->

    <!-- 这个是整个页面的容器 -->
    <div class="login_box">
      <!-- 这个是登录的白色的框的容器 -->
      <div class="login_image">
        <!-- 这个是登录的图片 -->
        <img src="../assets/cat.png" alt="" />
      </div>

      <div>
        <!-- 注意这里用到的组件，必须在plugins的elment.js里面导入 -->
        <el-form
          class="login_form"
          :model="loginForm"
          :rules="rules"
          ref="login_form_ref"
        >
          <!-- 账号 -->
          <el-form-item prop="name">
            <el-input
              prefix-icon="el-icon-user"
              v-model="loginForm.name"
            ></el-input>
          </el-form-item>
          <!-- 密码 -->
          <el-form-item prop="password">
            <el-input
              prefix-icon="el-icon-key"
              v-model="loginForm.password"
              type="password"
            ></el-input>
          </el-form-item>
          <!-- 按钮区域 -->
          <el-form-item class="login_btn">
            <el-button type="primary" @click="login">登录</el-button>
            <el-button type="info" @click="reset">重置为正确账号密码</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
// 这里为什么我们需要引入qs
// 原因是axios默认发送的是payload
// 后端需要接受FormData
// 所以下面我们做了qs.stringify(this.loginForm)
import qs from 'qs'

export default {
  data() {
    return {
      loginForm: {
        name: 'admin',
        password: 'admin'
      },
      rules: {
        name: [
          { required: true, message: '请输入账号名称', trigger: 'blur' },
          { min: 5, max: 18, message: '长度在 5 到 18 个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 5, max: 18, message: '长度在 5 到 18 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    reset() {
      // console.log(this)
      this.$refs.login_form_ref.resetFields()
    },
    login() {
      this.$refs.login_form_ref.validate(async valid => {
        // console.log(valid)
        if (!valid) return
        const result = await this.$http.post(
          'login',
          qs.stringify(this.loginForm)
        )
        console.log(result.data)
        if (result.data.data.status_code !== 200) {
          this.$message.error('登陆失败')
          return
        }
        // alert('登陆成功')
        this.$message({
          message: '成功登陆',
          type: 'success'
        })
        // 登录成功后，需要保存token至本地
        window.sessionStorage.setItem('token', result.data.data.token)
        // 编程式路由，跳转页面
        this.$router.push('/home')
      })
    }
  }
}
</script>

<style lang="less" scoped>
// lang=less，使得代码样式支持less语法
// scope限定样式只在当前组件生效

// 登录页面的整个容器
.login_container {
  background-color: #2b4b6b;
  height: 100%;
}

// 中间那个白色的框
.login_box {
  width: 450px;
  height: 300px;
  background-color: #fff;
  border-radius: 3px;
  // 位置采用绝对位置
  position: absolute;
  // 举例顶部和左边都有50%的差距
  left: 50%;
  top: 50%;
  // 然后再向左，向上位移50%
  transform: translate(-50%, -50%);
}

// 登录的图片
.login_image {
  width: 130px;
  height: 130px;
  // 给这个图片添加一个正方形的灰色边框
  border: solid #eee;
  // 把边框变成圆角
  border-radius: 50%;
  //图片和边框之间有间距，5px
  padding: 5px;
  // 添加阴影
  box-shadow: 0 0 10px #ddd;

  // 移动这个图片
  position: absolute;
  left: 50%;
  transform: translate(-50%);
  top: -25%;
  background-color: #fff;

  img {
    width: 100%;
    height: 100%;
    border-radius: 50%;
    background-color: #eee;
  }
}

// 登录整体表单
.login_form {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 0 20px;
  // 假如您需要并排放置两个带边框的框，可通过将 box-sizing 设置为 "border-box"
  box-sizing: border-box;
}

// 登录按钮
.login_btn {
  display: flex;
  justify-content: flex-end;
}
</style>
