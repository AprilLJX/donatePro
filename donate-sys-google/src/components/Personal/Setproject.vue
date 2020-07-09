<template>
  <div style="width:100%">

    <el-divider></el-divider>
    <span>捐赠项目发布</span>
    <el-divider></el-divider>

    <div class="donate_block">
      <div>
        <el-row style="margin-top: 5px">
          <el-col :span="8"><p>项目名称</p></el-col>
          <el-col :span="12">
            <el-input
              placeholder="请填写项目名称"
              v-model="project_detail.pro_name"
              minlength="5"
            maxlength="30">
            </el-input>
          </el-col>
        </el-row>
        <el-row style="margin-top: 5px">
          <el-col :span="8"><p>项目类型</p></el-col>
          <el-col :span="12" style="text-align: left">
              <el-switch v-model="project_detail.category"
                active-text="医疗" active-color="#ffb6c1" active-value='1'
                inactive-text="非医疗" inactive-color="#C0CCDA" inactive-value='0'
              style="height: 50px">
              </el-switch>
          </el-col>
        </el-row>
        <el-row style="margin-top: 5px">
          <el-col :span="8"><p>项目简介</p></el-col>
          <el-col :span="12" style="text-align: left">
            <el-input
              type="textarea" :rows="4"
              placeholder="请填写项目简介" minlength="20" maxlength="200"
              v-model="project_detail.introduction" style="width:400px">
            </el-input>
          </el-col>
        </el-row>
        <el-row style="margin-top: 5px">
          <el-col :span="8"> <p>项目截止时间</p> </el-col>
          <el-col :span="12" style="text-align: left">
            <el-date-picker
              v-model="project_detail.cut_off_time"
              type="date"
              placeholder="请选择日期">
            </el-date-picker>
          </el-col>
        </el-row>
        <el-row style="margin-top: 5px">
          <el-col :span="8"> <p>项目紧急度</p> </el-col>
          <el-col :span="12" style="text-align: left">
            <el-select v-model="project_detail.emergency_degree" placeholder="请选择">
              <el-option
                v-for="item in emergency_options"
                :key="item.value"
                :label="item.label"
                :value="item.value">
              </el-option>
            </el-select>
          </el-col>
        </el-row>

        <el-divider></el-divider>
        <el-row>
          <el-col :span="8"><p>详细地址</p></el-col>
          <el-col :span="12" style="text-align: left">
            <el-cascader
              style="width:400px"
              :options="options"
              v-model="form.selectedOptions"
              @change="handleChange2">
            </el-cascader>
            <el-input
              type="textarea"
              :rows="4"
              placeholder="详细地址"
              v-model="project_detail.rec_address" style="width:400px;margin-top: 5px">
            </el-input>
          </el-col>
        </el-row>
        <el-divider></el-divider>
        <el-row>
          <el-col :span="8">
            <p>项目缩略图</p>
          </el-col>
          <el-col :span="12" style="text-align: left">
            <el-upload ref="upload" action="#"
                       accept="image/png,image/gif,image/jpg,image/jpeg"
                       list-type="picture-card"
                       :limit=limitNum
                       :auto-upload="false"
                       :on-exceed="handleExceed"
                       :before-upload="handleBeforeUpload"
                       :on-preview="handlePictureCardPreview"
                       :on-remove="handleRemove">
              <i class="el-icon-plus"></i>
            </el-upload>
            <el-dialog :visible.sync="dialogVisible">
              <img width="100%" :src="dialogImageUrl" alt="">
            </el-dialog>
          </el-col>
        </el-row>
        <el-divider></el-divider>
        <p class="demander_detail">
          物资需求信息
          <i class="el-icon-plus" style="color: crimson" @click="donateSuppliesFormVisible = true"></i>
        </p>
        <el-divider></el-divider>
        <el-row>
          <el-col :span="8"><p>物资</p></el-col>
          <el-col :span="6"><p>参数</p></el-col>
          <el-col :span="6"><p>数量</p></el-col>
          <el-col :span="4"><p>删除</p></el-col>
        </el-row>
        <el-row v-for="(item,index) in supplies_list" :key="item.supply_id" class="supply_info">
          <el-col :span="8" class="supply_item" style="line-height: 30px;text-align: center">
            <p class="demander_detail">{{item.supply_name}}</p>
          </el-col>
          <el-col :span="6" class="supply_item" style="line-height: 30px;text-align: center">
            <p>{{item.supply_rule}}</p>
          </el-col>
          <el-col :span="6" class="supply_item">
            <el-input-number v-model="item.amount" @change="handleChange"
                             :min="1" label="描述文字"
                             :value="item.amount"></el-input-number>
          </el-col>
          <el-col :span="4" class="supply_item">
            <i class="el-icon-delete" style="color: red" @click="openDelete(index)"></i>
          </el-col>
        </el-row>


        <!--<el-button size="small" type="primary" @click="uploadFile">立即上传</el-button>
        <el-button size="small">取消</el-button>-->

        <el-divider></el-divider>

        <div style="margin-bottom: 3%;text-align: center">
          <el-button type="danger" round @click="gotoFinished">确认提交</el-button>
        </div>

      </div>

      <!--                弹窗-->
      <el-dialog title="添加捐赠物资" :visible.sync="donateSuppliesFormVisible" style="text-align: left">
        <el-form :model="form">
          <el-form-item v-if="project_detail.category==='1'" label="物资类别" :label-width="formLabelWidth">
            <el-select v-model="form.supply_type" placeholder="请选择" @change="selectSupplyType">
              <el-option v-for="i in supply_type_list" :key="i.supply_id" :label="i.type"
                         :value="i.supply_id"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="物资名称" :label-width="formLabelWidth">
            <el-input v-model="form.supply_name" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item v-if="project_detail.category==='1'" label="规格标准" :label-width="formLabelWidth">
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
              :on-remove="handleRemoveItem"
              :file-list="form.img_info"
              list-type="picture">
              <el-button size="small" type="primary">点击上传</el-button>
              <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
            </el-upload>
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="closeNoAdd">取消</el-button>
          <el-button type="primary" @click="closeAdd">添加</el-button>
        </div>
      </el-dialog>
<!--      <el-dialog title="添加捐赠物资" :visible.sync="donateSuppliesFormVisible">-->
<!--        <el-form :model="form">-->
<!--          <el-form-item label="物资名称" :label-width="formLabelWidth">-->
<!--            <el-select v-model="form.supply_id" placeholder="请选择" @change="selectSupply">-->
<!--              <el-option v-for="i in project_detail.demande_list.medical" :key="i.id" :label="i.name"-->
<!--                         :value="i.id"></el-option>-->
<!--              <el-option v-for="i in project_detail.demande_list.daliy" :key="i.id" :label="i.name"-->
<!--                         :value="i.id"></el-option>-->
<!--            </el-select>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="规格信息" :label-width="formLabelWidth">-->
<!--            <el-input v-model="form.rules" autocomplete="off"></el-input>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="捐赠数量" :label-width="formLabelWidth">-->
<!--            <el-input-number v-model="form.amount" @change="handleChange" :min="1" :max="form.supply.needy_amount"-->
<!--                             label="描述文字"></el-input-number>-->
<!--          </el-form-item>-->
<!--        </el-form>-->
<!--        <div slot="footer" class="dialog-footer">-->
<!--          <el-button type="primary" @click="close2">添加</el-button>-->
<!--        </div>-->
<!--      </el-dialog>-->
    </div>

  </div>
</template>

<script>
import {provinceAndCityData, CodeToText, TextToCode} from 'element-china-area-data';
import axios from 'axios'
var root_url = 'http://localhost:9090'

export default {
  name: 'Setproject',
  data() {
    return {
      dialogImageUrl: '',
      dialogVisible: false,
      formLabelWidth: '80px',
      limitNum: 3,
      options: provinceAndCityData,
      emergency_options:[
        { value: '1', label: '日常' },
        { value: '2', label: '紧急' },
        { value: '3', label: '很紧急' },
        { value: '4', label: '非常紧急' },
        { value: '5', label: '十万火急' }
        ],
      project_detail: {
        pro_name:'',
        category:'',
        introduction:'',
        rec_address:'',
        emergency_degree:'',
        cut_off_time:'',
      },
      donate_msg: '',
      donateSuppliesFormVisible: false,
      supplies_list: [],
      // form: {
      //   supply: {},
      //   supply_id: '',
      //   product: '',
      //   rules: '',
      //   amount: 0,
      //   add_identitfy: 0,
      // },
      form: {
        type:false,
        supply_name: '',
        supply_type: '',
        supply_rule:'',
        amount: 0,
        img_info: [],
        identitfy: false,
      },
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
    }
  },
  created() {
    this.getParams()
    console.log(this.header_info)
  },
  methods: {
    getParams() {

    },
    handleChange(value) {
      console.log(value);
      console.log(this._data)
    },
    handleRemoveItem(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      console.log(fileList);
    },
    selectSupply(value) {
      var i;
      // console.log(this.project_detail.demande_list.medical[0]);
      for (i of this.project_detail.demande_list.medical) {
        if (i.id === value) {
          this.form.supply = i;
        }
      }
      for (i of this.project_detail.demande_list.daliy) {
        if (i.id === value) {
          this.form.supply = i;
        }
      }
    },
    identifySupply() {
      console.log('This is identify supply fuction')
      this.form.identitfy = true
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
        this.supplies_list.splice(index, 1)
        console.log(this.supplies_list)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    },
    gotoDonateFinished() {
    },
    handleChange2(value) {
      console.log(value);
    },
    handleBeforeUpload(file) {
      console.log('before')
      if (!(file.type === 'image/png' || file.type === 'image/gif' || file.type === 'image/jpg' || file.type === 'image/jpeg')) {
        this.$notify.warning({
          title: '警告',
          message: '请上传格式为image/png, image/gif, image/jpg, image/jpeg的图片'
        })
      }
      let size = file.size / 1024 / 1024 / 2
      if (size > 2) {
        this.$notify.warning({
          title: '警告',
          message: '图片大小必须小于2M'
        })
      }
    },
    // 文件超出个数限制时的钩子
    handleExceed(files, fileList) {

    },
    // 文件列表移除文件时的钩子
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    // 点击文件列表中已上传的文件时的钩子
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    uploadFile() {
      this.$refs.upload.submit()
    },
    gotoFinished(){
      let req_data={
        "recipient_id":this.user_id,
        "pro_name":this.project_detail.pro_name,
        "introduction":this.project_detail.introduction,
        "category":this.project_detail.category,
        "materials":this.supplies_list,
        "rec_address":this.project_detail.rec_address,
        "cut_off_time":this.project_detail.cut_off_time,
        "emergency_degree":this.project_detail.emergency_degree,
      }
      console.log('project info',req_data)
      this.post_project_info_test(req_data)
    },
    post_project_info_test(req_data) {
      this.$message({
        type: 'success',
        message: '捐赠项目提交成功，请耐心等待审核通过~'
      });
      this.$router.push("/Maincontrol_g/" + this.user_id);
    },
    post_project_info(req_data) {
      //api请求方法
      // axios.post(`${this.$url}/test/testRequest`,data)
      axios.post(root_url + `/recipient/addDemandlist`, req_data)
        .then(res => {
          console.log('res=>', res);
          if (res.status === 200) {
            this.$message({
              type: 'success',
              message: '捐赠项目提交成功，请耐心等待审核通过~'
            });
            this.$router.push("/Maincontrol_g/" + this.user_id);
          } else {
            this.$message.error('提交失败，请重新提交~');
          }
        })
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
