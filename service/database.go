package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/onyas/newsAPI/model"
	"log"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "postgres"
	DsName := "postgres://postgres:root@127.0.0.1:5432/chat?sslmode=disable"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, DsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(false)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(2)

	//自动Sync
	err = DbEngin.Sync2(new(model.Conversation), new(model.Contact), new(model.User))
	if nil != err {
		log.Fatal(err.Error())
	}
	//DbEngin = dbengin
	fmt.Println("init data base ok")
}
