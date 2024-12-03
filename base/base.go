package base

type IMessageQueue interface {
	Producer(topic string, data []byte)
	Consumer(topic, channel string, ch chan []byte, f func(b []byte))
}

type MessageQueueTopic struct {
	MqType   string `json:"mqType"`
	Url      string `json:"url"`
	UserName string `json:"userName"`
	Pass     string `json:"pass"`
	// 端口
	AdminPort int16   `json:"adminPort" default:"15672"` // 端口
	QueueList []Queue `json:"queueList"`
}

type Queue struct {
	Topic       string `json:"topic"`                        //监控的队列
	Cron        string `json:"cron"`                         // 定时任务执行周期
	TaskType    string `json:"taskType" default:"consuming"` // 监控的类型，目前就是有没有在消费
	VirtualHost string `json:"virtualHost" default:"/"`
}

var MqServeList = make([]MessageQueueTopic, 10)
