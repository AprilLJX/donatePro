<template>
  <div>
    <el-header>
      <MainTop :header_info="header_info"></MainTop>
    </el-header>
    <!--如果有物流信息则显示当前物流情况，如无，则输入物流信息-->
    <el-main style="width:1440px;background:#F1F1F1">
      <ProjectCard :project_detail="project_info"></ProjectCard>
      <el-row style="margin-top: 5%">
        <el-col :span="18" :offset="3">
          <div style="background: white;text-align: left">
            <div style="padding: 5%">
              <div v-if="if_logistics">
                <div>
                  <p class="demander_detail">物流进度</p>
                  <p class="demander_info">物流公司&nbsp;&nbsp;&nbsp;{{logistics.company}}</p>
                  <p class="demander_info">物流单号&nbsp;&nbsp;&nbsp;{{logistics.id}}</p>
                  <div class="block" style="padding: 5%">
                    <el-timeline>
                      <el-timeline-item
                        v-for="(activity, index) in logistics_info"
                        :key="index"
                        :icon="activity.icon"
                        :type="activity.type"
                        :color="activity.color"
                        :size="activity.size"
                        :timestamp="activity.timestamp">
                        {{activity.content}}
                      </el-timeline-item>
                    </el-timeline>
                  </div>
                </div>

                <div style="margin-bottom: 3%;text-align: center">
                  <el-button type="danger" round @click="gotoPresonalCenter">进入个人中心</el-button>
                </div>
              </div>

              <div v-else="">
                <el-row style="text-align: center">
                  <div style="text-align: left;padding-bottom: 5%">
                    <p class="demander_detail">物流信息录入</p>
                  </div>
                  <p class="demander_info" >物流公司</p>
                  <el-autocomplete
                    class="inline-input"
                    v-model="logistics.company"
                    :fetch-suggestions="querySearch"
                    placeholder="请输入内容"
                    @select="handleSelect"
                  ></el-autocomplete>
                  <p class="demander_info">物流单号</p>
                  <el-input
                    placeholder="请输入物流单号"
                    v-model="logistics.id"
                    clearable style="width: auto">
                  </el-input>
                </el-row>

                <div style="padding: 3%;text-align: center">
                  <el-button type="danger" round @click="submitLogistics">提交物流信息</el-button>
                </div>
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
  components: {MainTop, MainBottom, SearchBar, ProjectCard},
  name: "DonateLogistics",
  data() {
    return {
      header_info:{
        height_line:-1,
        if_logo: false,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi:false
      },
      if_logistics: false,
      logistics: {
        company: '韵达快递',
        id: '242425325265343'
      },
      logistics_info: [
        {
          content: '支持使用图标',
          timestamp: '2018-04-12 20:46',
          size: 'large',
          type: 'primary',
          icon: 'el-icon-more'
        }, {
          content: '支持自定义颜色',
          timestamp: '2018-04-03 20:46',
          color: '#0bbd87'
        }, {
          content: '支持自定义尺寸',
          timestamp: '2018-04-03 20:46',
          size: 'large'
        }, {
          content: '默认样式的节点',
          timestamp: '2018-04-03 20:46'
        }],
      restaurants: [],
      state1: '',
      state2: ''
    }
  },
  created() {
    this.getParams()
    console.log(this.donater_info)
  },
  methods: {
    getParams() {
      // 取到路由带过来的参数
      console.log(this.$route.params)
      const routerParams = this.$route.params.jum
      console.log('准备数据中。。。。。')
      // 将数据放在当前组件的数据内
      this.donater_info = routerParams.donater_info;
      this.project_info = routerParams.project_info;
      this.header_info = routerParams.header_info
      console.log('数据已准备好！')
    },
    querySearch(queryString, cb) {
      var restaurants = this.restaurants;
      var results = queryString ? restaurants.filter(this.createFilter(queryString)) : restaurants;
      // 调用 callback 返回建议列表的数据
      cb(results);
    },
    createFilter(queryString) {
      return (restaurant) => {
        return (restaurant.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0);
      };
    },
    loadAll() {
      return [
        {"value": "三全鲜食（北新泾店）", "address": "长宁区新渔路144号"},
        {"value": "Hot honey 首尔炸鸡（仙霞路）", "address": "上海市长宁区淞虹路661号"},
        {"value": "新旺角茶餐厅", "address": "上海市普陀区真北路988号创邑金沙谷6号楼113"},
        {"value": "泷千家(天山西路店)", "address": "天山西路438号"},
        {"value": "胖仙女纸杯蛋糕（上海凌空店）", "address": "上海市长宁区金钟路968号1幢18号楼一层商铺18-101"},
        {"value": "贡茶", "address": "上海市长宁区金钟路633号"},
        {"value": "豪大大香鸡排超级奶爸", "address": "上海市嘉定区曹安公路曹安路1685号"},
        {"value": "茶芝兰（奶茶，手抓饼）", "address": "上海市普陀区同普路1435号"},
        {"value": "十二泷町", "address": "上海市北翟路1444弄81号B幢-107"},
        {"value": "星移浓缩咖啡", "address": "上海市嘉定区新郁路817号"},
        {"value": "阿姨奶茶/豪大大", "address": "嘉定区曹安路1611号"},
        {"value": "新麦甜四季甜品炸鸡", "address": "嘉定区曹安公路2383弄55号"},
        {"value": "Monica摩托主题咖啡店", "address": "嘉定区江桥镇曹安公路2409号1F，2383弄62号1F"},
        {"value": "浮生若茶（凌空soho店）", "address": "上海长宁区金钟路968号9号楼地下一层"},
        {"value": "NONO JUICE  鲜榨果汁", "address": "上海市长宁区天山西路119号"},
        {"value": "CoCo都可(北新泾店）", "address": "上海市长宁区仙霞西路"},
        {"value": "快乐柠檬（神州智慧店）", "address": "上海市长宁区天山西路567号1层R117号店铺"},
        {"value": "Merci Paul cafe", "address": "上海市普陀区光复西路丹巴路28弄6号楼819"},
        {"value": "猫山王（西郊百联店）", "address": "上海市长宁区仙霞西路88号第一层G05-F01-1-306"},
        {"value": "枪会山", "address": "上海市普陀区棕榈路"},
        {"value": "纵食", "address": "元丰天山花园(东门) 双流路267号"},
        {"value": "钱记", "address": "上海市长宁区天山西路"},
        {"value": "壹杯加", "address": "上海市长宁区通协路"},
        {"value": "唦哇嘀咖", "address": "上海市长宁区新泾镇金钟路999号2幢（B幢）第01层第1-02A单元"},
        {"value": "爱茜茜里(西郊百联)", "address": "长宁区仙霞西路88号1305室"},
        {"value": "爱茜茜里(近铁广场)", "address": "上海市普陀区真北路818号近铁城市广场北区地下二楼N-B2-O2-C商铺"},
        {"value": "鲜果榨汁（金沙江路和美广店）", "address": "普陀区金沙江路2239号金沙和美广场B1-10-6"},
        {"value": "开心丽果（缤谷店）", "address": "上海市长宁区威宁路天山路341号"},
        {"value": "超级鸡车（丰庄路店）", "address": "上海市嘉定区丰庄路240号"},
        {"value": "妙生活果园（北新泾店）", "address": "长宁区新渔路144号"},
        {"value": "香宜度麻辣香锅", "address": "长宁区淞虹路148号"},
        {"value": "凡仔汉堡（老真北路店）", "address": "上海市普陀区老真北路160号"},
        {"value": "港式小铺", "address": "上海市长宁区金钟路968号15楼15-105室"},
        {"value": "蜀香源麻辣香锅（剑河路店）", "address": "剑河路443-1"},
        {"value": "北京饺子馆", "address": "长宁区北新泾街道天山西路490-1号"},
        {"value": "饭典*新简餐（凌空SOHO店）", "address": "上海市长宁区金钟路968号9号楼地下一层9-83室"},
        {"value": "焦耳·川式快餐（金钟路店）", "address": "上海市金钟路633号地下一层甲部"},
        {"value": "动力鸡车", "address": "长宁区仙霞西路299弄3号101B"},
        {"value": "浏阳蒸菜", "address": "天山西路430号"},
        {"value": "四海游龙（天山西路店）", "address": "上海市长宁区天山西路"},
        {"value": "樱花食堂（凌空店）", "address": "上海市长宁区金钟路968号15楼15-105室"},
        {"value": "壹分米客家传统调制米粉(天山店)", "address": "天山西路428号"},
        {"value": "福荣祥烧腊（平溪路店）", "address": "上海市长宁区协和路福泉路255弄57-73号"},
        {"value": "速记黄焖鸡米饭", "address": "上海市长宁区北新泾街道金钟路180号1层01号摊位"},
        {"value": "红辣椒麻辣烫", "address": "上海市长宁区天山西路492号"},
        {"value": "(小杨生煎)西郊百联餐厅", "address": "长宁区仙霞西路88号百联2楼"},
        {"value": "阳阳麻辣烫", "address": "天山西路389号"},
        {"value": "南拳妈妈龙虾盖浇饭", "address": "普陀区金沙江路1699号鑫乐惠美食广场A13"}
      ];
    },
    handleSelect(item) {
      console.log(item);
    },
    gotoPresonalCenter() {
      this.$router.push('/Maincontrol');
    },
    submitLogistics(){
      console.log(this.logistics);
      // 如果提交成功则修改if_logistics
      this.if_logistics = true
    }
  },
  mounted() {
    this.restaurants = this.loadAll();
  }
}
</script>

<style scoped>
  .project_info > span {
    margin-right: 10px;
    text-align: center;
  }

  .demander_detail {
    font-size: 20px;
    line-height: 23px;
  }

</style>
