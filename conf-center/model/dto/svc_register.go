package dto

type SvcRegisterResp struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Config string `json:"config"`
}
