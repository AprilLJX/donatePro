<template>
<div style="width=100%">

    <el-divider></el-divider>
    <span>项目详情</span>
    <el-divider></el-divider>

    <table>
    <tr class="dataRow" width="100%">
				<td rowspan="2">
					<img src="../../assets/hospital.jpg" width="350px" height="350px" style="margin-left: 20px;"/>
				</td>
        <td width="30%">
					<p>项目名称</p>
        </td>
        <td>
					<p>XXXXXXXXXX</p>
        </td>
    </tr>
    <tr class="dataRow">
				<td width="30%">
					<p style="margin-top: -105px;">项目简介</p>
				</td>
        <td>
					<p>XXXXXXXXXXXXXXXXX</p>
					<p>XXXXXXXXXXXXXXXXX</p>
					<p>XXXXXXXXXXXXXXXXX</p>
					<p>XXXXXXXXXXXXXXXXX</p>
					<p>XXXXXXXXXXXXXXXXX</p>
					<p>XXXXXXXXXXXXXXXXX</p>
				</td>
    </tr>
    </table>

    <table class="quickTable" width="100%">
			<tr class="dataRow">
				<td width="10%"><p><font size="-2"><b>编号</b></font></p></td>
				<td width="10%"><p><font size="-2"><b>物资名称</b></font></p></td>
				<td width="15%"><p><font size="-2"><b>物资参数</b></font></p></td>
				<td width="35%" style="text-align:left;"><p><font size="-2"><b>筹集进度</b></font></p></td>
				<td width="15%"><p><font size="-2"><b>物流状态</b></font></p></td>
				<td width="15%"><p><font size="-2"><b>操作</b></font></p></td>
			</tr>
      <tr>
        <td colspan="6"><el-divider></el-divider><td/>
      </tr>
			<tr class="dataRow">
				<td width="10%"><p><font size="-2">1</font></p></td>
				<td width="10%"><p><font size="-2">口罩</font></p></td>
        <td width="15%"><p><font size="-2">XXXXXXX</font></p></td>
				<td width="35%" style="text-align:left;">
          <p>
            <font size="-2">XXX爱心人士XX个</font>
            <font size="-2">        </font>
            <font size="-2">筹集完毕</font>
          </p>
        </td>
				<td width="15%"><p><font size="-2">XXX爱心人士 已寄出</font></p></td>
				<td width="15%">
        <a href="#" style="margin-top: -30px;text-decoration: none;color:green;"><font size="-2">
								查看详细物流</font></a>
        </td>
			</tr>
		</table>
</div>
</template>

<script>

export default {
  name: 'Checkproject',
  data() {
    return {
    }
}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
