package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageChat(c *gin.Context) {
	c.JSON(http.StatusOK, &ResponseCommon{0, "成功"})
}
