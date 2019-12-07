package setting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type config struct {
	App      app      `json:"app"`
	Server   server   `json:"server"`
	Database database `json:"database"`
}

type app struct {
	PageSize       int    `json:"PageSize"`
	JwtSecret      int    `json:"JwtSecret"`
	ImageMaxSize   int    `json:"ImageMaxSize"`
	ImageAllowExts string `json:"ImageAllowExts"`
	LogSavePath    string `json:"LogSavePath"`
	LogSaveName    string `json:"LogSaveName"`
	LogFileExt     string `json:"LogFileExt"`
	TimeFormat     string `json:"TimeFormat"`
}

// AppSetting app global setting
var AppSetting = app{}

type server struct {
	// debug or release
	RunMode      string        `json:"RunMode"`
	HTTPPort     int           `json:"HTTPPort"`
	ReadTimeout  time.Duration `json:"ReadTimeout"`
	WriteTimeout time.Duration `json:"WriteTimeout"`
}

// ServerSetting app server setting
var ServerSetting = server{}

type database struct {
	Type     string `json:"Type"`
	User     string `json:"User"`
	Password string `json:"Password"`
	Host     string `json:"Host"`
	Name     string `json:"Name"`
}

// Databasesetting app database setting
var Databasesetting = database{}

// Setup init setting
func Setup() {
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config config
	json.Unmarshal([]byte(byteValue), &config)
	AppSetting = config.App
	ServerSetting = config.Server
	Databasesetting = config.Database

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	// RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	fmt.Println("Success to Load Setting!")
}
