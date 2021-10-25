package upiot

import (
	"log"
	"net/url"
	"testing"
)

const (
	api_key    = "aaaaaaaaaaa"
	api_secret = "bbbbbbbbbbbbbbb"
)

var cli = NewClient(api_key, api_secret)

func TestClient_GetCardNoList(t *testing.T) {
	resp, err := cli.GetCardNoList(url.Values{
		"bg_code":  []string{"test"},
		"page":     []string{"1"},
		"per_page": []string{"100"},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func TestClient_GetCardInfo(t *testing.T) {
	resp, err := cli.GetCardInfo("10648xxxx1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func TestClient_GetCardUsageInfo(t *testing.T) {
	resp, err := cli.GetCardUsageInfo(&ReqCardUsageInfo{
		Msisdns: []string{"89860412102090029371"},
		Month:   "202107",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func TestClient_GetCardMonthlyUsageLog(t *testing.T) {
	resp, err := cli.GetCardMonthlyUsageLog(url.Values{
		"bg_code":  []string{"1SHFD030M21"},
		"page":     []string{"1"},
		"per_page": []string{"1000"},
		"month":    []string{"202107"},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func TestClient_GetCardMonthlySms(t *testing.T) {
	resp, err := cli.GetCardMonthlySms(url.Values{
		"bg_code":  []string{"1SHFD030M21"},
		"page":     []string{"1"},
		"per_page": []string{"1000"},
		"month":    []string{"202107"},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func TestClient_GetCardLocate(t *testing.T) {
	resp, err := cli.GetCardLocate("89860412102090029371")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
