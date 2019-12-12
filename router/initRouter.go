package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onyas/newsAPI/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	// 添加 Get 请求路由
	router.GET("/", handler.IndexPage)
	router.GET("/chat/conversations/:userid", handler.ListConversations)

	return router
}
