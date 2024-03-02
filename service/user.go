package service

import (
	"douyin/entity"
	"douyin/midware"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type ResponseRegist struct {
	ResponseCommon
	Id    int64  `json:"user_id"`
	Token string `json:"token"`
}

func (r *ResponseRegist) Do(name, password string) error {
	passwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if _, err := entity.NewUserDaoInstance().QueryByName(name); err == nil {
		return errors.New("用户已存在")
	}
	id, err := entity.NewUserDaoInstance().Insert(name, string(passwd))
	if err != nil {
		return err
	}
	token, err := midware.CreateToken(id)
	if err != nil {
		return err
	}
	*r = ResponseRegist{
		ResponseCommon{0, "注册成功"},
		id,
		token,
	}
	return nil
}

type ResponseLogin struct {
	ResponseCommon
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func (r *ResponseLogin) Do(name, password string) error {
	user, err := entity.NewUserDaoInstance().QueryByName(name)
	if err != nil {
		return errors.New("用户名不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password)); err != nil {
		return errors.New("密码错误")
	}
	token, err := midware.CreateToken(user.Id)
	if err != nil {
		return err
	}
	*r = ResponseLogin{
		ResponseCommon{0, "登陆成功"},
		user.Id,
		token,
	}
	return nil
}

type ResonseUserInfo struct {
	*ResponseCommon
	User *entity.User `json:"user"`
}

func (r *ResonseUserInfo) QueryById(id int64) error {
	user, err := entity.NewUserDaoInstance().QueryById(id)
	if err != nil {
		return errors.New("用户不存在")
	}
	*r = ResonseUserInfo{
		&ResponseCommon{0, "查询用户信息成功"},
		&user,
	}
	return nil
}
