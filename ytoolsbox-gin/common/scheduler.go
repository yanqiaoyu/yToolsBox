package common

import (
	"github.com/robfig/cron/v3"
)

var Scheduler *cron.Cron

// 初始化
func InitScheduler() {
	// 初始化, 返回一个支持至 秒 级别的 cron
	scheduler := cron.New(cron.WithSeconds())
	scheduler.Start()

	// 赋值给全局变量
	Scheduler = scheduler
}

func GetScheduler() *cron.Cron {
	return Scheduler
}
