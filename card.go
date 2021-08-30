package upiot

import (
	"errors"
	"fmt"
	"net/url"
)

type RespCardInfo struct {
	Msisdn            string `json:"msisdn"`               // 10648xxxx1234
	Iccid             string `json:"iccid"`                // 8986xxxx1234
	Imsi              string `json:"imsi"`                 // 4600xxxxxxx0515
	Imsi2             string `json:"imsi2"`                // xxxx
	Carrier           string `json:"carrier"`              // 运营商
	SpCode            string `json:"sp_code"`              // 短信端口号
	DataPlan          string `json:"data_plan"`            // 套餐大小
	ExpiryDate        string `json:"expiry_date"`          // 计费结束日期
	DataUsage         string `json:"data_usage"`           // 当月流量
	AccountStatus     string `json:"account_status"`       // 卡状态
	Active            bool   `json:"active"`               // 激活/未激活
	TestValidDate     string `json:"test_valid_date"`      //  测试期起始日期
	SilentValidDate   string `json:"silent_valid_date"`    // 沉默期起始日期
	OutboundDate      string `json:"outbound_date"`        // 出库日期
	ActiveDate        string `json:"active_date"`          // 激活日期
	SupportSms        bool   `json:"support_sms"`          // 是否支持短信
	DataBalance       string `json:"data_balance"`         // 剩余流量
	TestUsedDataUsage string `json:"test_used_data_usage"` // 测试期已用流量
	SimType           string `json:"sim_type"`             // sim卡类型
	Accumulated       bool   `json:"accumulated"`          // 是否是累计卡boolean类型
	Code              string `json:"code"`                 // 计费组code
}

type RespCardNoList struct {
	Rows []struct {
		Msisdn string `json:"msisdn"`
		Iccid  string `json:"iccid"`
		Imsi   string `json:"imsi"`
		Imsi2  string `json:"imsi2"`
		Imei   string `json:"imei"`
	} `json:"rows"`
	Page     int64 `json:"page"`
	PerPage  int64 `json:"per_page"`
	NumPages int64 `json:"num_pages"`
}

type RespCardUsageInfo struct {
	Rows []struct {
		Msisdn          string `json:"msisdn"`            //msisdn",
		Iccid           string `json:"iccid"`             //iccid",
		Imsi            string `json:"imsi"`              //imsi",
		Imsi2           string `json:"imsi2"`             //imsi2",
		AccountStatus   string `json:"account_status"`    //卡状态,
		Carrier         int    `json:"carrier"`           //运营商,
		DataUsage       string `json:"data_usage"`        //流量用量,
		DataBalance     string `json:"data_balance"`      //剩余流量,
		DataPlan        string `json:"data_plan"`         //流量套餐,
		VoiceUsage      string `json:"voice_usage"`       //语音用量,
		VoiceBalance    string `json:"voice_balance"`     //剩余语音,
		VoiceMo         string `json:"voice_mo"`          //语音呼出时长(秒),
		VoiceMt         string `json:"voice_mt"`          //语音呼入时长(秒),
		VoicePlan       int    `json:"voice_plan"`        //语音套餐,
		SmsMo           string `json:"sms_mo"`            //上行短信数量,
		SmsMt           string `json:"sms_mt"`            //下行短信数量,
		ExpiryDate      string `json:"expiry_date"`       //计费结束日期,
		TestValidDate   string `json:"test_valid_date"`   //测试期起始日期,
		SilentValidDate string `json:"silent_valid_date"` //沉默期起始日期,
		TestDataUsage   string `json:"test_data_usage"`   //测试期已用流量,
		ActiveDate      string `json:"active_date"`       //激活日期,
		OutboundDate    string `json:"outbound_date"`     //出库日期,
		SupportSms      bool   `json:"support_sms"`       // 是否支持短信,
		SupportVoice    bool   `json:"support_voice"`     //是否支持语音,
		Month           string `json:"month"`             //201810,
		BgCode          string `json:"bg_code"`           //计费组code,
		UpdatedTime     string `json:"updated_time"`      //"流量更新时间，格式：2020-09-01 01:21"
	} `json:"rows"`

	Failed []string `json:"failed"`
}

type RespCardMonthlyUsageLog struct {
	Rows []struct {
		Msisdn    string  `json:"msisdn"`
		Iccid     string  `json:"iccid"`
		DataUsage float64 `json:"data_usage"`
	}
	NumPages int64
}

type RespCardMonthlySmsLog struct {
	Rows []struct {
		Msisdn string `json:"msisdn"`
		Iccid  string `json:"iccid"`
		Mo     int64  `json:"mo"` //指定月份上行短信数量
		Mt     int64  `json:"mt"` //指定月份下行短信数量
	}
}

type RespCardLocate struct {
	Lat    string `json:"lat"`
	Msg    string `json:"msg"`
	Status bool   `json:"status"`
	Lon    string `json:"lon"`
}

type ReqCardUsageInfo struct {
	Msisdns []string `json:"msisdns"`         //支持多个物联卡号(msisdn/iccid/imsi，可混合)，最多50个
	Month   string   `json:"month,omitempty"` //可选，无此参数返回当月数据，支持查询最近12个月用量信息
}

/*
param
bg_code string  //可选，无此参数返回账户内所有的卡，否则返回计费组内的卡
page    int //可选  请求的第几页，可选，默认为第1页
per_page int // 可选 每页显示数量，可选，默认每页显示100张
*/
func (c Client) GetCardNoList(param url.Values) (*RespCardNoList, error) {
	resp := new(RespCardNoList)
	err := c.get(c.url()+"/card_no_list", param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//获取卡信息 id可为 msisdn|iccid|imsi
func (c Client) GetCardInfo(id string) (*RespCardInfo, error) {
	resp := new(RespCardInfo)
	err := c.get(c.url()+"/card/"+id+"/", nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//获取卡流量使用信息
func (c Client) GetCardUsageInfo(req *ReqCardUsageInfo) (*RespCardUsageInfo, error) {
	resp := new(RespCardUsageInfo)
	err := c.post(c.url()+"/card_usage_info/", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
param
bg_code string //必填，计费组代码
page  int   //可选  请求的第几页， 默认为第1页
per_page int// 可选 每页显示数量， 默认每页显示1000张
month int //查询月份，默认返回当月数据
*/
//计费组物联卡月流量日志
func (c Client) GetCardMonthlyUsageLog(param url.Values) (*RespCardMonthlyUsageLog, error) {
	resp := new(RespCardMonthlyUsageLog)
	err := c.get(c.url()+"/card/monthly_usagelog/", param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
param
bg_code string //必填，计费组代码
page  int   //可选  请求的第几页， 默认为第1页
per_page int// 可选 每页显示数量， 默认每页显示1000张
month int //查询月份，默认返回当月数据
*/
//计费组物联卡月短信日志
func (c Client) GetCardMonthlySms(param url.Values) (*RespCardMonthlySmsLog, error) {
	resp := new(RespCardMonthlySmsLog)
	err := c.get(c.url()+"/card/monthly_sms/", param, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c Client) GetCardLocate(id string) (*RespCardLocate, error) {
	if len(id) == 0 {
		return nil, errors.New("ID cannot be empty")
	}
	resp := new(RespCardLocate)
	err := c.get(fmt.Sprintf("%s/card/%s/locate/", c.url(), id), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
