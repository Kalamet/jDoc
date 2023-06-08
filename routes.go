package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kalamet/jdoc/controller"
	"github.com/kalamet/jdoc/middleware"
)

func Routes(r *gin.Engine) *gin.Engine {
	//注册
	//r.Use(middleware.AuthMiddleware())
	r.POST("register", controller.Register)
	r.POST("login", controller.Login)
	r.GET("api/user/info", middleware.AuthMiddleware(), controller.UserInfo)
	return r
}
