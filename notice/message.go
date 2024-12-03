package notice

import (
	"github.com/CodyGuo/dingtalk"
	"github.com/CodyGuo/dingtalk/pkg/robot"
	"log"
)

func SendMessage(noticeInfo string, noticeType int8) {
	switch noticeType {
	case 1:
	default:
		sendDing(noticeInfo)

	}
}

var dingUrl = "https://oapi.dingtalk.com/robot/send?access_token=3ed647062cba95ff549d5a31cfbb760bcd29546313091aa3b4e95919cc94cf7c"
var phones = "15528374508"

func sendDing(noticeInfo string) {
	dt := dingtalk.New(dingUrl)
	// text类型
	textContent := noticeInfo
	atMobiles := robot.SendWithAtMobiles([]string{phones})
	if err := dt.RobotSendText(textContent, atMobiles); err != nil {
		log.Println(err)
	}
}

func sendMail(noticeInfo string) {

}
