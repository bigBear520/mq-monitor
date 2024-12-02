package main

import (
	"github.com/bigBear520/mq-monitor/executor"
	"github.com/bigBear520/mq-monitor/task"
)

func main() {

	//todo 解析配置文件

	task.AddTask()
	executor.CronExecutor.Run()
}
