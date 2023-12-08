package entity

import (
	"database/sql"
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
)

type ChatEntity struct {
	Id                  *string `gorm:"primaryKey,default:uuid_generate_v4()"`
	SessionChatEntityID string  `gorm:"index"`
	MessageType         string
	Message             string
	ReplyForChatId      sql.NullString
	CreatedAt           time.Time
	ReadAt              sql.NullTime
}

func (c *ChatEntity) FromModel(input model.ChatModel) {
	c.Id = input.Id
	c.SessionChatEntityID = input.SessionId
	c.MessageType = input.MessageType
	c.Message = input.Message
	c.CreatedAt = input.CreatedAt
	if input.ReplyForChatId != nil {
		c.ReplyForChatId = sql.NullString{
			Valid: false,
		}
	} else {
		c.ReplyForChatId = sql.NullString{
			Valid:  true,
			String: *input.ReplyForChatId,
		}
	}
	if input.ReadAt != nil {
		c.ReadAt = sql.NullTime{
			Valid: false,
		}
	} else {
		c.ReadAt = sql.NullTime{
			Valid: true,
			Time:  *input.ReadAt,
		}
	}
}

func (c *ChatEntity) ToModel() model.ChatModel {
	return model.ChatModel{
		Id:             c.Id,
		SessionId:      c.SessionChatEntityID,
		MessageType:    c.MessageType,
		Message:        c.Message,
		ReplyForChatId: &c.ReplyForChatId.String,
		CreatedAt:      c.CreatedAt,
		ReadAt:         &c.ReadAt.Time,
	}
}
