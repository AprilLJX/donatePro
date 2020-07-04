<template>
<div style="width=100%">
    <div>
    <el-divider></el-divider>
    <span>爱心值</span>
    <el-divider></el-divider>


        <table class="quickTable" id="quickTable" height="300px" width="100%">
					<tr class="dataRow">
						<td style="height: 20px;width:33%;margin-left: 0px;font-size:12px;">
							<span>7日内爱心值变化</span>
						</td>
            <td style="height: 20px;width:33%;margin-left: 0px;font-size:12px;">
							<span>30日内爱心值变化</span>
						</td>
            <td style="height: 20px;width:33%;margin-left: 0px;font-size:12px;">
							<span>现有爱心总值</span>
						</td>
					</tr>
					<tr class="dataRow" style="width:100%;">
						<td style="height: 20px;width:33%;margin-left: 0px;font-size:12px;">
							<span style="font-size: 35px">100</span>
						</td>
						<td style="height: 20px;width:33%;margin-left: 222px;font-size:12px;">
							<span style="font-size: 35px">1000</span>
						</td>
						<td style="height: 20px;width:33%;margin-left: 444px;font-size: 12px;">
							<span style="font-size: 35px">1000</span>
						</td>
					</tr>
				</table>

    <el-divider></el-divider>
    <span>7日内爱心值变化</span>
    <el-divider></el-divider>
    <el-calendar :range="['2020-06-08', '2020-06-14']">
      <template
       slot="dateCell"
       slot-scope="{date, data}">
        <!--自定义内容-->
          <div>
             <div class="calendar-day">{{ data.day.split('-').slice(2).join('-') }}</div>
                <div v-for="item in calendarData">
                   <div v-if="(item.months).indexOf(data.day.split('-').slice(1)[0])!=-1">
                     <div v-if="(item.days).indexOf(data.day.split('-').slice(2).join('-'))!=-1">
                          <el-tooltip class="item" effect="dark" :content="item.things" placement="right">
                               <div class="is-selected">{{item.things}}</div>
                          </el-tooltip>
                       </div>
                    <div v-else></div>
                   </div>
               <div v-else></div>
             </div>
          </div>
      </template>
    </el-calendar>

    <el-divider></el-divider>
    <span>28日内爱心值变化</span>
    <el-divider></el-divider>
    <el-calendar :range="['2020-05-18', '2020-06-14']">
      <template
       slot="dateCell"
       slot-scope="{date, data}">
        <!--自定义内容-->
          <div>
             <div class="calendar-day">{{ data.day.split('-').slice(2).join('-') }}</div>
                <div v-for="item in calendarData">
                   <div v-if="(item.months).indexOf(data.day.split('-').slice(1)[0])!=-1">
                     <div v-if="(item.days).indexOf(data.day.split('-').slice(2).join('-'))!=-1">
                          <el-tooltip class="item" effect="dark" :content="item.things" placement="right">
                               <div class="is-selected">{{item.things}}</div>
                          </el-tooltip>
                       </div>
                    <div v-else></div>
                   </div>
               <div v-else></div>
             </div>
          </div>
      </template>
    </el-calendar>
  </div>
</div>
</template>

<script>

export default {
  name: 'Love',
  data() {
    return {
                calendarData: [
                    { months: ['06'],days: ['12','13'],things: '+100' },
                    { months: ['05'],days: ['19','24'],things: '+200' }
                ],
                value: new Date()
    }
}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
