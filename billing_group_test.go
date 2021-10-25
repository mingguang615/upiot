package upiot

import (
	"log"
	"testing"
)

func TestClient_GetBillingGroup(t *testing.T) {
	resp, err := cli.GetBillingGroup()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(resp)
}
