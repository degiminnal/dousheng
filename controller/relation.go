package controller

import (
	"douyin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
		return
	}
	id := rawId.(int64)
	rawToUserId := c.Query("to_user_id")
	toUserId, err := strconv.ParseInt(rawToUserId, 10, 64)
	if err != nil {
		SendError(c, "获取to_user_id失败")
		return
	}
	if id == toUserId {
		SendError(c, "无法关注自己")
		return
	}
	rawActionType := c.Query("action_type")
	actionType, err := strconv.ParseInt(rawActionType, 10, 64)
	if err != nil {
		SendError(c, "获取action_type失败")
		return
	}
	response := service.ResonseRelation{}
	if actionType == 1 {
		if err := response.Follow(id, toUserId); err != nil {
			SendError(c, err.Error())
			return
		}
	} else if actionType == 2 {
		if err := response.UnFollow(id, toUserId); err != nil {
			SendError(c, err.Error())
			return
		}
	} else {
		SendError(c, "未识别的action_type")
		return
	}
	c.JSON(http.StatusOK, &response)
}

func FollowList(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
		return
	}
	id := rawId.(int64)
	response := service.ResonseRelationList{}
	if err := response.FollowList(id); err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}

func FollowerList(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
		return
	}
	id := rawId.(int64)
	response := service.ResonseRelationList{}
	if err := response.FollowerList(id); err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}

func FriendList(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
		return
	}
	id := rawId.(int64)
	response := service.ResonseRelationList{}
	if err := response.FriendList(id); err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}
