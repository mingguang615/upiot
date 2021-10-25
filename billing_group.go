package upiot

type RespBillingGroup struct {
	Rows []struct {
		Carrier  string `json:"carrier"`
		BgCode   string `json:"bg_code"`
		Name     string `json:"name"`
		DataPlan string `json:"data_plan"`
	}
}

func (c client) GetBillingGroup() (*RespBillingGroup, error) {
	resp := new(RespBillingGroup)
	err := c.get(c.url()+"/billing_group/", nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
