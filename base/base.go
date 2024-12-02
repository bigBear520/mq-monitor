package base

type IMessageQueue interface {
	Producer(topic string, data []byte)
	Consumer(topic, channel string, ch chan []byte, f func(b []byte))
}

type MessageQueueTopic struct {
	MqType      string
	Url         string
	UserName    string
	Pass        string
	Topic       string //监控的队列
	Cron        string // 定时任务执行周期
	TaskType    string `default:"consuming"` // 监控的类型，目前就是有没有在消费
	AdminPort   int16  `default:"15672"`     // 端口
	VirtualHost string `default:"/"`         // 端口
}
