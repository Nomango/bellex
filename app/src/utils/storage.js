const sessionStorage = window.sessionStorage

function setSession (key, value) {
  if (typeof value === 'object') {
    value = JSON.stringify(value)
  }
  sessionStorage.setItem(key, value)
}
function getSession (key) {
  const value = sessionStorage.getItem(key)
  try {
    return JSON.parse(value)
  } catch (e) {
    return value
  }
}
function removeSession (key) {
  sessionStorage.removeItem(key)
}
export {
  setSession,
  getSession,
  removeSession
}
