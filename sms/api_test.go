package sms

import (
	"fmt"
	"testing"
)

func TestClient_Send(t *testing.T) {
	c := NewSMSClient("00000", "xxxxxxxxx")
	if code, err := c.Send(1, "13800138000", "888888", "5"); err != nil {
		fmt.Println(code, err)
		t.Fatal("send sms failed")
	}
}
