<template>
  <div class= "findpassword">
	<div id="stage1" v-show="fir">
	<div style="position: absolute">
		<img src="./fp1.png" style="width: 1000px;height: 45px;margin-left: 220px;margin-top:80px"/>
	</div>
	
	<div style="position:absolute;top:200px;left:260px;width: 65%;height: 40%;background-color: white;border: 2px solid pink;border-radius: 5px">
	<table width="60%" align="center">
	<tr>
		<td colspan="2"><font size="6" style="margin-left: 100px">谷粒酬</font>找回密码</td>
	</tr>
	<tr>
		<td>手机号:</td>
		<td><input type="text" placeholder="11位手机号" class="in"/></td>
	</tr>
	<tr>
		<td>验证码:</td>
		<td><input type="text" placeholder="输入验证码" class="in"/></td> 
		<td><input type="button" value="获取验证码" style="height: 30px;background-color: coral;color: white;border: none;border-radius: 3px"/></td>
	</tr>
	<tr>
		<td></td>
		<td><input type="button" value="下一步" style="background-color: indianred;color:white;height:30px;width: 306px;border: none;border-radius: 3px;" @click="jump()"/></td>
	</tr>
	</table>
	</div>
	</div>
	
	<div id="stage2" v-show="sec">
	<div style="position: absolute">
		<img src="./fp2.png" style="width: 1000px;height: 45px;margin-left: 220px;margin-top:80px"/>
	</div>
	
	<div style="position:absolute;top:200px;left:260px;width: 65%;height: 40%;background-color: white;border: 2px solid pink;border-radius: 5px">
	<table width="60%" align="center">
	<tr>
		<td colspan="2"><font size="6" style="margin-left: 100px">谷粒酬</font>找回密码</td>
	</tr>
	<tr>
		<td>新密码:</td>
		<td><input type="password" placeholder="6-16位密码，区分大小写" class="in"/></td>
	</tr>
	<tr>
		<td>确认新密码:</td>
		<td><input type="password" placeholder="确认新密码" class="in"/></td>
	</tr>
	<tr>
		<td></td>
		<td><input type="button" value="下一步" style="background-color: indianred;color:white;height:30px;width: 306px;border: none;border-radius: 3px;" @click="jump2()"/></td>
	</tr>
	</table>
	</div>
	</div>
	
	<div id="stage3" v-show="thir">
	<div style="position: absolute">
		<img src="./fp3.png" style="width: 1000px;height: 45px;margin-left: 220px;margin-top:80px"/>
	</div>
	
	<div style="position:absolute;top:200px;left:260px;width: 65%;height: 40%;">
	<table width="60%" align="center">
		<tr>
			<td><span  style="margin-left: 150px;font-size: 50px">重置成功!</span></td>
		</tr>
	</table>
	</div>
	</div>		
	</div>
</template>

<script>
export default {
  name: 'Findpassword',
  data() {
	return{
	fir:true,
  	sec:false,
  	thir:false,
	}
  },
  methods: {
	  jump(){
		this.fir = false;
		this.sec = true;
	  },
	  jump2(){
		  this.sec = false;
		  this.thir = true;

	  }

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	.findpassword{
				margin: 0;
				border: 0px;
				padding: 0;
				background: url("./bg.png") no-repeat;
				background-size: cover;
				height: 100%;
				width: 100%;
				position: fixed;
			  }
</style>
