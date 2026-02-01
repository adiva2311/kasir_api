package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ViperConfigStruct struct {
	APIHost      string
	APIPort      string
	UsernameDB   string
	PasswordDB   string
	HostDB       string
	PortDB       string
	DatabaseName string
}

func ViperConfig() *ViperConfigStruct {
	// VIPER CONFIG
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	// READ CONFIG
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	appConfig := &ViperConfigStruct{
		APIHost:      viper.GetString("API_HOST"),
		APIPort:      viper.GetString("API_PORT"),
		UsernameDB:   viper.GetString("USERNAME_DB"),
		PasswordDB:   viper.GetString("PASSWORD_DB"),
		HostDB:       viper.GetString("HOST_DB"),
		PortDB:       viper.GetString("PORT_DB"),
		DatabaseName: viper.GetString("DATABASE_NAME"),
	}

	return appConfig
}
