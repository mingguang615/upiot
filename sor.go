package upiot

type RespSor struct {
}

type ReqSor struct {
	Number string `json:"number"`
	Type   string `json:"type"` // 停复机类型 分为两种类型。00 表示复机， 01表示停机
}

//设置停复机
func (c Client) SetSor(req *ReqSor) (*Response, error) {
	resp := new(Response)
	err := c.post(c.url()+"/sor/", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
