// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import jQuery from 'jquery'
// import VeeValidate from 'vee-validate'
// 引入ajax组件
import axios from './interceptor/http-axios'
// import Tether from 'tether'
// import('bootstrap')
// import('bootstrap/dist/css/bootstrap.css')
// window.Tether = Tether
window.jQuery = window.$ = jQuery
// 引入后台框架css
import('./vendor/bootstrap/css/bootstrap.min.css')
import('./vendor/metisMenu/metisMenu.min.css')
import('./vendor/custom/css/sb-admin-2.min.css')
import('./vendor/font-awesome/css/font-awesome.min.css')
// 引入后台框架JS
import('./vendor/bootstrap/js/bootstrap.min')
import('./vendor/metisMenu/metisMenu.min')
import('./vendor/custom/js/sb-admin-2.min')

// Vue.use(VeeValidate)
Vue.prototype.$http = axios
Vue.config.productionTip = false
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
