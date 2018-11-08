package rocketmq

type SendRequest struct {
	producerGroup string
	messageQueue  *MessageQueue
	nextOffset    int64
}

type SendMessageRequestHeader struct {
	ProducerGroup         string `json:"producerGroup"`
	Topic                 string `json:"topic"`
	DefaultTopic          string `json:"defaultTopic"`
	DefaultTopicQueueNums int    `json:"defaultTopicQueueNums"`
	QueueId               int32  `json:"queueId"`
	SysFlag               int    `json:"sysFlag"`
	BornTimestamp         int64  `json:"bornTimestamp"`
	Flag                  int32  `json:"flag"`
	Properties            string `json:"properties"`
	ReconsumeTimes        int    `json:"reconsumeTimes"`
	UnitMode              bool   `json:"unitMode"`
	MaxReconsumeTimes     int    `json:"maxReconsumeTimes"`
}

type SendMessageService struct {
	pushRequestQueue chan *SendRequest
	producer         *DefaultProducer
}

func NewSendMessageService() *SendMessageService {
	return &SendMessageService{
		pushRequestQueue: make(chan *SendRequest, 1024),
	}
}

func (s *SendMessageService) start() {
	//for {
	//	pushRequest := <-self.pushRequestQueue
	//	self.producer.sendMessage(pushRequest)
	//}
}

type SendMessageRequestHeaderV2 struct {
	A string `json:"a"`
	B string `json:"b"`
	C string `json:"c"`
	D int    `json:"d"`
	E int32  `json:"e"`
	F int    `json:"f"`
	G int64  `json:"g"`
	H int32  `json:"h"`
	I string `json:"i"`
	J int    `json:"j"`
	K bool   `json:"k"`
}

func (v2 *SendMessageRequestHeaderV2) createSendMessageRequestHeaderV2(v1 *SendMessageRequestHeader) {
	v2.A = v1.ProducerGroup
	v2.B = v1.Topic
	v2.C = v1.DefaultTopic
	v2.D = v1.DefaultTopicQueueNums
	v2.E = v1.QueueId
	v2.F = v1.SysFlag
	v2.G = v1.BornTimestamp
	v2.H = v1.Flag
	v2.I = v1.Properties
	v2.J = v1.ReconsumeTimes
	v2.K = v1.UnitMode
}
