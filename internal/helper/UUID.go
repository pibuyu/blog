package helper

import "github.com/google/uuid"

// GetUUID 生成UUID
func GetUUID() string {
	res, err := uuid.NewV6()
	if err != nil {
		return "生成uuid出错"
	}
	return res.String()
}
