package entity

import (
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
)

type ChatEntity struct {
	Id                  *string `gorm:"primaryKey,default:uuid_generate_v4()"`
	SessionChatEntityID string  `gorm:"index"`
	MessageType         string
	Message             string
	SenderId            string
	ReplyForChatId      *string
	CreatedAt           time.Time
	ReadAt              *time.Time
}

func (c *ChatEntity) FromModel(input model.ChatModel) {
	c.Id = input.Id
	c.SessionChatEntityID = input.SessionId
	c.MessageType = input.MessageType
	c.Message = input.Message
	c.CreatedAt = input.CreatedAt
	c.ReadAt = input.ReadAt
	c.SenderId = input.SenderId
	c.ReplyForChatId = input.ReplyForChatId
}

func (c *ChatEntity) ToModel() model.ChatModel {
	return model.ChatModel{
		Id:             c.Id,
		SessionId:      c.SessionChatEntityID,
		MessageType:    c.MessageType,
		Message:        c.Message,
		SenderId:       c.SenderId,
		ReplyForChatId: c.ReplyForChatId,
		CreatedAt:      c.CreatedAt,
		ReadAt:         c.ReadAt,
	}
}
