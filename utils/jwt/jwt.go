package jwt

import (
	"code/01/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var MySecret = []byte("app")


// GenToken 生成JWT
func GenToken(id string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		id, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    config.TokenIssuer,                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}