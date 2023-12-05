package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	user2 "kwseeker.top/chatgpt/chatgpt-api/domain/auth/user"
	"time"
)

const SecretKey = "#!@"

var Eng = New()

type Token struct {
	Content string
}

type Engine struct {
	method    jwt.SigningMethod
	secretKey string
}

func New() *Engine {
	return &Engine{
		method:    jwt.SigningMethodHS256,
		secretKey: SecretKey,
	}
}

func (e *Engine) GenerateWithUser(u *user2.User) (string, error) {
	data := jwt.MapClaims{
		"name": u.Name,
		"role": u.Role,
		"age":  u.Age,
		"desc": u.Desc,
	}
	return e.GenerateWithMapClaims(data)
}

func (e *Engine) GenerateWithMapClaims(data jwt.MapClaims) (string, error) {
	//添加超时时间
	data["exp"] = time.Now().Add(time.Hour * 2).Unix()
	//生成token
	token := jwt.NewWithClaims(e.method, data)
	// 将 JWT 编码为字符串
	tokenString, err := token.SignedString([]byte(e.secretKey))
	if err != nil {
		logrus.Error("Error signing token:", err)
		return "", err
	}
	logrus.Debug("Generated token: ", tokenString)
	return tokenString, nil
}
