package model

import "time"

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
