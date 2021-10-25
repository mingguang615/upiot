package upiot

import (
	"log"
	"testing"
)

func TestClient_GetBalance(t *testing.T) {
	resp, err := cli.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
