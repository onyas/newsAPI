package model

import "time"

type Contact struct {
	Id         int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	UserId     int64     `xorm:"bigint(20)" form:"userid" json:"userid"`
	ContactId  int64     `xorm:"bigint(20)" form:"contactid" json:"contactid"`
	Name       string    `xorm:"varchar(40)" form:"name" json:"name"`
	Avatar     string    `xorm:"varchar(100)" form:"avatar" json:"avatar"`
	SectionKey string    `xorm:"varchar(1)" form:"sectionkey" json:"sectionkey"`
	CreatedAt  time.Time `xorm:"created" form:"created" json:"created"`
}
