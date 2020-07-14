package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//常量key
const (
	SIGNED_KEY = "klklslklsm"
)

/***
创建token
*/
func CreateTocken(uid string) string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 3600),
		Issuer:    uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(SIGNED_KEY))
	if err != nil {
		logs.Error(err)
	}

	return tokenStr
}

/***
校验token
*/
func CheckToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SIGNED_KEY), nil
	})
	if err != nil {
		logs.Error("HS256的token解析错误，err:", err)
		return false
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logs.Error("ParseHStoken:claims类型转换失败")
		return false
	}
	return true
}
