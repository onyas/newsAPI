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

type Conversation struct {
	Id          int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	UserId      int64     `xorm:"bigint(20)" form:"userid" json:"userid"`
	Avatar      string    `xorm:"varchar(100)" form:"avatar" json:"avatar"`
	FromId      int64     `xorm:"bigint(20)" form:"fromid" json:"fromid"`
	Name        string    `xorm:"varchar(40)" form:"name" json:"name"`
	Message     string    `xorm:"varchar(150)" form:"message" json:"message"`
	MessageType string    `xorm:"varchar(10)" form:"messagetype" json:"messagetype"`
	UpdatedAt   time.Time `xorm:"updated" form:"updateat" json:"updateat"`
}

type User struct {
	Id        int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id,omitempty"`
	UserId    int64     `xorm:"bigint(20)" form:"userid" json:"userid"`
	Name      string    `xorm:"varchar(40)" form:"name" json:"name"`
	Avatar    string    `xorm:"varchar(100)" form:"avatar" json:"avatar,omitempty"`
	Email     string    `xorm:"varchar(50)" form:"email" json:"email,omitempty"`
	Company   string    `xorm:"varchar(50)" form:"company" json:"company,omitempty"`
	Followers int       `xorm:"int(20)" form:"followers" json:"followers,omitempty"`
	CreatedAt time.Time `xorm:"created" form:"created" json:"created"`
}

type Result struct {
	Code    int         `form:"code" json:"code"`
	Message string      ` form:"message" json:"message"`
	Data    interface{} `json:"data" `
}
