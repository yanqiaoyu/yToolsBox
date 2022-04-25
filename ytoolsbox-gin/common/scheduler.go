package common

import "github.com/jasonlvhit/gocron"

var Scheduler *gocron.Scheduler

// 初始化
func InitScheduler() {
	// 初始化
	scheduler := gocron.NewScheduler()
	scheduler.Start()

	// 赋值给全局变量
	Scheduler = scheduler
}

func GetScheduler() *gocron.Scheduler {
	return Scheduler
}
