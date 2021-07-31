package upiot

type RespSms struct {
	Rows []struct {
		Msisdn string `json:"msisdn"`
		SmsId  string `json:"sms_id"`
	}
	Failed []string `json:"failed"`
}

type ReqSms struct {
	Msisdns []string `json:"msisdns"`
	Content string   `json:"content,omitempty"`
	Binary  bool     `json:"binary"`
}

func (c Client) SendSms(req *ReqSms) (*RespSms, error) {
	resp := new(RespSms)
	err := c.post(c.url()+"/sms/", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
