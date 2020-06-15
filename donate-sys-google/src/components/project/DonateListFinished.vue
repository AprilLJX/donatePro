<template>

</template>

<script>
export default {
  name: "DonateListFinished",
  data() {
    return {
      donater_info: {
        id: "433333333333333333",
        name: "李粒粒",
        tel: "987654321",
      },
      demander_info: {
        contace_name: "谷小红",
        contact_tel: "123456789",
        description: "湖北红十字会致力于公益事业，巴拉巴拉巴拉巴拉，湖北红十字会致力于公益事业",
        name: "湖北红十字",
      },
      project_info: {
        id: '0100001',
        name: '武汉体育中心',
        place: '湖北省武汉市XXX区EEE街道',
        demander: {
          name: '湖北红十字',
          contace_name: "谷小红",
          contact_tel: "123456789",
          description: '湖北红十字会致力于公益事业，巴拉巴拉巴拉巴拉，湖北红十字会致力于公益事业',
        },
        receive_times: '1975',
        emergency: 4,
        needy_group: '新冠医护人员基本需求',
        demande_list: {
          medical: [
            {
              name: '口罩',
              id: '0200001',
              type: 'CN95',
              needy_amount: 2000,
              demander_amount: 10000,
              scal: '个',
              rules: 'WWWWWWWWWWWWWWWWWWW',
              img_url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
              if_need_identify: 1,
            }
          ],
          daliy: [
            {
              name: '帐篷',
              id: '0200002',
              type: '----',
              needy_amount: 200,
              demander_amount: 500,
              scal: '顶',
              if_need_identify: 0,
            },
            {
              name: '被子',
              id: '0200003',
              type: '1.5m',
              needy_amount: 200,
              demander_amount: 500,
              scal: '床',
              if_need_identify: 0,
            }
          ]
        },
        donate_footprint: [{
          user: "user1",
          id: '0300001',
          time: "2020年2月20日",
          message: "今生无悔如华夏，来生还做中华家。因为你们，支撑起万千中华儿女的信念。加油，同时也要保护好自己~",
          things: "被子4床，口罩20个"
        },
          {
            user: "user2",
            id: '0300002',
            time: "2020年2月7日",
            message: "",
            things: "口罩40个"
          },
          {
            user: "user3",
            id: '0300003',
            time: "2020年2月8日",
            message: "",
            things: "口罩100个"
          }
        ]

      },
      supplies_info: {
        add_identitfy: 2,
        amount: 1,
        img_info: {},
        product: "a",
        rules: "a",
        supply_id: "0200003",
        supply: {
          demander_amount: 500,
          id: "0200003",
          if_need_identify: 0,
          name: "被子",
          needy_amount: 200,
          scal: "床",
          type: "1.5m",
        },
        donate_msg: "",
      }

    }
  }
}
</script>

<style scoped>

</style>
