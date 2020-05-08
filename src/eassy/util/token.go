package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenStruct struct {
	Phone  string
	AccPwd string
}

type JwtToken struct {
	Token string `json:"token"`
}

func GenerateToken(ts *TokenStruct) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"phone": ts.Phone,
		"exp":   time.Now().Add(time.Hour * 2).Unix(), // 可以添加过期时间
	})

	return token.SignedString([]byte("secret")) //对应的字符串请自行生成，最后足够使用加密后的字符串
}

func CheckToken(Phone, tokenStr string) (ret bool) {
	if tokenStr == "" {
		ret = false
	} else {
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("not authorization")
			}
			return []byte("secret"), nil
		})
		if !token.Valid {
			ret = false
		} else {
			//验证AccName与Token
			finToken := token.Claims.(jwt.MapClaims)
			ret = (finToken["phone"] == Phone)
		}
	}
	return
}
