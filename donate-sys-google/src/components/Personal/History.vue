<template>
<div style="width=100%">
    <el-divider></el-divider>
    <span>搜索条件</span>
    <el-divider></el-divider>
    <el-row :gutter="30">
    <el-col :span="4"><el-input v-model="input" placeholder="请输入项目名称"></el-input></el-col>
    <el-col :span="4"><el-input v-model="input" placeholder="请输入物资类型"></el-input></el-col>
    <el-col :span="3">审核状态：</el-col>
    <el-col :span="4">
      <el-select v-model="value" placeholder="请选择">
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value">
      </el-option>
      </el-select>
    </el-col>
    <el-col :span="3">物流状态：</el-col>
    <el-col :span="4">
      <el-select v-model="value" placeholder="请选择">
      <el-option
        v-for="item in options"
        :key="item.value"
        :label="item.label"
        :value="item.value">
      </el-option>
      </el-select>
    </el-col>   
    </el-row>

    <el-divider></el-divider>
    <span>历史捐赠信息</span>
    <el-divider></el-divider>

    <table class="quickTable" width="100%">
			<tr class="dataRow">
				<td width="10%"><p><font size="-2"><b>编号</b></font></p></td>
				<td width="20%"><p><font size="-2"><b>项目名称</b></font></p></td>
				<td width="20%"><p><font size="-2"><b>项目缩略图</b></font></p></td>
				<td width="20%"><p><font size="-2"><b>物资类型</b></font></p></td>
				<td width="10%"><p><font size="-2"><b>审核状态</b></font></p></td>
				<td width="10%"><p><font size="-2"><b>物流状态</b></font></p></td>
				<td width="10%"><p><font size="-2"><b>操作</b></font></p></td>
			</tr>
      <tr>
        <td colspan="7"><el-divider></el-divider><td/>
      </tr>
			<tr class="dataRow">
				<td width="10%"><p><font size="-2">1</font></p></td>
				<td width="20%"><p><font size="-2">XXXXXX</font></p></td>
				<td width="20%"><img src="../../assets/hospital.jpg" width="100px" height="100px"/></td>
				<td width="20%"><p><font size="-2">口罩 XX个</font></p></td>
				<td width="10%"><p><font size="-2">审核通过</font></p></td>
				<td width="10%"><p><font size="-2">已发出</font></p></td>
				<td width="10%">
        <a href="#" style="margin-top: -30px;text-decoration: none;color:green;"><font size="-2">
								填写物流编号</font></a>
        <p></p>
        <p style="text-decoration: none;color:green;" @click="gotoFindpassword()"><font size="-2">
								查看项目进度</font></p></td>
			</tr>
		</table>
</div>

</template>

<script>

export default {
  name: 'History',
  data() {
	return{
	}
  },
  methods: {
    gotoFindpassword() {
      //直接跳转到找回密码
      this.$router.push('/checkproject');
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
