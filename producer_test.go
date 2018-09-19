package rocketmq

import (
	"fmt"
	"testing"
)

var pGroup, pTopic = "BasePayment", "ifp_fare"
var pConf = &Config{
	Namesrv: "10.100.159.200:9876;10.100.157.34:9876",
	// ClientIp:     "192.168.1.23",
	InstanceName: "DEFAULT",
}

func TestSendSync(t *testing.T) {
	producer, err := NewDefaultProducer(pGroup, pConf)
	if err != nil {
		t.Error(err)
		return
	}
	producer.Start()
	for i := 0; i < 100; i++ {
		msg := NewMessage(pTopic, []byte(fmt.Sprintf("Hello turboMQ(sync), %d", i)))
		if sendResult, err := producer.Send(msg); err != nil {
			t.Error("Sync sending fail!")
		} else {
			_ = sendResult
			t.Logf("Sync sending success, %d", i)
		}
	}

	t.Log("Sync sending success!")
}

func TestSendAsync(t *testing.T) {
	producer, err := NewDefaultProducer(pGroup, pConf)
	producer.Start()
	if err != nil {
		t.Error(err)
	}
	sendCallback := func() error {
		t.Log("I am callback")
		return nil
	}
	for i := 0; i < 100; i++ {
		msg := NewMessage(pTopic, []byte(fmt.Sprintf("Hello TurboMQ(async), %d", i)))
		if err := producer.SendAsync(msg, sendCallback); err != nil {
			t.Error(err)
		} else {
			t.Logf("Async sending success,%d", i)
		}
	}
	t.Log("Async sending success!")
}

func TestSendOneway(t *testing.T) {
	producer, err := NewDefaultProducer(pGroup, pConf)
	if err != nil {
		t.Fatalf("NewDefaultProducer err, %s", err)
	}
	producer.Start()
	for i := 0; i < 100; i++ {
		msg := NewMessage(pTopic, []byte(fmt.Sprintf("Hello TurboMQ(oneway), %d", i)))
		if err := producer.SendOneway(msg); err != nil {
			t.Fatalf("Oneway sending fail! %s", err.Error())
		} else {
			t.Logf("Oneway sending success, %d", i)
		}
	}

	t.Log("Oneway sending success!")
}
