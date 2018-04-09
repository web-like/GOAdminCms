import axios from 'axios'
// import $ from 'jquery'
// import Vue from 'vue'
import router from '../router'

var http = axios.create({
  baseURL: 'http://localhost:7070/v1/',
  timeout: 10000
})

// 返回结果拦截
http.interceptors.response.use(function (response) {
  if (response.data.login !== undefined) {
    router.push({
      path: '/login'
    })
  }
  return response.data
})

export default http
