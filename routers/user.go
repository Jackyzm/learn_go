package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learn_go/models"
)

// AddUser 新增用户
// @tags 用户
// @Summary 接口名称
// @Produce  json
// @Success 200
// @Failure 500
// @Router /user [put]
func AddUser(c *gin.Context) {
	models.SetUser()
	c.JSON(http.StatusOK, "success")
}
