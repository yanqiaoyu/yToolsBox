<template>
  <el-container class="home-container">
    <el-header>
      <div>
        <!-- <img src="../assets/cat.png" alt="" class="catImage" /> -->
        <!-- 需要require引入资源 -->
        <el-avatar
          :size="55"
          fit="fit"
          :src="require('../assets/cat.png')"
        ></el-avatar>
        <span>我的工具盒</span>
      </div>
      <el-button type="warning" @click="logout">退出</el-button>
    </el-header>

    <el-container>
      <el-aside :width="isCollapse ? '64px' : '200px'">
        <div class="toggle-button" @click="toggleCollapse">|||</div>
        <!-- unique-opened：一次只能有一个子模块 -->
        <el-menu
          class="el-menu-vertical-demo"
          background-color="#4d505b"
          text-color="#fff"
          active-text-color="#409eff"
          :unique-opened="true"
          :collapse="isCollapse"
          :collapse-transition="false"
          router
        >
          <!-- 把index从id改成path，配合组件的router属性，可以直接实现跳转 -->
          <template v-for="item in menuList">
            <el-submenu
              v-if="item.child.length"
              :index="'/' + item.path"
              :key="item.id"
            >
              <template slot="title">
                <i :class="iconList[item.id]"></i>
                <span>{{ item.authName }}</span>
              </template>
              <!-- 把index从id改成path，配合组件的router属性，可以直接实现跳转 -->
              <el-menu-item
                :index="'/' + child.path"
                v-for="child in item.child"
                :key="child.id"
              >
                <template slot="title">
                  <i :class="childIconList[child.id]"></i>
                  <span>{{ child.authName }}</span>
                </template>
              </el-menu-item>
            </el-submenu>

            <!-- 把index从id改成path，配合组件的router属性，可以直接实现跳转 -->
            <!-- <el-menu-item v-else :index="item.id + ''" :key="item.id"> -->
            <el-menu-item v-else :index="'/' + item.path" :key="item.id">
              <i :class="iconList[item.id]"></i>
              <span slot="title">{{ item.authName }}</span>
            </el-menu-item>
          </template>
        </el-menu>
      </el-aside>

      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      menuList: [],
      // 动态获取icon，当然这里写死了，如果后续新增了，这里也要改
      iconList: {
        '1': 'el-icon-s-home',
        '2': 'el-icon-view',
        '3': 'el-icon-s-goods',
        '4': 'el-icon-setting',
        '5': 'el-icon-info'
      },
      childIconList: {
        '401': 'el-icon-s-custom',
        '402': 'el-icon-cpu'
      },
      isCollapse: false
      // imgPath: "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
      // imgPath: "../assets/cat.png"
    }
  },
  created() {
    this.getMenuList()
  },
  methods: {
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
      if (res.meta.status_code !== 200)
        return this.$message.error('获取首页信息失败')
      this.menuList = res.data
    },
    toggleCollapse() {
      // console.log('展开与收起状态:', this.isCollapse)
      this.isCollapse = !this.isCollapse
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
  // 这个是为了解决子模块展开时对其的问题
  .el-menu {
    border-right: none;
  }
}

.el-main {
  background-color: #fff;
}

.catImage {
  height: 75%;
  display: flex;
  // 给这个图片添加一个正方形的灰色边框
  border: solid #eee;
  // 把边框变成圆角
  border-radius: 100%;
  //图片和边框之间有间距，5px
  padding: 5px;
  // 添加阴影
  box-shadow: 0 0 10px #ddd;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}

.toggle-button {
  background-color: #4a5064;
  font-size: 10px;
  line-height: 24px;
  color: #fff;
  text-align: center;
  letter-spacing: 0.2em;
  cursor: pointer;
}
</style>
