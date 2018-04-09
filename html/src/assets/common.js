function buildHttp(json) {
  let str = '?'
  for (var i in json) {
    if (str !== '?') {
      str += '&'
    }
    str += i + '=' + json[i]
  }
  return str
}
