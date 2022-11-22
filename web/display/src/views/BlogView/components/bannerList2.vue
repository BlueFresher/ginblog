<template>
  <div class="box">
    <div class="titleName"><h3>分类</h3></div>
    <div class="bannerBox">
      <v-card
      max-width="400"
      class="mx-auto"
    >
      <v-container>
        <v-row dense>
          <v-col
            v-for="(item, i) in cateList"
            :key="i"
            cols="12"
          >
            <v-card color="#952175" dark @click="gotoCate(item.id)">
              <div class="d-flex flex-no-wrap justify-space-between">
                <div>
                  <v-card-title
                    class="text-h5"
                    v-text="item.name"
                  ></v-card-title>
  
                </div>
  
              </div>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
    <div class="titleName"></div>
  </div>
</template>

<script>

export default {
data() {
  return {
    cateList:[],
  };
},
created(){
        this.GetCateList()
},
methods: {
    // 获取分类
    async GetCateList(){
      const {data:res} = await this.$http.get('category')
      this.cateList = res.data
    },
    gotoCate(cid) {
      this.$router.push(`/category/${cid}`).catch((err) => err)
    },
},
};
</script>

<style lang='scss'>
.box {
height:1000px;
box-shadow: 0 1px 10px rgb(0 0 0 / 10%);
margin-top: 30px;
border-radius: 36px 36px 36px 36px;

.titleName {
margin: 0 auto;
letter-spacing: 2px;
width: 80%;
font-size: 22px;
height: 50px;
line-height: 50px;
font-weight: 550;
}
.bannerBox {
height: 90%;
width:100%;
.slider-item {
  color: rgb(255, 255, 255);
  font-weight: 550;
  line-height: 100px;
  text-shadow: 2px 2px 2px #000;
  font-size: 30px;
}
}

}
</style>