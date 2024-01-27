package dao

type UserDao struct {
	ID   int
	Name string
	Age  int
	Sex  string
}

func (UserDao) TableName() string {
	return "user"
}
