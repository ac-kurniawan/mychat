package main

import (
	"context"
	"fmt"

	"github.com/ac-kurniawan/mychat/adaptor/gorm"
	"github.com/ac-kurniawan/mychat/adaptor/repository"
	"github.com/ac-kurniawan/mychat/adaptor/util"
	"github.com/ac-kurniawan/mychat/core"
	"github.com/ac-kurniawan/mychat/library"
)

type MychatApp struct {
	Env        string `mapstructure:"env"`
	AppName    string `mapstructure:"appName"`
	HttpServer struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"httpServer"`
}

func (t MychatApp) Init() {
	trace := library.NewAppTrace(context.Background(), false, "", "", "", "", "")
	log := library.NewAppLog(false)
	utilCore := util.NewUtil(util.Util{
		AppTrace: trace,
		AppLog:   log,
	})
	sqlite := library.NewGormSqliteConnection("test.db")
	DB := gorm.NewGormDB(gorm.GormDB{
		Gorm:  sqlite,
		Trace: trace,
	}, true)
	repository := repository.NewRepository(repository.Repository{
		IChatDB: DB,
	})
	service := core.NewMychatService(core.MychatService{
		Util:       utilCore,
		Repository: repository,
	})
	result, err := service.GetRoomChatBySessionId(context.Background(), "fe48e2af-b703-4582-af66-a01fe5c530c1", nil)
	if err != nil {
		log.LogError(context.Background(), err)
		return
	}
	log.LogInfo(context.Background(), fmt.Sprintf("%v", *result))
}
