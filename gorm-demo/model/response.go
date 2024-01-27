package model

type Response struct {
	Code  int         `json:"code"`
	ReqId string      `json:"reqId"`
	Ts    string      `json:"ts"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}
