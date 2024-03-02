package controller

import (
	"douyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Regist(c *gin.Context) {
	userName := c.Query("username")
	passWord := c.Query("password")
	response := service.ResponseRegist{}
	if err := response.Do(userName, passWord); err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}

func Login(c *gin.Context) {
	userName := c.Query("username")
	if userName == "" {
		SendError(c, "用户名不能为空")
		return
	}
	password := c.Query("password")
	response := service.ResponseLogin{}
	err := response.Do(userName, password)
	if err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}

func GetUserInfo(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
		return
	}
	id := rawId.(int64)
	respons := service.ResonseUserInfo{}
	err := respons.QueryById(id)
	if err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &respons)
}
