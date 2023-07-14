package main

import (
	"fmt"

	"github.com/gin-blog/faye-wong/configs"
	"github.com/gin-blog/faye-wong/global"
	"github.com/gin-blog/faye-wong/internal/api"
	"github.com/gin-blog/faye-wong/internal/server"
	"github.com/gin-blog/faye-wong/internal/service"
	"github.com/gin-blog/faye-wong/model"
)

func main() {
	// 1.加载配置文件
	configs.AutomaticEnv()
	// 2.加载数据库连接
	db, err := model.NewDB()
	if err != nil {
		fmt.Println("创建数据库连接失败:", err.Error())
		return
	}
	global.DBEngine = db
	if err := start(); err != nil {
		fmt.Println("服务启动失败")
		return
	}
}

func start() error {
	user := service.NewUser(global.DBEngine)
	apiUser := api.NewUser(user)
	apIs := api.NewAPIS(apiUser)

	newServer := server.NewServer(apIs)
	newServer.StartServe()
	fmt.Println("服务启动成功...")
	return nil
}
