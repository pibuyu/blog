package define

import "github.com/golang-jwt/jwt/v4"

const SECRET_KEY = "huhaifeng_key"

const REDIS_CONN_IP = "localhost:6379"
const REDIS_CONN_PWD = ""
const REDIS_VERI_CODE_PRE = "REGISTER_VERIFY_CODE_OF_"

type UserClaim struct {
	Id       string
	Identity string
	Name     string
	jwt.RegisteredClaims
}
