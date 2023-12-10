package gorm

import (
	"time"

	"github.com/ac-kurniawan/mychat/adaptor/gorm/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var participant1 = "6282000000001"
var participant2 = "6282000000002"

func CreateParticipantGrop(g *gorm.DB) error {
	id1 := uuid.NewString()
	id2 := uuid.NewString()
	entities := []entity.ParticipantGroupEntity{
		{
			Id:               &id1,
			ParticipantGroup: "GROUP1",
			Participant:      participant1,
		},
		{
			Id:               &id2,
			ParticipantGroup: "GROUP1",
			Participant:      participant2,
		},
	}

	result := g.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(entities, 2)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateRoomChatRecords(g *gorm.DB) error {
	roomChatId := uuid.New().String()
	sessionChatId := "7418a734-f232-43ab-b53e-c491836fde8d"

	var chatIds []string

	for i := 0; i < 5; i++ {
		id := uuid.New().String()
		chatIds = append(chatIds, id)
	}

	var chats []entity.ChatEntity
	for index, val := range chatIds {
		var participant string
		if index%2 == 0 {
			participant = participant2
		} else {
			participant = participant1
		}
		newId := val
		chat := entity.ChatEntity{
			Id:                  &newId,
			SessionChatEntityID: sessionChatId,
			MessageType:         "TEXT",
			Message:             "Hi",
			SenderId:            participant,
			ReplyForChatId:      nil,
			CreatedAt:           time.Now(),
		}
		chats = append(chats, chat)
	}

	entity := entity.RoomChatEntity{
		Id:               &roomChatId,
		ParticipantGroup: "GROUP1",
		CreatedAt:        time.Now(),
		SessionChats: []entity.SessionChatEntity{
			{
				Id:               &sessionChatId,
				RoomChatEntityID: roomChatId,
				Status:           "ACTIVE",
				CreatedAt:        time.Now(),
			},
		},
	}

	result := g.Clauses(clause.OnConflict{DoNothing: true}).Create(&entity)

	if result.Error != nil {
		return result.Error
	}

	result = g.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(chats, len(chats))
	if result.Error != nil {
		return result.Error
	}

	return nil
}
