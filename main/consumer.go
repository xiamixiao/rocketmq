package main

import (
	"fmt"
	"rocketmq"
)

var cGroup, cTopic = "consumerGroup", "golang_ifp_fare"
var cConf = &rocketmq.Config{
	Namesrv: "10.100.41.47:9876",
	// ClientIp:     "192.168.1.23",
	InstanceName: "DEFAULT",
}

func main() {
	consumer, err := rocketmq.NewDefaultConsumer(cGroup, cConf)
	if err != nil {
		fmt.Println(err)
	}
	consumer.Subscribe(cTopic, "*")
	consumer.RegisterMessageListener(
		func(msgs []*rocketmq.MessageExt) error {
			for i, msg := range msgs {
				fmt.Println("-----------------------------------------------------")
				fmt.Println("msg", i, msg.Topic, msg.Flag, msg.Properties, string(msg.Body))
			}
			fmt.Println("Consume success!")
			return nil
		})
	consumer.Start()

	// time.Sleep(30 * time.Second)
	select {}
}
