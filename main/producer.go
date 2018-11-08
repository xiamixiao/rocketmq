package main

import (
	"fmt"
	"rocketmq"
)

var pGroup, pTopic = "BasePayment", "golang_ifp_fare"
var pConf = &rocketmq.Config{
	Namesrv: "10.100.159.200:9876;10.100.157.34:9876",
	// ClientIp:     "192.168.1.23",
	InstanceName: "DEFAULT",
}

var pConfBatch = &rocketmq.Config{
	Namesrv: "10.100.41.47:9876",
	// ClientIp:     "192.168.1.23",
	InstanceName: "DEFAULT",
}

func main() {
	producer, err := rocketmq.NewDefaultProducer(pGroup, pConfBatch)
	if err != nil {
		fmt.Println(err)
		return
	}
	producer.Start()
	var msgs []*rocketmq.Message
	for i := 0; i < 100; i++ {
		msg := rocketmq.NewMessage(pTopic, []byte(fmt.Sprintf("Hello turboMQ(10-25 11:05)(batch sync), %d", i)))
		msg.Properties["WAIT"] = "true"
		msgs = append(msgs, msg)
	}

	if sendResult, err := producer.SendBatchSync(msgs); err != nil {
		fmt.Println("Sync sending fail!")
	} else {
		fmt.Println(*sendResult)
	}

	fmt.Println("Sync sending success!")
}
