package rocketmq

import (
	"testing"
	"time"
)

var cGroup, cTopic = "consumerGroup", "ifp_fare"
var cConf = &Config{
	Namesrv: "10.100.159.200:9876;10.100.157.34:9876",
	// ClientIp:     "192.168.1.23",
	InstanceName: "DEFAULT",
}

func TestConsume(t *testing.T) {
	consumer, err := NewDefaultConsumer(cGroup, cConf)
	if err != nil {
		t.Error(err)
	}
	consumer.Subscribe(cTopic, "*")
	consumer.RegisterMessageListener(
		func(msgs []*MessageExt) error {
			for i, msg := range msgs {
				t.Log("msg", i, msg.Topic, msg.Flag, msg.Properties, string(msg.Body))
			}
			t.Log("Consume success!")
			return nil
		})
	consumer.Start()

	time.Sleep(60 * time.Second)
}
