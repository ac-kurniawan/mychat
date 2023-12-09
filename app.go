package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/ac-kurniawan/mychat/adaptor/gorm"
	"github.com/ac-kurniawan/mychat/adaptor/repository"
	"github.com/ac-kurniawan/mychat/adaptor/util"
	"github.com/ac-kurniawan/mychat/core"
	"github.com/ac-kurniawan/mychat/interface/http"
	"github.com/ac-kurniawan/mychat/library"
)

type MychatApp struct {
	Env        string `mapstructure:"env"`
	AppName    string `mapstructure:"appName"`
	HttpServer struct {
		Port string `mapstructure:"port"`
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

	httpServer := http.NewHttpServer(http.HttpServer{
		Port:    t.HttpServer.Port,
		Service: service,
		Trace:   trace,
	})

	httpServer.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shut down the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	httpServer.Stop(ctx)
}
