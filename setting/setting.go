package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig App
type AppConfig struct {
	Release      bool   `ini:"release"`
	Port         int    `ini:"port"`
	WatchDb      string `ini:"watch_db"`
	*MySQLConfig `ini:"mysql"`
}

// MySQLConfig Mysql
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
