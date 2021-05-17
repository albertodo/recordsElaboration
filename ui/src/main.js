import Vue from 'vue'
//import Loading from 'vue-loading-overlay'
import 'vue-loading-overlay/dist/vue-loading.css'
import App from './App.vue'
import router from './router'
//import store from './store'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import Vuetify from 'vuetify';
import vuetify from './plugins/vuetify'
import VueExcelXlsx from "vue-excel-xlsx";

Vue.config.productionTip = false

Vue.use(Vuetify)
Vue.use(VueExcelXlsx);

new Vue({
  router,
  Vuetify : new Vuetify(),
  vuetify,

  //store,
  render: h => h(App)
}).$mount('#app')