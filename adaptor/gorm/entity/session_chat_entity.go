package entity

import (
	"database/sql"
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
)

type SessionChatEntity struct {
	Id         *string `gorm:"primaryKey,default:uuid_generate_v4()"`
	RoomChatID string  `gorm:"index"`
	Status     string
	CreatedAt  time.Time
	EndedAt    sql.NullTime
	Chats      []ChatEntity
}

func (s *SessionChatEntity) FromModel(input model.ChatSessionModel) {
	s.Id = input.Id
	s.RoomChatID = input.RoomChatId
	s.Status = input.Status
	s.CreatedAt = input.CreatedAt
	if input.EndedAt != nil {
		s.EndedAt = sql.NullTime{
			Valid: false,
		}
	} else {
		s.EndedAt = sql.NullTime{
			Valid: true,
			Time:  *input.EndedAt,
		}
	}
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
		RoomChatId: s.RoomChatID,
		Status:     s.Status,
		CreatedAt:  s.CreatedAt,
		EndedAt:    &s.EndedAt.Time,
		Chats:      chats,
	}
}