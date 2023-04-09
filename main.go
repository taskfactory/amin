package main

import (
	"context"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/taskfactory/admin/repo/mysql"
	"github.com/taskfactory/admin/services/admin"

	proto "github.com/taskfactory/admin/tars-protocol/admin"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	if err := InitDB(); err != nil {
		panic(fmt.Sprintf("failed to init db, err:%v", err))
	}
	// New servant imp
	serv := admin.NewServ()
	err := serv.Init(context.Background())
	if err != nil {
		panic(fmt.Sprintf("TaskServImp init fail, err:(%s)\n", err))
	}
	// New servant
	app := new(proto.AdminService)
	// Register Servant
	app.AddServantWithContext(serv, cfg.App+"."+cfg.Server+".admin")

	// Run application
	tars.Run()
}

// InitDB 初始化数据库配置
func InitDB() error {
	servCfg := tars.GetServerConfig()
	remoteConf := tars.NewRConf(servCfg.App, servCfg.Server, servCfg.BasePath)
	cfgStr, err := remoteConf.GetConfig("mysql.yaml")
	if err != nil {
		return err
	}
	return mysql.Init(cfgStr)
}
