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
	producer.Start()
	if err != nil {
		glog.Errorln(err)
	}
	msg := rocketmq.NewMessage(topic, []byte("Hello RocketMQ!"))
	sendCallback := func() error {
		fmt.Println("I am callback")
		return nil
	}
	i := 0
	for {
		i++
		if err := producer.SendAsync(msg, sendCallback); err != nil {
			glog.Errorln(err)
		} else {
			fmt.Println("Async sending success!")
		}
		if i == 10 {
			break
		}
	}
}
