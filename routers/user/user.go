package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	model "learn_go/model"
	response "learn_go/pkg/response"
	"learn_go/pkg/utils"
	"learn_go/pkg/validate"
)

// Register 注册
// @tags 用户
// @Summary 注册
// @Produce  json
// @Param Account body model.Account true "登陆"
// @Success 200 "成功"
// @Failure 400 "参数错误"
// @Router /register [post]
func Register(c *gin.Context) {
	var account model.Account
	response := response.Gin{C: c}

	err := c.ShouldBindJSON(&account)

	if err != nil {
		errstr := validate.ValidatorError(err)
		log.Println("错误：" + errstr)
		response.SetBadResponse(http.StatusBadRequest, errstr)
		return
	}

	err = account.Register()

	if err != nil {
		log.Println("错误：" + err.Error())
		response.SetBadResponse(http.StatusInternalServerError, err.Error())
		return
	}

	var userInfo model.UserInfo
	userInfo, _ = account.GetUser()

	token, err := utils.GenerateToken(userInfo)
	if err != nil {
		log.Println("error:" + err.Error())
		response.SetBadResponse(http.StatusInternalServerError, err.Error())
		return
	}

	response.SetSuccessResponse(http.StatusOK, token)
}

// Login 登陆
// @tags 用户
// @Summary 登陆
// @Produce  json
// @Param Account body model.Account true "登陆"
// @Success 200 {object} model.UserInfo "成功"
// @Failure 400 "参数错误"
// @Router /login [post]
func Login(c *gin.Context) {
	var account model.Account
	response := response.Gin{C: c}

	err := c.ShouldBindJSON(&account)

	// 敏感信息 不打印
	// reqBody, _ := json.Marshal(account)
	// log.Println(string(reqBody))

	if err != nil {
		errstr := validate.ValidatorError(err)
		log.Println("error:" + errstr)
		response.SetBadResponse(http.StatusBadRequest, errstr)
		return
	}

	userInfo, accountDB, err := account.Login()

	if err != nil {
		log.Println("error:" + err.Error())
		response.SetBadResponse(http.StatusInternalServerError, err.Error())
		return
	}

	if account.Password != accountDB.Password {
		log.Println("error:密码错误")
		response.SetBadResponse(http.StatusInternalServerError, "密码错误")
		return
	}

	token, err := utils.GenerateToken(userInfo)
	if err != nil {
		log.Println("error:" + err.Error())
		response.SetBadResponse(http.StatusInternalServerError, err.Error())
		return
	}

	response.SetSuccessResponse(http.StatusOK, token)
}
