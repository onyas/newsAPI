package service

import (
	"github.com/onyas/newsAPI/model"
	"log"
)

type ContactService struct {
}

func (service *ContactService) SearchContactsByIds(userId int64) []model.Contact {
	contacts := make([]model.Contact, 0)
	err := DbEngin.Where("user_id = ? ", userId).Find(&contacts)
	if nil != err {
		log.Fatal(err)
	}
	return contacts
}
