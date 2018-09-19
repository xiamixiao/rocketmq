package main

import (
	"fmt"
	"rocketmq"

	"github.com/golang/glog"
)

func main() {
	var group, topic = "broker-c", "goclienttopic"
	conf := &rocketmq.Config{
		Namesrv: "10.100.159.200:9876;10.100.157.34:9876",
		// ClientIp:     "192.168.1.23",
		InstanceName: "DEFAULT",
	}
	producer, err := rocketmq.NewDefaultProducer(group, conf)
	producer.Start()
	if err != nil {
		glog.Errorln(err)
	}
	sendCallback := func() error {
		fmt.Println("I am callback")
		return nil
	}
	i := 0
	for {
		i++
		msg := rocketmq.NewMessage(topic, []byte(fmt.Sprintf("Hello TurboMQ!%d", i)))
		if err := producer.SendAsync(msg, sendCallback); err != nil {
			glog.Errorln(err)
		} else {
			fmt.Println("Async sending success!")
		}
		if i == 100 {
			break
		}
	}
}
