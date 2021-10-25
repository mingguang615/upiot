package upiot

type RespBalance struct {
	Data float64
}

func (c client) GetBalance() (float64, error) {
	balance := float64(0.0)
	err := c.get(c.url()+"/balance/", nil, &balance)
	if err != nil {
		return balance, err
	}
	return balance, nil
}
