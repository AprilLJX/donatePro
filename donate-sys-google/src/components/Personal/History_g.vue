<template>
<div style="width:100%">
<!--    <el-divider></el-divider>-->
<!--    <span>搜索条件</span>-->
<!--    <el-divider></el-divider>-->
<!--    <el-row :gutter="30">-->
<!--    <el-col :span="4"><el-input v-model="input" placeholder="请输入项目名称"></el-input></el-col>-->
<!--    <el-col :span="4"><el-input v-model="input" placeholder="请输入物资类型"></el-input></el-col>-->
<!--    <el-col :span="3">审核状态：</el-col>-->
<!--    <el-col :span="4">-->
<!--      <el-select v-model="value" placeholder="请选择">-->
<!--      <el-option-->
<!--        v-for="item in options"-->
<!--        :key="item.value"-->
<!--        :label="item.label"-->
<!--        :value="item.value">-->
<!--      </el-option>-->
<!--      </el-select>-->
<!--    </el-col>-->
<!--    <el-col :span="3">物流状态：</el-col>-->
<!--    <el-col :span="4">-->
<!--      <el-select v-model="value" placeholder="请选择">-->
<!--      <el-option-->
<!--        v-for="item in options"-->
<!--        :key="item.value"-->
<!--        :label="item.label"-->
<!--        :value="item.value">-->
<!--      </el-option>-->
<!--      </el-select>-->
<!--    </el-col>-->
<!--    </el-row>-->

    <el-divider></el-divider>
    <span>历史捐赠信息</span>
    <el-divider></el-divider>

    <table class="quickTable" width="100%">
			<tr class="dataRow">
        <td width="10%"><p><b>编号</b></p></td>
        <td width="20%"><p><b>项目名称</b></p></td>
        <td width="20%"><p><b>项目缩略图</b></p></td>
        <td width="30%"><p><b>物资详情</b></p></td>
<!--        <td width="10%"><p><b>审核状态</b></p></td>-->
<!--        <td width="10%"><p><b>物流状态</b></p></td>-->
        <td width="10%"><p><b>操作</b></p></td>
			</tr>
      <tr>
        <td colspan="7"><el-divider></el-divider><td/>
      </tr>
      <tr class="dataRow" v-for="(i,index) in proList" :key="index">
        <td width="10%"><p>{{i.demandId}}</p></td>
        <td width="20%"><p>{{i.proName}}</p></td>
        <td width="20%"><img src="../../assets/hospital.jpg" width="100px" height="100px"/></td>
        <td width="30%" style="text-align: left;padding-left: 5%;">
          <p v-for="n in i.materials.split('；')">{{n}}</p>
        </td>
<!--        <td width="10%">-->
<!--          <p v-if="i.ifAudit==='1'">审核通过</p>-->
<!--          <p v-else-if="i.ifAudit==='0'">待审核</p>-->
<!--          <p v-else-if="i.ifAudit==='2'">审核未通过</p>-->
<!--        </td>-->
<!--        <td width="10%"><p>已发出</p></td>-->
        <td width="10%">
        <p style="text-decoration: none;color:green;" @click="gotoDonateDetail()">查看项目进度</p></td>
			</tr>
		</table>
</div>

</template>

<script>
import axios from 'axios'
var root_url = 'http://localhost:9090'
export default {
  name: 'History_g',
  data() {
	return{
	}
  },
  created() {
    this.user_id = this.$route.params.user_id
    console.log('loading histroy center', this.user_id)
    this.get_history_info_test()

  },
  methods: {
    get_history_info_test() {
      var res ={
        "msg": "查询成功",
        "proList": [
          {
            "RecAddress": "广州",
            "demandId": "2",
            "intro": "养老院a现需要电热毯，热水袋若干，为老人的冬天带来温暖",
            "materials": "电热毯：50；热水袋：50；保暖衣：100",
            "proName": "为老人献爱心"
          }
        ],
        "status": 200
      }
      console.log('login status,',window.sessionStorage.getItem('if_login'))
      if (res.status === 200 && window.sessionStorage.getItem('if_login')) {
        this.proList = res.proList
      } else {
        this.$message.error('获取信息失败，请重新登录~');
        this.$router.push('/login');
      }
    },
    get_history_info() {
      //api请求方法
      // let data = {"recipient_id": this.user_id};
      // axios.post(`${this.$url}/test/testRequest`,data)
      axios.get(root_url + `/recipient/personalCenter`, {
        params: {
          recipient_id: this.user_id,
        }
      })
        .then(res => {
          console.log('res=>', res);
          if (res.status === 200 && window.sessionStorage.getItem('if_login')) {
            //登陆成功，直接跳转到个人中心
            this.proList = res.proList
          } else {
            this.$message.error('获取信息失败，请重新登录~');
            this.$router.push('/login');
          }
        })
    },
    gotoDonateDetail() {
      //跳转到捐赠单详情
      this.$router.push({
        name: '定向捐赠单填写完成',
        // name: 'mallList',
        params: {jum:request_data}
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
