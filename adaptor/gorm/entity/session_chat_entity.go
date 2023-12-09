package entity

import (
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
)

type SessionChatEntity struct {
	Id               *string `gorm:"primaryKey,default:uuid_generate_v4()"`
	RoomChatEntityID string  `gorm:"index"`
	Status           string
	CreatedAt        time.Time
	EndedAt          *time.Time
	Chats            []ChatEntity
}

func (s *SessionChatEntity) FromModel(input model.ChatSessionModel) {
	s.Id = input.Id
	s.RoomChatEntityID = input.RoomChatId
	s.Status = input.Status
	s.CreatedAt = input.CreatedAt
	s.EndedAt = input.EndedAt
}

func (s *SessionChatEntity) ToModel() model.ChatSessionModel {
	var chats []model.ChatModel
	if s.Chats != nil {
		for _, val := range s.Chats {
			chats = append(chats, val.ToModel())
		}
	}
	return model.ChatSessionModel{
		Id:         s.Id,
		RoomChatId: s.RoomChatEntityID,
		Status:     s.Status,
		CreatedAt:  s.CreatedAt,
		EndedAt:    s.EndedAt,
		Chats:      chats,
	}
}
