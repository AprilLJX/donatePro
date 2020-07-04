<template>
<div style="width=70%;">
    <el-divider></el-divider>
    <span>消息</span>
    <el-divider></el-divider>
    <table class="quickTable" width="100%">
			<tr class="dataRow">
				<td width="20%"><p><font size="-2"><b>消息来源</b></font></p></td>
				<td width="60%" style="text-align:left;"><p><font size="-2"><b>消息内容</b></font></p></td>
				<td width="10%"><p><font size="-2"><b>时间</b></font></p></td>
				<td width="10%"><p><font size="-2"><b>操作</b></font></p></td>
			</tr>
      <tr>
        <td colspan="7"><el-divider></el-divider><td/>
      </tr>
			<tr class="dataRow">
				<td width="10%"><p><font size="-2">系统消息</font></p></td>
				<td width="20%" style="text-align:left;"><p><font size="-2">XXXXXXXXX项目捐赠意向单已审核通过！请尽快补充物流信息吧！</font></p></td>
				<td width="10%"><p><font size="-2">2020-06-10</font></p></td>
				<td width="10%">
        <a href="./checkproject.html" style="text-decoration: none;color:red;"><font size="-2">
								删除消息</font></a></td>
			</tr>
		</table>
</div>
</template>

<script>

export default {
  name: 'Message',
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

</style>
