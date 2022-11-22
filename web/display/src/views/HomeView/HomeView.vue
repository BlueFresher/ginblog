<!--
 * @Author: chaichai chaichai@cute.com
 * @Date: 2022-09-26 08:29:56
 * @LastEditors: chaichai chaichai@cute.com
 * @LastEditTime: 2022-11-09 10:38:10
 * @FilePath: \blog3.0\src\views\HomeView\HomeView.vue
 * @Description: 
 * 
 * Copyright (c) 2022 by CQUCC-4-433, All Rights Reserved. 
-->
<template>
  <div class="home">
    <div class="coverBox">
      <div class="centerBox">
        <!-- 头像 -->
        <v-avatar
          :size="150"
          class="picT"
        ></v-avatar>
        <!-- 作者姓名 -->
        <div class="ahtuorName">Windy</div>
        <v-divider class="divider"></v-divider>
        <!-- 内容 -->
        <div class="title" v-for="(v, k) in content" :key="k">
          <!-- <typewriter class="str" :str="v"></typewriter> -->
        </div>
        <!-- 引导按钮 -->
        <div class="btnBox">
          <v-btn rounded @click="clickHome" class="btnStyle">首页</v-btn>
          <v-btn rounded @click="$router.push('/blog')" class="btnStyle">博客</v-btn>
          <v-btn rounded @click="clickAdmin" class="btnStyle">后台</v-btn>
        </div>
        <!-- 链接图标模块 -->
        <div class="continueBox">
          <img src="@/assets/QQ.png" alt="" class="logoimg" @click="goQQ" />
          <img
            src="@/assets/github.png"
            alt=""
            class="logoimg"
            @click="jump"
          />
        </div>
      </div>
      <!-- 页脚 -->
      <div class="footerBox">
        &copy; 2022 枫叶荻花<br />
      </div>                                             
    </div>
  </div>
</template>


<script>
import Popup from "@/components/dialogView/index.vue"; //弹窗组件                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       /dialogView/index.vue"; //弹窗组件
import typewriter from "./components/typewriter.vue";
import "./css/HomeView.scss";
import { Url } from '../../plugins/http';
export default {
  name: "HomeView",
  components: { Popup, typewriter },
  data() {
    return {
      content: [],
      str: "我是文字",
      //文章内容
      words: [
        "嗨 欢迎来到chaichai.top",
        "励志成为优秀且花里胡哨的程序员，并在努力奋斗	",
        "CQUCC-4-433正在找寻志同道合的小伙伴，欢迎前端、后端、UI加入我们！",
      ],
      dialogVisible: false, // 弹框的出现与否
      dialogTitle: "GitHub", //弹窗标题
    };
  },
  mounted() {
    this.setTime();
  },
  methods: {
    setTime() {
      var arr = this.words;
      var that = this;
      arr.forEach(function (v, k) {
        setTimeout(function () {
          that.content.push(v);
        }, k * 2500);
      });
    },
    updateVisible(val) {
      this.dialogVisible = val;
    },
    // 点击取消的事件
    resetPopupData() {
      //  这里可重置数据
      this.dialogVisible = false;
    },
    // 点击确定的按钮
    async submitPopupData() {
      this.dialogVisible = false;
    },
    // 关闭弹框（头部的X）
    handleClose() {
      this.dialogVisible = false;
    },
    clickHome() {
      this.$router.push("/about");
    },
    clickAdmin() {
      window.location.href= Url + ":3001/admin/"
    },
    goQQ() {
      //qq弹窗内容
      this.$message.success("QQ:3053239377");
    },
    jump() {
      window.location.href = "https://github.com/BlueFresher"
    }
  },
};
</script>