package utils

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")

	err := viper.ReadInConfig()

	if err != nil {
		panic("")
	}

	fmt.Println("配置文件 server:", viper.Get("server.port"))
	fmt.Println("配置文件 mysql:", viper.Get("mysql"))
}
