<!--
 * @Author: chaichai chaichai@cute.com
 * @Date: 2022-09-27 10:56:50
 * @LastEditors: chaichai chaichai@cute.com
 * @LastEditTime: 2022-09-27 16:25:25
 * @FilePath: \blog3.0\src\views\404View.vue
 * @Description: 
 * 
 * Copyright (c) 2022 by CQUCC-4-433, All Rights Reserved. 
-->
<template>
<div align="center" style='font-family:"微软雅黑"'>
    <img src="@/assets/404.jpeg" border="0" height="" style="margin:10vh 0 0 0"/>
    <br>
    <h2>页面加载出错！</h2><br><br>
    <v-btn @click="handleBack" color="red"><h1>返回导航页</h1></v-btn>
    <br><br> &copy; 枫叶荻花
</div>
</template>

<script>
export default{
    methods:{
        handleBack() {
            this.$router.push("/")
        }
    }
}
</script>