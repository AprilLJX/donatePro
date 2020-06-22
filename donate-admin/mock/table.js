const Mock = require('mockjs')

// 捐赠方信息
const donator_data = Mock.mock({
  'items|30': [{
    id: '@id',
    'status|1': ['published', 'draft', 'deleted'],
    account: 'name',
    donate_times: '@integer(300, 5000)'
  }]
})
// 受赠方信息
const demander_data = Mock.mock({
  'items|10': [{
    id: '@id',
    'status|1': ['审核通过', '待审核', '审核未通过'],
    account: 'name',
    name: '@cname',
    credit_code: '@integer(300, 5000)',
    company:'红十字会',
    com_category:'医疗慈善机构',
    com_address:'@province()'+'@city()',
    com_profile:'@csentence(2,20)'
  }]
})
// 物资信息
const supplies_data = Mock.mock({
  'items|50': [{
    id: '@id',
    'status|1': ['生活物资', '医疗物资'],
    name: '@cword(2,5)',
    standard:'@csentence(2,10)',
  }]
})
// 项目信息
const project_data = Mock.mock({
  'items|50': [{
    id: '@id',
    pro_name:'@cword(2,7)',
    demander_id:'@integer(1,100)',
    introduction:'@csentence(2,20)',
    'category|1': ['生活', '医疗','教育'],
    rec_address:'@province()'+'@city()'+'@cword(2,3)'+'街',
    materials: '@cword(2,5)'+'@integer(1,50)'+'个；'+'@cword(2,5)'+'@integer(1,50)'+'个；'+'@cword(2,5)'+'@integer(1,50)'+'个',
    init_time:'@date(yy-mm-dd)'+' 00:00:00',
    end_time:'@date(yy-mm-dd)'+ ' 23:59:59',
    'status|1': ['已完成', '进行中','待审核','待发布'],
  }]
})
// 捐赠单
const target_donation_data = Mock.mock({
  'items|50': [{
    id: '@id',
    taget_id:'@id',
    donator_id:'@id',
    project_id:'@id',
    // 'if_anonymous|1':['Yes','No'],
    if_anonymous:'@integer(0,1)',
    donate_materials: '@cword(2,5)'+'@integer(1,50)'+'个；'+'@cword(2,5)'+'@integer(1,50)'+'个；'+'@cword(2,5)'+'@integer(1,50)'+'个',
    message:'@csentence(5,50)',
    donate_time:'@date(yy-mm-dd)'+'@time(hh-mm-ss)',
    'status|1': ['已完成', '进行中','待审核'],
  }]
})

const no_target_donation_data = Mock.mock({
  'items|50': [{
    id: '@id',
    taget_id:'@id',
    donator_id:'@id',
    if_anonymous:'@integer(0,1)',
    donate_materials: '@cword(2,5)'+'@integer(1,50)'+'个；'+'@cword(2,5)'+'@integer(1,50)'+'个；'+'@cword(2,5)'+'@integer(1,50)'+'个',
    message:'@csentence(5,50)',
    donate_time:'@date(yy-mm-dd)'+'@time(hh-mm-ss)',
    'status|1': ['待匹配','待审核'],
  }]
})


module.exports = [
  {
    url: '/donate-admin-sys/table/donatorList',
    type: 'get',
    response: config => {
      const items = donator_data.items
      return {
        code: 20000,
        data: {
          total: items.length,
          items: items
        }
      }
    }
  },
  {
    url: '/donate-admin-sys/table/demanderList',
    type: 'get',
    response: config => {
      const items = demander_data.items
      return {
        code: 20000,
        data: {
          total: items.length,
          items: items
        }
      }
    }
  },
  {
    url: '/donate-admin-sys/table/suppliesList',
    type: 'get',
    response: config => {
      const items = supplies_data.items
      return {
        code: 20000,
        data: {
          total: items.length,
          items: items
        }
      }
    }
  },
  {
    url: '/donate-admin-sys/table/projectList',
    type: 'get',
    response: config => {
      const items = project_data.items
      return {
        code: 20000,
        data: {
          total: items.length,
          items: items
        }
      }
    }
  },
  {
    url: '/donate-admin-sys/table/targetDonationList',
    type: 'get',
    response: config => {
      const items = target_donation_data.items
      return {
        code: 20000,
        data: {
          total: items.length,
          items: items
        }
      }
    }
  },
  {
    url: '/donate-admin-sys/table/noTargetDonationList',
    type: 'get',
    response: config => {
      const items = no_target_donation_data.items
      return {
        code: 20000,
        data: {
          total: items.length,
          items: items
        }
      }
    }
  },
]


