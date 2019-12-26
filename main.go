package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"learn_go/model"
	"learn_go/pkg/setting"
	"learn_go/routers"
)

func init() {
	setting.Setup()
	model.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	file, _ := os.Create(fmt.Sprintf("%s/%s.%s",
		setting.AppSetting.LogSavePath,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	))

	gin.DefaultWriter = io.MultiWriter(file)
	log.SetOutput(gin.DefaultWriter)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	routersInit := routers.GetRouters()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	port := fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           port,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	fmt.Println(fmt.Sprintf("http server listening localhost:%d", setting.ServerSetting.HTTPPort))

	// routersInit.Run(fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	server.ListenAndServe()
	// routers.GetRouters()
}
