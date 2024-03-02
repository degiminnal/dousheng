package controller

import (
	"douyin/midware"
	"douyin/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func VideoFeed(c *gin.Context) {
	var id int64
	var err error
	var latestTime int64
	tokenStr := c.Query("token")
	if tokenStr != "" {
		_, claims, err := midware.ParseToken(tokenStr)
		if err != nil {
			id = 0
		} else {
			id = claims.UserId
		}
	} else {
		id = 0
	}
	rawTime, ok := c.GetQuery("latest_time")
	if ok && rawTime != "" {
		latestTime, err = strconv.ParseInt(rawTime, 10, 64)
		if err != nil {
			SendError(c, err.Error())
			return
		}
	} else {
		latestTime = int64(time.Now().Unix())
	}
	response := service.ResponseFeed{}
	if err := response.Do(latestTime, id); err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}
