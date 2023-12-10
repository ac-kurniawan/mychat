package dto

type GetRoomChatByParticipantIdDto struct {
	Participant   string  `query:"participant"`
	SessionStatus *string `query:"session_status"`
}
