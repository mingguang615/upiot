package upiot

import "net/url"

type RespCellTrack struct {
	Result []struct {
		Id      string `json:"id"`      //纬度
		Lng     string `json:"lng"`     //经度
		Radius  string `json:"radius"`  //覆盖半径参考值
		Address string `json:"address"` //靠近地址描述
		Roads   string `json:"roads"`   //靠近道路描述
		// 以下为 MCC = 460 时列表独有结果
		Rid  string `json:"rid,omitempty"`  //6位行政区划代码
		Rids string `json:"rids,omitempty"` //12位区划代码
		Lats string `json:"lats,omitempty"` //附带纬度
		Lngs string `json:"lngs,omitempty"` //附带经度
	}
	Latitude  string `json:"latitude"`  //纬度, 多基站定位请求时为综合计算结果
	Longitude string `json:"longitude"` //经度, 多基站定位请求时为综合计算结果
}

//curl http://ec.upiot.net/api/v2/9516a58d24cd6b76382f27c8ab4292b0151b2702/celltrack/?_sign=c28270f9eb6a22d3563800e7af461a4b\&bs=460,01,40977,2205409
//bs参数： GSM/UMTS: MCC,MNC,LAC,CI 或 MCC,MNC,LAC,CI,dBm LTE: MCC,MNC,TAC,CI 或 MCC,MNC,TAC,CI,dBm CDMA: MCC,SID,NID,BID 或 MCC,MNC,TAC,CI,dBm 多基站分隔符号: |
func (c Client) GetCellTrack(bs string) (*RespCellTrack, error) {
	resp := new(RespCellTrack)
	param := url.Values{
		"bs": []string{bs},
	}
	err := c.get(c.url()+"/celltrack/", param, resp)
	if err != nil{
		return nil, err
	}
	return resp, nil
}
