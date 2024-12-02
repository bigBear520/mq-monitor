package task

import (
	"encoding/json"
	"fmt"
	"github.com/bigBear520/mq-monitor/executor"
	"github.com/bigBear520/mq-monitor/notice"
	"io"
	"net/http"
	"strings"
)
import "github.com/bigBear520/mq-monitor/base"

func AddTask(topic base.MessageQueueTopic) {
	monitor := newMonitor(topic)

	// 定时任务
	_ = executor.CronExecutor.AddJob(topic.Cron, monitor)
}

func newMonitor(topic base.MessageQueueTopic) *Monitor {
	// todo 根据传入的建立连接
	m := &Monitor{
		mqConnect: topic,
	}

	return m
}

type Monitor struct {
	mqConnect interface{}
	topic     base.MessageQueueTopic
}

func (r *Monitor) Run() {
	monitorRabbitMq(r.topic)
}

func monitorRabbitMq(topic base.MessageQueueTopic) {
	info := getRabbitmqInfo(topic)
	unackConut := info["messages_unacknowledged"]
	i := info["messages_count"]
	messageCount := i.(int64)

	// 看当前unack是否为0， message是否为0，上一次的 unack是否为0 ，
	if unackConut == 0 && messageCount > 0 {
		noticeInfo := " 队列消费异常，请检查：{}"
		notice.SendMessage(noticeInfo, 1)
	}

	return
}
func getRabbitmqInfo(topic base.MessageQueueTopic) map[string]interface{} {
	var result map[string]interface{}

	url := topic.Url + "/api/queues/" + topic.VirtualHost

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("获取队列监控失败：queue=", topic.Topic, err)
		return nil
	}
	body := response.Body
	defer body.Close()
	all, err := io.ReadAll(body)
	// 将响应的body转成结构体
	json.Unmarshal(all, &result)
	return result
}
