package dto

type AuthenticateReq struct {
	JwtToken string `json:"jwtToken"`
}

type AuthenticateResp struct {
	Uid   int    `json:"uid"`
	Uname string `json:"uname"`
}
