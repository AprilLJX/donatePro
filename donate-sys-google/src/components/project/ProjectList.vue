<template>
  <div>
    <el-header>
      <MainTop :header_info="header_info"></MainTop>
    </el-header>
    <SearchBar></SearchBar>
    <el-main style="width: 1440px">
      <div v-for="i in 4" :key="i" class="project_card">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>{{'武汉体育中心 ' + i }}</span>
            <el-button style="float: right; padding: 3px 0" type="text" @click="gotoProjectDetail(i)">查看详情</el-button>
          </div>
          <el-row>
            <el-col :span="10">
              <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png" class="image">
            </el-col>
            <el-col :span="10">
              <div v-for="o in 4" :key="o" class="text item">
                {{'列表内容 ' + o }}
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
          :total="500">
        </el-pagination>
      </div>
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

export default {
  components: {MainTop, MainBottom,SearchBar},
  name: "ProjectList",
  data(){
    return{
      value:true,
      project_detail: {},
      header_info:{
        height_line:2,
        if_logo: false,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi:true
      },
    };
  },
  created(){
    this.getParams()
    console.log(this.header_info)
  },
  methods: {
    getParams(){
      // 取到路由带过来的参数
      console.log('准备数据中。。。。。')
      // 将数据放在当前组件的数据内
      this.header_info  = this.$route.params
      this.header_info.height_line = 2
      console.log('数据已准备好！')
    },
    gotoProjectDetail(i) {
      //直接跳转
      window.console.log("查询成功", i);
      console.log(this.header_info)
      var request_data={}
      this.$set(request_data,'header_info',this.header_info)
      this.$set(request_data,'project_detail',this.project_detail)
      this.$router.push({
        name: '项目详情',
        path: '/projectList/projectDetail',
        params: {jum:request_data}
      });
    }
  }
};
</script>

<style scoped>
  .box-card {
    width: 60%;
  }
  .project_card{
    text-align: left;
    margin-left: 20%;
  }
  .pages_element{
    margin-top: 20px;
  }
</style>
