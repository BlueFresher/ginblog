import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = "http://101.42.137.141:3001/api/v1"
// 启用网络请求插件
Vue.prototype.$http = axios