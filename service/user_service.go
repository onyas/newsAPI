package service

import (
	"github.com/onyas/newsAPI/model"
	"log"
)

type UserService struct {
}

func (service *UserService) Save(user *model.User) int64 {
	count, error := DbEngin.InsertOne(user)
	if error != nil {
		log.Fatal(error)
	}
	return count
}

func (service *UserService) SelectByUserId(userId int64) model.User {
	user := model.User{}
	_, err := DbEngin.Where("user_id = ? ", userId).Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
