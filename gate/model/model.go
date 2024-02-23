package model

type Service struct {
	ID     int                    `mapstructure:"id"`
	Name   string                 `mapstructure:"name"`
	Config map[string]interface{} `mapstructure:"config"`
}
