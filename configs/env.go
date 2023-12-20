package configs

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	RootPath   = filepath.Join(filepath.Dir(b), "../")
	MysqlDSN   string
	ListenAddr string
	Debug      bool
)

func AutomaticEnv() {
	fmt.Println("RootPath:", RootPath)
	if err := godotenv.Load(filepath.Join(RootPath, ".env")); err != nil {
		fmt.Println("加载环境变量失败:", err.Error())
		return
	}
	viper.AutomaticEnv()
	Debug = viper.GetBool("DEBUG")
	MysqlDSN = viper.GetString("MYSQL_DSN")
	ListenAddr = viper.GetString("LISTEN_ADDR")
}
