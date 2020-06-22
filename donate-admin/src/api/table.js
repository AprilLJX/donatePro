import request from '@/utils/request'
// user management
export function getDonatorList(params) {
  return request({
    url: '/donate-admin-sys/table/donatorList',
    method: 'get',
    params
  })
}
export function getDemanderList(params) {
  return request({
    url: '/donate-admin-sys/table/demanderList',
    method: 'get',
    params
  })
}
// supplies management
export function getSuppliesList(params) {
  return request({
    url: '/donate-admin-sys/table/suppliesList',
    method: 'get',
    params
  })
}

// project management
export function getProjectList(params) {
  return request({
    url: '/donate-admin-sys/table/projectList',
    method: 'get',
    params
  })
}

// donation management
export function getTargetDonationList(params) {
  return request({
    url: '/donate-admin-sys/table/targetDonationList',
    method: 'get',
    params
  })
}
export function getNoTargetDonationList(params) {
  return request({
    url: '/donate-admin-sys/table/noTargetDonationList',
    method: 'get',
    params
  })
}

export function getList(params) {
  return request({
    url: '/donate-admin-sys/table/List',
    method: 'get',
    params
  })
}
