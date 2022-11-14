import Vue from 'vue'
import axios from 'axios'

let Url = 'http://localhost:3001/api/v1/'
axios.defaults.baseURL = Url
// 添加请求拦截器
axios.interceptors.request.use(config => {
  // 为请求头对象,添加Token验证的Authorization字段,才可以调用有权限的API
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  // 在最后必须return config
  return config
})

Vue.prototype.$http = axios

export { Url }