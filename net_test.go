package rocketmq

import (
	"fmt"
	"testing"
)

func TestGetLocalIp4(t *testing.T) {
	fmt.Println(GetLocalIp4())
}
