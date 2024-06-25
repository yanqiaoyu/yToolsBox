export function restoreservertime(){
    // 修改系统时间为正常时间（精确到分，秒的话就设置为30）  2022-2-22 19:45:00
    // 可以直接调用 Linux 命令： ntpdate -u ntp.api.bz   实现，但试了不行
    var date = new Date();
    var year = date.getFullYear() // 年
    var month = date.getMonth() + 1; // 月
    var day = date.getDate(); // 日
    var hour = date.getHours(); // 时
    var minutes = date.getMinutes(); // 分
    var seconds = date.getSeconds() //秒
    // var weekArr = ['星期一', '星期二', '星期三', '星期四', '星期五', '星期六', '星期天'];
    // var week = weekArr[date.getDay()];
    // 给一位数的数据前面加 “0”
    if (month >= 1 && month <= 9) {
    month = "0" + month;
    }
    if (day >= 0 && day <= 9) {
        day = "0" + day;
    }
    if (hour >= 0 && hour <= 9) {
    hour = "0" + hour;
    }
    if (minutes >= 0 && minutes <= 9) {
    minutes = "0" + minutes;
    }
    if (seconds >= 0 && seconds <= 9) {
    seconds = "0" + seconds;
    }

    var now = year + "-" + month + "-" + day + " " + hour + ":" + minutes + ":" + seconds
    const {data : resChangeTime} = await this.$http.get(
    'custom/modifydate?time=normaltime&date=' + now
    )
    // 时间是否恢复成功
    if (resChangeTime.meta.date != now) {
        return 400
    }

    return 200
}