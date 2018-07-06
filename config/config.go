package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}


func Init() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func Get(name string) string {
	return viper.GetString(name)
}
