package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onyas/newsAPI/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	// 添加 Get 请求路由
	router.GET("/", handler.IndexPage)

	conversationRouter := router.Group("/chat/conversations")
	{
		conversationRouter.GET("/:userid", handler.ListConversations)
	}

	contactRouter := router.Group("/chat/contacts")
	{
		contactRouter.GET("/:userid", handler.ListContacts)
	}

	return router
}
