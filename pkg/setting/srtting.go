package setting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

// Appsetting app global setting
var Appsetting = app{}

type server struct {
	// debug or release
	RunMode      string `json:"RunMode"`
	HTTPPort     int    `json:"HTTPPort"`
	ReadTimeout  int    `json:"ReadTimeout"`
	WriteTimeout int    `json:"WriteTimeout"`
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
	Appsetting = config.App
	ServerSetting = config.Server
	Databasesetting = config.Database

	fmt.Println("Success to Load Setting!")
}
