package helper

import (
	"blog/rpc/internal/define"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// GenerateToken 生成token
func GenerateToken(id string, identity string, name string) (token string, err error) {
	//设置token为7天过期
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		},
	}
	//加密
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	//盐值加密
	signedToken, err := t.SignedString([]byte(define.SECRET_KEY)) // []byte(define.SECRET_KEY):define.SECRET_KEY转为byte数组类型
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
