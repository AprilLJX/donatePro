<template>
  <div>
    <div style="background: white;text-align: left">
      <div style="padding: 5%">
        <div class="donate_block">
          <p class="demander_detail">捐赠方信息</p>
          <p class="demander_info">捐赠方姓名&nbsp;&nbsp;&nbsp;{{donor_info.nickname}}</p>
<!--          <p class="demander_info">联系方式&nbsp;&nbsp;&nbsp;{{donor.phone}}</p>-->
          <p class="demander_info">身份证号&nbsp;&nbsp;&nbsp;{{donor_info.account}}</p>
          <p class="demander_info">是否匿名&nbsp;&nbsp;
            <span v-if="if_anonymous">是，我是雷锋&nbsp;</span>
            <span v-else>实名捐赠&nbsp;</span>
            <el-switch
              v-model="if_anonymous"
              active-color="#ff4949"
              inactive-color="#C0CCDA">
            </el-switch>
          </p>
          <p class="demander_info">非定向捐赠类型&nbsp;&nbsp;
            <span v-if="form_type">医疗物资捐赠&nbsp;</span>
            <span v-else>生活物资捐赠&nbsp;</span>
            <el-switch
              v-model="form_type"
              active-color="#ffb6c1"
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
                <p class="demander_detail">{{item.supply_name}}</p>
                <p>{{item.supply_rule}}</p>
              </el-col>
              <el-col :span="6" class="supply_item">
                <el-input-number v-model="item.amount" @change="handleChange"
                                 :min="1" label="描述文字"
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
              <el-form-item v-if="form.type" label="物资类别" :label-width="formLabelWidth">
                <el-select v-model="form.supply_type" placeholder="请选择" @change="selectSupplyType">
                  <el-option v-for="i in supply_type_list" :key="i.supply_id" :label="i.type"
                             :value="i.supply_id"></el-option>
                  </el-select>
              </el-form-item>
              <el-form-item label="物资名称" :label-width="formLabelWidth">
                <el-input v-model="form.supply_name" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item v-if="form.type" label="规格标准" :label-width="formLabelWidth">
                <el-input v-model="form.supply_rule" autocomplete="off"></el-input>
              </el-form-item>
              <el-form-item label="捐赠数量" :label-width="formLabelWidth">
                <el-input-number v-model="form.amount" @change="handleChange" :min="1" label="描述文字"></el-input-number>
              </el-form-item>
              <el-form-item label="" :label-width="formLabelWidth">
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
                <el-button type="primary" @click="identifySupply(form.supply_name)">物资验证</el-button>
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
          <el-button type="danger" round @click="gotoDonateFinished">提交非定向捐赠单</el-button>
        </div>


      </div>
    </div>


  </div>
</template>

<script>
import MainTop from "../MainTop";
import MainBottom from "../MainBottom";
import SearchBar from "../SearchBar";
import ProjectCard from "../project/ProjectCard"
import axios from 'axios'
var root_url = 'http://localhost:9090'

export default {
  name: "NoRedictDonateList",
  data() {
    return {
      donate_msg: '',
      donateSuppliesFormVisible: false,
      supplies_list: [],
      form: {
        type:false,
        supply_name: '',
        supply_type: '',
        supply_rule:'',
        amount: 0,
        img_info: [{
          name: 'food.jpeg',
          url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
        }],
        identitfy: false,
      },
      formLabelWidth: '120px',
      supply_type_list: [
        {supply_id: '10001', type: '医用护目镜', rules: 'GB 14866 2006 <个人用眼护具技术要求》'},
        {supply_id: '10002', type: '医用手套', rules: 'GB0213-200/I10 119399 1:2002《一次性使用医用橡胶检查手套》'},
        {supply_id: '10003', type: '医用手术帽', rules: 'YY/T 1642-2019 《一次性使用医用防护帽》'},
        {supply_id: '10004', type: '医用鞋套', rules: 'YY/T 1633-2019 《一次性使用医用防护鞋套》'},
        {supply_id: '10005', type: '医用消毒巾', rules: 'GB 15979-2002 《一次性使用卫生用品卫生标准》'},
        {supply_id: '10006', type: '医用消毒剂', rules: 'WS 628 2018《消毒产品卫生安全评价技术要求》卫生部'},
        {supply_id: '10007', type: '红外测温仪', rules: 'GB/T 21417 1-2008《医用红外体温计 第1部分，耳腔式》'},
        {supply_id: '10008', type: '医用口罩', rules: 'YY/T 0969 2013 《一次性使用医用口罩》'},
        {supply_id: '10009', type: '医用防护服', rules: 'GB 19082-2009 《医用一次性防护服技术要求》'},
      ],
      donor_info: {},
      form_type: false,
      if_anonymous:false
    }
  },
  created() {
    this.user_id = this.$route.params.user_id
    console.log('donor_info!', this.user_id)
    this.get_personal_info()
    // this.if_anonymous=false
    console.log(this.donor_info)

  },
  methods: {
    get_personal_info() {
      if (window.sessionStorage.getItem('if_login')) {
        this.donor_info = JSON.parse(window.sessionStorage.getItem('donor_info'))
        console.log('json data', this.donor_info)
      } else {
        this.$message.error('获取信息失败，请重新登录~');
        this.$router.push('/login');
      }
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
    selectSupplyType(value) {
      var i;
      for (i of this.supply_type_list) {
        if (i.supply_id === value) {
          this.form.supply_type = i.type;
          this.form.supply_rule = i.rules;
        }
      }
    },
    closeNoAdd() {
      this.donateSuppliesFormVisible = false;
    },
    closeAdd() {
      this.donateSuppliesFormVisible = false;
      this.supplies_list.push(this.form);
      // 初始化
      console.log(this.form);
      this.form={
        type:false,
          supply_name: '',
          supply_type: '',
          supply_rule:'',
          amount: 0,
          img_info: [{
          name: 'food.jpeg',
          url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100'
        }],
          identitfy: false,
      }
      console.log(this.form);
    },
    editSupplyInfo(index) {
      // 修改添加的物资信息
      console.log('this is edit supply function~')
      console.log(this.supplies_list[index])
      this.donateSuppliesFormVisible = true
      this.form = this.supplies_list[index]
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
        this.supplies_list.splice(index, 1)
        console.log(this.supplies_list)

      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    },
    identifySupply() {
      console.log('This is identify supply fuction')
      this.form.identitfy = true
    },
    gotoDonateFinished() {
      this.$message({
        type: 'success',
        message: '提交成功'
      });
      this.$router.push("/Maincontrol/" + this.user_id);

      // // 验证：
      // var finished = false
      // //登陆验证
      // // 1.是否未添加物资
      // if (this.supplies_list.length === 0) {
      //   finished = false
      //   this.$message.error('请先添加物资~');
      // }
      // // 2.是否有未验证的信息
      // else {
      //   for (var i of this.supplies_list) {
      //     if (i.identitfy) {
      //       finished = true
      //     } else {
      //       finished = false
      //       this.$message.error('请先完成验证~');
      //       break
      //     }
      //   }
      // }
      // // 所有验证判断通过
      // if (finished) {
      //   // 合并数据
      //   var request_data = {}
      //   this.$set(request_data, 'donater_info', this.user_info)
      //   this.$set(request_data, 'demander_info', this.project_detail.demander)
      //   this.$set(request_data, 'project_info', this.project_detail)
      //   this.$set(request_data, 'supplies_info', this.supplies_list)
      //   this.$set(request_data, 'donate_msg', this.donate_msg)
      //   this.$set(request_data, 'header_info', this.header_info)
      //   console.log(request_data)
      //   // 跳转
      //   this.$router.push({
      //     name: '定向捐赠单填写完成',
      //     // name: 'mallList',
      //     params: {jum: request_data}
      //   });
      // }

    }
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
