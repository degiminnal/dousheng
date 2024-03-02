package service

import (
	"douyin/entity"
	"errors"
)

type ResponseFavoriteList struct {
	*ResponseCommon
	VideoList []entity.Video `json:"video_list"`
}

func (r *ResponseFavoriteList) Do(userId int64) error {
	videoList, err := entity.NewFavoriteDaoInstance().QueryVideoListByUserId(userId)
	if err != nil {
		return err
	}
	for idx := range videoList {
		videoList[idx].IsFavorite = true
		videoList[idx].SetUrl()
	}
	*r = ResponseFavoriteList{
		&ResponseCommon{0, "查询喜欢列表成功"},
		videoList,
	}
	return nil
}

type ResponseFavoriteAction struct {
	*ResponseCommon
}

func (r *ResponseFavoriteAction) Do(userId, videoId, actionType int64) error {
	switch actionType {
	case 1:
		exist, err := entity.NewFavoriteDaoInstance().Exist(userId, videoId)
		if err != nil {
			return err
		}
		if exist {
			return errors.New("请勿重复点赞")
		}
		if err := entity.NewFavoriteDaoInstance().Create(userId, videoId); err != nil {
			return err
		}
	case 2:
		exist, err := entity.NewFavoriteDaoInstance().Exist(userId, videoId)
		if err != nil {
			return err
		}
		if !exist {
			return errors.New("尚未点赞，无法取消")
		}
		if err := entity.NewFavoriteDaoInstance().Delete(userId, videoId); err != nil {
			return err
		}
	default:
		return errors.New("未识别的action_type")
	}
	*r = ResponseFavoriteAction{
		&ResponseCommon{0, "点赞/取消成功"},
	}
	return nil
}
