<template>
    <v-container>
        <div v-if="total == 0 && isLoad" class="d-flex justify-center align-center">
            <div>
                <v-alert class="ma-5" dense outlined type="error">抱歉，暂无数据！</v-alert>
            </div>
        </div>
        <v-sheet>
            <v-card class="ma-3"
            v-for="item in artList"
            :key="item.id"
            link
            @click="$router.push(`/article/detail/${item.ID}`)">
            <v-row no-gutters class="d-flex align-center">
                <v-avatar class="ma-3 hidden-sm-and-down" size="125" tile>
                    <v-img :src="item.img"></v-img>
                </v-avatar>
                <v-col>
                    <!-- 标题 -->
                    <v-card-title>
                        <v-chip color="purple" outlined label class="mr-3 white--text">{{item.Category.name}}</v-chip>
                        <div>{{item.title}}</div>
                    </v-card-title>
                    <!-- 描述 -->
                    <v-card-subtitle class="mt-1" v-text="item.desc"></v-card-subtitle>
                    
                    <v-divider class="mx-4"></v-divider>
                    <!-- 日期 -->
                    <v-card-text class="d-flex align-center">
                        <div class="d-flex align-center">
                            <v-icon class="mr-1" small>{{ 'mdi-calendar-month' }}</v-icon>
                        <span>{{item.CreatedAt | dateformat('YYYY-MM-DD HH:MM')}}</span>
                        </div>
                    </v-card-text>

                </v-col>
            </v-row>
            </v-card>
            <!-- 显示页数 -->
            <v-col>
                <div class="text-center">
                  <v-pagination
                    total-visible="7"
                    v-model="queryParam.pagenum"
                    :length="Math.ceil(total / queryParam.pagesize)"
                    @input="getArtList()"
                  ></v-pagination>
                </div>
              </v-col>
        </v-sheet>
    </v-container>
</template>
<script>
export default {
    // 点击页面传递进来的参数
    props:['cid'],
    // 自定义的数据，用于业务处理和显示
    data(){
        return {
            artList:[],
            queryParam:{
                pagesize:5,
                pagenum:1
            },
            total: 0,
            isLoad: false
        }
    },
    created(){
        this.getArtList()
    },
    mounted() {
        this.getArtList()
    },
    methods: {
        // 获取文章列表
        async getArtList() {
        const { data: res } = await this.$http.get(`article/list/${this.cid}`, {
            params: {
            pagesize: this.queryParam.pagesize,
            pagenum: this.queryParam.pagenum
            }
        })
            this.artList = res.data
            this.total = res.total
            this.isLoad = true
            console.log(res)
        }
    }
}
</script>

<style scoped>
.nodate {
  width: 100%;
  height: 100%;
}
</style>