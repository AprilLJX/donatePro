<template>
  <div>
    <el-header>
      <MainTop :header_info="header_info"></MainTop>
    </el-header>
    <SearchBar :navi_info="navi_info"></SearchBar>
    <el-main style="width:1440px;background:#F1F1F1">
      <ProjectCard :project_detail="project_detail"></ProjectCard>
      <el-row style="margin-top: 5%">
        <el-col :span="18" :offset="3">
          <div style="background: white;text-align: left">
            <div style="padding: 5%">
              <div class="donate_block">
                <p class="demander_detail">受赠方信息</p>
                <p class="demander_info">捐赠项目&nbsp;&nbsp;&nbsp;{{project_detail.proName}}</p>
                <p class="demander_info">受赠机构&nbsp;&nbsp;&nbsp;{{recipient.company}}</p>
                <p class="demander_info">负责人姓名&nbsp;&nbsp;&nbsp;{{recipient.name}}</p>
                <p class="demander_info">联系方式&nbsp;&nbsp;&nbsp;{{recipient.phone}}</p>
                <p class="demander_info">联系地址&nbsp;&nbsp;&nbsp;{{recipient.com_address}}</p>
              </div>

              <div class="donate_block">
                <p class="demander_detail">捐赠方信息</p>
                <p v-if="donation_info.IfAnonymous" class="demander_info">捐赠方姓名&nbsp;&nbsp;&nbsp;爱心小谷粒</p>
                <p v-else="" class="demander_info">捐赠方姓名&nbsp;&nbsp;&nbsp;{{donor.name}}</p>
                <p class="demander_info">联系方式&nbsp;&nbsp;&nbsp;{{donor.phone}}</p>
                <p class="demander_info">身份证号&nbsp;&nbsp;&nbsp;{{donor.id_number}}</p>
              </div>

              <div class="donate_block">
                <p class="demander_detail">捐赠物资信息</p>
                <div>
                  <el-row v-for="(item,index) in supplies_info" :key="item.supply_id" class="supply_info">
                    <el-col :span="4" class="supply_item" style="line-height: normal">
                      <img :src="supply_img_url" style="width: 100%">
                    </el-col>
                    <el-col :span="8" class="supply_item" style="line-height: 30px;text-align: left">
                      <p class="demander_detail">{{item.type}}</p>
                      <p>{{item.rules}}</p>
                    </el-col>
                    <el-col :span="6" class="supply_item">
                      <p class="demander_detail">{{item.num}}</p>
                    </el-col>
                    <el-col :span="3" class="supply_item">
                      <i class="el-icon-check" style="color: green">已验证</i>
                    </el-col>
                  </el-row>
                </div>
              </div>

              <div>
                <p class="demander_detail">爱心寄语</p>
                <p class="demander_info">
                  {{donation_info.Message}}
                </p>
              </div>
              <div style="margin-bottom: 3%;text-align: center">
                <el-button type="danger" round @click="gotoLogistics">填写物流单</el-button>
                <el-button type="danger" round @click="gotoPresonalCenter">进入个人中心</el-button>
                <el-button type="danger" round @click="gotoProjectList">继续捐赠</el-button>
              </div>


            </div>
          </div>
        </el-col>
      </el-row>
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
import ProjectCard from "../project/ProjectCard"
import axios from "axios";

export default {
  components: {MainTop, MainBottom, SearchBar, ProjectCard},
  name: "DonateListFinished",
  data() {
    return {
      header_info: {
        height_line: -1,
        if_logo: false,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi: false
      },
      navi_info: {
        if_searchBar: false,
        navi_list: [
          {name: '首页', path: '/'},
          {name: '项目列表', path: '/projectList'},
          {name: '项目详情', path: '/projectList/projectDetail'}
        ],
        now_place: '定向捐赠单填写完成'
      },
      project_detail: {},
      donor:{},
      recipient:{},
      supplies_info:{},
      donation_info:{},
      supply_img_url:'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
    }
  },
  created() {
    // this.getParams()
    this.project_detail = JSON.parse(window.sessionStorage.getItem('pro_detail'))
    var test = this.project_detail.materials.split('；')
    this.project_detail.materials = test
    console.log('现在是完成页', this.project_detail)
    this.get_donate_info_test()

    this.navi_info.navi_list.pop()
    this.navi_info.navi_list.push({
      name: this.project_detail.proName,
      path: '/projectList/projectDetail/' + this.project_detail.proId
    })
    console.log('donate list this.navi_info', this.navi_info)
  },
  methods: {
    get_donate_info_test() {
      //请求捐赠单信息
      var res1 = {
        "donotion": {
          "TargetId": 22,
          "DonorId": 1,
          "Category": 0,
          "DonateMaterials": "电热毯:30;手套:100;",
          "IfStandard": 1,
          "IfAudit": 1,
          "DonateTime": "2020-07-02T22:40:12.911391+08:00",
          "MatchPro": 2,
          "IfAnonymous": 0,
          "Message": "加油",
          "IfFinish": 0,
          "Reasons": {
            "String": "",
            "Valid": false
          }
        },
        "materailDetail": [
          {
            "num": "30",
            "type": "电热毯"
          },
          {
            "num": "100",
            "type": "手套"
          }
        ],
        "msg": "新建捐赠单成功",
        "status": 200
      }
      this.donation_info = res1.donotion
      this.supplies_info = res1.materailDetail
      //根据捐赠单返回的受赠方id和捐赠方id获取其相关信息
      var res2 = {
        "donor": {
          "id_number": "123456",
          "name": "alice",
          "phone": "123456"
        },
        "msg": "返回成功",
        "recipient": {
          "com_address": "北京",
          "company": "医院a",
          "name": "医院a",
          "phone": "123456"
        },
        "status": 200
      }
      this.donor = res2.donor
      this.recipient = res2.recipient

    },
    get_donate_info() {
      //请求捐赠单信息

      //api请求方法
      let data = {"donor_id ": this.user_id, "recipient_id": this.project_detail.demander_id};
      axios.post(root_url + `/donor/donateList`, data)
        .then(res => {
          console.log('res=>', res);
          if (res.status === 200) {
            //根据捐赠单返回的受赠方id和捐赠方id获取其相关信息
            this.donor = res.donor
            this.recipient = res.recipient
          } else {
            this.$message.error('获取信息失败~');
          }
        })
    },
    getParams() {
      // 取到路由带过来的参数
      console.log(this.$route.params)
      const routerParams = this.$route.params.jum
      console.log('准备数据中。。。。。')
      // 将数据放在当前组件的数据内
      this.donater_info = routerParams.donater_info;
      this.demander_info = routerParams.demander_info;
      // this.project_info = routerParams.project_info;
      this.supplies_info = routerParams.supplies_info;
      this.donate_msg = routerParams.donate_msg
      this.header_info = routerParams.header_info
      console.log('数据已准备好！')
    },
    gotoProjectList() {
      console.log(this.header_info)
      // 跳转
      this.$router.push({
        name: '项目列表',
        path: '/projectList',
        params: {jum: this.header_info}
      });
    },
    gotoPresonalCenter() {
      var test = window.sessionStorage.getItem('donor_info')
      var user_id = JSON.parse(test).donor_id
      this.$router.push("/Maincontrol/" + user_id);
      // this.$router.push({
      //   name: '个人中心',
      //   path: '/Maincontrol',
      //   params: this.header_info
      // });
      // this.$router.push('/Maincontrol');
    },
    gotoLogistics() {
      // 合并数据
      var request_data = {}
      this.$set(request_data, 'donater_info', this.donater_info)
      this.$set(request_data, 'project_info', this.project_info)
      this.$set(request_data, 'donate_msg', this.header_info)
      console.log(request_data)
      // 跳转
      this.$router.push({
        name: '物流信息',
        path: '/logistics',
        params: {jum: request_data}
      });
    }
  },
  watch: {
    '$route': 'getParams'
  }


}
</script>

<style scoped>
  .project_info > span {
    margin-right: 10px;
  }

  .demander_detail {
    font-size: 20px;
    line-height: 23px;
  }

  .donate_block {
    border-bottom: 1px solid gray;
  }

  .supply_info {
    background: #F1F1F1;
    margin: 2%;
  }

  .supply_item {
    padding: 1%;
    line-height: 100px;
    text-align: center;
  }
</style>

