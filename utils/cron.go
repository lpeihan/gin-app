package utils

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// 每隔5秒执行一次：*/5 * * * * ?

// 每隔1分钟执行一次：0 */1 * * * ?

// 每天23点执行一次：0 0 23 * * ?

// 每天凌晨1点执行一次：0 0 1 * * ?

// 每月1号凌晨1点执行一次：0 0 1 1 * ?

// 每周一和周三晚上22:30: 00 30 22 * * 1,3

// 在26分、29分、33分执行一次：0 26,29,33 * * * ?

// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?

// 每年三月的星期四的下午14:10和14:40:  00 10,40 14 ? 3 4

func InitCron() {
	crontab := cron.New(cron.WithSeconds())

	crontab.AddFunc("0 */1 * * * ?", func() {
		fmt.Println("hello world", time.Now())
	})

	crontab.Start()
}
