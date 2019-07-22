package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

// H 用于存储 jwt 的数据结构
type H map[string]interface{}

// JwtParse 解析 tokenString 到 H
func JwtParse(tokenString string) (*jwt.Token, H, error) {

	// sample token string taken from the New example
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("0okm3edC"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		retH := H{}
		for k, v := range claims {
			retH[k] = v
		}
		return token, retH, nil
	}
	return nil, nil, err
}

// JwtNew 返回一个新的 jwt？
func JwtNew(h H) (string, error) {
	mapClaims := jwt.MapClaims{}
	for k, v := range h {
		mapClaims[k] = v
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("buzhidao."))

	return tokenString, err
}

// GetInt64 jwt 中的 和 outTime 有关 具体不知道
func GetInt64(outTime interface{}) int64 {

	switch inst := outTime.(type) {

	case int64:
		return inst
	case float64:
		return int64(inst)
	default:
		log.Printf("%#v", inst)
		log.Println("GetInt64")
	}

	return 0
}
