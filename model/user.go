package model

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// UserInfo Binding from JSON
type UserInfo struct {
	UserName string `json:"userName"` // 用户名
	Mobile   string `json:"mobile" `  // 手机号
	ID       string `json:"id"`       // id
}

// Login Binding from JSON
type Login struct {
	Mobile   string `json:"mobile" binding:"required"`   // 手机号
	Password string `json:"password" binding:"required"` // 密码
}

// Register SetPi
func (c *Login) Register() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result UserInfo
	userCol.FindOne(ctx, bson.M{"mobile": c.Mobile}).Decode(&result)
	log.Println(result)
	if result != (UserInfo{}) {
		return errors.New("手机号码已存在")
	}

	_, err := userCol.InsertOne(ctx, bson.M{
		"mobile":   c.Mobile,
		"password": c.Password,
		// "id":       uuid.New(),
	})

	if err != nil {
		return err
	}

	return nil
}
