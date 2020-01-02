package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"learn_go/model"
	"learn_go/pkg/setting"
)

var signature = []byte(setting.AppSetting.JwtSecret)

// Claims jwt payload 部分
type Claims struct {
	UserInfo *model.UserInfo `json:"userInfo"`
	jwt.StandardClaims
}

// Token 返回前端token定义
type Token struct {
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
}

// GenerateToken 设置jwt
func GenerateToken(UserInfo model.UserInfo) (Token, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		&model.UserInfo{
			Mobile:   UserInfo.Mobile,
			UserName: UserInfo.UserName,
			ID:       UserInfo.ID,
		},
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	AccessToken, err := tokenClaims.SignedString(signature)

	resToken := Token{
		AccessToken: AccessToken,
		ExpiresIn:   3 * time.Hour / 1e9,
	}

	return resToken, err
}

// ParseToken 解析jwt
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return signature, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
