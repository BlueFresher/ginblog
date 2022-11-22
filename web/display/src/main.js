import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import './plugins/http'

// import ElementUI from 'element-ui';
// import 'amfe-flexible'
// import animated from 'animate.css'
// Vue.use(animated)

// Vue.use(ElementUI);

import day from 'dayjs'

Vue.filter('dateformat', function(indate, outdate) {
  return day(indate).format(outdate)
})

Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
