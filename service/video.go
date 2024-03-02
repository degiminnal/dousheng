package service

import (
	"douyin/entity"
	"douyin/utils"
	"errors"
	"fmt"
	"time"
)

type ResponseVideoPublish struct {
	*ResponseCommon
}

func (r *ResponseVideoPublish) Do(userId int64, title, filename string) error {
	videoName := fmt.Sprintf("%d_%s", userId, filename)
	coverName := fmt.Sprintf("%d_%s.jpg", userId, filename)
	videoPath := fmt.Sprintf("./public/videos/%s", videoName)
	coverPath := fmt.Sprintf("./public/covers/%s", coverName)
	if err := utils.SnapShotFromVideo(videoPath, coverPath, 1); err != nil {
		return errors.New("截取封面失败")
	}
	return entity.NewVideoDaoInstance().Insert(userId, title, videoName, coverName)
}

func VideoPublish(userId int64, title, videoName, coverName string) error {
	return entity.NewVideoDaoInstance().Insert(userId, title, videoName, coverName)
}

type ResponsePublishList struct {
	ResponseCommon
	VideoList []entity.Video `json:"video_list"`
}

func (r *ResponsePublishList) QueryByUserId(userId int64) error {
	videos, err := entity.NewVideoDaoInstance().QueryPublishListByUserId(userId)
	if err != nil {
		return err
	}
	for idx := range videos {
		videos[idx].SetUrl()
	}
	*r = ResponsePublishList{ResponseCommon{0, "获取发布列表成功"}, videos}
	return nil
}

type ResponseFeed struct {
	*ResponseCommon
	NextTime  int64          `json:"next_time"`
	VideoList []entity.Video `json:"video_list"`
}

func (r *ResponseFeed) Do(latestTime, userId int64) error {
	videos, err := entity.NewVideoDaoInstance().QueryFeedListByTime(latestTime)
	if err != nil {
		return err
	}
	var nextTime int64
	if len(videos) == 0 {
		nextTime = int64(time.Now().Unix())
	} else {
		nextTime = videos[len(videos)-1].PublishTime
	}
	for idx := range videos {
		videos[idx].SetUrl()
		videos[idx].User.SetIsFollow(userId)
		videos[idx].SetIsFavorite(userId)
	}
	*r = ResponseFeed{
		&ResponseCommon{0, "成功返回视频流信息"},
		nextTime,
		videos,
	}
	return nil
}
