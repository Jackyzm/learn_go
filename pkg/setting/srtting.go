package setting

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type config struct {
	App      app      `yaml:"app"`
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
}

type app struct {
	PageSize       int    `yaml:"PageSize"`
	JwtSecret      int    `yaml:"JwtSecret"`
	ImageMaxSize   int    `yaml:"ImageMaxSize"`
	ImageAllowExts string `yaml:"ImageAllowExts"`
	LogSavePath    string `yaml:"LogSavePath"`
	LogFileExt     string `yaml:"LogFileExt"`
	TimeFormat     string `yaml:"TimeFormat"`
}

// AppSetting app global setting
var AppSetting = app{}

type server struct {
	RunMode      string        `yaml:"RunMode"`
	HTTPPort     int           `yaml:"HTTPPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

// ServerSetting app server setting
var ServerSetting = server{}

type database struct {
	Type     string `yaml:"Type"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Name     string `yaml:"Name"`
}

// DatabaseSetting app database setting
var DatabaseSetting = database{}

// Setup init setting
func Setup() {
	content, err := ioutil.ReadFile("config/config.yaml")

	if err != nil {
		fmt.Println(err)
	}
	var config config

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		fmt.Println(err)
	}

	AppSetting = config.App
	ServerSetting = config.Server
	DatabaseSetting = config.Database

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	// RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	fmt.Println("Success to Load Setting!")
}
