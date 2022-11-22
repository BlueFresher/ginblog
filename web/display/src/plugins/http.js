import Vue from 'vue'
import axios from 'axios'

// axios.defaults.baseURL = "http://localhost:3001/api/v1/"

let protocol = window.location.protocol; //协议
let host = window.location.host; //主机
// let reg = /^localhost+/;
let Url="";

// if(reg.test(host)) {
//    //若本地项目调试使用
//     axios.defaults.baseURL = 'http://localhost:3001/api/v1/';
//     Url = 'http://localhost';
// } else {
    //动态请求地址             协议               主机
    axios.defaults.baseURL = protocol + "//" + host.slice(0,9)  +"/api/v1/";
    Url = protocol + "//" + host.slice(0,9);
// }
// console.log(axios.defaults.baseURL)

// 启用网络请求插件
Vue.prototype.$http = axios

export { Url }