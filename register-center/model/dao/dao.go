package dao

type SvcRegister struct {
	ID     int
	Name   string
	Config string
}

func (SvcRegister) TableName() string {
	return "svc_register"
}
