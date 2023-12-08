package dto

type GetRoomChatBySessionIdDto struct {
	SessionId     string  `query:"session_id"`
	SessionStatus *string `query:"session_status"`
}
