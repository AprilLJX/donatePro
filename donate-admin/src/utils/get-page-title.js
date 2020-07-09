import defaultSettings from '@/settings'

const title = defaultSettings.title || '谷粒捐 Donation Admin System'

export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`
  }
  return `${title}`
}
