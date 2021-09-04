<template>
  <div class="login_container">
      <!-- 登录界面 -->

      <!-- 这个是整个页面的容器 -->
      <div class="login_box">
        <!-- 这个是登录的白色的框的容器 -->
        <div class="login_image">
          <!-- 这个是登录的图片 -->
          <img src="../assets/logo.png" alt="">
        </div>
      </div>
  </div>
</template>

<script>
export default {

}
</script>

<style lang="less" scoped>
// lang=less，使得代码样式支持less语法
// scope限定样式只在当前组件生效

// 登录页面的整个容器
.login_container{
  background-color: #2b4b6b;
  height: 100%;
}

// 中间那个白色的框
.login_box{
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
  transform: translate(-50%,-50%);
}

.login_image{
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

  img{
    width: 100%;
    height: 100%;
    border-radius: 50%;
    background-color: #eee;
  }
}
</style>
