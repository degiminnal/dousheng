package main

import (
	"douyin/config"
	"douyin/controller"
	"douyin/entity"
	"douyin/midware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	if err := entity.Init(); err != nil {
		fmt.Println("entity init error")
		return
	}
	r := gin.Default()
	r.Static("/static/", "./public")
	apiRouter := r.Group("/douyin/")

	apiRouter.POST("/user/register/", controller.Regist)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.GET("/user/", midware.TokenParse(), controller.GetUserInfo)

	apiRouter.POST("/publish/action/", midware.TokenParse(), controller.PublishAction)
	apiRouter.GET("/publish/list/", midware.TokenParse(), controller.PublishList)

	apiRouter.GET("/feed/", controller.VideoFeed)

	apiRouter.GET("/favorite/list/", midware.TokenParse(), controller.FavoriteList)
	apiRouter.POST("/favorite/action/", midware.TokenParse(), controller.FavoriteAction)

	apiRouter.POST("/comment/action/", midware.TokenParse(), controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	apiRouter.POST("/relation/action/", midware.TokenParse(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", midware.TokenParse(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", midware.TokenParse(), controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", midware.TokenParse(), controller.FriendList)

	apiRouter.GET("/message/chat/", midware.TokenParse(), controller.MessageChat)

	r.Run(":8080")
}
