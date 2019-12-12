package service

import (
	"github.com/onyas/newsAPI/model"
	"log"
)

type ConversationService struct {
}

func (service *ConversationService) SearchConverByIds(userId int64) []model.Conversation {
	conversations := make([]model.Conversation, 0)
	err := DbEngin.Where("user_id = ? ", userId).Find(&conversations)
	if nil != err {
		log.Fatal(err)
	}
	return conversations
}
