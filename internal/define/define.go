package define

import "github.com/golang-jwt/jwt/v4"

const SECRET_KEY = "huhaifeng_key"

const REDIS_CONN_IP = "localhost:6379"
const REDIS_CONN_PWD = ""

type UserClaim struct {
	Id       string
	Identity string
	Name     string
	jwt.RegisteredClaims
}
