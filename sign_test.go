package upiot

import (
	"log"
	"net/url"
	"testing"
)

func TestSign(t *testing.T) {
	urls := url.Values{
		"key1": []string{"value1"},
		"key2": []string{"value"},
	}
	sign := client.sign(urls)
	log.Println(sign)
}
