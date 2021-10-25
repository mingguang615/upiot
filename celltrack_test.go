package upiot

import (
	"log"
	"testing"
)

func TestClient_GetCellTrack(t *testing.T) {
	resp, err := cli.GetCellTrack("460,01,40977,2205409")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
