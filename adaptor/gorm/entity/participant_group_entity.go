package entity

import "github.com/ac-kurniawan/mychat/core/model"

type ParticipantGroupEntity struct {
	Id               *string `gorm:"primaryKey,default:uuid_generate_v4()"`
	ParticipantGroup string  `gorm:"index"`
	Participant      string  `gorm:"index"`
}

func (p *ParticipantGroupEntity) ToModel() model.ParticipantGroupModel {
	return model.ParticipantGroupModel{
		Id:               p.Id,
		ParticipantGroup: p.ParticipantGroup,
		Participant:      p.Participant,
	}
}
