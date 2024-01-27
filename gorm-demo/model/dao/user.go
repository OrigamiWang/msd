package dao

type UserDao struct {
	ID   int
	Name string
	age  int
	sex  string
}

func (UserDao) TableName() string {
	return "user"
}
