package helper

import (
	"blog/rpc/internal/define"
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

// AnalyseToken 解析token，返回user对象，向请求头里放userId,username和identity
func AnalyseToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err

}
