<template>
    <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="信息查看" name="first">
    <table class="quickTable" id="rtable" width="500px" height="600px">
					<tr class="dataRow">
					  <td width="250px" style="margin-top: 50px">用户账号：</td>
					  <td width="250px" style="margin-top: 50px; margin-left: 200px">XXX-XXXX-XXXX(手机号）</td>
					</tr>
					<tr class="dataRow">
					  <td width="500px" style="margin-top: 100px">姓名：</td>
					  <td width="500px" style="margin-top: 100px; margin-left: 200px">XXX</td>
					</tr>
					<tr class="dataRow">
					  <td width="500px" style="margin-top: 150px">身份证号：</td>
					  <td width="500px" style="margin-top: 150px; margin-left: 200px">（18位）</td>
					</tr>
					<tr class="dataRow">
					  <td width="500px" style="margin-top: 200px">地址：</td>
					  <td width="500px" style="margin-top: 200px; margin-left: 200px">XXX省   XX市</td>
					</tr>
					<tr class="dataRow">
					  <td width="500px" style="margin-top: 250px">详细地址：</td>
					  <td width="500px" style="margin-top: 250px; margin-left: 200px">
						  <p >xx区</p>
						  <p style="margin-top: 30px">xx路</p>
						  <p style="margin-top: 60px">xx小区</p>
						  <p style="margin-top: 90px">xx栋xx号</p>
					  </td>
					</tr>
				</table>
    </el-tab-pane>
    <el-tab-pane label="信息修改" name="second">
        <el-form ref="form" :model="form" label-width="100px">
          <el-form-item label="用户账号：">
            <el-col :span="6"><el-input v-model="form.id"></el-input></el-col>
          </el-form-item>
          <el-form-item label="姓名：">
            <el-col :span="6"><el-input v-model="form.name"></el-input></el-col>
          </el-form-item>
          <el-form-item label="身份证号：">
            <el-col :span="6"><el-input v-model="form.ID"></el-input></el-col>
          </el-form-item>
          <el-form-item label="地址：">
            <el-col :span="6"><el-select v-model="form.province" placeholder="请选择省份">
              <el-option label="区域一" value="shanghai"></el-option>
              <el-option label="区域二" value="beijing"></el-option>
            </el-select></el-col>
            <el-col :span="6":offset="3"><el-select v-model="form.city" placeholder="请选择城市">
              <el-option label="区域一" value="shanghai"></el-option>
              <el-option label="区域二" value="beijing"></el-option>
            </el-select></el-col>
          </el-form-item>
          <el-form-item label="详细地址">
            <el-col :span="6"><el-input type="textarea" v-model="form.address"></el-input></el-col>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">立即创建</el-button>
            <el-button>取消</el-button>
          </el-form-item>
        </el-form>
    </el-tab-pane>
  </el-tabs>
  
</template>

<script>

export default {
  name: 'Personalcenter',
  data() {
    return {
        form: {
          id: '',
          name: '',
          ID: '',
          province: '',
          city: '',
          address: '',
        },
        activeName: 'first'
      };
  },
  methods: {
    handleClick(tab, event) {
        console.log(tab, event);
      },
    onSubmit() {
        console.log('submit!');
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
