package service

import (
	"douyin/entity"
	"time"
)

type ResponseComment struct {
	*ResponseCommon
	Comment *entity.Comment `json:"comment,omitempty"`
}

func (r *ResponseComment) Create(userId, videoId int64, commentText string) error {
	createTime := time.Now()
	comment, err := entity.NewCommentDaoInstance().Create(userId, videoId, commentText, createTime)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	*r = ResponseComment{
		&ResponseCommon{0, "评论成功"},
		comment,
	}
	return nil
}

func (r *ResponseComment) Delete(commentId int64) error {
	if err := entity.NewCommentDaoInstance().Delete(commentId); err != nil {
		return err
	}
	*r = ResponseComment{
		&ResponseCommon{0, "删除评论成功"},
		nil,
	}
	return nil
}

type ResponseCommentList struct {
	*ResponseCommon
	CommentList []entity.Comment `json:"comment_list"`
}

func (r *ResponseCommentList) ByVideoId(videoId int64) error {
	commentList, err := entity.NewCommentDaoInstance().QueryByVideoId(videoId)
	if err != nil {
		return err
	}
	*r = ResponseCommentList{
		&ResponseCommon{0, "查询评论列表成功"},
		commentList,
	}
	return nil
}
