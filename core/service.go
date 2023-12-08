package core

import (
	"context"

	"github.com/ac-kurniawan/mychat/core/model"
)

type IMychatService interface {
	GetRoomChatsByParticipant(ctx context.Context, particitpant string, sessionStatus *string) ([]model.RoomChatModel, error)
	GetRoomChatBySessionId(ctx context.Context, sessionid string, sessionStatus *string) (*model.RoomChatModel, error)
}

type MychatService struct {
	Repository IRepository
	Util       IUtil
}

func (m *MychatService) validateSessionStatus(sessionStatus *string) error {
	if sessionStatus != nil {
		isExist := false
		allowedSessionStatus := []string{ChatSessionStatusActive, ChatSessionStatusEnded}
		for _, val := range allowedSessionStatus {
			if *sessionStatus == val {
				isExist = true
				break
			}
		}

		if !isExist {
			return ErrSessionStatusInvalid
		}
	}
	return nil
}

// GetRoomChatByParticipant get room list which include given participant.
func (m *MychatService) GetRoomChatsByParticipant(ctx context.Context, participant string, sessionStatus *string) ([]model.RoomChatModel, error) {
	ctx, span := m.Util.StartTrace(ctx, "SERVICE - GetRoomChatByParticipant")
	defer m.Util.EndTrace(span)
	err := m.validateSessionStatus(sessionStatus)
	if err != nil {
		m.Util.LogError(ctx, err)
		m.Util.TraceError(span, err)
		return nil, err
	}
	participantGroups, err := m.Repository.GetParticipantGroupsByParticipant(ctx, participant)
	if err != nil {
		m.Util.LogError(ctx, err)
		m.Util.TraceError(span, err)
		return nil, ErrFailedGetParticipantGroup
	}
	var participantGroupList []string
	for _, val := range participantGroups {
		participantGroupList = append(participantGroupList, val.ParticipantGroup)
	}
	listRoomChat, err := m.Repository.GetRoomChatByParticipantGroups(ctx, participantGroupList, sessionStatus, 10)
	if err != nil {
		m.Util.LogError(ctx, err)
		m.Util.TraceError(span, err)
		return nil, ErrFailedGetRoomChat
	}
	return listRoomChat, nil
}

// GetRoomChatBySessionId get room chat by session id.
func (m *MychatService) GetRoomChatBySessionId(ctx context.Context, sessionid string, sessionStatus *string) (*model.RoomChatModel, error) {
	ctx, span := m.Util.StartTrace(ctx, "SERVICE - GetRoomChatBySessionId")
	defer m.Util.EndTrace(span)
	err := m.validateSessionStatus(sessionStatus)
	if err != nil {
		m.Util.LogError(ctx, err)
		m.Util.TraceError(span, err)
		return nil, err
	}
	roomChat, err := m.Repository.GetRoomChatBySessionId(ctx, sessionid, sessionStatus, 10)
	if err != nil {
		m.Util.LogError(ctx, err)
		m.Util.TraceError(span, err)
		return nil, ErrFailedGetRoomChat
	}
	return roomChat, nil
}

func NewMychatService(module MychatService) IMychatService {
	return &module
}
