package config

import (
	"os"

	"github.com/spf13/viper"
)

func Init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("configure")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
