<template>
  <el-row>
    <el-col :span="12" :offset="3">
      <div style="background: white;height: 270px">
        <el-row>
          <el-col :span="12">
            <div style="width: 90%;margin: 5%;text-align:center;vertical-align:middle;">
              <img src='https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg' class="image" style="width: 100%;height:60%">
            </div>
          </el-col>
          <el-col :span="12">
            <div style="margin: 5%;">
              <div class="need_staff" style="text-align: left">
                <p class="project_name">{{project_detail.name}}</p>
                <p class="project_info">
                      <span>
                        <i class="el-icon-location" style="color: red"></i>
                        {{project_detail.place}}
                      </span>
                </p>
                <p class="project_info">
                  <span>{{'发起方： '+project_detail.demander.name}}</span>
                </p>
                <p class="project_info">
                      <span>
                        紧急度：
                        <i v-for="i in project_detail.emergency" :key="i" class="el-icon-star-on" style="color: red"></i>
                      </span>
                </p>
                <p class="project_info">
                  <span>{{"参与度： "+project_detail.receive_times +" 人次"}}</span>
                </p>
                <li v-for="(item,idx) in needy_list" style="color: crimson;line-height: 20px">
                  <span>{{item[0]+': '+ item[1]+' '+ item[2]}}</span>
                </li>
              </div>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-col>
    <el-col :span="6">
      <div style="background: white;height: 270px"></div>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "ProjectCard",
  props: {
    project_detail: {
      type: Object,//type为Array,default为函数
      default() {
      }
    },//用props属性进行传数据，此时子组件已经获取到list的数据了

  }
}
</script>

<style scoped>

</style>
