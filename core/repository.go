package core

import (
	"context"

	"github.com/ac-kurniawan/mychat/core/model"
)

type IChatDB interface {
	GetParticipantGroupsByParticipant(ctx context.Context, participant string) ([]model.ParticipantGroupModel, error)
	GetRoomChatByParticipantGroups(ctx context.Context, participantGroups []string, sessionStatus *string, initialChatNumber uint) ([]model.RoomChatModel, error)
	GetRoomChatByParticipant(ctx context.Context, participant string, sessionStatus *string, initialChatNumber uint) ([]model.RoomChatModel, error)
	GetRoomChatBySessionId(ctx context.Context, sessionid string, sessionStatus *string, initialChatNumber uint) (*model.RoomChatModel, error)
}

type IRepository interface {
	IChatDB
}
