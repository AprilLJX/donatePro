<template>
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="信息查看" name="first">
        <table class="quickTable" id="rtable" width="500px" height="600px">
<!--          <tr class="dataRow">-->
<!--            <td width="250px" style="margin-top: 50px">用户账号：</td>-->
<!--            <td width="250px" style="margin-top: 50px; margin-left: 200px">{{recipient_info.Account}}</td>-->
<!--          </tr>-->
<!--          <tr class="dataRow">-->
<!--            <td width="500px" style="margin-top: 100px">姓名：</td>-->
<!--            <td width="500px" style="margin-top: 100px; margin-left: 200px">{{recipient_info.Nickname}}</td>-->
<!--          </tr>-->
<!--          <tr class="dataRow">-->
<!--            <td width="500px" style="margin-top: 150px">身份证号：</td>-->
<!--            <td width="500px" style="margin-top: 150px; margin-left: 200px">{{recipient_info.IdNumber}}</td>-->
<!--          </tr>-->
          <tr class="dataRow">
            <td width="500px" style="margin-top: 200px">类别：</td>
            <td width="500px" style="margin-top: 200px; margin-left: 200px">{{recipient_info.category}}</td>
          </tr>
          <tr class="dataRow">
            <td width="500px" style="margin-top: 200px">机构名称：</td>
            <td width="500px" style="margin-top: 200px; margin-left: 200px">{{recipient_info.company}}</td>
          </tr>
          <tr class="dataRow">
            <td width="500px" style="margin-top: 200px">地址：</td>
            <td width="500px" style="margin-top: 200px; margin-left: 200px">{{recipient_info.com_address}}</td>
          </tr>
<!--          <tr class="dataRow">-->
<!--            <td width="500px" style="margin-top: 250px">详细地址：</td>-->
<!--            <td width="500px" style="margin-top: 250px; margin-left: 200px">-->
<!--              <p>{{donor_info.com_address}}</p>-->
<!--            </td>-->
<!--          </tr>-->
          <tr class="dataRow">
            <td width="500px" style="margin-top: 250px">机构简介：</td>
            <td width="500px" style="margin-top: 250px; margin-left: 200px">
              <p>{{recipient_info.profile}}</p>
            </td>
          </tr>
        </table>
      </el-tab-pane>
<!--    <el-tab-pane label="信息修改" name="second">-->
<!--        <el-form ref="form" :model="form" label-width="100px">-->
<!--          <el-form-item label="用户账号：">-->
<!--            <el-col :span="6"><el-input v-model="form.id"></el-input></el-col>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="法定代表人：">-->
<!--            <el-col :span="6"><el-input v-model="form.name"></el-input></el-col>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="身份证号：">-->
<!--            <el-col :span="6"><el-input v-model="form.ID"></el-input></el-col>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="所属区域：">-->
<!--            <el-col :span="6"><el-cascader-->
<!--							style="width:240px"-->
<!--							:options="options"-->
<!--							v-model="form.selectedOptions"-->
<!--							@change="handleChange2">-->
<!--					  </el-cascader></el-col>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="详细地址">-->
<!--            <el-col :span="6"><el-input type="textarea" v-model="form.address"></el-input></el-col>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="单位名称：">-->
<!--            <el-col :span="6"><el-input v-model="form.id"></el-input></el-col>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="单位类别：">-->
<!--            <el-col :span="6"><el-input v-model="form.id"></el-input></el-col>-->
<!--          </el-form-item>-->
<!--          <el-form-item label="统一信用代码：">-->
<!--            <el-col :span="6"><el-input v-model="form.id"></el-input></el-col>-->
<!--          </el-form-item>-->
<!--          <el-form-item>-->
<!--            <el-button type="primary" @click="onSubmit">确认修改</el-button>-->
<!--            <el-button>取消</el-button>-->
<!--          </el-form-item>-->
<!--        </el-form>-->
<!--    </el-tab-pane>-->
  </el-tabs>

</template>

<script>
import { provinceAndCityData,CodeToText,TextToCode } from 'element-china-area-data';
import axios from 'axios'
var root_url = 'http://localhost:9090'
export default {
  name: 'Personalcenter_g',
  data() {
    return {
        options: provinceAndCityData,
        form: {
          id: '',
          name: '',
          ID: '',
          province: '',
          city: '',
          address: '',
        },
        activeName: 'first',
      recipient_info:{}
      };
  },
  created() {
    this.user_id = this.$route.params.user_id
    console.log('recipient_id!',this.user_id)
    this.get_personal_info()
    // console.log(this.recipient_info)
  },
  methods: {
    get_personal_info() {
      console.log('creat function')
      console.log('if_login',window.sessionStorage.getItem('if_login'))
      console.log('recipient_info',window.sessionStorage.getItem('recipient_info'))

      if(window.sessionStorage.getItem('if_login')){
        this.recipient_info = JSON.parse(window.sessionStorage.getItem('recipient_info'))
        console.log('json data',this.recipient_info)
      }else{
        this.$message.error('获取信息失败，请重新登录~');
        this.$router.push('/login');
      }
    },
    handleClick(tab, event) {
        console.log(tab, event);
      },
    onSubmit() {
        console.log('submit!');
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
