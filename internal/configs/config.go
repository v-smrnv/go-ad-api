package configs

import (
	"github.com/spf13/viper"
)

var Cfg *viper.Viper

func InitConfig(path string) {
	Cfg = viper.New()
	Cfg.AddConfigPath(path)
	Cfg.SetConfigName("config")
	Cfg.SetConfigType("yaml")

	if err := Cfg.ReadInConfig(); err != nil {
		panic("ошибка при чтении конфигурации: " + err.Error())
	}
}
