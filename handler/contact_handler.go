package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/onyas/newsAPI/service"
	"log"
	"net/http"
	"strconv"
)

var contactService service.ContactService

func ListContacts(context *gin.Context) {
	cid := context.Param("userid")
	userid, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	contacts := contactService.SearchContactsByIds(int64(userid))
	context.JSON(http.StatusOK, gin.H{
		"contacts": contacts,
	})
}
