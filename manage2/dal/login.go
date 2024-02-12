package dal

import "github.com/OrigamiWang/msd/manage2/model/dao"

func Login(name, pswd string) (*dao.UserDao, error) {
	user := &dao.UserDao{}
	result := conn.Where("name = ?", name).Where("pswd = ?", pswd).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
