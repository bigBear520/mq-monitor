package message

import (
	"github.com/CodyGuo/dingtalk"
	"github.com/CodyGuo/dingtalk/pkg/robot"
	"log"
)

func SendMessage(noticeInfo string, noticeType int8) {
	switch noticeType {
	case 1:

	}
}

var dingUrl = "https://oapi.dingtalk." +
	"com/robot/send?access_token=be12aebc1c01283c16a6366b12b4c74b265344d2f09470e91e02d357e4713d72"
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
