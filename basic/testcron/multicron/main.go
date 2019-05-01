package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

//单个任务结构体
type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {

	//创建任务调度表
	scheduleTable := make(map[string]*CronJob)

	//当前时间
	now := time.Now()

	//创建cronjob
	expr, err := cronexpr.Parse("*/5 * * * * * *")
	if err != nil {
		fmt.Println(err)
	}

	cronJob := &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}

	//注册任务到调度表
	scheduleTable["job1"] = cronJob

	//第二个任务
	expr, err = cronexpr.Parse("*/7 * * * * * *")
	if err != nil {
		fmt.Println(err)
	}
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"] = cronJob

	//启动调度协程
	go func() {
		for {
			now = time.Now()
			for jobName, cronJob := range scheduleTable {
				//判断任务是否过期，判断job的下个时间在当前时间之前
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {

					go func(jobName string) {
						fmt.Println("running", jobName)
					}(jobName)

					//计算下次调度时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println(jobName, "next running time:", cronJob.nextTime)

				}
			}

			//睡眠100毫秒
			select {
				case <-time.NewTimer(100 * time.Millisecond).C:
			}
		}
	}()

	time.Sleep(100 * time.Second)

}
