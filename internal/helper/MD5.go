package helper

import (
	"crypto/md5"
	"fmt"
)

func Md5(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
