package controller

import (
	"douyin/service"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
	}
	id := rawId.(int64)
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		SendError(c, err.Error())
	}

	filename := filepath.Base(data.Filename)
	videoPath := fmt.Sprintf("./public/videos/%d_%s", id, filename)
	if err := c.SaveUploadedFile(data, videoPath); err != nil {
		SendError(c, err.Error())
	}

	response := service.ResponseVideoPublish{}
	err = response.Do(id, title, filename)
	if err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &ResponseCommon{0, "发布成功"})
}

func PublishList(c *gin.Context) {
	rawId, ok := c.Get("userId")
	if !ok {
		SendError(c, "解析token失败")
		return
	}
	id := rawId.(int64)
	response := service.ResponsePublishList{}
	err := response.QueryByUserId(id)
	if err != nil {
		SendError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, &response)
}
