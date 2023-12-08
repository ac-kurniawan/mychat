package repository

import "github.com/ac-kurniawan/mychat/core"

type Repository struct {
	core.IChatDB
}

func NewRepository(module Repository) core.IRepository {
	return &module
}
