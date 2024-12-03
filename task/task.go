package task

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bigBear520/mq-monitor/executor"
	"github.com/bigBear520/mq-monitor/notice"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)
import "github.com/bigBear520/mq-monitor/base"

func AddTask(topic base.MessageQueueTopic) {
	queues := topic.QueueList
	if len(queues) == 0 {
		return
	}
	//遍历当前mq服务器下的所有队列
	for i := range queues {
		queue := queues[i]
		monitor := newMonitor(topic, queue)
		// 定时任务
		_ = executor.CronExecutor.AddJob(queue.Cron, monitor)
		monitor.Run()
	}

}

func newMonitor(topic base.MessageQueueTopic, queue base.Queue) *Monitor {
	handler := getHandler(topic.MqType, queue.TaskType)
	m := &Monitor{
		topic:   topic,
		queue:   queue,
		handler: handler,
	}
	return m
}

type Monitor struct {
	mqConnect string
	queue     base.Queue
	topic     base.MessageQueueTopic
	handler   func(topic base.MessageQueueTopic, queue base.Queue)
}

func (r *Monitor) Run() {
	r.handler(r.topic, r.queue)
}

func getHandler(MqType string, TaskType string) func(topic base.MessageQueueTopic, queue base.Queue) {
	//todo 目前只支持rabbitmq
	return monitorRabbitMq
}

func monitorRabbitMq(topic base.MessageQueueTopic, queue base.Queue) {
	info := getRabbitmqInfo(topic, queue)
	if info == nil {
		log.Println("获取rabbitmq信息失败失败")
		return
	}
	unackConut := info["messages_unacknowledged"]
	i := info["messages"] //todo 或者用 messages_ready？
	messageCount := int64(i.(float64))
	consumers := info["consumers"]
	log.Println("consumers:", consumers)

	// 看当前unack是否为0， message是否为0，上一次的 unack是否为0 ，
	if unackConut == 0 && messageCount > 0 {
		log.Println("开始通知异常情况")
		noticeInfo := "队列消费异常，请检查：{}"
		notice.SendMessage(noticeInfo, 1)
	}

	return
}
func getRabbitmqInfo(topic base.MessageQueueTopic, queue base.Queue) map[string]interface{} {
	var result map[string]interface{}
	virtualHost := strings.ReplaceAll(queue.VirtualHost, "/", "%2F")

	url := topic.Url + ":" + strconv.Itoa(int(topic.AdminPort)) + "/api/queues/" + virtualHost + "/" + queue.Topic

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	auth := base64.StdEncoding.EncodeToString([]byte(topic.UserName + ":" + topic.Pass))
	// 设置Authorization头
	req.Header.Add("Authorization", "Basic "+auth)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("获取队列监控失败：queue=", queue.Topic, err)
		return nil
	}
	body := response.Body
	defer body.Close()
	all, err := io.ReadAll(body)
	// 将响应的body转成结构体
	json.Unmarshal(all, &result)
	return result
}
