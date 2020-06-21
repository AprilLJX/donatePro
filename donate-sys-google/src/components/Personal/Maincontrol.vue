<template>
<div>
    <el-header>
      <div id="top">
        <MainTop v-bind:if_logo="true" v-bind:user_type="'0'"></MainTop>
      </div>
	  <div id="top" style="margin-top:85px">
	  <el-divider></el-divider>
	  </div>
    </el-header>

	<el-container>
	<el-aside width="30%">
	<el-col :span="18":offset="2">
	<el-menu v-on:select="handleSelect" default-active="/love" router
                background-color="rgba(70, 76, 91, 1)" text-color="#ccc"  style="border-radius: 2px; box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);margin-bottom:60px">
          <el-menu-item index="/love">爱心值</el-menu-item>
          <el-menu-item index="/history">历史信息查询</el-menu-item>
		  <el-menu-item index="/personalcenter">个人信息</el-menu-item>
          <el-menu-item index="#">捐赠意向提交</el-menu-item>
		  <el-menu-item index="/message">消息列表</el-menu-item>
    </el-menu>
	</el-col>
	</el-aside>

	<el-main width="70%">
	<el-col :span="22" :offset="1" style="border-radius: 2px; box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);margin-top:-28px;margin-bottom:60px;min-height:700px">
	<router-view></router-view>
	</el-col>
	</el-main>
	</el-container>

	<el-footer>
	  <el-col :span="24"></el-col>
      <MainBottom></MainBottom>
    </el-footer>
</div>
</template>

<script>
import '../../assets/baseStyle.css'
import MainTop from "../MainTop";
import MainBottom from "../MainBottom";

export default {
  name: 'Maincontrol',
  components: {MainTop, MainBottom},
  data() {
	return{
	}
  },
  methods: {

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
	.findpassword{
				background-color: #F1F1F1;
	}
</style>
