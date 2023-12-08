package entity

import (
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
)

type RoomChatEntity struct {
	Id               *string `gorm:"primaryKey,default:uuid_generate_v4()"`
	ParticipantGroup string  `gorm:"index"`
	CreatedAt        time.Time
	SessionChats     []SessionChatEntity
}

func (r *RoomChatEntity) FromModel(input model.RoomChatModel) {
	r.Id = input.Id
	r.ParticipantGroup = input.ParticipantGroup
	r.CreatedAt = input.CreatedAt
}

func (r *RoomChatEntity) ToModel() model.RoomChatModel {
	var sessionChats []model.ChatSessionModel
	if r.SessionChats != nil {
		for _, val := range r.SessionChats {
			sessionChats = append(sessionChats, val.ToModel())
		}
	}
	return model.RoomChatModel{
		Id:               r.Id,
		ParticipantGroup: r.ParticipantGroup,
		CreatedAt:        r.CreatedAt,
		SessionChats:     sessionChats,
	}
}
