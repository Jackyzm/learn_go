package model

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// UserInfo Binding from JSON
type UserInfo struct {
	Mobile   string `json:"mobile"`   // 手机号
	UserName string `json:"userName"` // 用户名
	ID       string `json:"id"`       // id
}

// Account Binding from JSON
type Account struct {
	Mobile   string `json:"mobile" binding:"required,mobile"` // 手机号
	Password string `json:"password" binding:"required"`      // 密码
}

// Register user register
func (c *Account) Register() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result Account
	userCol.FindOne(ctx, bson.M{"mobile": c.Mobile}).Decode(&result)

	if result != (Account{}) {
		return errors.New("手机号码已存在")
	}

	_, err := userCol.InsertOne(ctx, bson.M{
		"mobile":   c.Mobile,
		"password": c.Password,
		"id":       uuid.New().String(),
	})

	if err != nil {
		return err
	}

	return nil
}

// GetUser find user
func (c *Account) GetUser() (UserInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userInfo UserInfo
	userCol.FindOne(ctx, bson.M{"mobile": c.Mobile}).Decode(&userInfo)

	if userInfo == (UserInfo{}) {
		return UserInfo{}, errors.New("用户不存在")
	}

	return userInfo, nil
}

// Login login user
func (c *Account) Login() (UserInfo, Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var userInfo UserInfo
	var account Account
	res := userCol.FindOne(ctx, bson.M{"mobile": c.Mobile})
	res.Decode(&userInfo)
	res.Decode(&account)

	if userInfo == (UserInfo{}) {
		return UserInfo{}, Account{}, errors.New("用户不存在")
	}

	return userInfo, account, nil
}
