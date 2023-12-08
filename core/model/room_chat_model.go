package model

import "time"

type RoomChatModel struct {
	Id               *string
	ParticipantGroup string
	CreatedAt        time.Time
	SessionChats     []ChatSessionModel
}
