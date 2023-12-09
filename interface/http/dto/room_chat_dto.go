package dto

import (
	"time"

	"github.com/ac-kurniawan/mychat/core/model"
)

type RoomChatDto struct {
	Id               *string          `json:"id"`
	ParticipantGroup string           `json:"participantGroup"`
	CreatedAt        time.Time        `json:"createdAt"`
	SessionChats     []SessionChatDto `json:"sessionChats,omitempty"`
}

func (r *RoomChatDto) FromModel(input model.RoomChatModel) {
	r.Id = input.Id
	r.ParticipantGroup = input.ParticipantGroup
	r.CreatedAt = input.CreatedAt

	var sessionChats []SessionChatDto
	if input.SessionChats != nil {
		for _, val := range input.SessionChats {
			var dto SessionChatDto
			dto.FromModel(val)
			sessionChats = append(sessionChats, dto)
		}
	}
	r.SessionChats = sessionChats
}

func (r *RoomChatDto) ToModel() model.RoomChatModel {
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
