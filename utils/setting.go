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
	DbName     string

	AccessKey   string
	Secretkey   string
	Bucket      string
	QiNiuServer string

	ReHost string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误．请检查文件路径：", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiNiu(file)
	LoadRedist(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3001")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")

}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbHost = file.Section("database").Key("DbHost").MustString("localost")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("hyq5201314")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func LoadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	Secretkey = file.Section("qiniu").Key("Secretkey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiNiuServer = file.Section("qiniu").Key("QiNiuServer").String()
}
func LoadRedist(file *ini.File) {
	ReHost = file.Section("redis").Key("ReHost").String()
}
