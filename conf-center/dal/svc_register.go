package dal

import "github.com/OrigamiWang/msd/conf-center/model/dao"

func GetSvcByName(name string) (*dao.SvcRegister, error) {
	s := &dao.SvcRegister{}
	result := conn.Where(&dao.SvcRegister{Name: name}).First(s)
	return s, result.Error
}

func GetAllSvc() (*[]dao.SvcRegister, error) {
	s := &[]dao.SvcRegister{}
	result := conn.Find(&s)
	return s, result.Error
}
