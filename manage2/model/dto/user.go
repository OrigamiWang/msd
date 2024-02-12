package dto

// 等同于dto

type UserReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
	Pswd string `json:"pswd"`
}

type LoginReq struct {
	Name string `json:"name"`
	Pswd string `json:"pswd"`
}
