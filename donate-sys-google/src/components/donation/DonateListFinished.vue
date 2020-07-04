<template>
  <div>
    <el-header>
      <MainTop :header_info="header_info"></MainTop>
    </el-header>
    <SearchBar></SearchBar>
    <el-main style="width:1440px;background:#F1F1F1">
      <ProjectCard :project_detail="project_info"></ProjectCard>
      <el-row  style="margin-top: 5%">
        <el-col :span="18" :offset="3">
          <div style="background: white;text-align: left">
            <div style="padding: 5%">
              <div class="donate_block">
                <p class="demander_detail" >受赠方信息</p>
                <p class="demander_info">捐赠项目&nbsp;&nbsp;&nbsp;{{project_info.name}}</p>
                <p class="demander_info">受赠机构&nbsp;&nbsp;&nbsp;{{demander_info.name}}</p>
                <p class="demander_info">负责人姓名&nbsp;&nbsp;&nbsp;{{demander_info.contace_name}}</p>
                <p class="demander_info">联系方式&nbsp;&nbsp;&nbsp;{{demander_info.contact_tel}}</p>
                <p class="demander_info">联系地址&nbsp;&nbsp;&nbsp;{{project_info.place}}</p>
              </div>

              <div class="donate_block">
                <p class="demander_detail" >捐赠方信息</p>
                <p v-if="donater_info.if_anonymous" class="demander_info">捐赠方姓名&nbsp;&nbsp;&nbsp;爱心小谷粒</p>
                <p v-else="" class="demander_info">捐赠方姓名&nbsp;&nbsp;&nbsp;{{donater_info.name}}</p>
                <p class="demander_info">联系方式&nbsp;&nbsp;&nbsp;{{donater_info.tel}}</p>
                <p class="demander_info">身份证号&nbsp;&nbsp;&nbsp;{{donater_info.id}}</p>
              </div>

              <div class="donate_block">
                <p class="demander_detail">捐赠物资信息</p>
                <div>
                  <el-row v-for="(item,index) in supplies_info" :key="item.supply_id" class="supply_info">
                    <el-col :span="4" class="supply_item" style="line-height: normal">
                      <img :src="item.img_info[0].url" style="width: 100%">
                    </el-col>
                    <el-col :span="8" class="supply_item" style="line-height: 30px;text-align: left">
                      <p class="demander_detail">{{item.supply.name}}</p>
                      <p>{{item.rules}}</p>
                    </el-col>
                    <el-col :span="6" class="supply_item">
                      <p class="demander_detail">{{item.amount + item.supply.scal}}</p>
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
                  {{donate_msg}}
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

export default {
  components: {MainTop, MainBottom,SearchBar, ProjectCard},
  name: "DonateListFinished",
  data() {
    return {
      header_info:{
        height_line:-1,
        if_logo: false,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi:false
      },
    }
  },
  created(){
    this.getParams()
    console.log(this.donater_info)
  },
  methods :{
    getParams(){
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
    gotoProjectList(){
      console.log(this.header_info)
      // 跳转
      this.$router.push({
        name: '项目列表',
        path: '/projectList',
        params: {jum:this.header_info}
      });
    },
    gotoPresonalCenter(){
      this.$router.push({
        name: '个人中心',
        path:'/Maincontrol',
        params: this.header_info
      });
      // this.$router.push('/Maincontrol');
    },
    gotoLogistics(){
      // 合并数据
      var request_data={}
      this.$set(request_data,'donater_info',this.donater_info)
      this.$set(request_data,'project_info',this.project_info)
      this.$set(request_data,'donate_msg',this.header_info)
      console.log(request_data)
      // 跳转
      this.$router.push({
        name: '物流信息',
        path:'/logistics',
        params: {jum:request_data}
      });
    }
  },
  watch: {
      '$route': 'getParams'
  }


}
</script>

<style scoped>
  .project_info > span{
    margin-right: 10px;
  }
  .demander_detail{
    font-size: 20px;
    line-height: 23px;
  }
  .donate_block{
    border-bottom: 1px solid gray;
  }
  .supply_info{
    background: #F1F1F1;
    margin: 2%;
  }
  .supply_item{
    padding: 1%;
    line-height: 100px;
    text-align: center;
  }
</style>

