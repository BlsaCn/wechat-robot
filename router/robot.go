package router

import (
	"github.com/gin-gonic/gin"
	"wechat-robot/controller"
)

func Router() (router *gin.Engine) {
	router = gin.Default()
	// 回调地址
	router.POST("/api/robot/notice", controller.Notice)
	// 手动触发测试地址
	router.GET("/api/robot/test", controller.Test)

	return
}
