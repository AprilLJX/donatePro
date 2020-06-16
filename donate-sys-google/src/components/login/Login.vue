<template>
  <div class= "login">
				
				<div style="position:absolute;top:200px;left:850px;border: 2px solid pink;width: 28%;height: 35%;background-color: white;border-radius: 5px">
					<table width="60%" align="center">
						<tr>
							<td colspan="2"><font size="4" style="position:absolute;top:30px;margin-left: 55px">谷粒酬登录</font></td>
							
						</tr>
						<tr>
							<td style="position:absolute;top:70px">用户名:</td>
							<td style="position:absolute;top:70px">
								<input type="text" placeholder="请输入手机号" style="border: 1px solid pink;height: 30px;width: 220px;border-radius: 3px"/></td>
						</tr>
						<tr>
							<td style="position:absolute;top:110px">密码:</td>
							<td style="position:absolute;top:110px">
								<input type="password" placeholder="请输入密码"  style="border: 1px solid pink;height: 30px;width: 220px;border-radius: 3px"/></td>
						</tr>
						<tr>
							<td style="position:absolute;top:155px">
								<input type="submit" value="登录" style="background-color: indianred;color: white;border: none;width: 224px;height: 30px;border-radius: 3px" @click = "gotoPersonal()"/></td>
						</tr>
						<tr>
							<td style="position:absolute;top:195px">
								<span style="font-size: 12px;text-decoration: none; color: indianred" @click = "gotoRegister_g()">受赠方注册</span>
							</td>
							<td style="position:absolute;top:195px;left:150px">
								<span style="font-size: 12px;text-decoration: none; color: indianred" @click = "gotoRegister_d()">捐赠方注册</span>
							</td>
							<td style="position:absolute;top:195px;left: 250px">
								<span style="font-size: 12px;text-decoration: none; color: indianred" @click = "gotoFindpassword()">忘记密码</span>
							</td>
						</tr>
					</table>
				</div>
				
			</div>
</template>

<script>
export default {
  name: 'Login',
  data () {
    return {
      msg: 'Login Interface',
    }
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
      this.$router.push('/personalcenter');
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
				background: url("./login.png") no-repeat;
				background-size: cover;
				height: 100%;
				width: 100%;
				position: fixed;
			  }
</style>
