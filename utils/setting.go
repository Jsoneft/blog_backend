package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	Dbname     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("zpBpgBdkg")

}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("cdb-gbwa3hlu.cd.tencentcdb.com")
	DbPort = file.Section("database").Key("DbPort").MustString("10086")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassword = file.Section("database").Key("DbPassword").MustString("zjxstcstc")
	Dbname = file.Section("database").Key("Dbname").MustString("blogs")

}
