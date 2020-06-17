<template>
	<div>
	<el-header>
      <div id="top">
        <MainTop v-bind:if_logo="true" v-bind:user_type="'0'"></MainTop>
      </div>
    </el-header>
	<div id="top" style="margin-top:-50px">
	  <el-divider></el-divider>
	</div>
    <el-main>
  	<div class= "login" style="width:100%;height:600px;margin-top:-25px">
			<el-col :span="8" :offset="15">
				<div style="border: 2px solid pink;width: 100%;height: 70%;background-color: white;border-radius: 5px;margin-top:180px;box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);">
					<el-form ref="form" :rules="rules" :model="form" label-width="80px" class="login-form">
						<h2 class="login-title">谷粒酬</h2>
						<el-form-item label="用户名" prop="username" style="width:200px; margin-left: 80px">
						<el-input v-model="form.username" style="width:160px"></el-input>
						</el-form-item>
						<el-form-item label="密码" prop="password" style="width:200px; margin-left: 80px">
						<el-input v-model="form.password" style="width:160px"></el-input>
						</el-form-item>
						<el-form-item style="margin-left: -60px">
						<el-button type="primary;" style="background-color: indianred; color: white; width: 120px" @click="submitForm('form')">登录</el-button>
						</el-form-item>
					</el-form>
					<table width="100%">
						<tr>
							<td style="text-align:right; width:33%">
								<span style="font-size: 12px;text-decoration: none; color: indianred" @click = "gotoRegister_g()">受赠方注册</span>
							</td>
							<td style="text-align:center; width:33%">
								<span style="font-size: 12px;text-decoration: none; color: indianred" @click = "gotoRegister_d()">捐赠方注册</span>
							</td>
							<td style="text-align:left; width:33%">
								<span style="font-size: 12px;text-decoration: none; color: indianred" @click = "gotoFindpassword()">忘记密码</span>
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

export default {
  name: 'Login',
  components: {MainTop, MainBottom},
  data () {
    return {
         form: {
           username: "",
           password: ""
         },
         rules: {
       }
       };
  },
  methods:{
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
	gotoPersonal() {
	  //直接跳转到个人中心
      this.$router.push('/Maincontrol');
	},
	submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            if(this.form.username !="root" || this.form.password!="123456"){
              this.$message.error('账号密码不正确');
              return false;
            }else{//真正项目中登录成功之后，就可以用路由跳转页面
               this.$router.push('/Maincontrol');
			   return true;
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
	.login{
				margin: 0;
				border: 0px;
				padding: 0;
				background: url("../../assets/login.png") no-repeat;
				background-size: cover;
				height: 100%;
				width: 100%;
			  }
</style>
