import Vue from 'vue'
import axios from 'axios'

// let Url = 'http://101.42.137.141:3001/api/v1'
// axios.defaults.baseURL = Url
let protocol = window.location.protocol; //协议
let host = window.location.host; //主机
// let reg = /^localhost+/;
// if(reg.test(host)) {
//    //若本地项目调试使用
//     axios.defaults.baseURL = 'http://localhost:3001/api/v1/';
//     Url = 'http://localhost:3001/api/v1';
// } else {
    //动态请求地址             协议               主机
    axios.defaults.baseURL = "http://101.42.137.141:3001/api/v1/";
    let Url = "http://101.42.137.141:3001";
// }
// 添加请求拦截器
axios.interceptors.request.use(config => {
  // 为请求头对象,添加Token验证的Authorization字段,才可以调用有权限的API
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  // 在最后必须return config
  return config
})

Vue.prototype.$http = axios

export { Url }