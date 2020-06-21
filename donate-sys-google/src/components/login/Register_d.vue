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
  <div class= "register_d" style="height: 600px">
	<div id="stage1" v-show="fir">
	<el-col :span="8" :offset="8">
	<div style="border: 1px solid pink;width: 100%;height: 100%;background-color: white;border-radius: 5px;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);">
	<el-form ref="formRD" :rules="rules" :model="formRD" label-width="100px" class="login-form">
						<h2 class="login-title">谷粒酬 捐赠方注册</h2>
			<div style="margin-left: -50px">
						<el-form-item label="手机号" prop="username" style="width:200px; margin-left: 80px">
						<el-input v-model="formRD.username" style="width:160px"></el-input>
						</el-form-item>

						<el-row :gutter="10">
                		<el-col :span="17">
						<el-form-item label="获取验证码" prop="code" style="width:260px; margin-left: 80px">
							<el-input v-model="formRD.code" placeholder="请输入验证码" style="width:160px"></el-input>
						</el-form-item>
						</el-col>
                		<el-col :span="7">
							<el-button type="button" @click="sendcode" :disabled="disabled" v-if="disabled==false">发送验证码
							</el-button>
							<el-button type="button" @click="sendcode" :disabled="disabled" v-if="disabled==true">{{btntxt}}
							</el-button>
						</el-col>
              			</el-row>

						<el-form-item label="密码" prop="password" style="width:200px; margin-left: 80px">
						<el-input v-model="formRD.password" style="width:160px"></el-input>
						</el-form-item>
						<el-form-item label="确认密码" prop="cpassword" style="width:200px; margin-left: 80px">
						<el-input v-model="formRD.cpassword" style="width:160px"></el-input>
						</el-form-item>
						<el-form-item label="真实姓名" prop="name" style="width:200px; margin-left: 80px">
						<el-input v-model="formRD.name" style="width:160px"></el-input>
						</el-form-item>
						<el-form-item label="身份证号" prop="id" style="width:200px; margin-left: 80px">
						<el-input v-model="formRD.id" style="width:160px"></el-input>
						</el-form-item>
						<el-form-item label="省份城市" prop="procity" style="width:200px; margin-left:80px">
							<el-cascader
							style="width:160px"
							:options="options"
							v-model="formRD.selectedOptions"
							@change="handleChange">
							</el-cascader>
						</el-form-item>
			</div>
						<el-form-item style="margin-left: -60px">
						<el-button type="primary;" style="background-color: indianred; color: white; width: 120px" @click="submitFormRD('formRD')">注册</el-button>
						</el-form-item>
	</el-form>
	</div>
	</el-col>
	</div>

	<div id="stage1" v-show="sec">
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
import { provinceAndCityData,CodeToText,TextToCode } from 'element-china-area-data';

export default {
  name: 'Register_d',
  components: {MainTop, MainBottom},
  data() {
	return{
    header_info:{
      height_line:-1,
      if_logo: false,
      user_type: '0', // 0 is donator, 1 is reciver
      if_show_navi:false
    },
	fir: true,
	sec: false,
	options: provinceAndCityData,
	formRD: {
           username: "",
		   code: "",
           password: "",
		   cpassword: "",
		   name: "",
		   id: "",
		   selectedOptions: [],
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
             {required: true, message: "真实姓名不能为空", trigger: 'blur'},
           ],
		   id: [
			 {required: true, message: "身份证号不能为空", trigger: 'blur'},
			 {min: 18, max: 18, message: "身份证号18位", trigger: 'blur'},
		   ]
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
                if (this.formRD.username == '') {
                    this.$message({
                        message:'手机号不能为空',
                        center: true
                    })
                    return
                }
                if (!reg.test(this.formRD.username)) {
                    this.$message({
                        message:'请输入正确的手机号',
                        center:true
                    })
                    return
                } else {
                    console.log(this.formRD.username);
                    this.$message({
                        message: '发送成功',
                        type: 'success',
                        center:true
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
	  submitFormRD(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {//真正项目中登录成功之后，就可以用路由跳转页面
               this.fir = false;
			   this.sec = true;
			   return true;
		  }
        });
      },
	  resetForm(formName) {
		  this.$refs[formName].resetFields();
	  },
	  handleChange (value) {
        console.log(value);
      }

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	.register_d{
				margin: 0;
				border: 0px;
				padding: 0;
			  }
</style>
