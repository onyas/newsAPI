package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/onyas/newsAPI/service"
	"log"
	"net/http"
	"strconv"
)

var conversationService service.ConversationService

func ListConversations(context *gin.Context) {
	cid := context.Param("userid")
	userid, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	conversations := conversationService.SearchConverByIds(int64(userid))
	context.JSON(http.StatusOK, gin.H{
		"conversations": conversations,
	})
}
