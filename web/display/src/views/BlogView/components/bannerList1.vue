<template>
      <div class="box">
        <div class="titleName"><h3>轮播图</h3></div>
        <div class="bannerBox">

        <slider ref="slider" :options="options" @slide="slider">
          <slideritem
            v-for="(item, index) in someList"
            :key="index"
            :style="item.style"
            @click="clickSlider(item,index)"
            >{{ nowDate }}<br />
            {{ nowWeek }}<br />
            {{ nowTime }}</slideritem
          >
        
        </slider>
      </div>
        <div class="titleName"></div>
      </div>
</template>

<script>
import { slider, slideritem } from "vue-concise-slider";
import imgSrc11 from "../../../assets/bg11.jpg"
import imgSrc12 from "../../../assets/bg12.jpg"
import imgSrc13 from "../../../assets/bg13.jpg"
export default {
  components: { slider, slideritem },
  data() {
    return {
      cover:"cover",
      someList: [
        {
          html: "Vue入门到入土",
          athuer: "chaichai",
          creatTime: "2021-10-19",
          style: {
            background:'url(' + imgSrc12 + ')',
            backgroundSize: "cover",
          },
        },
        {
          html: "Vue框架详解",
          athuer: "chaichai",
          creatTime: "2021-10-19",

          style: {
            background:
              'url(' + imgSrc11 + ')',
            backgroundSize: "cover",
          },
        },
        {
          html: "Vue个人博客案例",
          athuer: "chaichai",
          creatTime: "2021-10-19",

          style: {
            background:'url(' + imgSrc13 + ')',
            backgroundSize: "cover",
          },
        },
      ],
      options: {
        currentPage: 0,
        direction: "vertical",
        slidesToScroll: 1, //每次滑动项数
        effect: "fade",
        thresholdDistance: 10,
        loop: true,
      },
      timer: "", //定义一个定时器
      nowDate: "", // 当前日期时间
      nowTime:"",
      nowWeek:"",
    };
  },
  mounted(){
  this.getTime();
  },
  methods: {
      slider(data) {
        console.log(data,'silder1');
      },
      randomRgb() {
                let colorList = [
                  'url("http://chaichaiimage.oss-cn-hangzhou.aliyuncs.com/blog3.0/bg12.jpg")',
                  'url("http://chaichaiimage.oss-cn-hangzhou.aliyuncs.com/blog3.0/bg11.jpg")',
                ]
                let index = Math.floor(Math.random()*colorList.length)
                this.bgColor =  colorList[index]
                return colorList[index]
            },
      getTime() {
        this.timer = setInterval(() => {
        let timeDate = new Date();
        let year = timeDate.getFullYear();
        let mounth = timeDate.getMonth() + 1;
        let day = timeDate.getDate();
        let hours = timeDate.getHours();
        hours = hours >= 10 ? hours : "0" + hours;
        let minutes = timeDate.getMinutes();
        minutes = minutes >= 10 ? minutes : "0" + minutes;
        let seconds = timeDate.getSeconds();
        seconds = seconds >= 10 ? seconds : "0" + seconds;
        let week = timeDate.getDay();
        let weekArr = [
          "星期日",
          "星期一",
          "星期二",
          "星期三",
          "星期四",
          "星期五",
          "星期六",
        ];

        this.nowDate = `${year}-${mounth}-${day}`;
        this.nowWeek = `${weekArr[week]}`;
        this.nowTime = `${hours}:${minutes}:${seconds}`;
      }, 1000);
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