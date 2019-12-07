package user

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

// Login Binding from JSON
type Login struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// SuccessRes Res
type SuccessRes struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

// BadRes Res
type BadRes struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

// Response setting gin.JSON
// func (c *gin.Context) Response(httpCode, errCode int, data interface{}) {
// 	c.JSON(httpCode, Response{
// 		Code: errCode,
// 		Msg:  e.GetMsg(errCode),
// 		Data: data,
// 	})
// 	return
// }

// Register 注册
// @tags 用户
// @Summary 注册
// @Produce  json
// @Success 200
// @Failure 500
// @Router /register [put]
func Register(c *gin.Context) {
	var json Login
	// c.BindJSON(&json)
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, BadRes{
			Code:    http.StatusBadRequest,
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	// fmt.Println(json.Mobile, json.Password, json)

	c.JSON(http.StatusOK, SuccessRes{
		Code:    http.StatusOK,
		Success: true,
		Data: map[string]interface{}{
			"mobile":   json.Mobile,
			"password": json.Password,
		},
	})
	// mobile := c.PostForm("mobile")
	// password := c.PostForm("password")
	// fmt.Println(mobile, password)
}
