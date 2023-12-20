package cmd

import (
	"fmt"
	"os"

	"github.com/gin-blog/faye-wong/configs"
	"github.com/gin-blog/faye-wong/global"
	"github.com/gin-blog/faye-wong/internal/api"
	"github.com/gin-blog/faye-wong/internal/server"
	"github.com/gin-blog/faye-wong/internal/service"
	"github.com/gin-blog/faye-wong/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "faye",
	Short: "faye",
	Long:  `faye is the English name of singer WangFei'`,
	Run: func(cmd *cobra.Command, args []string) {
		configs.AutomaticEnv()
		db, err := model.NewDB()
		if err != nil {
			fmt.Println("创建数据库连接失败:", err.Error())
			return
		}
		global.DBEngine = db
		if configs.Debug {
			fmt.Println("当前的日志级别是debug")
		}
		if err := start(); err != nil {
			fmt.Println("服务启动失败")
			return
		}
		fmt.Println("服务启动成功...")
	},
}

func start() error {
	user := service.NewUser(global.DBEngine)
	apiUser := api.NewUser(user)
	apiArticles := service.NewArticles(global.DBEngine)
	articles := api.NewArticles(apiArticles)
	apIs := api.NewAPIS(apiUser, articles)

	newServer := server.NewServer(apIs)
	newServer.StartServe()
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
	os.Exit(1)
}
