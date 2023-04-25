package dao

import (
	"face/config"
	"fmt"

	"github.com/spf13/viper"
)

var MySQLConfig *config.MySQLConfig = &config.MySQLConfig{}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	dev := GetEnvInfo("DEV")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-dev.yaml", configFilePrefix)
	if !dev {
		configFileName = fmt.Sprintf("%s-pro.yaml", configFilePrefix)
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(MySQLConfig); err != nil {
		panic(err)
	}
}
