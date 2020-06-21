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
      <div class="register_d" style="height: 800px">
        <div id="stage1" v-show="fir">
          <el-steps :space="500" :active="0" finish-status="success" align-center style="margin-top:30px">
            <el-step title="进行中" description="实名制认证"></el-step>
            <el-step title="步骤 2" description="信息填写"></el-step>
            <el-step title="步骤 3" description="注册成功"></el-step>
          </el-steps>
          <el-col :span="8" :offset="8">
            <div
              style="border: 1px solid pink;width: 100%;height: 100%;background-color: white;border-radius: 5px;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);margin-top:60px">
              <el-form ref="formo" :rules="rules" :model="formRGO" label-width="160px" class="login-form">
                <h2 class="login-title">谷粒酬 受赠方注册</h2>
                <div style="margin-left: -80px">
                  <el-form-item label="手机号" prop="username" style="width:260px; margin-left: 80px">
                    <el-input v-model="formRGO.username" style="width:160px"></el-input>
                  </el-form-item>

                  <el-row :gutter="10">
                    <el-col :span="17">
                      <el-form-item label="获取验证码" prop="code" style="width:260px; margin-left: 80px">
                        <el-input v-model="formRGO.code" placeholder="请输入验证码" style="width:160px"></el-input>
                      </el-form-item>
                    </el-col>
                    <el-col :span="7">
                      <el-button type="button" @click="sendcode" :disabled="disabled" v-if="disabled==false">发送验证码
                      </el-button>
                      <el-button type="button" @click="sendcode" :disabled="disabled" v-if="disabled==true">{{btntxt}}
                      </el-button>
                    </el-col>
                  </el-row>

                  <el-form-item label="密码" prop="password" style="width:260px; margin-left: 80px">
                    <el-input v-model="formRGO.password" style="width:160px"></el-input>
                  </el-form-item>
                  <el-form-item label="确认密码" prop="cpassword" style="width:260px; margin-left: 80px">
                    <el-input v-model="formRGO.cpassword" style="width:160px"></el-input>
                  </el-form-item>
                  <el-form-item label="法定代表人" prop="name" style="width:260px; margin-left: 80px">
                    <el-input v-model="formRGO.name" style="width:160px"></el-input>
                  </el-form-item>
                  <el-form-item label="法定代表人身份证号" prop="id" style="width:260px; margin-left: 80px">
                    <el-input v-model="formRGO.id" style="width:160px"></el-input>
                  </el-form-item>
                </div>
                <el-form-item>
                  <el-button type="primary"
                             style="background-color: indianred; color: white; width: 120px; margin-left: -130px"
                             @click="submitFormRG('formo')">下一步
                  </el-button>
                </el-form-item>
              </el-form>
            </div>
          </el-col>
        </div>

        <div id="stage1" v-show="sec">
          <el-steps :space="500" :active="1" finish-status="success" align-center>
            <el-step title="已完成" description="实名制认证"></el-step>
            <el-step title="进行中" description="信息填写"></el-step>
            <el-step title="步骤 3" description="注册成功"></el-step>
          </el-steps>
          <el-col :span="8" :offset="8">
            <div
              style="border: 1px solid pink;width: 100%;height: 100%;background-color: white;border-radius: 5px;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);">
              <el-form ref="formb" :rules="rulerg" :model="formrg" label-width="120px" class="login-form">
                <h2 class="login-title">谷粒酬 受赠方注册</h2>
                <el-form-item label="省份城市" prop="procity" style="width:200px; margin-left: 80px">
                  <el-cascader
                    style="width:160px"
                    :options="options"
                    v-model="formrg.selectedOptions"
                    @change="handleChange">
                  </el-cascader>
                </el-form-item>
                <el-form-item label="具体市内地址" prop="address" style="width:200px; margin-left: 80px">
                  <el-input v-model="formrg.address" style="width:160px" type="textarea"></el-input>
                </el-form-item>
                <el-form-item label="单位名称" prop="hname" style="width:200px; margin-left: 80px">
                  <el-input v-model="formrg.hname" style="width:160px"></el-input>
                </el-form-item>
                <el-form-item label="单位类别" prop="utype" style="width:200px; margin-left: 80px">
                  <el-input v-model="formrg.utype" style="width:160px"></el-input>
                </el-form-item>
                <el-form-item label="统一信用代码" prop="tycode" style="width:200px; margin-left: 80px">
                  <el-input v-model="formrg.tycode" style="width:160px"></el-input>
                </el-form-item>
                <el-form-item label="单位简介" prop="info" style="width:200px; margin-left: 80px">
                  <el-input v-model="formrg.info" style="width:160px" type="textarea"></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary;" style="background-color: indianred; color: white; width: 120px"
                             @click="submitFormrg('formb')">注册
                  </el-button>
                </el-form-item>
              </el-form>
            </div>
          </el-col>
        </div>

        <div id="stage1" v-show="thir">
          <el-steps :space="500" :active="3" finish-status="success" align-center>
            <el-step title="已完成" description="实名制认证"></el-step>
            <el-step title="已完成" description="信息填写"></el-step>
            <el-step title="注册成功"></el-step>
          </el-steps>
          <h1 style="font-size:80px;margin-top: 200px">注册成功！</h1>
        </div>

      </div>


    </el-main>

    <el-footer>
      <MainBottom></MainBottom>
    </el-footer>
  </div>
</template>

<script>
import '../../assets/baseStyle.css';
import MainTop from "../MainTop";
import MainBottom from "../MainBottom";
import {provinceAndCityData, CodeToText, TextToCode} from 'element-china-area-data';

export default {
  name: 'Register_g',
  components: {MainTop, MainBottom},
  data() {
    return {
      header_info:{
        height_line:-1,
        if_logo: false,
        user_type: '0', // 0 is donator, 1 is reciver
        if_show_navi:false
      },
      fir: true,
      sec: false,
      thir: false,
      options: provinceAndCityData,
      formrg: {
        selectedOptions: [],
        address: "",
        hname: "",
        utype: "",
        tycode: "",
        info: "",
      },
      rulerg: {
        selectedOptions: [{
          type: 'Array',
          required: true,
          message: '请选择',
          trigger: 'blur'
        }],
        hname: [
          {required: true, message: "真实姓名不能为空", trigger: 'blur'},
        ],
        address: [
          {required: true, message: "地址不能为空", trigger: 'blur'},
        ],
        utype: [
          {required: true, message: "单位类别不能为空", trigger: 'blur'},
        ],
        info: [
          {required: true, message: "单位简介不能为空", trigger: 'blur'},
        ],
        tycode: [
          {required: true, message: "统一信用代码不能为空", trigger: 'blur'},
        ]
      },
      formRGO: {
        username: "",
        code: "",
        password: "",
        cpassword: "",
        name: "",
        id: "",

      },
      rules: {
        username: [
          {required: true, message: "手机号不能为空", trigger: 'blur'},
          {min: 11, max: 11, message: "手机号11位", trigger: 'blur'},
        ],
        code: [
          {required: true, message: "验证码不能为空", trigger: 'blur'},
        ],
        password: [
          {required: true, message: "密码不能为空", trigger: 'blur'},
          {min: 6, max: 10, message: "密码6-10位", trigger: 'blur'},
        ],
        cpassword: [
          {required: true, message: "密码不能为空", trigger: 'blur'},
          {min: 6, max: 10, message: "密码6-10位", trigger: 'blur'},
        ],
        name: [
          {required: true, message: "法定代表人不能为空", trigger: 'blur'},
        ],
        id: [
          {required: true, message: "法定代表人身份证号不能为空", trigger: 'blur'},
          {min: 18, max: 18, message: "身份证号18位", trigger: 'blur'},
        ],

      },
      disabled: false,
      time: 0,
      btntxt: "重新发送",
    }
  },
  methods: {
    //手机验证发送验证码
    sendcode() {
      const reg = 11 && /^((13|14|15|17|18)[0-9]{1}\d{8})$/
      if (this.formRGO.username == '') {
        this.$message({
          message: '手机号不能为空',
          center: true
        })
        return
      }
      if (!reg.test(this.formRGO.username)) {
        this.$message({
          message: '请输入正确的手机号',
          center: true
        })
        return
      } else {
        console.log(this.formRGO.username);
        this.$message({
          message: '发送成功',
          type: 'success',
          center: true
        });
        this.time = 60;
        this.disabled = true;
        this.timer();
      }
    },
    //60S倒计时
    timer() {
      if (this.time > 0) {
        this.time--;
        this.btntxt = this.time + "s后重新获取";
        setTimeout(this.timer, 1000);
      } else {
        this.time = 0;
        this.btntxt = "获取验证码";
        this.disabled = false;
      }
    },
    submitFormRG(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {//真正项目中登录成功之后，就可以用路由跳转页面
          this.fir = false;
          this.sec = true;
          return true;
        }
      });
    },
    submitFormrg(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {//真正项目中登录成功之后，就可以用路由跳转页面
          this.sec = false;
          this.thir = true;
          return true;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    handleChange(value) {
      console.log(value);
    }

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .register_g {
    margin: 0;
    border: 0px;
    padding: 0;
  }
</style>
