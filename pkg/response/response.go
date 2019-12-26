package response

import (
	"github.com/gin-gonic/gin"
)

// Gin get ctx
type Gin struct {
	C *gin.Context
}

// SuccessResponse SuccessResponse
type SuccessResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

// BadResponse BadResponse
type BadResponse struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

// SetBadResponse bad response
func (G *Gin) SetBadResponse(httpCode int, errMsg string) {
	G.C.JSON(httpCode, BadResponse{
		Code:    httpCode,
		Msg:     errMsg,
		Success: false,
	})
}

// SetSuccessResponse success response
func (G *Gin) SetSuccessResponse(httpCode int, data interface{}) {
	G.C.JSON(httpCode, SuccessResponse{
		Code:    httpCode,
		Data:    data,
		Success: true,
	})
}
