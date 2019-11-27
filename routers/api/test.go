package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPing 测试
// @tags 测试
// @Summary 接口名称
// @Produce  json
// @Success 200
// @Failure 500
// @Router /api/ping [get]
func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "abc",
	})
}

// GetPing1 测试
// @tags 测试时
// @Summary 接口名称
// @Produce  json
// @Success 200
// @Failure 500
// @Router /api/ping1 [get]
func GetPing1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "abc",
	})
}
