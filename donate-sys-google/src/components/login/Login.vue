<template>
  <div>
    <el-header>
      <div>
        <MainTop :header_info="header_info"></MainTop>
      </div>
    </el-header>
    <div id="top" style="margin-top:-50px">
      <el-divider></el-divider>
    </div>
    <el-main>
      <div class="login" style="width:100%;height:600px;margin-top:-25px">
        <el-col :span="8" :offset="15">
          <div
            style="border: 2px solid pink;width: 100%;height: 70%;background-color: white;border-radius: 5px;margin-top:180px;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);">
            <el-form ref="form" :rules="rules" :model="form" label-width="80px" class="login-form">
              <h2 class="login-title">谷粒酬</h2>
              <el-form-item label="用户名" prop="username" class="form_item_title">
                <el-input v-model="form.username" style="width:160px"></el-input>
              </el-form-item>
              <el-form-item label="密码" prop="password" class="form_item_title">
                <el-input v-model="form.password" show-password style="width:160px"></el-input>
              </el-form-item>
              <el-form-item label="登陆身份" class="form_item_title">
                <p class="demander_info" style="width: 160px;margin: 0;text-align: left">
                  <span v-if="form.user_type">受赠方&nbsp;</span>
                  <span v-else>捐赠方&nbsp;</span>
                  <el-switch
                    v-model="form.user_type"
                    active-color="#ffb6c1"
                    inactive-color="#C0CCDA">
                  </el-switch>
                </p>
              </el-form-item>
              <el-form-item style="margin-left: -60px">
                <el-button type="primary;" style="background-color: indianred; color: white; width: 120px"
                           @click="submitForm('form')">登录
                </el-button>
              </el-form-item>
            </el-form>
            <table width="100%">
              <tr>
                <td style="text-align:right; width:33%">
                  <span style="font-size: 12px;text-decoration: none; color: indianred"
                        @click="gotoRegister_g()">受赠方注册</span>
                </td>
                <td style="text-align:center; width:33%">
                  <span style="font-size: 12px;text-decoration: none; color: indianred"
                        @click="gotoRegister_d()">捐赠方注册</span>
                </td>
                <td style="text-align:left; width:33%">
                  <span style="font-size: 12px;text-decoration: none; color: indianred"
                        @click="gotoFindpassword()">忘记密码</span>
                </td>
              </tr>
            </table>
          </div>
        </el-col>
      </div>
    </el-main>

    <el-footer>
      <MainBottom></MainBottom>
    </el-footer>
  </div>

</template>

<script>
import '../../assets/baseStyle.css'
import MainTop from "../MainTop";
import MainBottom from "../MainBottom";

import axios from 'axios'

var root_url = 'http://localhost:9090'

export default {
  name: 'Login',
  components: {MainTop, MainBottom},
  data() {
    return {
      form: {
        username: "",
        password: "",
        user_type: false,//false为捐赠方登陆
      },
      rules: {},
      header_info: {
        height_line: -1,
        if_login: false,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi: false
      },
    };
  },
  methods: {
    gotoFindpassword() {
      //直接跳转到找回密码
      this.$router.push('/login/Findpassword');
    },
    gotoRegister_d() {
      //直接跳转到捐赠方注册
      this.$router.push('/login/Register_d');
    },
    gotoRegister_g() {
      //直接跳转到受赠方注册
      this.$router.push('/login/Register_g');
    },
    // gotoPersonal() {}
    submitForm1() {
      //api请求方法
      let data = {"account": this.form.username, "password": this.form.password};
      // axios.post(`${this.$url}/test/testRequest`,data)
      axios.get(root_url + `/donor/login`, {
        params: {
          account: this.form.username,
          password: this.form.password
        }
      })
        .then(res => {
          console.log('res=>', res);
          if (res.status === 200) {
            //登陆成功，直接跳转到个人中心
            var test_data = JSON.stringify(res.donor)
            console.log('json data',test_data)
            window.sessionStorage.setItem('if_login',true)
            window.sessionStorage.setItem('user_typr','0')
            window.sessionStorage.setItem('donor_info',test_data)
            this.$router.push("/Maincontrol/" + res.donor.donor_id);

            // this.$router.push({
            //   name: '个人中心',
            //   path: '/Maincontrol',
            //   params: {jum:request_data}
            // })
            // this.$router.push('/Maincontrol');
            // this.donor_info = res.data
          }
          if (res.status === '400') {
            this.$message.error('账号密码不正确');
          }
        })

    },
    submitForm2() {
      //api请求方法
      let data = {"account": this.form.username, "password": this.form.password};
      // axios.post(`${this.$url}/test/testRequest`,data)
      axios.get(root_url + `/donor/login`, {
        params: {
          account: this.form.username,
          password: this.form.password
        }
      })
        .then(res => {
          console.log('res=>', res);
          if (res.status === 200) {
            //登陆成功，直接跳转到个人中心
            window.sessionStorage.setItem('if_login',true)
            window.sessionStorage.setItem('user_typr','1')
            window.sessionStorage.setItem('recipient_info',res.recipientInfo)
            this.$router.push('/Maincontrol_g/'+res.recipientId);

            // this.$router.push('/Maincontrol_g');
            // this.donor_info = res.data
          }
          if (res.status === '400') {
            this.$message.error('账号密码不正确');
          }
        })

    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          if (this.form.username == "root" && this.form.password == "123456" && !this.form.user_type) {
            // this.$router.push('/Maincontrol');
            // var request_data={}
            var donor_info_response = {
              "donor": {
                "account": "123456",
                "avatar": "",
                "city": "北京",
                "cur_residence": "北京",
                "donor_id": 1,
                "id_number": "123456",
                "love_value": "",
                "nickname": "alice",
                "profile": ""
              },
              "msg": "登录成功",
              "status": 200
            }
            // 保存信息
            var test_data = JSON.stringify(donor_info_response.donor)
            console.log('json data',test_data)
            window.sessionStorage.setItem('if_login',true)
            window.sessionStorage.setItem('user_typr','0')
            window.sessionStorage.setItem('donor_info',test_data)

            // this.$set(request_data,'header_info',this.header_info)
            // this.$set(request_data,'donor_info',donor_info_response.donor)
            // this.$router.push({
            //   name: '个人中心',
            //   path: '/Maincontrol',
            //   params: {jum:request_data}
            // })

            this.$router.push("/Maincontrol/" + donor_info_response.donor.donor_id);
            return true;
          } else if (this.form.username == "root2" && this.form.password == "123456" && this.form.user_type) {
            var recipient_info_response = {
              "msg": "登录成功",
              "recipientId": 1,
              "recipientInfo": {
                "category": "医院",
                "com_address": "北京",
                "company": "医院a",
                "profile": "医院"
              },
              "status": 200
            }
            var test_data = JSON.stringify(recipient_info_response.recipientInfo)
            window.sessionStorage.setItem('if_login',true)
            window.sessionStorage.setItem('user_typr','1')
            window.sessionStorage.setItem('recipient_info',test_data)
            this.$router.push('/Maincontrol_g/'+recipient_info_response.recipientId);
            return true;
          } else {//真正项目中登录成功之后，就可以用路由跳转页面
            if(this.form.user_type){
              //api 受赠方请求方法
              // this.submitForm2()
            }else{
              //api 捐赠方请求方法
              // this.submitForm1()

            }
            // //直接跳转
            // window.console.log("查询成功", i);
            // // console.log(this.header_info)
            // var request_data={}
            // this.$set(request_data,'header_info',this.header_info)
            // this.$set(request_data,'donor_info',this.project_detail)
            // this.$router.push({
            //   name: '项目详情',
            //   path: '/projectList/projectDetail',
            //   params: {jum:request_data}
            // });

            this.$message.error('账号密码不正确');
            return false;
          }
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .login {
    margin: 0;
    border: 0px;
    padding: 0;
    background: url("../../assets/login.png") no-repeat;
    background-size: cover;
    height: 100%;
    width: 100%;
  }
  .form_item_title{
    width:200px;
    margin-left: 80px;
  }
</style>
