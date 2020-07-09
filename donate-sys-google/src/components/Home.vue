<template>
  <div>
    <el-header>
      <div id="top">
        <MainTop :header_info="header_info"></MainTop>
      </div>
    </el-header>
    <el-main>
      <div class="group"
           style="left: 0px; top: 120px; width: 1440px; height: 600px; z-index: 20; opacity: 1;">
        <div class="widget image_view"
             style="width: 1440px; height: 600px; left: 0px; z-index: 5; background-color: rgb(229, 229, 229); border-color: rgb(187, 187, 187); border-width: 0px; border-style: solid; font-weight: normal; font-style: normal; opacity: 1;">
          <div class="image-wrapper">
            <div class="wrapper" style="width: 100%;height: 100%">
              <img src="../assets/home1.png" draggable="false" style="display: block; height: auto;">
            </div>
          </div>
        </div>
        <div class="widget"
             style="width: 830px; height: 600px; left: 0px; z-index: 6; background-color: rgb(56, 148, 255); border: none; font-size: 14px; padding: 0px; text-align: center; line-height: 20px; font-weight: normal; font-style: normal; opacity: 0.5;">
          <div class="text" style="padding: 0px;"><p>&nbsp;</p></div>
          <div class="region gesture "></div>
        </div>
        <div class="widget"
             style="width: 361px; height: 89px; left: 120px; top: 275px; z-index: 7; font-size: 14px; text-align: left; font-weight: normal; font-style: normal; opacity: 1;">
          <div class="rich-text">
            <p style="line-height:0; margin-bottom:5px;">
                <span class="--mb--rich-text" data-boldtype="0"
                      style="font-family:SourceHanSansSC; font-weight:400; font-size:60px; color:rgb(255, 255, 255); font-style:normal; letter-spacing:0px; line-height:89px; text-decoration:none;">
            让公益更简单
            </span>
            </p>
          </div>
        </div>
        <div class="widget"
             style="width: 445px; height: 63px; left: 120px; top: 389px; z-index: 8; font-size: 14px; text-align: left; font-weight: normal; font-style: normal; opacity: 1;">
          <div class="rich-text">
            <p style="line-height:0; margin-bottom:5px;">
                <span class="--mb--rich-text" data-boldtype="0"
                      style="font-family:SourceHanSansSC; font-weight:400; font-size:20px; color:rgba(255, 255, 255, 1); font-style:normal; letter-spacing:0px; line-height:29px; text-decoration:none;">
                优质服务值得您信赖
            </span>
            </p>
          </div>
        </div>
        <div class="widget learn_more"
             style="left: 120px; top: 504px; z-index: 9; border: none; background: white">
          <div class="text" style="padding: 0px;">
            <p>了解更多</p>
          </div>
        </div>
        <div class="widget rounded_rect hcenter vmiddle"
             style="width: 50px; height: 5px; left: 665px; top: 675px; z-index: 10; background-color: rgb(174, 174, 174);border: none; ">
          <div class="text" style="padding: 0px;">
            <p>&nbsp;</p>
          </div>
        </div>
        <div class="widget rounded_rect hcenter vmiddle"
             style="width: 50px; height: 5px; left: 725px; top: 675px; z-index: 10; background-color: rgb(174, 174, 174); border: none;">
          <div class="text" style="padding: 0px;">
            <p>&nbsp;</p>
          </div>
        </div>
      </div>
      <div>
        <div style="top: 720px; width: 1440px;">
          <div class="card">
            <div class="card_info">
              <div style="height: 41px; font-size: 14px;text-align: left">
                <p style="line-height:0; margin-bottom:5px;">
                <span style="font-weight: 400; font-size: 28px; line-height: 41px;">
                  XXXXXXXXTEST1
                </span>
                </p>
              </div>
              <div
                style="width: 171px; height: 27px; left: 120px; top: 233px; z-index: 7; font-size: 14px; text-align: left; font-weight: normal; font-style: normal; opacity: 1;">
                <p style="line-height:0; margin-bottom:5px;">
                <span data-boldtype="0"
                      style="font-family:SourceHanSansSC; font-weight:400; font-size:18px; color:rgba(174, 174, 174, 1); font-style:normal; letter-spacing:0px; line-height:27px; text-decoration:none;">
                  XXXXXXXXTEST2
                </span>
                </p>
              </div>
              <div
                style="width: 269px; height: 91px; left: 120px; top: 285px; z-index: 8; font-size: 14px; text-align: left; font-weight: normal; font-style: normal; opacity: 1;">
                <p style="line-height:0; margin-bottom:5px;">
                            <span data-boldtype="0"
                                  style="font-family:SourceHanSansSC; font-weight:400; font-size:18px; color:rgb(16, 16, 16); font-style:normal; letter-spacing:0px; line-height:27px; text-decoration:none;">
                                XXXXXXXXXXXXXXXXXXXTEST3
                            </span>
                </p>
              </div>
              <div class="widget rounded_rect hcenter vmiddle learn_more">
                <div class="text" style="padding: 0px;">
                  <p>了解更多</p>
                </div>
              </div>
            </div>
            <div class="card_img"
                 style="border-width: 0px;">
              <div class="image-wrapper">
                <div class="wrapper">
                  <img src="../assets/home2.png" draggable="false"
                       style="display: block; height: auto;">
                </div>
              </div>
            </div>
          </div>
          <div class="card">
            <div class="card_img"
                 style="border-width: 0px;">
              <div class="image-wrapper">
                <div class="wrapper">
                  <img src="../assets/home2.png" draggable="false"
                       style="display: block; height: auto;">
                </div>
              </div>
            </div>
            <div class="card_info" style="text-align: left">
              <div>
                <p style="line-height:0; margin-bottom:5px;">
                <span class="--mb--rich-text" data-boldtype="0"
                      style="font-family:SourceHanSansSC; font-weight:400; font-size:28px; color:rgb(16, 16, 16); font-style:normal; letter-spacing:0px; line-height:41px; text-decoration:none;">
                    XXXXXXXX
                </span>
                </p>
              </div>
              <div>
                <p style="line-height:0; margin-bottom:5px;">
                            <span data-boldtype="0"
                                  style="font-family:SourceHanSansSC; font-weight:400; font-size:18px; color:rgba(174, 174, 174, 1); font-style:normal; letter-spacing:0px; line-height:27px; text-decoration:none;">
                                XXXXXXXX,
                            </span>
                </p>
              </div>
              <div>
                <p style="line-height:0; margin-bottom:5px;">
                            <span data-boldtype="0"
                                  style="font-family:SourceHanSansSC; font-weight:400; font-size:18px; color:rgb(16, 16, 16); font-style:normal; letter-spacing:0px; line-height:27px; text-decoration:none;">
                                XXXXXXXXXXXXXXXXXXXXXXXX
                            </span>
                </p>
              </div>
              <div class="widget rounded_rect hcenter vmiddle learn_more">
                <div class="text" style="padding: 0px;">
                  <p>了解更多</p>
                </div>
              </div>
            </div>

          </div>
        </div>
      </div>
    </el-main>

    <el-footer>
      <MainBottom></MainBottom>
    </el-footer>

  </div>

</template>


<script>
import '../assets/baseStyle.css'
import MainTop from "./MainTop";
import MainBottom from "./MainBottom";

export default {
  name: "Home",
  components: {MainTop, MainBottom},
  data() {
    return {
      header_info: {
        height_line: 0,
        if_login: false,
        user_type: '0', // 0 is donator, 1 is reciver,
        if_show_navi: true
      },
    }
  }
}

</script>

<style scoped>
  .learn_more {
    width: 160px;
    height: 50px;
    border-color: crimson;
    border-width: 1px;
    border-style: solid;
    color: crimson;
    border-radius: 5px;
    font-size: 14px;
    text-align: center;
    line-height: 20px;

  }
</style>
