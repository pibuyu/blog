package helper

import (
	"fmt"
	"math/rand"
	"time"
)

// GenVerifyCode 生成6位随机验证码
func GenVerifyCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
