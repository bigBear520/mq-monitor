package main

import (
	"github.com/bigBear520/mq-monitor/base"
	"github.com/bigBear520/mq-monitor/task"
	"github.com/spf13/viper"
	"log"
)

func main() {
	//配置文件 暂定json格式
	list := base.MqServeList
	log.Println(list)
	for i := range list {
		task.AddTask(list[i])
	}
	//entries := executor.CronExecutor.Entries()
	//if len(entries) > 0 {
	//	executor.CronExecutor.Run()
	//} else {
	//	log.Println("no usable task in config files, please check the config file ")
	//}

}

func init() {
	viper.SetConfigName("config")   // 配置文件名称(无扩展名)
	viper.SetConfigType("json")     // 或viper.SetConfigType("YAML")
	viper.AddConfigPath("./config") // 配置文件路径,从项目开始
	err := viper.ReadInConfig()     // 查找并读取配置文件
	if err != nil {                 // 处理读取配置文件的错误
		log.Printf("reslove config file err %v\n", err)
	} else {
		err := viper.UnmarshalKey("config", &base.MqServeList)
		if err != nil {
			log.Printf("resolve config file err %v\n", err)
			panic("resolve config file fail")
		}
	}

}
