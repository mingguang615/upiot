package upiot

import (
	"log"
	"testing"
)

func TestClient_SendSms(t *testing.T) {
	resp, err := client.SendSms(&ReqSms{
		Msisdns: []string{"aaa"},
		Content: "bbb",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
