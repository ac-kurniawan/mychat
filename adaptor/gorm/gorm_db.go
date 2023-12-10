package gorm

import (
	"context"

	"github.com/ac-kurniawan/mychat/adaptor/gorm/entity"
	"github.com/ac-kurniawan/mychat/core"
	"github.com/ac-kurniawan/mychat/core/model"
	"github.com/ac-kurniawan/mychat/library"
	gorm2 "gorm.io/gorm"
)

type GormDB struct {
	Gorm  *gorm2.DB
	Trace library.AppTrace
}

// GetParticipantGroupsByParticipant implements core.IChatDB.
func (g *GormDB) GetParticipantGroupsByParticipant(ctx context.Context, participant string) ([]model.ParticipantGroupModel, error) {
	ctx, span := g.Trace.StartTrace(ctx, "DATABASE - GetParticipantGroupsByParticipant")
	defer g.Trace.EndTrace(span)

	var participantGroups []entity.ParticipantGroupEntity
	result := g.Gorm.WithContext(ctx).Where("participant = ?", participant).Find(&participantGroups)
	if result.Error != nil {
		g.Trace.TraceError(span, result.Error)
		return nil, result.Error
	}

	var output []model.ParticipantGroupModel
	for _, val := range participantGroups {
		output = append(output, val.ToModel())
	}

	return output, nil
}

// GetRoomChatByParticipant implements core.IChatDB.
func (g *GormDB) GetRoomChatByParticipant(ctx context.Context, participant string, sessionStatus *string, initialChatNumber uint) ([]model.RoomChatModel, error) {
	panic("unimplemented")
}

// GetRoomChatByParticipantGroups implements core.IChatDB.
func (g *GormDB) GetRoomChatByParticipantGroups(ctx context.Context, participantGroups []string, sessionStatus *string, initialChatNumber uint) ([]model.RoomChatModel, error) {
	ctx, span := g.Trace.StartTrace(ctx, "DATABASE - GetRoomChatByParticipantGroups")
	defer g.Trace.EndTrace(span)

	var roomChats []entity.RoomChatEntity
	tx := g.Gorm.WithContext(ctx).Where("participant_group IN ?", participantGroups)

	if sessionStatus != nil {
		tx.Preload("SessionChats", "session_status = ?", *sessionStatus)
	} else {
		tx.Preload("SessionChats")
	}
	tx.Preload("SessionChats.Chats", func(db *gorm2.DB) *gorm2.DB {
		return db.Order("created_at ASC").Limit(int(initialChatNumber))
	})

	result := tx.Find(&roomChats)
	if result.Error != nil {
		g.Trace.TraceError(span, result.Error)
		return nil, result.Error
	}

	var output []model.RoomChatModel
	for _, val := range roomChats {
		output = append(output, val.ToModel())
	}

	return output, nil
}

// GetRoomChatBySessionId implements core.IChatDB.
func (g *GormDB) GetRoomChatBySessionId(ctx context.Context, sessionid string, sessionStatus *string, initialChatNumber uint) (*model.RoomChatModel, error) {
	ctx, span := g.Trace.StartTrace(ctx, "DATABASE - GetRoomChatBySessionId")
	defer g.Trace.EndTrace(span)

	var sessionChat entity.SessionChatEntity
	tx := g.Gorm.Where("id = ?", sessionid)
	if sessionStatus != nil {
		tx.Where("session_status = ", *sessionStatus)
	}
	result := tx.WithContext(ctx).Find(&sessionChat)
	if result.Error != nil {
		g.Trace.TraceError(span, result.Error)
		return nil, result.Error
	}

	var roomChat entity.RoomChatEntity
	txRoomChat := g.Gorm.Where("id = ?", sessionChat.RoomChatEntityID)
	if sessionStatus != nil {
		txRoomChat.Preload("SessionChats", "session_status = ?", *sessionStatus)
	} else {
		txRoomChat.Preload("SessionChats")
	}
	txRoomChat.Preload("SessionChats.Chats", func(db *gorm2.DB) *gorm2.DB {
		return db.Order("created_at ASC").Limit(int(initialChatNumber))
	})
	resultRoomChat := txRoomChat.WithContext(ctx).Find(&roomChat)
	if resultRoomChat.Error != nil {
		g.Trace.TraceError(span, result.Error)
		return nil, result.Error
	}

	output := roomChat.ToModel()

	return &output, nil
}

func NewGormDB(module GormDB, enableAutoMigration bool) core.IChatDB {
	if enableAutoMigration {
		module.Gorm.AutoMigrate(&entity.ParticipantGroupEntity{}, &entity.RoomChatEntity{}, &entity.SessionChatEntity{}, &entity.ChatEntity{})
		CreateParticipantGrop(module.Gorm)
		CreateRoomChatRecords(module.Gorm)
	}
	return &module
}
