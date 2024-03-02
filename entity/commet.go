package entity

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id         int64     `json:"id"`
	VideoId    int64     `json:"-"`
	UserId     int64     `json:"-"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"-"`
	CreateDate string    `json:"create_date"`
	User       User      `json:"user"`
	Video      Video     `json:"-"`
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(
		func() {
			commentDao = &CommentDao{}
		},
	)
	return commentDao
}

func (*CommentDao) Create(userId, videoId int64, commentText string, createTime time.Time) (*Comment, error) {
	comment := Comment{
		UserId:     userId,
		VideoId:    videoId,
		Content:    commentText,
		CreateTime: createTime,
		CreateDate: createTime.Format("2006-01-02"),
	}
	if err := DB.Model(&Comment{}).Create(&comment).Error; err != nil {
		return &comment, err
	}
	if err := DB.Model(&User{}).Find(&comment.User, "id = ?", userId).Error; err != nil {
		return &comment, err
	}
	if err := DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		return &comment, err
	}
	comment.setCreateDate()
	return &comment, nil
}

func (*CommentDao) Delete(commentId int64) error {
	comment := Comment{}
	if err := DB.Model(&Comment{}).Where("id = ?", commentId).Take(&comment).Delete(&comment).Error; err != nil {
		return err
	}
	if err := DB.Model(&Video{}).Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (*CommentDao) QueryByVideoId(videoId int64) ([]Comment, error) {
	comments := []Comment{}
	if err := DB.Model(&Comment{}).Preload("User").Where("video_id = ?", videoId).Order("create_time desc").Find(&comments).Error; err != nil {
		return comments, err
	}
	for idx := range comments {
		comments[idx].setCreateDate()
	}
	return comments, nil
}

func (c *Comment) setCreateDate() {
	duration := time.Now().Unix() - c.CreateTime.Unix()
	if duration < 30 {
		c.CreateDate = "刚刚"
	} else if duration < 60 {
		c.CreateDate = fmt.Sprintf("%d秒前", duration)
	} else if duration/60 < 60 {
		c.CreateDate = fmt.Sprintf("%d分钟前", duration/60)
	} else if duration/60/60 < 24 {
		c.CreateDate = fmt.Sprintf("%d小时前", duration/60/60)
	} else if duration/60/60/24 <= 3 {
		c.CreateDate = fmt.Sprintf("%d天前", duration/60/60/24)
	}
}
