package upiot

import (
	"log"
	"testing"
)

func TestClient_SetSor(t *testing.T) {
	resp, err := client.SetSor(&ReqSor{
		Number: "aaa",
		Type:   "00",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
