package models

import (
	"gorm.io/gorm"
)

type MessageType string

const (
	UserType   MessageType = "user"
	BotType    MessageType = "bot"
	SystemType MessageType = "system"
)

type Message struct {
	gorm.Model
	SessionID   uint        `gorm:"column:session_id"`
	UserID      uint        `gorm:"column:user_id"`
	IPAddress   string      `gorm:"column:ip_address"`
	MessageText string      `gorm:"column:message_text"`
	Response    string      `gorm:"column:response"`
	IsResponded bool        `gorm:"column:is_responded"`
	MessageType MessageType `gorm:"column:message_type"`
}

func (Message) TableName() string {
	return "messages"
}
