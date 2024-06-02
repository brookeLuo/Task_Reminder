package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

var jwtKey = []byte("task_reminder") // 将这个密钥保密

func GenerateToken(id uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 设置 token 的有效期为 24 小时

	// 创建一个新的 token 对象，指定签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   strconv.Itoa(int(id)),
	})

	// 创建 JWT 字符串
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractUserIDFromToken(tokenString string) (int, error) {
	// 解析 JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥
		return jwtKey, nil
	})

	if err != nil {
		return 0, err // 解析错误
	}

	// 检查 Token 是否有效
	if !token.Valid {
		return 0, errors.New("Invalid token")
	}

	// 从 Token 的声明中提取用户 ID
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Failed to parse claims")
	}

	// 提取用户 ID
	idStr, ok := claims["sub"].(string)
	if !ok {
		return 0, errors.New("user ID not found in claims")
	}

	// 将用户 ID 转换为整数
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("Failed to convert user ID to integer")
	}

	return id, nil
}
