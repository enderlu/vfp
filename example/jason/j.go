package main

import (
	"encoding/json"
	"fmt"
	. "github.com/enderlu/vfp"
	"html/template"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	Init()
}

type ConfigStruct struct {
	Host              string `json:"host"`
	Port              int    `json:"port"`
	DB                string `json:"db"`
	CookieSecret      string `json:"cookie_secret"`
	SmtpUsername      string `json:"smtp_username"`
	SmtpPassword      string `json:"smtp_password"`
	SmtpHost          string `json:"smtp_host"`
	SmtpAddr          string `json:"smtp_addr"`
	FromEmail         string `json:"from_email"`
	Superusers        string `json:"superusers"`
	TimeZoneOffset    int64  `json:"time_zone_offset"`
	AnalyticsFile     string `json:"analytics_file"`
	StaticFileVersion int    `json:"static_file_version"`
	QiniuAccessKey    string `json:"qiniu_access_key"`
	QiniuSecretKey    string `json:"qiniu_secret_key"`
}

var Config ConfigStruct
var analyticsCode template.HTML // 网站统计分析代码
var goVersion = runtime.Version()

func Init() {
	file, err := os.Open("./config.json.default")
	if err != nil {
		fmt.Println("配置文件读取失败:", err.Error())
		os.Exit(1)
	}

	fmt.Println(Upper("tt"))
	defer file.Close()

	dec := json.NewDecoder(file)

	err = dec.Decode(&Config)
	fmt.Println("jason数据:", Config.Host)
	if err != nil {
		fmt.Println("配置文件解析失败:", err.Error())
		os.Exit(1)
	}

	if Config.AnalyticsFile != "" {
		content, err := ioutil.ReadFile(Config.AnalyticsFile)

		if err != nil {
			fmt.Println("统计分析文件没有找到:", err.Error())
			os.Exit(1)
		}

		analyticsCode = template.HTML(string(content))
	}
}
