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
	DbPassWord string
	DbName     string

	AccessKey    string
	SecretKey    string
	Bucket       string
	AliyunServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}

	LoadServer(file)
	LoadData(file)
	LoadAliyun(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("abc")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("AppMode").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("Aa123456")
	DbName = file.Section("database").Key("DbName").MustString("baby_blog")
}

func LoadAliyun(file *ini.File) {
	AccessKey = file.Section("servealiyunr").Key("AccessKey").String()
	SecretKey = file.Section("aliyun").Key("SecretKey").String()
	Bucket = file.Section("aliyun").Key("Bucket").String()
	AliyunServer = file.Section("aliyun").Key("AliyunServer").String()
}
