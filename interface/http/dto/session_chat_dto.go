package dto

import (
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
)

type SessionChatDto struct {
	Id         *string    `json:"id"`
	RoomChatId string     `json:"roomChatId"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"createdAt"`
	EndedAt    *time.Time `json:"endedAt"`
	Chats      []ChatDto  `json:"chats,omitempty"`
}

func (s *SessionChatDto) FromModel(input model.ChatSessionModel) {
	s.Id = input.Id
	s.RoomChatId = input.RoomChatId
	s.Status = input.Status
	s.CreatedAt = input.CreatedAt
	s.EndedAt = input.EndedAt

	var chats []ChatDto
	if input.Chats != nil {
		for _, val := range input.Chats {
			var dto ChatDto
			dto.FromModel(val)
			chats = append(chats, dto)
		}
	}
	s.Chats = chats
}

func (s *SessionChatDto) ToModel() model.ChatSessionModel {
	var chats []model.ChatModel
	if s.Chats != nil {
		for _, val := range s.Chats {
			chats = append(chats, val.ToModel())
		}
	}
	return model.ChatSessionModel{
		Id:         s.Id,
		RoomChatId: s.RoomChatId,
		Status:     s.Status,
		CreatedAt:  s.CreatedAt,
		EndedAt:    s.EndedAt,
		Chats:      chats,
	}
}
