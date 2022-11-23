<!--
 * @Author: chaichai chaichai@cute.com
 * @Date: 2022-09-26 08:29:56
 * @LastEditors: chaichai chaichai@cute.com
 * @LastEditTime: 2022-10-10 17:13:00
 * @FilePath: \blog3.0\src\views\FirstView\FirstView.vue
 * @Description:  [CQUCC-4-433](https://github.com/4-433) 正在找寻志同道合的小伙伴，欢迎前端、后端、UI加入我们！
 * 
 * Copyright (c) 2022 by CQUCC-4-433, All Rights Reserved. 
-->

<template>
  <div class="aboutBox">
    <bannerView
      :imgUrl="this.img"
      :titleName="this.title"
      ref="banner"
    ></bannerView>
    <div class="mainBox">
      <div class="contentBox">
        <div class="contentTitle">
          <div class="markdown-body">
            <markdown />
          </div>
        </div>
      </div>
      <div :class="locked ? 'asideBoxFix' : 'asideBox'">
        <!-- <div class="asideImg"> -->
          <!-- 头像 -->
          <v-img :src="profileInfo.img">
            <v-col class="avatar">
              <v-avatar :size="size" >
                <img :src="profileInfo.avatar" alt />
              </v-avatar>
            </v-col>
          </v-img>
        <!-- </div> -->
        <div class="title"><h2>{{profileInfo.name}}</h2></div>
        <br/>
        <div class="asideTile1">{{profileInfo.desc}}</div>
        <v-list nav dense>
          <v-list-item>
              <v-list-item-icon class="ma-3">
                  <v-icon color="blue darken-2">{{'mdi-qqchat'}}</v-icon>
              </v-list-item-icon>
              <v-list-item-content class="grey-text">{{profileInfo.qq_chat}}</v-list-item-content>
          </v-list-item>

          <v-list-item>
              <v-list-item-icon class="ma-3">
                  <v-icon color="indigo">{{'mdi-email'}}</v-icon>
              </v-list-item-icon>
              <v-list-item-content class="grey-text">{{profileInfo.email}}</v-list-item-content>
          </v-list-item>

          <v-list-item>
              <v-list-item-icon class="ma-3">
                  <v-icon color="grey darken-1">{{'mdi-github'}}</v-icon>
              </v-list-item-icon>
              <v-list-item-content class="grey-text">{{profileInfo.email}}</v-list-item-content>
          </v-list-item>
      </v-list>
        <v-divider>🌴</v-divider>
        <!-- 侧边栏底部图片 -->
        <img src="@/assets/aside3.gif" alt="" class="bottomImg" />
      </div>
      <div v-if="btnFlag" class="go-top" @click="backTop">
        <!-- 返回顶部图标 -->
        <img src="@/assets/backTop.png" alt="" class="backTopbtn" /> 
      </div>
    </div>
    <footerView></footerView>
  </div>
</template>

<script>
import bannerView from "@/components/bannerView/index";
import footerView from "@/components/footerView/index.vue";
import img16 from "../../assets/bg16.jpg"
// md文件地址
import markdown from "../home.md";
import "./css/FirstView.scss";
import "highlight.js/styles/github.css";
import "github-markdown-css";
export default {
  name:'FirstView',
  components: { bannerView, markdown, footerView },
  created() {
    this.getProfileInfo()
  },
  mounted() {
    window.addEventListener("scroll", this.scrollToTop);
    this.$nextTick(() => {
      let $ele = this.$refs.banner;
      this.bannerH = $ele.$el.offsetHeight;
    });
  },
  destroyed() {
    window.removeEventListener("scroll", this.scrollToTop);
  },
  data() {
    return {
      //侧边栏头像大小
      size: 90,
      bannerH: 0,
      locked: false,
      btnFlag: false,
      //导航背景图片
      img: img16,
      // 导航文字说明
      title: "首页",
      profileInfo: {
        id: 1
      }
    };
  },
  methods: {
    // 获取个人设置
    async getProfileInfo() {
        const { data: res } = await this.$http.get(
            `profile/${this.profileInfo.id}`)
      this.profileInfo = res.data
      this.$root.$emit('msg', res.data.icp_record)
    },

    backTop() {
      const that = this;
      let timer = setInterval(() => {
        let ispeed = Math.floor(-that.scrollTop / 5);
        document.documentElement.scrollTop = document.body.scrollTop =
          that.scrollTop + ispeed;
        if (that.scrollTop === 0) {
          clearInterval(timer);
        }
      }, 16);
    },
    scrollToTop() {
      const that = this;
      let scrollTop =
        window.pageYOffset ||
        document.documentElement.scrollTop ||
        document.body.scrollTop;
      that.scrollTop = scrollTop;
      that.locked = that.btnFlag = that.scrollTop > that.bannerH
      if (that.scrollTop > that.bannerH) {
        that.locked = true;
        that.btnFlag = true;
      } else {
        that.locked = false;
        that.btnFlag = false;
      }
    },
  },
};
</script>

<style lang="scss">
.avatar {
  margin-top: 100px;
}
.title {
  margin-top: 20px;
}
</style>
