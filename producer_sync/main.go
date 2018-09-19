package main

import (
	"fmt"
	"rocketmq"

	"github.com/golang/glog"
)

func main() {
	group := "yyyw-monitor-dev.novalocal"
	topic := "TopicTest"
	conf := &rocketmq.Config{
		Namesrv: "10.100.158.212:9876",
		// ClientIp:     "192.168.1.23",
		InstanceName: "DEFAULT",
	}

	producer, err := rocketmq.NewDefaultProducer(group, conf)
	if err != nil {
		fmt.Println("-------------end-------------")
		glog.Errorln(err)
		return
	}
	producer.Start()
	fmt.Println("------------start--------------")
	i := 0
	for {
		i++
		msg := rocketmq.NewMessage(topic, []byte("Hello RocketMQ!"))
		if sendResult, err := producer.Send(msg); err != nil {
			glog.Errorln("Sync sending fail!")
		} else {
			fmt.Println("Sync sending success!, ", sendResult)
		}
		if i == 100 {
			break
		}
	}
	fmt.Println("-------------end-------------")
}
