package upiot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (resp *Response) decode(i interface{}) error {
	if resp == nil {
		return errors.New("response is empty")
	}
	bs, err := json.Marshal(resp.Data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bs, &i); err != nil {
		return err
	}
	return nil
}

func (resp *Response) parse(data interface{}) error {
	if resp.Code != 200 {
		return errors.New(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, resp.Code, resp.Msg))
	}

	if _, ok := data.(*Response); ok {
		data = resp
		return nil
	}
	if err := resp.decode(&data); err != nil {
		return err
	}
	return nil
}

func get(url string) (*Response, error) {
	log.Println(url)
	cli := &http.Client{Timeout: time.Second * 30}
	res, err := cli.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	log.Println(string(body))
	resp2 := &Response{}
	if err := json.Unmarshal(body, &resp2); err != nil {
		return nil, err
	}

	return resp2, nil
}

func post(url string, reqBS []byte) (*Response, error) {
	cli := &http.Client{Timeout: time.Second * 30}
	res, err := cli.Post(url, "application/json", bytes.NewBuffer(reqBS))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))
	resp2 := &Response{}
	if err := json.Unmarshal(body, &resp2); err != nil {
		return nil, err
	}

	return resp2, nil
}
