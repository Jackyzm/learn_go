package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	model "learn_go/model"
	response "learn_go/pkg/response"
	validate "learn_go/pkg/validate"
)

// Register 注册
// @tags 用户
// @Summary 注册
// @Produce  json
// @Param Login body model.Login true "登陆"
// @Success 200
// @Failure 500
// @Router /register [post]
func Register(c *gin.Context) {
	var login model.Login
	response := response.Gin{C: c}

	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.SetBadResponse(http.StatusBadRequest, err.Error())
		return
	}

	reqBody, _ := json.Marshal(login)
	log.Println(string(reqBody))

	err = validate.VerifyMobile(login.Mobile)
	if err != nil {
		response.SetBadResponse(http.StatusBadRequest, err.Error())
		return
	}

	err = validate.VerifyPwd(login.Password)
	if err != nil {
		response.SetBadResponse(http.StatusBadRequest, err.Error())
		return
	}

	err = login.Register()

	if err != nil {
		log.Println(err.Error())
		response.SetBadResponse(http.StatusInternalServerError, err.Error())
		return
	}

	response.SetSuccessResponse(http.StatusOK, "注册成功")
}
