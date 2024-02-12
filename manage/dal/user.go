package dal

import (
	"github.com/OrigamiWang/msd/manage/model/dao"
	"github.com/OrigamiWang/msd/manage/model/dto"
	"github.com/OrigamiWang/msd/micro/auth/crypto"
	dao2 "github.com/OrigamiWang/msd/micro/model/dao"
	logutil "github.com/OrigamiWang/msd/micro/util/log"
	"gorm.io/gorm"
)

var DATABSE_KEY = "sample_mysql1"

var conn *gorm.DB

func init() {
	var err error
	conn, err = dao2.MySQL(DATABSE_KEY)
	if err != nil {
		logutil.Error("can not connect mysql, database_key: %v, err: %v", DATABSE_KEY, err)
	}
}

func GetUserById(id string) (*dao.UserDao, error) {
	user := &dao.UserDao{}
	result := conn.Where("id = ?", id).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetAllUser() (*[]dao.UserDao, error) {
	users := &[]dao.UserDao{}
	result := conn.Find(users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func UpdateUser(id string, userReq *dto.UserReq) (*dao.UserDao, error) {
	tx := conn.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	encodedPswd := crypto.Md5Encode(userReq.Pswd)
	user := &dao.UserDao{
		Name: userReq.Name,
		Age:  userReq.Age,
		Sex:  userReq.Sex,
		Pswd: encodedPswd,
	}
	err := tx.Model(&dao.UserDao{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		logutil.Error("gorm: update user failed, err: %v", err)
		tx.Rollback()
		return nil, err
	}
	err = tx.Where("id = ?", id).First(user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return user, nil
}

func AddUser(userReq *dto.UserReq) (*dao.UserDao, error) {
	tx := conn.Begin()
	encodedPswd := crypto.Md5Encode(userReq.Pswd)
	user := &dao.UserDao{
		Name: userReq.Name,
		Age:  userReq.Age,
		Sex:  userReq.Sex,
		Pswd: encodedPswd,
	}
	err := tx.Model(&dao.UserDao{}).Create(user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Last(user).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return user, nil
}

func GetUserByName(name string) (*dao.UserDao, error) {
	user := &dao.UserDao{}
	result := conn.Where("name = ?", name).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
