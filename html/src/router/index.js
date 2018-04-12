import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'

import Roles from '@/components/admin/Roles'
import User from '@/components/admin/Users'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld
    },
    {
      path: '/admin/roles',
      name: '分组管理',
      component: Roles
    },
    {
      path: '/admin/users',
      name: '用户管理',
      component: User
    }
  ]
})
