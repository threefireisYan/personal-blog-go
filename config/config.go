package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	//	定义初始化方法
	Cfg = new(tomlConfig)

	Cfg.System.AppName = "personal-blog-go"
	Cfg.System.Version = 1.0
	currentdir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentdir
	/*
		//这里不能直接传入文件路径名，需要先读取config文件的内容，然后将内容进行传入
		openFile, err := os.Open("./config/config.toml")
		if err != nil {
			log.Println("打开文件失败")
			panic(err)
		}
		//关闭文件
		defer openFile.Close()
		//读取文件
		content, err := ioutil.ReadAll(openFile)
		if err != nil {
			log.Println("读取文件失败")
			panic(err)
		}
		//这里不能直接传入文件路径名，需要先读取config文件的内容，然后将内容进行传入
		_, err = toml.Decode(string(content), &Cfg)*/
	_, err := toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		fmt.Println(Cfg.System.CurrentDir)
		panic(err)
	}
}
