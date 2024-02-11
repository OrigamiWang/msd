package errcode

const (
	Success     = 0
	Unknown     = 1
	ServerError = 500
	ServerPanic = 501
	WrongArgs   = 400
	WrongJwt    = 401
	MysqlErr    = 1000
	RedisErr    = 1100
)
