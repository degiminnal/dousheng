package entity

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type Relation struct {
	Id       int64 `json:"id"`
	UserId   int64
	ToUserId int64
	User     User `json:"-" gorm:"foreignKey:UserId"`
	ToUser   User `json:"-" gorm:"foreignKey:ToUserId"`
}

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once

func NewRelationDaoInstance() *RelationDao {
	relationOnce.Do(
		func() {
			relationDao = &RelationDao{}
		},
	)
	return relationDao
}

func (*RelationDao) Follow(userId, toUserId int64) error {
	var cnt int64
	if err := DB.Model(&Relation{}).Where("user_id = ? and to_user_id = ?", userId, toUserId).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("请勿重复关注")
	}
	relation := Relation{UserId: userId, ToUserId: toUserId}
	if err := DB.Model(&Relation{}).Create(&relation).Error; err != nil {
		return err
	}
	if err := DB.Model(&User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		return err
	}
	if err := DB.Model(&User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (*RelationDao) UnFollow(userId, toUserId int64) error {
	var cnt int64
	if err := DB.Model(&Relation{}).Where("user_id = ? and to_user_id = ?", userId, toUserId).Count(&cnt).Error; err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("尚未关注")
	}
	relation := Relation{}
	if err := DB.Model(&Relation{}).Where("user_id = ? and to_user_id = ?", userId, toUserId).Delete(&relation).Error; err != nil {
		return err
	}
	if err := DB.Model(&User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		return err
	}
	if err := DB.Model(&User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (*RelationDao) QueryFollowList(userId int64) (*[]User, error) {
	var relations []Relation
	if err := DB.Model(&relations).Preload("ToUser").Find(&relations, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	var users []User
	for idx := range relations {
		relations[idx].ToUser.SetIsFollow(userId)
		users = append(users, relations[idx].ToUser)
	}
	return &users, nil
}

func (*RelationDao) QueryFollowerList(userId int64) (*[]User, error) {
	var relations []Relation
	if err := DB.Model(&relations).Preload("User").Find(&relations, "to_user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	var users []User
	for idx := range relations {
		relations[idx].User.SetIsFollow(userId)
		users = append(users, relations[idx].User)
	}
	return &users, nil
}

func (*RelationDao) QueryFriendList(userId int64) (*[]User, error) {
	var relations []Relation
	if err := DB.Model(&relations).Preload("User").Find(&relations, "to_user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	var users []User
	for idx := range relations {
		var cnt int64
		if err := DB.Model(&Relation{}).Where("user_id = ? and to_user_id = ?", userId, relations[idx].UserId).Count(&cnt).Error; err != nil {
			return nil, err
		}
		if cnt == 0 {
			continue
		}
		relations[idx].User.SetIsFollow(userId)
		users = append(users, relations[idx].User)
	}
	return &users, nil
}
