export function isvalidUsername (str) {
  const validMap = ['admin', 'editor']
  return validMap.indexOf(str.trim()) >= 0
}

export function isExternal (path) {
  return /^(https?:|mailto:|tel:)/.test(path)
}

export function isPhone (str) {
  const reg = /^1[3|4|5|7|8][0-9]\d{8}$/
  return reg.test(str)
}
export function translateTime (str) {
  const d = new Date(str)
  return {
    y: d.getFullYear(),
    ym: d.getFullYear() + '-' + (d.getMonth() + 1),
    ymd: d.getFullYear() + '-' + (d.getMonth() + 1) + '-' + d.getDate(),
    hms: d.getHours() + ':' + d.getMinutes() + ':' + d.getSeconds(),
    all: d.getFullYear() + '-' + (d.getMonth() + 1) + '-' + d.getDate() + ' ' + d.getHours() + ':' + d.getMinutes() + ':' + d.getSeconds()
  }
}
