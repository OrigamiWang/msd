package dto

type SvcConfReq struct {
	Desc string `json:"desc"`
	Env  string `json:"env"`
	Conf string `json:"conf"`
}
