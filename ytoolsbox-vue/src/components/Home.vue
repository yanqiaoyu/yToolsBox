<template>
  <el-container class="home-container">
    <el-header>
      <div>
        <img src="../assets/cat.png" alt="" class="catImage" />
        <span>我的工具盒</span>
      </div>
      <el-button type="warning" @click="logout">退出</el-button>
    </el-header>

    <el-container>
      <el-aside width="200px">
        <el-menu
          background-color="#4d505b"
          text-color="#fff"
          active-text-color="#ffd04b"
        >
          <el-submenu index="1">
            <template slot="title">
              <i class="el-icon-s-goods"></i>
              <span>工具盒</span>
            </template>
            <el-menu-item index="1-1">第一个工具</el-menu-item>
          </el-submenu>

          <el-menu-item index="2">
            <i class="el-icon-setting"></i>
            <span slot="title">配置</span>
          </el-menu-item>

          <el-menu-item index="3">
            <i class="el-icon-info"></i>
            <span slot="title">关于我</span>
          </el-menu-item>
        </el-menu>
      </el-aside>

      <el-main>Main</el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      menuList: []
    }
  },
  created() {
    this.getMenuList()
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath)
    },
    logout() {
      window.sessionStorage.clear()
      this.$router.push('/login')
      this.$message({
        message: '成功退出',
        type: 'success'
      })
    },
    async getMenuList() {
      const { data: res } = await this.$http.get('menus')
      console.log(res)
      if (res.meta.status_code !== 200)
        return this.$message.error('获取首页信息失败')
      this.menuList = res.data
    }
  }
}
</script>

<style lang="less" scoped>
.home-container {
  height: 100%;
}

.el-header {
  color: #fff;
  background: #0388e5;
  background: -webkit-gradient(
    linear,
    left top,
    right top,
    from(#0388e5),
    to(#07bdf4)
  );
  background: linear-gradient(90deg, #0388e5 0, #07bdf4);
  display: flex;
  justify-content: space-between;
  padding-left: 10;
  font-size: 25px;
  align-items: center;
  > div {
    display: flex;
    align-items: center;
    height: 100%;
    span {
      margin-left: 15px;
    }
  }
}

.el-aside {
  background: #4d505b;
  background: -webkit-gradient(
    linear,
    left top,
    left bottom,
    from(#4d505b),
    to(#3b3e47)
  );
  background: linear-gradient(180deg, #4d505b 0, #3b3e47);
}

.el-main {
  background-color: #fff;
}

.catImage {
  height: 100%;
  display: flex;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
