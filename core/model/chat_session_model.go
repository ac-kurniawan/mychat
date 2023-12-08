package model

import "time"

type ChatSessionModel struct {
	Id         *string
	RoomChatId string
	Status     string
	CreatedAt  time.Time
	EndedAt    *time.Time
	Chats      []ChatModel
}
