package utils

import (
	"encoding/base64"
	"errors"
)

func EncodeScriptContent(content string) string {
	return base64.StdEncoding.EncodeToString([]byte(content))
}

func DecodeScriptContent(encoded string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", errors.New("脚本内容 Base64 解码失败: " + err.Error())
	}
	return string(data), nil
}
