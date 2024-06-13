package define

import "github.com/golang-jwt/jwt/v4"

const SECRET_KEY = "huhaifeng_key"

type UserClaim struct {
	Id       string
	Identity string
	Name     string
	jwt.RegisteredClaims
}
