package main

type MychatApp struct {
	Env        string `mapstructure:"env"`
	AppName    string `mapstructure:"appName"`
	HttpServer struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"httpServer"`
}

func (t MychatApp) Init() {

}