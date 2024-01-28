package dao

type SvcConfig struct {
	ID      int
	SvcName string
	Desc    string
	Env     string
	Conf    string
}

func (SvcConfig) TableName() string {
	return "svc_config"
}
