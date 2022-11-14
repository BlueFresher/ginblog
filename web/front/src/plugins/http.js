import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = "http://localhost:3001/api/v1"
// 启用网络请求插件
Vue.prototype.$http = axios