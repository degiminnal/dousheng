package controller

import (
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
		return
	}
	id := rawId.(int64)
	rawVideoId := c.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		SendError(c, err.Error())
		return
	}
	rawActionType := c.Query("action_type")
	actionType, err := strconv.ParseInt(rawActionType, 10, 64)
	if err != nil {
		SendError(c, err.Error())
		return
	}
	response := service.ResponseComment{}
	if actionType == 1 {
		commentText := c.Query("comment_text")
		if err := response.Create(id, videoId, commentText); err != nil {
			SendError(c, err.Error())
			return
		}
	} else if actionType == 2 {
		rawCommentId := c.Query("comment_id")
		commentId, err := strconv.ParseInt(rawCommentId, 10, 64)
		if err != nil {
			SendError(c, err.Error())
			return
		}
		if err := response.Delete(commentId); err != nil {
			SendError(c, err.Error())
			return
		}
	} else {
		SendError(c, "未识别的 action_type")
		return
	}
	c.JSON(http.StatusOK, &response)
}

func CommentList(c *gin.Context) {
	rawVideoId := c.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		SendError(c, err.Error())
		return
	}
	response := service.ResponseCommentList{}
	if err := response.ByVideoId(videoId); err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}
