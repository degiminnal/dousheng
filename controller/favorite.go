package controller

import (
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FavoriteList(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
		return
	}
	id := rawId.(int64)
	response := service.ResponseFavoriteList{}
	err := response.Do(id)
	if err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}

func FavoriteAction(c *gin.Context) {
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
	response := service.ResponseFavoriteAction{}
	if err := response.Do(id, videoId, actionType); err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}
