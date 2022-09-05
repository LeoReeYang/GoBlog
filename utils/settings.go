package utils

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DB          string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassWords string
	DBName      string
)

func init() {
	file, err := ini.Load("config/config.ini")

	if err != nil {
		log.Fatalf("can't open config file,go and check err: %s", err)
	}

	LoadServer(file)

	LoadDataBase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
}

func LoadDataBase(file *ini.File) {
	DB = file.Section("database").Key("DB").MustString("mysql")
	DBHost = file.Section("database").Key("DBHost").MustString("localhost")
	DBPort = file.Section("database").Key("DBPort").MustString("33060")
	DBUser = file.Section("database").Key("DBUser").MustString("root")
	DBPassWords = file.Section("database").Key("DBPassWords").MustString("123")
	DBName = file.Section("database").Key("DBName").MustString("test")
}
