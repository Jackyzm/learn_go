package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"learn_go/models"
	"learn_go/pkg/setting"
	"learn_go/routers"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	t := time.Now().Format("2006-01-02") // 据说是go诞生时间 6-1-2-3-4-5 2006-01-02 15:04:05

	// Logging to a file.
	f, _ := os.Create(fmt.Sprintf("./log/gin-%s.log", t))
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)

	routersInit := routers.GetRouters()
	routersInit.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// routers.GetRouters()
}
