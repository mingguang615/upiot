package upiot

import (
	"encoding/json"
	"net/url"
)

const api_v2 = "http://ec.upiot.net/api/v2"

type Client struct {
	apiKey    string
	apiSecret string
}

func NewClient(key, secret string) Client {
	return Client{
		apiKey:    key,
		apiSecret: secret,
	}
}

func (c Client) url() string {
	return api_v2 + "/" + c.apiKey
}

func (c Client) get(url string, req url.Values, data interface{}) error {
	url += "?_sign=" + c.sign(req)
	if len(req) > 0 {
		url += "&" + req.Encode()
	}
	resp, err := get(url)
	if err != nil {
		return err
	}
	return resp.parse(data)
}

func (c Client) post(url string, req interface{}, data interface{}) error {
	reqBS, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := post(url+"?_sign="+md5Hex(string(reqBS)+c.apiSecret), reqBS)
	if err != nil {
		return err
	}
	return resp.parse(data)
}
