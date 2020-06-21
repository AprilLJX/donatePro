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
  	<div class= "findpassword" style="height: 500px">
	<div id="stage1" v-show="fir">
	<el-col :span="8" :offset="8">
	<div style="border: 1px solid pink;width: 100%;height: 100%;background-color: white;border-radius: 5px;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);margin-top:60px">
	<el-form ref="formp" :rules="rules" :model="formp" label-width="160px" class="login-form">
						<h2 class="login-title">谷粒酬 重置密码</h2>
			<div style="margin-left: -80px">
						<el-form-item label="手机号" prop="username" style="width:260px; margin-left: 80px">
						<el-input v-model="formp.username" style="width:160px"></el-input>
						</el-form-item>

						<el-row :gutter="10">
                		<el-col :span="17">
						<el-form-item label="获取验证码" prop="code" style="width:260px; margin-left: 80px">
							<el-input v-model="formp.code" placeholder="请输入验证码" style="width:160px"></el-input>
						</el-form-item>
						</el-col>
                		<el-col :span="7">
							<el-button type="button" @click="sendcode" :disabled="disabled" v-if="disabled==false">发送验证码
							</el-button>
							<el-button type="button" @click="sendcode" :disabled="disabled" v-if="disabled==true">{{btntxt}}
							</el-button>
						</el-col>
              			</el-row>

						<el-form-item label="新密码" prop="password" style="width:260px; margin-left: 80px">
						<el-input v-model="formp.password" style="width:160px"></el-input>
						</el-form-item>
						<el-form-item label="确认新密码" prop="cpassword" style="width:260px; margin-left: 80px">
						<el-input v-model="formp.cpassword" style="width:160px"></el-input>
						</el-form-item>
			</div>
						<el-form-item>
						<el-button type="primary" style="background-color: indianred; color: white; width: 120px; margin-left: -130px" @click="submitFormp('formp')">确认重置</el-button>
						</el-form-item>
	</el-form>
	</div>
	</el-col>
	</div>

	<div id="stage2" v-show="sec">
		<h1 style="font-size:80px;padding-top: 150px">重置成功！</h1>
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

export default {
  name: 'Findpassword',
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
	formp: {
           username: "",
		   code: "",
           password: "",
		   cpassword: "",
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
                if (this.formp.username == '') {
                    this.$message({
                        message:'手机号不能为空',
                        center: true
                    })
                    return
                }
                if (!reg.test(this.formp.username)) {
                    this.$message({
                        message:'请输入正确的手机号',
                        center:true
                    })
                    return
                } else {
                    console.log(this.formp.username);
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
	  submitFormp(formName) {
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
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	.findpassword{
				margin: 0;
				border: 0px;
				padding: 0;
			  }
</style>
