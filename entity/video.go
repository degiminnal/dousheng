package entity

import (
	"errors"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Video struct {
	Id            int64  `json:"id"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `json:"title"`
	IsFavorite    bool   `json:"is_favorite"`
	User          User   `json:"author"`
	UserId        int64  `json:"-"`
	PublishTime   int64  `json:"-"`
}

func (v *Video) SetIsFavorite(userId int64) error {
	var cnt int64
	if err := DB.Model(&Favorite{}).Where("video_id = ? and user_id = ?", v.Id, userId).Count(&cnt).Error; err != nil {
		return err
	}
	v.IsFavorite = cnt > 0
	return nil
}

func (v *Video) SetUrl() error {
	prefix := viper.GetString("video.urlPrefix")
	if prefix == "" {
		return errors.New("获取urlPrefix失败")
	}
	v.PlayUrl = prefix + "videos/" + v.PlayUrl
	v.CoverUrl = prefix + "covers/" + v.CoverUrl
	return nil
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		},
	)
	return videoDao
}

func (*VideoDao) Insert(userId int64, title, videoName, coverName string) error {
	video := Video{
		UserId:      userId,
		Title:       title,
		PlayUrl:     videoName,
		CoverUrl:    coverName,
		PublishTime: int64(time.Now().Unix()),
	}
	err := DB.Model(&Video{}).Create(&video).Error
	if err != nil {
		return err
	}
	err = DB.Model(&User{}).Where("id = ?", userId).Update("work_count", gorm.Expr("work_count + ?", 1)).Error
	return err
}

func (*VideoDao) QueryPublishListByUserId(userId int64) ([]Video, error) {
	var user User
	err := DB.Model(&User{}).Preload("Videos.User").Where("id = ?", userId).Take(&user).Error
	for idx := range user.Videos {
		user.Videos[idx].User.SetIsFollow(userId)
	}
	return user.Videos, err
}

func (*VideoDao) QueryFeedListByTime(latest_time int64) ([]Video, error) {
	var videos []Video
	err := DB.Model(&Video{}).Preload("User").Where("publish_time < ?", latest_time).Order("publish_time desc").Limit(10).Find(&videos).Error
	return videos, err
}
