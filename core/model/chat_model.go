package model

import "time"

type ChatModel struct {
	Id             *string
	SenderId       string
	SessionId      string
	MessageType    string
	Message        string
	ReplyForChatId *string
	CreatedAt      time.Time
	ReadAt         *time.Time
}
