package controller

import (
	"douyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCommon service.ResponseCommon

func SendError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, &ResponseCommon{1, msg})
}
