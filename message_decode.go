package rocketmq

import (
	"bytes"
	"encoding/binary"
)

func encodeMessages(msgs []*Message) []byte {
	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.BigEndian, int32(len(msgs))); err != nil {
		return buffer.Bytes()
	}
	var body []byte
	for _, msg := range msgs {
		body = encodeMessage(msg)
		buffer.Write(body)
	}
	res := buffer.Bytes()
	// fmt.Println(res)
	return res
}

func encodeMessage(msg *Message) []byte {
	var topicBytes []byte
	var topicLength int
	if topicBytes = []byte(msg.Topic); len(topicBytes) > 0 {
		topicLength = len(topicBytes)
	} else {
		topicLength = 0
	}

	properties := properties2String(msg.Properties)
	propertiesBytes := []byte(properties)
	propertiesLength := len(propertiesBytes)

	bodyLength := len(msg.Body)

	messageStoreSize := calcMessageLengthWhenBatch(topicLength, bodyLength, propertiesLength)

	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, int32(messageStoreSize))
	binary.Write(buffer, binary.BigEndian, msg.Flag)
	binary.Write(buffer, binary.BigEndian, int32(bodyLength))
	buffer.Write(msg.Body)
	buffer.WriteByte(byte(topicLength))
	buffer.Write(topicBytes)
	binary.Write(buffer, binary.BigEndian, int16(propertiesLength))
	buffer.Write(propertiesBytes)

	return buffer.Bytes()
}

const (
	NAME_VALUE_SEPARATOR int8 = 1
	PROPERTY_SEPARATOR   int8 = 2
)

func properties2String(properties map[string]string) string {
	buffer := new(bytes.Buffer)
	if properties != nil && len(properties) > 0 {
		for name, value := range properties {
			buffer.WriteString(name)
			binary.Write(buffer, binary.BigEndian, NAME_VALUE_SEPARATOR)
			buffer.WriteString(value)
			binary.Write(buffer, binary.BigEndian, PROPERTY_SEPARATOR)
		}
	}
	return buffer.String()
}

func calcMessageLengthWhenBatch(topicLength, bodyLength, propertiesLength int) int {
	return 4 + 4 + 4 + bodyLength + 1 + topicLength + 2 + propertiesLength
}
