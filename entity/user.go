package entity

import (
	"errors"
	"sync"
)

type User struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	PassWord      string  `json:"-"`
	FollowCount   int64   `json:"follow_count"`
	FollowerCount int64   `json:"follower_count"`
	IsFollow      bool    `json:"is_follow"`
	Videos        []Video `json:"-"`
	FavoriteCount int64   `json:"favorite_count"`
	WorkCount     int64   `json:"work_count"`
	Avatar        string  `json:"avatar"`
}

func (u *User) SetIsFollow(id int64) error {
	var cnt int64
	if err := DB.Model(&Relation{}).Where("user_id = ? and to_user_id = ?", id, u.Id).Count(&cnt).Error; err != nil {
		return err
	}
	u.IsFollow = cnt > 0
	return nil
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		},
	)
	return userDao
}

func (*UserDao) Insert(name, password string) (int64, error) {
	user := User{
		Name:     name,
		PassWord: password,
	}
	if err := DB.Model(&User{}).Create(&user).Error; err != nil {
		return -1, errors.New("插入数据失败")
	}
	DB.Model(&User{}).Find(&user, "name = ?", name)
	return user.Id, nil
}

func (*UserDao) QueryByName(name string) (User, error) {
	var user User
	err := DB.Model(&User{}).Where("name = ?", name).Take(&user).Error
	return user, err
}

func (*UserDao) QueryById(id int64) (User, error) {
	var user User
	err := DB.Model(&User{}).Where("id = ?", id).Take(&user).Error
	return user, err
}
