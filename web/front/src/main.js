import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import './plugins/http'

// import moment from 'moment'
import day from 'dayjs'

Vue.filter('dateformat', function(indate, outdate) {
  return day(indate).format(outdate)
})
Vue.config.productionTip = false

new Vue({ // 页面入口
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
