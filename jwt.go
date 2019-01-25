package goutil

import (
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// Token JwtToken
type Token struct {
	Token   string
	ExpTime time.Time
}

// CreateJwtTokenByParam 生成JWT TOKEN
func CreateJwtTokenByParam(params map[string]string, secret string, exp int64, fun func(token string) error) (*Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	expTime := time.Now().Add(time.Second * time.Duration(exp)).Unix()
	claims["exp"] = expTime
	claims["iat"] = time.Now().Unix()

	for key, value := range params {
		claims[key] = value
	}

	token.Claims = claims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	if fun != nil {
		err = fun(tokenString)
		if err != nil {
			return nil, err
		}
	}
	return &Token{Token: tokenString, ExpTime: time.Unix(expTime, 1)}, nil
}

func getToken(req *http.Request, secret string) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(req, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return token, err
}

// ValidToken 验证token
func ValidToken(req *http.Request, secret string, fun func() bool) bool {
	token, err := getToken(req, secret)
	if err != nil || !token.Valid {
		return false
	}
	if fun != nil {
		return fun()
	}
	return true
}

// GetInfoFromJwtToken 从Token解密获取信息
func GetInfoFromJwtToken(req *http.Request, secret string) (map[string]interface{}, error) {
	token, err := getToken(req, secret)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token验证未通过")
	}

	claims := token.Claims
	result := make(map[string]interface{}, len(claims.(jwt.MapClaims)))
	for key, value := range claims.(jwt.MapClaims) {
		result[key] = value
	}
	return result, nil
}

// GetValueFromJwtToken 从Token解密获取某个字段信息
func GetValueFromJwtToken(req *http.Request, key string, secret string) (interface{}, error) {
	token, err := getToken(req, secret)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token验证未通过")
	}

	claims := token.Claims
	return claims.(jwt.MapClaims)[key], nil
}
