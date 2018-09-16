package main

import (
	"github.com/robfig/cron"
	"fmt"
	"time"
)

func main() {
	c := cron.New()
	// 注意一定要空格
	c.AddFunc("0 * * * * *", func() {
		fmt.Println("每一分钟0秒一次")
	})


	c.Start()
	time.Sleep(2 * time.Minute)
}