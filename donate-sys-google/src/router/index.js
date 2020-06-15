// import Vue from 'vue'
import Vue from 'vue/dist/vue.common'
import Router from 'vue-router'
//  HelloWorldimport from '@/components/HelloWorld'
import Home from "@/components/Home";

import ProjectList from "@/components/project/ProjectList";
import ProjectDetail from "@/components/project/ProjectDetail";
import DonateList from "@/components/project/DonateList";
import DonateListFinished from "@/components/project/DonateListFinished";

import Login from "@/components/login/Login";
import Findpassword from "@/components/login/Findpassword";
import Register_d from "@/components/login/Register_d";
import Register_g from "@/components/login/Register_g";

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: ()=>import('../components/Home')
    },
    {
      path: '/home',
      name: '谷粒捐物资捐赠系统-首页',
      component: Home
    },
    {
      path: '/projectList',
      name: '谷粒捐物资捐赠系统-项目列表',
      component: ProjectList
    },
    {
      path:'/projectDetail',
      name:'谷粒捐物资捐赠系统-项目详情',
      component: ProjectDetail,
      children:[]
    },
    {
      path:'/projectDetail/donateList',
      name:'谷粒捐物资捐赠系统-定向捐赠单填写',
      component: DonateList
    },
    {
      path:'/projectDetail/donateList/finished',
      name:'谷粒捐物资捐赠系统-定向捐赠单填写',
      component: DonateListFinished
    },
    {
      path:'/login',
      name:'登录',
      component: Login
     },
     {
       path:'/login/Findpassword',
       name:'重置密码',
       component: Findpassword,
      },
      {
        path:'/login/Register_d',
        name:'捐赠方注册',
        component: Register_d,
       },
      {
        path:'/login/Register_g',
        name:'受赠方注册',
        component: Register_g,
       }

  ]
})
