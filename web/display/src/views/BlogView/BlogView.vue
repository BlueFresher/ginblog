<!--
 * @Author: chaichai chaichai@cute.com
 * @Date: 2022-09-26 08:54:27
 * @LastEditors: chaichai chaichai@cute.com
 * @LastEditTime: 2022-09-27 16:25:29
 * @FilePath: \blog3.0\src\views\BlogView.vue
 * @Description: 
 * 
 * Copyright (c) 2022 by CQUCC-4-433, All Rights Reserved. 
-->
<template>
  <div class="aboutBox">
    <bannerView :imgUrl="this.img" :titleName="this.title"></bannerView>
    <!-- <v-row style="height:50px"> -->
    <!-- </v-row> -->
    <div class="mainBox">
      <!-- <v-container> -->
      <v-row>
          <v-col cols="12" md="2">
            <bannerList1></bannerList1>
          </v-col>
          <v-col cols="12" md="1">
            <bannerList2></bannerList2>
          </v-col>
        <v-col cols="12" md="9">
          <v-sheet rounded="lg">
            <router-view :key="$route.fullPath"></router-view>
          </v-sheet>
        </v-col>
      </v-row>
    <!-- </v-container> -->
      <!-- <bannerList2></bannerList2>
      <bannerList3></bannerList3> -->
    </div>
    <footerView></footerView>
  </div>
</template>

<script>
import bannerView from "@/components/bannerView/index.vue";
import footerView from "@/components/footerView/index.vue";
import bannerList1 from "./components/bannerList1.vue";
import bannerList2 from "./components/bannerList2.vue";
import bannerList3 from "./components/bannerList3.vue";
import img from "../../assets/bgg.jpeg"
export default {
  name:'BlogView',
  components: { bannerView, footerView, bannerList1, bannerList2, bannerList3 },
  data() {
    return {
      title: "博客",
      img:img
    };
  },
};
</script>

 <style lang="scss">
.aboutBox {
  height: 100%;
  // 背景图片
  background: url("@/assets/cover6.gif");
  background-repeat: no-repeat;
  background-size: 500px;
  background-position: 110% 100%;
  background-attachment: fixed;

  .mainBox {
    width:85%;
    margin-left:100px;
    margin-right:300px;
    margin-top: 10px;
  }
}
</style>

