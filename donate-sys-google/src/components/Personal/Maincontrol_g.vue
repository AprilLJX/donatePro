<template>
  <div>
    <el-header>
      <MainTop :header_info="header_info"></MainTop>
    </el-header>
    <SearchBar :navi_info="navi_info"></SearchBar>
    <div style="margin-top:85px">
      <el-divider></el-divider>
    </div>

    <el-container>
      <el-aside width="30%">
        <el-col :span="18" :offset="2">
          <el-menu default-active="$route.path" router
                   background-color="rgba(70, 76, 91, 1)" text-color="#ccc"
                   style="border-radius: 2px; box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);margin-bottom:60px">
            <el-menu-item index="presonal_g" @click="$router.push('/Maincontrol_g/personalcenter_g/'+user_id)">个人信息</el-menu-item>
            <el-menu-item index="history_g" @click="$router.push('/Maincontrol_g/history_g/'+user_id)">历史信息查询</el-menu-item>
            <el-menu-item index="setproject_g" @click="$router.push('/Maincontrol_g/setproject_g/'+user_id)">捐赠项目发布</el-menu-item>
            <el-menu-item index="message_g" @click="$router.push('/Maincontrol_g/message_g/'+user_id)">消息列表</el-menu-item>
          </el-menu>
        </el-col>
      </el-aside>

      <el-main width="70%">
        <el-col :span="22" :offset="1"
                style="border-radius: 2px; box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);margin-top:-28px;margin-bottom:60px;min-height:700px">
          <router-view></router-view>
        </el-col>
      </el-main>
    </el-container>

    <el-footer>
      <el-col :span="24"></el-col>
      <MainBottom></MainBottom>
    </el-footer>
  </div>
</template>

<script>
import '../../assets/baseStyle.css'
import MainTop from "../MainTop";
import MainBottom from "../MainBottom";
import SearchBar from "../SearchBar";
import axios from 'axios'
var root_url = 'http://localhost:9090'
export default {
  name: 'Maincontrol_g',
  components: {SearchBar, MainTop, MainBottom},
  data() {
    return {
      header_info: {
        height_line: -1,
        if_login: true,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi: false
      },
      navi_info: {
        if_searchBar: false,
        navi_list: [
          {name: '首页', path: '/'},
          // {name: '项目列表',path: ''}
        ],
        now_place: '个人中心'
      },
    }
  },
  created() {
    this.user_id = this.$route.params.user_id
    console.log('loading personal center', this.user_id)
    this.get_center_info_test()

  },
  methods: {
    get_center_info_test() {
      var res = {
        donor_id: 2,
        data: {
          "donorHistory": [
            {
              "TargetId": 2,
              "DonorId": 2,
              "Category": 2,
              "DonateMaterials": "电热毯：20",
              "IfStandard": 1,
              "IfAudit": 1,
              "DonateTime": "2020-06-29T16:20:58.177748+08:00",
              "MatchPro": 2,
              "IfAnonymous": 0,
              "Message": "要开心"
            },
            {
              "TargetId": 3,
              "DonorId": 2,
              "Category": 3,
              "DonateMaterials": "蜡笔：10",
              "IfStandard": 1,
              "IfAudit": 1,
              "DonateTime": "2020-06-29T16:21:02.502261+08:00",
              "MatchPro": 3,
              "IfAnonymous": 0,
              "Message": "快快乐乐！"
            }
          ],
          "donorInfo": {
            "DonorID": 2,
            "Account": "123457",
            "Password": "111",
            "Nickname": "bob",
            "Name": "bob",
            "IdNumber": "123457",
            "CurResidence": "湖南长沙",
            "City": "长沙",
            "Avatar": "",
            "LoveValue": "100",
            "Profile": ""
          },
          "msg": "查询成功",
          "proList": [
            {
              "materials": "电热毯：50；热水袋：50；保暖衣：100",
              "proId": "2",
              "proIntro": "养老院a现需要电热毯，热水袋若干，为老人的冬天带来温暖",
              "proName": "为老人献爱心",
              "rec_donation_num": "2"
            },
            {
              "materials": "水彩笔：30套；蜡笔：30套；画本：100本",
              "proId": "3",
              "proIntro": "给自闭症儿童一只展现自己的画笔",
              "proName": "孩子们的是艺术道路",
              "rec_donation_num": "1"
            }
          ],
          "status": 200
        }
      }
      this.proList = res.data.proList
      this.donorInfo = res.data.donorInfo
      this.donorHistory = res.data.donorHistory
    },
    get_center_info() {
      //api请求方法
      let data = {"donor_id ": this.user_id};
      // axios.post(`${this.$url}/test/testRequest`,data)
      axios.post(root_url + `/donor/personalCenter`, data)
        .then(res => {
          console.log('res=>', res);
          if (res.status === 200) {
            //登陆成功，直接跳转到个人中心
            this.proList = res.proList
            this.donorInfo = res.donorInfo
            this.donorHistory = res.donorHistory
          } else {
            this.$message.error('获取信息失败~');
          }
        })
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .findpassword {
    background-color: #F1F1F1;
  }
</style>
