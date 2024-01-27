package model

type Response struct {
	Code int         `json:"code"`
	Ts   string      `json:"ts"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
