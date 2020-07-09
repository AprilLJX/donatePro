<template>
  <div>
    <el-header>
      <MainTop :header_info="header_info"></MainTop>
    </el-header>
    <SearchBar :navi_info="navi_info"></SearchBar>
    <el-main style="width: 1440px">
      <el-col :span="12" :offset="6">


      <div v-for="i in project_detail" :key="i.proID" class="project_card">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>{{ i.proName }}</span>
            <el-button style="float: right; padding: 3px 0" type="text" @click="gotoProjectDetail(i)">查看详情</el-button>
          </div>
          <el-row>
            <el-col :span="6">
<!--              <img :src="require('../../assets/project.png')" class="image">-->
<!--              <img src="src/assets/project.png" class="image">-->
              <img :src="img_url">
            </el-col>
            <el-col :span="14">
              <div class="text item">
                {{ i.proIntro }}
              </div>
            </el-col>
          </el-row>
        </el-card>
      </div>
<!--      10 item for one page-->
      <div class="pages_element">
        <el-pagination
          background :hide-on-single-page="value"
          layout="prev, pager, next"
          :total="project_list_response.proNum">
        </el-pagination>
      </div>
      </el-col>
    </el-main>
    <el-footer>
      <MainBottom></MainBottom>
    </el-footer>

  </div>
</template>

<script>
import MainTop from "../MainTop";
import MainBottom from "../MainBottom";
import SearchBar from "../SearchBar";

import project from '../../assets/project.png';

import axios from 'axios'
var root_url = 'http://localhost:9090'

export default {
  components: {MainTop, MainBottom,SearchBar},
  name: "ProjectList",
  data(){
    return{
      value:true,
      project_detail: {},
      img_url:project,
      header_info:{
        height_line:2,
        if_login: false,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi:true
      },
      navi_info:{
        if_searchBar:false,
        navi_list:[
          {name: '首页',path:'/'},
          // {name: '项目列表',path: ''}
        ],
        now_place:'项目列表'
      }
    };
  },
  created(){
    this.getParams()
    console.log(this.header_info)
    this.get_project_list_test()
    // 获取列表数据
    // this.get_project_list()
  },
  mounted(){
    //api GET方法：localhost:9090/projects/prolist
    // this.get_project_list()
  },
  methods: {
    get_project_list_test(){
      this.project_list_response={
        "msg": "查询成功",
        "proList": [
          {
            "proID": 1,
            "proIntro": "医院a现急需口罩等医疗物资",
            "proName": "急需口罩志愿"
          },
          {
            "proID": 2,
            "proIntro": "养老院a现需要电热毯，热水袋若干，为老人的冬天带来温暖",
            "proName": "为老人献爱心"
          },
          {
            "proID": 3,
            "proIntro": "给自闭症儿童一只展现自己的画笔",
            "proName": "孩子们的是艺术道路"
          }
        ],
        "proNum": 3,
        "status": 200
      }
      this.project_detail = this.project_list_response.proList
    },
    get_project_list(){
      axios.get(root_url+'/projects/prolist')
        .then((response)=>{
          console.log(response.data);
          this.re_status = response.status;
          this.show = response.data;
          console.log(root_url+'/projects/prolist')
          console.log(this.show)
          // this.project_list_response = response.data
          // this.project_detail = this.project_list_response.proList
        })
        .catch((error)=> {
          console.log(error);
        });
    },
    // gotoPage(path){
    //   console.log(path)
    // },
    getParams(){
      // 取到路由带过来的参数
      console.log('准备数据中。。。。。')
      // 将数据放在当前组件的数据内
      this.header_info  = this.$route.params
      // this.navi_info = this.$route.params
      this.header_info.height_line = 2
      console.log('数据已准备好！')
    },
    gotoProjectDetail(i) {
      //直接跳转
      window.console.log("查询成功", i);
      this.$router.push('/projectList/projectDetail/'+i.proID);
    }
  }
};
</script>

<style scoped>
  .box-card {
    /*width: 60%;*/
  }
  .project_card{
    text-align: left;
    /*margin-left: 20%;*/
  }
  .pages_element{
    margin-top: 20px;
  }
</style>
