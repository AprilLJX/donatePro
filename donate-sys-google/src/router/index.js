// import Vue from 'vue'
import Vue from 'vue/dist/vue.common'
import Router from 'vue-router'
//  HelloWorldimport from '@/components/HelloWorld'
import Home from "@/components/Home";

import ProjectList from "@/components/project/ProjectList";
import ProjectDetail from "@/components/project/ProjectDetail";

import DonateList from "@/components/donation/DonateList";
import DonateListFinished from "@/components/donation/DonateListFinished";
import DonateLogistics from "@/components/donation/DonateLogistics";
import NoRedictDonateCard from "@/components/noRedictDonation/NoRedictDonateCard";

import Login from "@/components/login/Login";
import Findpassword from "@/components/login/Findpassword";
import Register_d from "@/components/login/Register_d";
import Register_g from "@/components/login/Register_g";

import Maincontrol from "@/components/Personal/Maincontrol";
import Love from "@/components/Personal/Love";
import History from "@/components/Personal/History";
import Personalcenter from "@/components/Personal/Personalcenter";
import Message from "@/components/Personal/Message";
import Checkproject from "@/components/Personal/Checkproject";

import Maincontrol_g from "@/components/Personal/Maincontrol_g";
import Personalcenter_g from "@/components/Personal/Personalcenter_g";
import Setproject from "@/components/Personal/Setproject";
import History_g from "@/components/Personal/History_g";

Vue.use(Router)

export default new Router({
  mode:'history',
  routes: [
    {
      path: '/',
      name: 'index',
      component: () => import('../components/Home')
    },
    {
      path: '/home',
      name: '首页',
      component: Home
    },
    {
      path: '/projectList',
      name: '项目列表',
      component: ProjectList
    },
    {
      path: '/projectList/projectDetail/:pro_id/',
      name: '项目详情',
      component: ProjectDetail,
      children: []
    },
    {
      path: '/donateList/input/:pro_id/:user_id/',
      name: '定向捐赠单填写',
      component: DonateList
    },
    {
      path: '/donateList/finished/:tar_don_id',
      name: '定向捐赠单填写完成',
      component: DonateListFinished
    },
    {
      path: '/logistics/:jum',
      name: '物流信息',
      component: DonateLogistics
    },
    {
      path: '/login',
      name: '登录',
      component: Login
    },
    {
      path: '/login/Findpassword',
      name: '重置密码',
      component: Findpassword,
    },
    {
      path: '/login/Register_d',
      name: '捐赠方注册',
      component: Register_d,
    },
    {
      path: '/login/Register_g',
      name: '受赠方注册',
      component: Register_g,
    },
    {
      path: '/Maincontrol/:user_id/',
      name: "个人中心",
      component: Maincontrol,
      children: [
        {
          path: '/Maincontrol/love/:user_id/',//以“/”开头的嵌套路径会被当作根路径，所以子路由上不用加“/”;在生成路由时，主路由上的path会被自动添加到子路由之前，所以子路由上的path不用在重新声明主路由上的path了。
          name: 'Love',
          component: Love
        },
        {
          path: '/Maincontrol/history/:user_id/',//以“/”开头的嵌套路径会被当作根路径，所以子路由上不用加“/”;在生成路由时，主路由上的path会被自动添加到子路由之前，所以子路由上的path不用在重新声明主路由上的path了。
          name: 'History',
          component: History
        },
        {
          path: '/Maincontrol/personalcenter/:user_id/',//以“/”开头的嵌套路径会被当作根路径，所以子路由上不用加“/”;在生成路由时，主路由上的path会被自动添加到子路由之前，所以子路由上的path不用在重新声明主路由上的path了。
          name: 'Personalcenter',
          component: Personalcenter
        },
        {
          path: '/Maincontrol/message/:user_id/',//以“/”开头的嵌套路径会被当作根路径，所以子路由上不用加“/”;在生成路由时，主路由上的path会被自动添加到子路由之前，所以子路由上的path不用在重新声明主路由上的path了。
          name: 'Message',
          component: Message
        },
        {
          path: '/Maincontrol/noRedictDonate/:user_id/',
          name: 'NoRedictDonate',
          component: NoRedictDonateCard
        }

      ]
    },
      {
        path: '/Maincontrol_g/:user_id/',
        name: "受赠方个人中心",
        component: Maincontrol_g,
        children: [
          {
            path: '/Maincontrol_g/history_g/:user_id/',//以“/”开头的嵌套路径会被当作根路径，所以子路由上不用加“/”;在生成路由时，主路由上的path会被自动添加到子路由之前，所以子路由上的path不用在重新声明主路由上的path了。
            name: 'History_g',
            component: History_g
          },
          {
            path: '/Maincontrol_g/personalcenter_g/:user_id/',//以“/”开头的嵌套路径会被当作根路径，所以子路由上不用加“/”;在生成路由时，主路由上的path会被自动添加到子路由之前，所以子路由上的path不用在重新声明主路由上的path了。
            name: 'Personalcenter_g',
            component: Personalcenter_g
          },
          {
            path: '/Maincontrol_g/message_g/:user_id/',//以“/”开头的嵌套路径会被当作根路径，所以子路由上不用加“/”;在生成路由时，主路由上的path会被自动添加到子路由之前，所以子路由上的path不用在重新声明主路由上的path了。
            name: 'Message',
            component: Message
          },
          {
            path: '/Maincontrol_g/setproject_g/:user_id/',
            name: 'Setproject',
            component: Setproject
          },
              {
                path: '/checkproject_g',
                name: 'Checkproject',
                component: Checkproject
              }

        ]
      }

  ]
})
