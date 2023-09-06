package common

import "math/rand"

// 生成字符串随机数
func GenerateRandomString(length int) string {
	// 首先定义随机范围
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 创建切片 用以存储生成随机数
	result := make([]byte, length)
	// 遍历 随机生成随机数
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	// 返回
	return string(result)
}
