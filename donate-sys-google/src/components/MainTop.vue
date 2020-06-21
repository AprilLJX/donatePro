<template id="top_elem">
  <div data-type="top_title">
    <div data-type="top_logo">
        <span class="icon svg-icon-wrap svg-antFill svg-antFill-antFill-heart-wrap"
              style="width: 64px; height: 64px; min-width: 64px; min-height: 64px;">
            <svg xmlns="http://www.w3.org/2000/svg" class="svg-icon" viewBox="64 64 896 896" width="64" height="64"
                 style="fill: rgb(238, 11, 11);">
                <path
                  d="M923 283.6a260.04 260.04 0 0 0-56.9-82.8 264.4 264.4 0 0 0-84-55.5A265.34 265.34 0 0 0 679.7 125c-49.3 0-97.4 13.5-139.2 39-10 6.1-19.5 12.8-28.5 20.1-9-7.3-18.5-14-28.5-20.1-41.8-25.5-89.9-39-139.2-39-35.5 0-69.9 6.8-102.4 20.3-31.4 13-59.7 31.7-84 55.5a258.44 258.44 0 0 0-56.9 82.8c-13.9 32.3-21 66.6-21 101.9 0 33.3 6.8 68 20.3 103.3 11.3 29.5 27.5 60.1 48.2 91 32.8 48.9 77.9 99.9 133.9 151.6 92.8 85.7 184.7 144.9 188.6 147.3l23.7 15.2c10.5 6.7 24 6.7 34.5 0l23.7-15.2c3.9-2.5 95.7-61.6 188.6-147.3 56-51.7 101.1-102.7 133.9-151.6 20.7-30.9 37-61.5 48.2-91 13.5-35.3 20.3-70 20.3-103.3.1-35.3-7-69.6-20.9-101.9z">

                </path>
            </svg>
        </span>
    </div>
    <div>
        <span data-type="top_note1" class="--mb--rich-text" data-boldtype="0">
            谷粒捐物资捐赠系统
        </span>
    </div>
    <div v-if="header_info.if_show_navi" style="margin-left: auto">
      <div data-type="top_title2">
<!--        <p :class="coms===1?'hear1':'hear'"></p>-->
<!--        判断coms是否等于1，如果等于就追加hear1类名否则追加hear类名-->
        <span data-type="top_note2" :class="[header_info.height_line===0?'heightline':'','--mb--rich-text']" data-boldtype="0" @click="gotoHome()">
          首页
        </span>
      </div>
      <div data-type="top_title2">
        <span data-type="top_note2" :class="header_info.height_line===1?'--mb--rich-text heightline':'--mb--rich-text'" data-boldtype="0">
          平台介绍
        </span>
      </div>
      <div data-type="top_title2">
        <span v-if="header_info.user_type==='1'" data-type="top_note2" :class="header_info.height_line===2?'--mb--rich-text heightline':'--mb--rich-text'" data-boldtype="0">
          项目发布
        </span>
        <span v-else="" data-type="top_note2" :class="header_info.height_line===2?'--mb--rich-text heightline':'--mb--rich-text'" data-boldtype="0" @click="gotoProjectList()">
          捐赠项目
        </span>
      </div>
      <div data-type="top_title2">
        <span data-type="top_note2" :class="header_info.height_line===3?'--mb--rich-text heightline':'--mb--rich-text'" data-boldtype="0">
          关于我们
        </span>
      </div>
    </div>
      <div :style="{'margin-left':header_info.if_show_navi ? '':'auto'}">
      <div data-type="top_title2" style="display: table">
        <div class="widget svg_icon_path clickable" style="display: table-cell">
          <span class="icon svg-icon-wrap svg-fa5 svg-fa5-sign-in-alt-fas-wrap"
                style="width: 24px; height: 24px; min-width: 24px; min-height: 24px;" @click="gotoLogin()">
            <svg v-if="header_info.if_login" xmlns="http://www.w3.org/2000/svg" class="svg-icon" viewBox="64 64 896 896" width="27"
                 height="27"
                 style="fill: rgb(16, 16, 16);">
              <path
                d="M858.5 763.6a374 374 0 0 0-80.6-119.5 375.63 375.63 0 0 0-119.5-80.6c-.4-.2-.8-.3-1.2-.5C719.5 518 760 444.7 760 362c0-137-111-248-248-248S264 225 264 362c0 82.7 40.5 156 102.8 201.1-.4.2-.8.3-1.2.5-44.8 18.9-85 46-119.5 80.6a375.63 375.63 0 0 0-80.6 119.5A371.7 371.7 0 0 0 136 901.8a8 8 0 0 0 8 8.2h60c4.4 0 7.9-3.5 8-7.8 2-77.2 33-149.5 87.8-204.3 56.7-56.7 132-87.9 212.2-87.9s155.5 31.2 212.2 87.9C779 752.7 810 825 812 902.2c.1 4.4 3.6 7.8 8 7.8h60a8 8 0 0 0 8-8.2c-1-47.8-10.9-94.3-29.5-138.2zM512 534c-45.9 0-89.1-17.9-121.6-50.4S340 407.9 340 362c0-45.9 17.9-89.1 50.4-121.6S466.1 190 512 190s89.1 17.9 121.6 50.4S684 316.1 684 362c0 45.9-17.9 89.1-50.4 121.6S557.9 534 512 534z">

              </path>
            </svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" class="svg-icon" viewBox="0 0 512 512" width="24" height="24"
                 style="fill: rgb(16, 16, 16);">
              <path
                d="M416 448h-84c-6.6 0-12-5.4-12-12v-40c0-6.6 5.4-12 12-12h84c17.7 0 32-14.3 32-32V160c0-17.7-14.3-32-32-32h-84c-6.6 0-12-5.4-12-12V76c0-6.6 5.4-12 12-12h84c53 0 96 43 96 96v192c0 53-43 96-96 96zm-47-201L201 79c-15-15-41-4.5-41 17v96H24c-13.3 0-24 10.7-24 24v96c0 13.3 10.7 24 24 24h136v96c0 21.5 26 32 41 17l168-168c9.3-9.4 9.3-24.6 0-34z">

              </path>
            </svg>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "MainTop",
  props: {
    header_info: {
      type: Object,//type为Array,default为函数
      default() {
      }
    },//用props属性进行传数据，此时子组件已经获取到list的数据了
  },
  created() {
    console.log(this.header_info)
  },
  methods: {
    gotoProjectList() {
      console.log(this.header_info.height_line)
      if(this.header_info.height_line!==2){
        //直接跳转
        // 修改heightline下标，将信息传值给下一个页面
        console.log(this.header_info)
        // 跳转
        this.$router.push({
          name: '项目列表',
          path: '/projectList',
          params: this.header_info
        });
      }
    },
    gotoHome() {
      // console.log(this.header_info.height_line)
      if(this.header_info.height_line!==0){
        //直接跳转
        // 修改heightline下标，将信息传值给下一个页面
        console.log(this.header_info)
        // 跳转
        this.$router.push({
          name: 'index',
          path: '/',
          params: this.header_info
        });
      }
    },
    gotoLogin() {
      //直接跳转到个人中心
      this.$router.push('/login');
    }
  }
}
</script>

<style scoped>
  div {
    display: flex;
  }

  div[data-type="top_title"] {
    left: 0px;
    top: 0px;
    width: 1440px;
    height: 100px;
    position: absolute;

  }

  div[data-type="top_title2"] {
    height: 27px;
    width: 72px;
    top: 47px;
    font-size: 14px;
    text-align: left;
    margin: 20px;
    margin-top: 35px;
  }

  div[data-type="top_logo"] {
    margin: 20px;
    width: 64px;
    height: 64px;
    left: 39px;
    top: 33px;
  }

  span[data-type="top_note1"] {
    font-size: 36px;
    color: rgb(174, 174, 174);
    line-height: 54px;
    margin-top: 20px;
  }

  span[data-type="top_note2"] {
    font-family: SourceHanSansSC;
    font-weight: 400;
    font-size: 18px;
    color: rgb(16, 16, 16);
    font-style: normal;
    letter-spacing: 0px;
    line-height: 27px;
    text-decoration: none;
  }
  .heightline{
    color: crimson!important;
  }
</style>
