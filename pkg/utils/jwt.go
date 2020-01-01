package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"learn_go/model"
	"learn_go/pkg/setting"
)

var signature = []byte(setting.AppSetting.JwtSecret)

// Claims Claims
type Claims struct {
	UserInfo *model.UserInfo `json:"userInfo"`
	jwt.StandardClaims
}

// GenerateToken 设置jwt
func GenerateToken(UserInfo model.UserInfo) (string, error) {
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
	token, err := tokenClaims.SignedString(signature)

	return token, err
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
