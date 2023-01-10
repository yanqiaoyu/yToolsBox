import moment from 'moment'

const commonFun = {
  // 转换时间戳
  FormatDate(unixtime) {
    return moment.unix(unixtime).format('YYYY-MM-DD HH:mm:ss')
  }
}

export default commonFun
