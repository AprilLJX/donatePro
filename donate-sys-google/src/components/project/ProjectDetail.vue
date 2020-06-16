<template>
  <div>
    <el-header>
      <MainTop v-bind:if_logo="true" v-bind:user_type="'0'"></MainTop>
    </el-header>
    <SearchBar></SearchBar>

    <el-main style="width:1440px;background:#F1F1F1 ">
      <el-row>
        <el-col :span="12" :offset="3">
          <div style="background: white;height: 320px">
            <el-row>
              <el-col :span="12">
                <div style="width: 90%;margin: 5%;text-align:center;vertical-align:middle;">
                  <img src='https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg' class="image" style="width: 100%;height:60%">
                </div>
              </el-col>
              <el-col :span="12">
                <div style="margin: 5%;">
                  <div class="need_staff" style="text-align: left">
                    <p class="project_name">{{project_detail.name}}</p>
                    <p class="project_info">
                      <span>
                        <i class="el-icon-location" style="color: red"></i>
                        {{project_detail.place}}
                      </span>
                    </p>
                    <p class="project_info">
                      <span>{{'发起方： '+project_detail.demander.name}}</span>
                    </p>
                    <p class="project_info">
                      <span>
                        紧急度：
                        <i v-for="i in project_detail.emergency" :key="i" class="el-icon-star-on" style="color: red"></i>
                      </span>
                    </p>
                    <p class="project_info">
                      <span>{{"参与度： "+project_detail.receive_times +" 人次"}}</span>
                    </p>
                    <li v-for="(item,idx) in needy_list" style="color: crimson;line-height: 20px">
                      <span>{{item[0]+': '+ item[1]+' '+ item[2]}}</span>
                    </li>
                  </div>
                </div>
              </el-col>
            </el-row>
            <el-row>
              <div style="margin-bottom: 3%">
                <el-button type="danger" round @click="gotoDonate()">我要捐赠</el-button>
              </div>

            </el-row>

          </div>
        </el-col>
        <el-col :span="5" :offset="1">
          <div style="background: white;height: 320px">

              <div class="block">
                <el-row>
                  <el-avatar :size="100" :src="circleUrl" style="margin: 5%"></el-avatar>
                </el-row>
                <el-row>
                  <span>{{project_detail.demander.name}}</span>
                </el-row>
              </div>
            <el-row>
              <p style="margin-right: 5%;margin-left: 5%;text-align: left;line-height: 30px">
                {{project_detail.demander.description}}
              </p>
            </el-row>

          </div>
        </el-col>
      </el-row>
      <el-row style="margin-top: 5%">
        <el-col :span="12" :offset="3">
          <div style="background: white;text-align: left;">
            <div style="padding: 5%">
              <p class="demander_detail" >物资应用人群</p>
              <p class="demander_info">{{project_detail.needy_group}}</p>
              <p class="demander_detail">物资详情</p>
              <p class = "demander_type">医疗物资</p>
              <div class="demander_info">
                <div v-for="(item,idx) in project_detail.demande_list.medical">
                  <p>{{'【 '+item.name+' 】'}}</p>
                  <p>{{'型号：'+item.type}}</p>
                  <p>{{'需求数量：'+item.demander_amount}}</p>
                  <p>{{'剩余需求：'+item.needy_amount}}</p>
                  <p>{{'规格：'+item.rules}}</p>
                  <p>参考图片：</p>
                  <img :src="item.img_rul">
                </div>
              </div>
              <p class = "demander_type">生活物资</p>
              <div class="demander_info">
                <div v-for="(item,idx) in project_detail.demande_list.daliy">
                  <p>{{'【 '+item.name+' 】'}}</p>
                  <p>{{'型号：'+item.type}}</p>
                  <p>{{'需求数量：'+item.demander_amount}}</p>
                  <p>{{'剩余需求：'+item.needy_amount}}</p>
                </div>
              </div>
            </div>

          </div>
        </el-col>
        <el-col :span="5" :offset="1">
          <div style="background: white;text-align: left">
            <div class="footprint_titel">
              爱心足迹
            </div>
            <div v-for="l in project_detail.donate_footprint" :key="l.id" style="padding: 7%">
              <el-row>
                <el-col :span="4">
                  <div>
                    <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                  </div>
                </el-col>
                <el-col :span="20">
                  <p class="footprint_info">{{l.user}}</p>
                  <p class="footprint_info">{{l.things}}</p>
                </el-col>
              </el-row>
              <el-row>
                <div v-if="l.message">
                  <p>{{l.message}}</p>
                </div>
              </el-row>
              <el-row>
                <p class="footprint_info">{{l.time}}</p>
              </el-row>
            </div>
            <el-collapse v-model="activeNames" @change="handleChange">
              <el-collapse-item title="查看更多" name="1" style="padding: 7%">
                <div v-for="n in project_detail.donate_footprint" :key="n.id">
                  <el-row>
                    <el-col :span="4">
                      <div>
                        <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"></el-avatar>
                      </div>
                    </el-col>
                    <el-col :span="20">
                      <p class="footprint_info">{{n.user}}</p>
                      <p class="footprint_info">{{n.things}}</p>
                    </el-col>
                  </el-row>
                  <el-row>
                    <div v-if="n.message">
                      <p>{{n.message}}</p>
                    </div>
                  </el-row>
                  <el-row>
                    <p class="footprint_info">{{n.time}}</p>
                  </el-row>
                </div>
              </el-collapse-item>
            </el-collapse>
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
export default {
  components: {MainTop, MainBottom,SearchBar},
  name: "ProjectDetail",
  data () {
    return {
      circleUrl: "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png",
      img_url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
      project_detail:{
        id:'0100001',
        name:'武汉体育中心',
        place:'湖北省武汉市XXX区EEE街道',
        demander:{
          name:'湖北红十字',
          description:'湖北红十字会致力于公益事业，巴拉巴拉巴拉巴拉，湖北红十字会致力于公益事业',
        },
        receive_times:'1975',
        emergency:4,
        needy_group:'新冠医护人员基本需求',
        demande_list:{
          medical:[
            {
              name:'口罩',
              id:'0200001',
              type:'CN95',
              needy_amount:2000,
              demander_amount:10000,
              scal:'个',
              rules:'WWWWWWWWWWWWWWWWWWW',
              img_url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
            }
            ],
          daliy:[
            {
              name:'帐篷',
              id:'0200002',
              type:'----',
              needy_amount:200,
              demander_amount:500,
              scal:'顶',
            },
            {
              name:'被子',
              id:'0200003',
              type:'1.5m',
              needy_amount:200,
              demander_amount:500,
              scal:'顶',
            }
          ]
        },
        donate_footprint: [{
          user:"user1",
          id:'0300001',
          time:"2020年2月20日",
          message:"今生无悔如华夏，来生还做中华家。因为你们，支撑起万千中华儿女的信念。加油，同时也要保护好自己~",
          things:"被子4床，口罩20个"
          },
          {
            user:"user2",
            id:'0300002',
            time:"2020年2月7日",
            message:"",
            things:"口罩40个"
          },
          {
            user:"user3",
            id:'0300003',
            time:"2020年2月8日",
            message:"",
            things:"口罩100个"
          }
        ]

      },
      needy_list:[
        ['口罩','1000','个'],
        ['帐篷','200','顶'],
        ['被子','200','床']
      ],
      activeNames: ['0']

    }
  },
  methods:{
      handleChange(val) {
        console.log(val);
      },
      gotoDonate() {
        //直接跳转
        window.console.log("跳转填写页");
        this.$router.push('/projectDetail/donateList/finished');
      }
  }
}
</script>

<style scoped>
  .project_name{
    font-size: 20px;
    line-height: 26px;
    color: crimson;
  }
  .project_info{
    font-size: 12px;
    line-height: 12px;
    color: gray;
  }
  .project_info > span{
    margin-right: 10px;
  }
  .demander_detail{
    font-size: 20px;
    line-height: 23px;
  }
  .demander_type{
    color: crimson;
  }
  .footprint_titel{
    line-height: 23px;
    font-size: 20px;
    font-size: 20px;
    color: crimson;
    text-align: center!important;
    padding: 5%;
  }
  .footprint_info{
    color: gray;
    font-size: 14px;
    margin-top: 0;
    margin-bottom: 0;
    line-height: 15px;
  }
</style>
