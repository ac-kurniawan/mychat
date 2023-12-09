package dto

import (
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
)

type ChatDto struct {
	Id             *string    `json:"id"`
	SessionId      string     `json:"sessionId"`
	MessageType    string     `json:"messageType"`
	Message        string     `json:"message"`
	ReplyForChatId *string    `json:"replyforChatId"`
	CreatedAt      time.Time  `json:"createdAt"`
	ReadAt         *time.Time `json:"readAt"`
}

func (c *ChatDto) FromModel(input model.ChatModel) {
	c.Id = input.Id
	c.SessionId = input.SessionId
	c.MessageType = input.MessageType
	c.Message = input.Message
	c.CreatedAt = input.CreatedAt
	c.ReplyForChatId = input.ReplyForChatId
	c.ReadAt = input.ReadAt
}

func (c *ChatDto) ToModel() model.ChatModel {
	return model.ChatModel{
		Id:             c.Id,
		SessionId:      c.SessionId,
		MessageType:    c.MessageType,
		Message:        c.Message,
		ReplyForChatId: c.ReplyForChatId,
		CreatedAt:      c.CreatedAt,
		ReadAt:         c.ReadAt,
	}
}
