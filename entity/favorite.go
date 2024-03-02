package entity

import (
	"sync"

	"gorm.io/gorm"
)

type Favorite struct {
	Id      int64
	UserId  int64
	VideoId int64
	User    User  `json:"-"`
	Video   Video `json:"-"`
}

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(
		func() {
			favoriteDao = &FavoriteDao{}
		},
	)
	return favoriteDao
}

func (*FavoriteDao) Exist(userId, videoId int64) (bool, error) {
	var cnt int64
	if err := DB.Model(&Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Count(&cnt).Error; err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (*FavoriteDao) Create(userId, videoId int64) error {
	favorite := Favorite{UserId: userId, VideoId: videoId}
	err := DB.Model(&Favorite{}).Create(&favorite).Error
	if err != nil {
		return err
	}
	err = DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	if err != nil {
		return err
	}
	err = DB.Model(&User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	return err
}

func (*FavoriteDao) Delete(userId, videoId int64) error {
	favorite := Favorite{UserId: userId, VideoId: videoId}
	err := DB.Model(&Favorite{}).Where("user_id = ? and video_id = ?", userId, videoId).Delete(&favorite).Error
	if err != nil {
		return err
	}
	err = DB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	if err != nil {
		return err
	}
	err = DB.Model(&User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	return err
}

func (*FavoriteDao) QueryVideoListByUserId(userId int64) ([]Video, error) {
	var videoList []Video
	var favoriteList []Favorite
	if err := DB.Model(&Favorite{}).Where("user_id = ?", userId).Find(&favoriteList).Error; err != nil {
		return videoList, err
	}
	var videoIds []int64
	for _, v := range favoriteList {
		videoIds = append(videoIds, v.VideoId)
	}
	err := DB.Model(&Video{}).Where("id in ?", videoIds).Find(&videoList).Error
	return videoList, err
}
