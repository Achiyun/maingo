package routers

import (
	"maingo/controllers/douyin"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	// public directory is used to serve static resources
	// r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	// basic apis
	apiRouter.GET("/feed/", douyin.Feed)

	apiRouter.GET("/user/", douyin.UserInfo)
	apiRouter.POST("/user/register/", douyin.Register)
	// apiRouter.GET("/user/login/", douyin.Login)
	apiRouter.POST("/user/login/", douyin.UserController{}.DoLogin)
	apiRouter.POST("/publish/action/", douyin.Publish)
	apiRouter.GET("/publish/list/", douyin.PublishList)
	// apiRouter.GET("/dologin", douyin.UserController{}.DoLogin)
	// extra apis - I
	apiRouter.POST("/favorite/action/", douyin.FavoriteAction)
	apiRouter.GET("/favorite/list/", douyin.FavoriteList)
	apiRouter.POST("/comment/action/", douyin.CommentAction)
	apiRouter.GET("/comment/list/", douyin.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", douyin.RelationAction)
	apiRouter.GET("/relation/follow/list/", douyin.FollowList)
	apiRouter.GET("/relation/follower/list/", douyin.FollowerList)

}
