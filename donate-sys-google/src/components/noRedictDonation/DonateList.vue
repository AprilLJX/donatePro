<template>
  <div>
    <el-header>
      <MainTop :header_info="header_info"></MainTop>
    </el-header>
    <SearchBar :navi_info="navi_info"></SearchBar>
    <el-main style="width:1440px;background:#F1F1F1">
      <ProjectCard :project_detail="project_detail"></ProjectCard>
      <el-row  style="margin-top: 5%">
        <el-col :span="18" :offset="3">
          <div style="background: white;text-align: left">
            <div style="padding: 5%">
              <div class="donate_block">
                <p class="demander_detail" >受赠方信息</p>
                <p class="demander_info">捐赠项目&nbsp;&nbsp;&nbsp;{{project_detail.name}}</p>
                <p class="demander_info">受赠机构&nbsp;&nbsp;&nbsp;{{recipient.company}}</p>
                <p class="demander_info">负责人姓名&nbsp;&nbsp;&nbsp;{{recipient.name}}</p>
                <p class="demander_info">联系方式&nbsp;&nbsp;&nbsp;{{recipient.phone}}</p>
                <p class="demander_info">联系地址&nbsp;&nbsp;&nbsp;{{recipient.com_address}}</p>
              </div>

              <div class="donate_block">
                <p class="demander_detail" >捐赠方信息</p>
                <p class="demander_info">捐赠方姓名&nbsp;&nbsp;&nbsp;{{donor.name}}</p>
                <p class="demander_info">联系方式&nbsp;&nbsp;&nbsp;{{donor.phone}}</p>
                <p class="demander_info">身份证号&nbsp;&nbsp;&nbsp;{{donor.id_number}}</p>
                <p class="demander_info">是否匿名捐赠&nbsp;&nbsp;
                  <el-switch
                    v-model="donor.if_anonymous"
                    active-color="#ff4949"
                    inactive-color="#C0CCDA">
                  </el-switch>
                </p>
              </div>

              <div class="donate_block">
                <p class="demander_detail">
                  捐赠物资信息
                  <i class="el-icon-plus" style="color: crimson" @click="donateSuppliesFormVisible = true"></i>
                </p>
                <div>
                  <el-row v-for="(item,index) in supplies_list" :key="item.supply_id" class="supply_info">
                    <el-col :span="4" class="supply_item" style="line-height: normal">
                      <img :src="item.img_info[0].url" style="width: 100%">
                    </el-col>
                    <el-col :span="8" class="supply_item" style="line-height: 30px;text-align: left">
                      <p class="demander_detail">{{item.supply.name}}</p>
                      <p>{{item.rules}}</p>
                    </el-col>
                    <el-col :span="6" class="supply_item">
                      <el-input-number v-model="item.amount" @change="handleChange"
                                       :min="1" :max="item.supply.needy_amount" label="描述文字"
                                       :value="item.amount"></el-input-number>
                    </el-col>
                    <el-col :span="3" class="supply_item">
                      <i v-if="item.identitfy" class="el-icon-check" style="color: green">已验证</i>
                      <i v-else class="el-icon-close" style="color: red">未验证</i>
                    </el-col>
                    <el-col :span="3" class="supply_item">
                      <i class="el-icon-edit" @click="editSupplyInfo(index)" style="margin-right: 10%"></i>
                      <i class="el-icon-delete" style="color: red" @click="openDelete(index)"></i>

                    </el-col>
                  </el-row>
                </div>

                <!--                弹窗-->
                <el-dialog title="添加捐赠物资" :visible.sync="donateSuppliesFormVisible">
                  <el-form :model="form">
                    <el-form-item label="物资名称" :label-width="formLabelWidth">
                      <el-select v-model="form.supply_id" placeholder="请选择" @change="selectSupply">
                        <el-option v-for="i in project_detail.demande_list.medical" :key="i.id" :label="i.name" :value="i.id" :disabled="typeof i.if_select!='undefined' && i.if_select"></el-option>
                        <el-option v-for="i in project_detail.demande_list.daliy" :key="i.id" :label="i.name" :value="i.id" :disabled="typeof i.if_select!='undefined'&& i.if_select"></el-option>
                      </el-select>
                    </el-form-item>
                    <el-form-item label="物资类别" :label-width="formLabelWidth">
                      <el-input v-model="form.supply_type" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="规格标准" :label-width="formLabelWidth">
                      <el-input v-model="form.rules" autocomplete="off"></el-input>
                    </el-form-item>
                    <el-form-item label="捐赠数量" :label-width="formLabelWidth">
                      <el-input-number v-model="form.amount" @change="handleChange" :min="1" :max="form.supply.needy_amount" label="描述文字"></el-input-number>
                    </el-form-item>
                    <el-form-item v-if="form.supply.if_need_identify" label="" :label-width="formLabelWidth">
                      <el-upload
                        class="upload-demo"
                        action="https://jsonplaceholder.typicode.com/posts/"
                        :on-preview="handlePreview"
                        :on-remove="handleRemove"
                        :file-list="form.img_info"
                        list-type="picture">
                        <el-button size="small" type="primary">点击上传</el-button>
                        <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
                      </el-upload>
                    </el-form-item>
                    <el-form-item label="" :label-width="formLabelWidth">
                      <el-button type="primary" @click="identifySupply(form.supply_id)">物资验证</el-button>
                      <i v-if="form.identitfy" class="el-icon-check" style="color: green"></i>
                      <i v-else class="el-icon-close" style="color: red"></i>
                    </el-form-item>
                  </el-form>
                  <div slot="footer" class="dialog-footer">
                    <el-button @click="closeNoAdd">取消</el-button>
                    <el-button type="primary" @click="closeAdd">添加</el-button>
                  </div>
                </el-dialog>
              </div>

              <div>
                <p class="demander_detail">爱心寄语</p>
                <p class="demander_info">
                  <el-input
                    type="textarea"
                    :rows="4"
                    placeholder="请输入内容"
                    v-model="donate_msg">
                  </el-input>
                </p>
              </div>
              <div style="margin-bottom: 3%;text-align: center">
                <el-button type="danger" round @click="gotoDonateFinished">捐赠单</el-button>
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
  components: {MainTop, MainBottom,SearchBar,ProjectCard},
  name: "DonateList",
  data () {
    return {
      header_info:{
        height_line:-1,
        if_logo: false,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi:false
      },
      navi_info:{
        if_searchBar:false,
        navi_list:[
          {name: '首页',path:'/'},
          {name: '项目列表',path: '/projectList'},
          {name:'项目详情',path:'/projectList/projectDetail'}
        ],
        now_place:'定向捐赠单填写'
      },
      // donateList_info_response:{},
      recipient:{},
      donor:{},
      project_detail: {
        id: '0100001',
        name: '武汉体育中心',
        place: '湖北省武汉市XXX区EEE街道',
        demander: {
          name: '湖北红十字',
          contace_name:"谷小红",
          contact_tel:"123456789",
          description: '湖北红十字会致力于公益事业，巴拉巴拉巴拉巴拉，湖北红十字会致力于公益事业',
        },
        receive_times: '1975',
        emergency: 4,
        needy_group: '新冠医护人员基本需求',
        demande_list: {
          medical: [
            {
              name: '口罩',
              id: '0200001',
              type: 'CN95',
              needy_amount: 2000,
              demander_amount: 10000,
              scal: '个',
              rules: 'WWWWWWWWWWWWWWWWWWW',
              img_url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
              if_need_identify:1,
            }
          ],
          daliy: [
            {
              name: '帐篷',
              id: '0200002',
              type: '----',
              needy_amount: 200,
              demander_amount: 500,
              scal: '顶',
              if_need_identify:0,
            },
            {
              name: '被子',
              id: '0200003',
              type: '1.5m',
              needy_amount: 200,
              demander_amount: 500,
              scal: '床',
              if_need_identify:0,
            }
          ]
        },
        donate_footprint: [
          {
          user: "user1",
          id: '0300001',
          time: "2020年2月20日",
          message: "今生无悔如华夏，来生还做中华家。因为你们，支撑起万千中华儿女的信念。加油，同时也要保护好自己~",
          things: "被子4床，口罩20个"
        },
          {
            user: "user2",
            id: '0300002',
            time: "2020年2月7日",
            message: "",
            things: "口罩40个"
          },
          {
            user: "user3",
            id: '0300003',
            time: "2020年2月8日",
            message: "",
            things: "口罩100个"
          }
        ]

      },
      donate_msg:'',
      donateSuppliesFormVisible: false,
      supplies_list: [],
      form: {
        supply: {},
        supply_id:'',
        supply_type:'',
        rules: '',
        amount: 0,
        img_info:[],
        identitfy:false,
      },
      formLabelWidth: '120px',
      supply_type_list:[
        {supply_id: '10001',type:'医用护目镜',rules:'GB 14866 2006 <个人用眼护具技术要求》'},
        {supply_id: '10002',type:'医用手套',rules:'GB0213-200/I10 119399 1:2002《一次性使用医用橡胶检查手套》'},
        {supply_id: '10003',type:'医用手术帽',rules:'YY/T 1642-2019 《一次性使用医用防护帽》'},
        {supply_id: '10004',type:'医用鞋套',rules:'YY/T 1633-2019 《一次性使用医用防护鞋套》'},
        {supply_id: '10005',type:'医用消毒巾',rules:'GB 15979-2002 《一次性使用卫生用品卫生标准》'},
        {supply_id: '10006',type:'医用消毒剂',rules:'WS 628 2018《消毒产品卫生安全评价技术要求》卫生部'},
        {supply_id: '10007',type:'红外测温仪',rules:'GB/T 21417 1-2008《医用红外体温计 第1部分，耳腔式》'},
        {supply_id: '10008',type:'医用口罩',rules:'YY/T 0969 2013 《一次性使用医用口罩》'},
        {supply_id: '10009',type:'医用防护服',rules:'GB 19082-2009 《医用一次性防护服技术要求》'},
      ]
    }
  },
  created(){
    this.user_id = this.$route.params.user_id
    this.getParams()
    console.log(this.header_info)
    this.donor = this.donateList_info_response.donor
    this.donor.if_anonymous=false
    this.recipient = this.donateList_info_response.recipient
  },
  methods: {
    get_donate_list_test() {
      var res ={
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

      this.donor = res.donor
      this.recipient = res.recipient
    },
    get_donate_list() {
      //api请求方法
      let data = {"donor_id ": this.user_id,"recipient_id":this.u};
      axios.post(root_url + `/projects/prodetails`, data)
        .then(res => {
          console.log('res=>', res);
          if (res.status === 200) {
            //登陆成功，直接跳转到个人中心
            this.project_detail = res.proDetails
            this.demander_info = res.recipientInfo
            this.donate_footprint = res.donorInfo
          } else {
            this.$message.error('获取信息失败~');
          }
        })
    },
    getParams(){
      // 取到路由带过来的参数
      console.log('准备数据中。。。。。')
      // 将数据放在当前组件的数据内
      const routerParams = this.$route.params.jum
      this.header_info = routerParams.header_info
      this.project_detail = routerParams.project_detail
      this.header_info.height_line = -1//哪一个link 块被选中，即表示当前页
      this.header_info.if_show_navi = false
      console.log('数据已准备好！')
    },
    handleChange(value) {
      console.log(value);
      console.log(this._data)
    },
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      console.log(fileList);
    },
    selectSupply(value){
      var i;
       // console.log(this.project_detail.demande_list.medical[0]);
      for(i of this.project_detail.demande_list.medical){
        if(i.id === value){
          this.form.supply = i;
          this.form.supply.if_need_identify = true
        }
      }
      for(i of this.project_detail.demande_list.daliy){
        if(i.id === value){
          this.form.supply = i;
          this.form.supply.if_need_identify = false
        }
      }

    },
    closeNoAdd(){
      this.donateSuppliesFormVisible = false;
    },
    closeAdd(){
      this.donateSuppliesFormVisible = false;
      this.supplies_list.push(this.form);
      this.noteSelect(this.form.supply.id, true)
      // 初始化
      console.log(this.form);
      this.form = {
        supply: {},
        supply_id:'',
        product: '',
        rules: '',
        amount: 0,
        img_info:[{
          name: 'food.jpeg',
          url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
        }],
        identitfy:false,
      }
      console.log(this.form);
    },
    noteSelect(value, t){
      var i;
      for(i of this.project_detail.demande_list.medical){
        if(i.id === value){
          i.if_select = t
        }
      }
      for(i of this.project_detail.demande_list.daliy){
        if(i.id === value){
          i.if_select = t
        }
      }
    },
    editSupplyInfo(index){
      // 修改添加的物资信息
      console.log('this is edit supply function~')
      console.log(this.supplies_list[index])
      this.donateSuppliesFormVisible = true
      this.form = this.supplies_list[index]
      // this.noteSelect(this.supplies_list[index].id,false)
      // this.supplies_list.splice(index,1)
      console.log(this.supplies_list)
    },
    openDelete(index) {
      this.$confirm('此操作将永久删除该物资, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$message({
          type: 'success',
          message: '删除成功!'
        });
        console.log(this.supplies_list[index])
        this.noteSelect(this.supplies_list[index].id,false)
        this.supplies_list.splice(index,1)
        console.log(this.supplies_list)

      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    },
    identifySupply(){
      console.log('This is identify supply fuction')
      this.form.identitfy = true
    },
    gotoDonateFinished(){
      // 验证：
      var finished = false
      //登陆验证
      // 1.是否未添加物资
      if(this.supplies_list.length===0){
        finished = false
        this.$message.error('请先添加物资~');
      }
      // 2.是否有未验证的信息
      else{
        for(var i of this.supplies_list){
          if(i.identitfy){
            finished = true
          }else{
            finished = false
            this.$message.error('请先完成验证~');
            break
          }
        }
      }
      // 所有验证判断通过
      if(finished){
        // 合并数据
        var request_data={}
        this.$set(request_data,'donater_info',this.user_info)
        this.$set(request_data,'demander_info',this.project_detail.demander)
        this.$set(request_data,'project_info',this.project_detail)
        this.$set(request_data,'supplies_info',this.supplies_list)
        this.$set(request_data,'donate_msg',this.donate_msg)
        this.$set(request_data,'header_info',this.header_info)
        console.log(request_data)
        // 跳转
        this.$router.push({
          name: '定向捐赠单填写完成',
          // name: 'mallList',
          params: {jum:request_data}
        });
      }

    }
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
