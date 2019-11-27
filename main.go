package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"learn_go/routers"
)

func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	t := time.Now().Format("2006-01-02 15:04:05") // 据说是go诞生时间 1-2-3-4-5

	// Logging to a file.
	f, _ := os.Create(fmt.Sprintf("./log/gin-%s.log", t))
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)

	routers.GetRouters()
}
