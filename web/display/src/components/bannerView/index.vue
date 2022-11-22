<!--
 * @Author: chaichai chaichai@cute.com
 * @Date: 2022-09-26 08:42:49
 * @LastEditors: chaichai chaichai@cute.com
 * @LastEditTime: 2022-10-09 10:39:13
 * @FilePath: \blog3.0\src\components\bannerView\index.vue
 * @Description: 
 * 
 * Copyright (c) 2022 by CQUCC-4-433, All Rights Reserved. 
-->
<template>
  <div>
    <div
      class="bannerBox"
      :style="{ background: 'url(' + imgUrl + ')', backgroundSize: 'cover' }"
    >
      <div class="coverBox">
        <div class="navBox">
          <div class="topTitle">Windy.top</div>
        
        <div class="el-menu-demo" background-color="#0000001D" text-color="#fff">
          <v-btn text @click="$router.push('/')"><h1>导航页</h1></v-btn>
          <v-btn text index="about" @click="$router.push('/about')"><h1>首页</h1></v-btn>
          <v-btn text index="blog" @click="$router.push('/blog')"><h1>博客</h1></v-btn>
          <v-btn text index="back" @click="clickAdmin"><h1>后台</h1></v-btn>
          <v-dialog max-width="1000">
            <template v-slot:activator="{on, attrs}" text-color="#fff" background-color="#0000001D">
              <v-btn class="hidden-md-and-down" v-if="headers.username" text @click="loginout"><h1>退出</h1></v-btn>
              <v-btn v-if="headers.username" text dark><h1>欢迎你{{headers.username}}</h1></v-btn>
              <v-btn text v-if="!headers.username" v-bind="attrs" v-on="on"><h1>登录&注册</h1></v-btn>
            </template>

            <template v-slot:default="dialog" class="ground">
              <div class="main-box">
                <div :class="['container', 'container-register', { 'is-txl': isLogin }]">
                  <v-form ref="registerformRef" v-model="registerformvalid">
                    <h2 class="title">请注册</h2>
                    <!-- <span class="text">or use username for registration</span> -->
                    <input class="form__input" type="text" placeholder="Username" v-model="formdata.username" :rules="nameRules"/>
                    <input class="form__input" type="password" placeholder="Password" v-model="formdata.password" :rules="passwordRules"/>
                    <!-- <input class="form__input" type="password" placeholder="checkPassword" v-model="checkPassword" :rules="checkPasswordRules"/> -->
                    <div class="primary-btn" @click="register">立即注册</div>
                  </v-form>
                </div>
                <div
                  :class="['container', 'container-login', { 'is-txl is-z200': isLogin }]"
                >
                  <v-form ref="loginFormRef" v-model="valid">
                    <h2 class="title">请登录</h2>
                    <!-- <span class="text">or use username for registration</span> -->
                    <input class="form__input" type="text" placeholder="Username" v-model="formdata.username" :rules="nameRules"/>
                    <input class="form__input" type="password" placeholder="Password" v-model="formdata.password" :rules="passwordRules"/>
                    <div class="primary-btn" @click="login">立即登录</div>
                  </v-form>
                </div>
                <div :class="['switch', { login: isLogin }]">
                  <div class="switch__circle"></div>
                  <div class="switch__circle switch__circle_top"></div>
                  <div class="switch__container">
                    <h2>{{ isLogin ? 'Hello Friend !' : 'Welcome Back !' }}</h2>
                    <div class="primary-btn" @click="isLogin = !isLogin">
                      {{ isLogin ? '立即注册' : '立即登录' }}
                    </div>
                  </div>
                </div>
              </div>
            </template>
            
          </v-dialog>
        </div>
        </div>
        <div class="centerTile">{{ titleName }}</div>
        <div class="icon">﹀</div>
        <div class="search">
            <v-responsive color="white">
              <v-text-field 
                height="60" 
                dense 
                flat 
                hide-details 
                solo-inverted 
                rounded 
                dark
                placeholder="请输入文章标题查找"
                append-icon="mdi-text-search"
                v-model="searchName"
                @click:append="searchTitle(searchName)"
                @change="searchTitle(searchName)">
              </v-text-field>
            </v-responsive>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Url } from '../../plugins/http';
export default {
  name:'bannerView',

  props: {
    imgUrl: {
      required: true,
    },
    titleName: {
      required: true,
    },
  },
  data() {
    return {
      searchName: '',
      activeIndex: "1",
      valid: true,
      registerformvalid: true,
      formdata: {
            username:'',
            password:''
      },
      checkPassword: '',
      dialog: false,
      headers: {
            Authorization: '',
            username: ''
      },
      isLogin: true,
      nameRules:[
            (v) => !!v || '用户名不能为空',
            (v) =>
              (v && v.length >= 4 && v.length <= 12) ||
              '用户名必须在4到12个字符之间'
          ],
          passwordRules: [
            (v) => !!v || '密码不能为空',
            (v) =>
              (v && v.length >= 6 && v.length <= 20) || '密码必须在6到20个字符之间',
            (v) => v == this.checkPassword || '密码两次输入不一致，请检查'
          ],
          checkPasswordRules: [
            (v) => !!v || '密码不能为空',
            (v) =>
              (v && v.length >= 6 && v.length <= 20) || '密码必须在6到20个字符之间',
            (v) => v == this.formdata.password || '密码两次输入不一致，请检查'
          ]
    };
  },
  mounted() {
    this.headers = {
          Authorization: `Bearer ${window.sessionStorage.getItem('token')}`,
          username: window.sessionStorage.getItem('username')
        }
  },
  methods: {
    clickAdmin() {
      window.location.href= Url + ":3001/admin/"
    },
    // 通过标题查找文章
    searchTitle(title) {
      if(title.length == 0) return this.$message.error('请输入查找文章标题')
      this.$router.push(`/search/${title}`)
    },
    async login() {
      if (!this.$refs.loginFormRef.validate())
            return this.$message.error('输入数据非法，请检查输入的用户名和密码')
          const { data: res } = await this.$http.post('loginfront', this.formdata)
          console.log(res)
          console.log(this.formdata)
          if (res.status != 200) return this.$message.error(res.message)
          window.sessionStorage.setItem('username', res.data)
          window.sessionStorage.setItem('user_id', res.id)
          this.$message.success('登录成功')
          // this.$router.go(val) => 在history记录中前进或者后退val步，当val为0时刷新当前页面。
          this.$router.go(0)
    },
    async register() {
      if (!this.$refs.registerformRef.validate()) 
            return this.$message.error('输入数据非法，请检查输入的用户名和密码')
            const { data: res } = await this.$http.post('user/add', {
              username: this.formdata.username,
              password: this.formdata.password,
              role: 2
            })
            if (res.status !== 200) return this.$message.error(res.message)
            this.$message.success('注册成功')
            this.$router.go(0)
    },
      // 退出
    loginout() {
      window.sessionStorage.clear('token')
      window.sessionStorage.clear('username')
      this.$message.success('退出成功')
      this.$router.go(0)
    },
  },
};
</script>

<style lang="scss">
.ground{
  display: none;
  
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  width: 400px;
  height: 250px;
  border-radius: 20px;
  border: 2px solid #fff;
  background: #fff;
  margin: auto;
  position: absolute;
  top: -200px;
  left: 0;
  right: 0;
  bottom: 0;
}

.main-box {
  position: relative;
  width: 1000px;
  min-width: 1000px;
  min-height: 600px;
  height: 600px;
  padding: 25px;
  background-color: #ecf0f3;
  box-shadow: 10px 10px 10px #d1d9e6, -10px -10px 10px #f9f9f9;
  border-radius: 12px;
  overflow: hidden;
  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    width: 600px;
    height: 100%;
    padding: 25px;
    background-color: #ecf0f3;
    transition: all 1.25s;
    form {
      display: flex;
      justify-content: center;
      align-items: center;
      flex-direction: column;
      width: 100%;
      height: 100%;
      color: #a0a5a8;
      .title {
        font-size: 34px;
        font-weight: 700;
        line-height: 3;
        color: #181818;
      }
      .text {
        margin-top: 30px;
        margin-bottom: 12px;
      }
      .form__input {
        width: 350px;
        height: 40px;
        margin: 4px 0;
        padding-left: 25px;
        font-size: 13px;
        letter-spacing: 0.15px;
        border: none;
        outline: none;
        // font-family: 'Montserrat', sans-serif;
        background-color: #ecf0f3;
        transition: 0.25s ease;
        border-radius: 8px;
        box-shadow: inset 2px 2px 4px #d1d9e6, inset -2px -2px 4px #f9f9f9;
        &::placeholder {
          color: #a0a5a8;
        }
      }
    }
  }
  .container-register {
    z-index: 100;
    left: calc(100% - 600px);
  }
  .container-login {
    left: calc(100% - 600px);
    z-index: 0;
  }
  .is-txl {
    left: 0;
    transition: 1.25s;
    transform-origin: right;
  }
  .is-z200 {
    z-index: 200;
    transition: 1.25s;
  }
  .switch {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    width: 400px;
    padding: 50px;
    z-index: 200;
    transition: 1.25s;
    background-color: #ecf0f3;
    overflow: hidden;
    box-shadow: 4px 4px 10px #d1d9e6, -4px -4px 10px #f9f9f9;
    color: #a0a5a8;
    .switch__circle {
      position: absolute;
      width: 500px;
      height: 500px;
      border-radius: 50%;
      background-color: #ecf0f3;
      box-shadow: inset 8px 8px 12px #d1d9e6, inset -8px -8px 12px #f9f9f9;
      bottom: -60%;
      left: -60%;
      transition: 1.25s;
    }
    .switch__circle_top {
      top: -30%;
      left: 60%;
      width: 300px;
      height: 300px;
    }
    .switch__container {
      display: flex;
      justify-content: center;
      align-items: center;
      flex-direction: column;
      position: absolute;
      width: 400px;
      padding: 50px 55px;
      transition: 1.25s;
      h2 {
        font-size: 34px;
        font-weight: 700;
        line-height: 3;
        color: #181818;
      }
      p {
        font-size: 14px;
        letter-spacing: 0.25px;
        text-align: center;
        line-height: 1.6;
      }
    }
  }
  .login {
    left: calc(100% - 400px);
    .switch__circle {
      left: 0;
    }
  }
  .primary-btn {
    width: 180px;
    height: 50px;
    border-radius: 25px;
    margin-top: 50px;
    text-align: center;
    line-height: 50px;
    font-size: 14px;
    letter-spacing: 2px;
    background-color: #4b70e2;
    color: #f9f9f9;
    cursor: pointer;
    box-shadow: 8px 8px 16px #d1d9e6, -8px -8px 16px #f9f9f9;
    &:hover {
      box-shadow: 4px 4px 6px 0 rgb(255 255 255 / 50%),
        -4px -4px 6px 0 rgb(116 125 136 / 50%),
        inset -4px -4px 6px 0 rgb(255 255 255 / 20%),
        inset 4px 4px 6px 0 rgb(0 0 0 / 40%);
    }
  }
}

.bannerBox {
  width: 100vw;
  height: 100vh;
  background-size: cover;
  overflow: hidden;
}
.coverBox {
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.4);

  .navBox {
    height: 60px;
    display: flex;
    flex-wrap: nowrap;
    line-height: 60px;
    justify-content: space-between;
    padding: 20px 35px 0 35px;

    .topTitle {
      width: 300px;
      text-align: left;
      color: #fff;
      font-size: 38px;
      font-weight: 900;
      text-transform: uppercase;
    }
    .el-menu {
      background-color: rgb(0, 0, 0, 0) !important;
      border: 0px;
    }

    .el-menu-item {
      font-size: 18px;
      font-weight: 600;
      background-color: rgb(0, 0, 0, 0) !important;
    }
    .el-submenu__title {
      font-size: 18px !important;
      font-weight: 600;
      background-color: rgb(0, 0, 0, 0) !important;
    }
  }
  .centerTile {
    width: 100%;
    line-height: 70vh;
    color: #fff;
    font-size: 38px;
    font-weight: 900;
    letter-spacing: 8px;
  }
  .icon {
    margin-top: 30px;
    z-index: 99999;
    font-weight: 900;
    font-size: 35px;
    color: #fff;
  }
  .search {
    width: 700px;
    height: 200px;
    position: absolute;
    top: 50%;
    left: 50%;
    margin-left: -350px;
    margin-top: -650px;
  }
}
</style>
<style>
.friendIco {
  display: inline-block;
  width: 25px;
  margin: 0 3px;
  vertical-align: middle;
}
.friendList {
  height: 80px !important;
  line-height: 80px !important;
  margin: 5px auto;
  text-align: center;
}
</style>