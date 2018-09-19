package main

import (
	"fmt"
	"rocketmq"
	"time"

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

	consumer, err := rocketmq.NewDefaultConsumer(group, conf)
	if err != nil {
		glog.Errorln(err)
	}
	consumer.Subscribe(topic, "*")
	consumer.RegisterMessageListener(
		func(msgs []*rocketmq.MessageExt) error {
			for i, msg := range msgs {
				fmt.Println("msg", i, msg.Topic, msg.Flag, msg.Properties, string(msg.Body))
			}
			fmt.Println("Consume success!")
			return nil
		})
	consumer.Start()

	time.Sleep(10000 * time.Second)
}
