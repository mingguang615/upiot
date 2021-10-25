package upiot

import (
	"net/url"
	"testing"
)

func TestSign(t *testing.T) {
	urls := url.Values{
		"key1": []string{"value1"},
		"key2": []string{"value"},
	}
	sign := cli.sign(urls)
	t.Log(sign)
}
