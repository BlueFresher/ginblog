<template>
    <v-col>
        <v-card class="ma-3" hover raised rounded v-for="item in artList" :key="item.id" link @click="$router.push(`article/detail/${item.ID}`)">
            <v-system-bar light></v-system-bar>
            <v-divider
            class="mx-4"
          ></v-divider>
            <v-row np-gutters>
                <v-col class="d-flex justify-center align-center mx-3" cols="1">
                    <v-img max-height="100" max-width="85" :src="item.img"></v-img>
                </v-col>
                <v-col>
                    <v-card-title class="my-2">
                        <v-chip color="pink" label class="mr-3 white--text">{{item.Category.name}}</v-chip>
                    {{item.title}}
                    </v-card-title>
                    <v-card-subtitle class="desc" v-text="item.desc"></v-card-subtitle>
                    <v-divider></v-divider>
                    <v-card-text>
                        <v-icon class="mr-1" small>{{'mdi-calendar-month'}}</v-icon>
                        <span>{{item.CreatedAt | dateformat('YYYY-MM-DD')}}</span>
                    </v-card-text>
                </v-col>
            </v-row>
        </v-card>
        <div class="text-center">
            <v-pagination
              color="indigo"
              total-visible="7"
              v-model="queryParam.pagenum"
              :length="Math.ceil(total / queryParam.pagesize)"
              @input="getArtList()"
            ></v-pagination>
          </div>
    </v-col>
</template>

<script>
// :style="{background:randomRgb()}"
export default {
    data(){
       return {
        artList:[],
        queryParam:{
            pagesize:5,
            pagenum:1,
        },
        total:0,
        bgColor:''
       } 
    },
    created(){
    },
    mounted(){
        this.getArtList()
        // this.randomRgb()
    },
    methods:{
        // 获取文章列表
        async getArtList(){
            const {data:res} = await this.$http.get('article', {params:{
                pagesize: this.queryParam.pagesize,
                pagenum: this.queryParam.pagenum
            }})
            this.artList = res.data
            this.total = res.total
            console.log(res)
        },

        randomRgb() {
                let colorList = [
                    "linear-gradient(120deg, #0E2101 0%, #66a6ff 100%)",
                    "linear-gradient(120deg, #22E1FF 0%, #1D8FE1 48%, #625EB1 100%)",
                    "linear-gradient(120deg, #3D4E81 0%, #5753C9 48%, #6E7FF3 100%)",
                    "linear-gradient(120deg, #209cff 0%, #68e0cf 100%)",
                    "linear-gradient(120deg, #000000 0%, #50c9c3 100%)",
                    "linear-gradient(120deg, #007adf 0%, #00ecbc 100%)",
                    "linear-gradient(120deg, #6a11cb 0%, #2575fc 100%)"
                ]
                let index = Math.floor(Math.random()*colorList.length)
                this.bgColor =  colorList[index]
                return colorList[index]
            },
    }
}
</script>

<style lang="scss">
.desc {
    color: rgb(255, 255, 255);
    font-weight: 550;
    line-height: 50px;
    text-shadow: 2px 2px 2px #000;
    font-size: 20px;
}
</style>