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

function format (result) {
  if (result >= 0 && result < 10) {
    result = '0' + result
  }
  return result || 0
}
export function translateTime (str) {
  const d = new Date(str)
  let y = d.getFullYear()
  let mon = format(d.getMonth() + 1)
  let day = format(d.getDate())
  let h = format(d.getHours())
  let m = format(d.getMinutes())
  let s = format(d.getSeconds())
  return {
    y,
    ym: `${y}-${mon}`,
    ymd: `${y}-${mon}-${day}`,
    hm: `${h}:${m}`,
    all: `${y}-${mon}-${day} ${h}:${m}:${s}`
  }
}
