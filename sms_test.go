package upiot

import (
	"log"
	"testing"
)

func TestClient_SendSms(t *testing.T) {
	resp, err := cli.SendSms(&ReqSms{
		Msisdns: []string{"aaa"},
		Content: "bbb",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
