package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Charset  string `json:"charset"`
}

func NewMysqlConfig() *MysqlConfig {

	viper.SetConfigFile("./config/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file : %s \n", err))
	}

	fmt.Println("IP = ", viper.GetString("mysql.host"))

	return &MysqlConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
		Charset:  viper.GetString("mysql.charset"),
	}
}
