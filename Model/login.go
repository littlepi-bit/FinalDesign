package Model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	ErrorServerBusy = "server is busy"
	ErrorReLogin    = "relogin"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID   int    `json:"userId"`
	Password string `json:"password"`
	Username string `json:"username"`
}

var (
	//自定义的token秘钥
	Secret = []byte("16849841325189456f487")
	//该路由下不校验token
	noVerify = []string{"/loginCheck", "/ping", "/signIn", "/"}
	//token有效时间（纳秒）
	effectTime = 2 * time.Hour
)

// 生成token
func GenerateToken(claims *JWTClaims) string {
	//设置token有效期
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	// 生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(Secret)
	if err != nil {
		panic(err)
	}
	return sign
}

//验证token
func JwtVerfiy(c *gin.Context) {
	//过滤是否验证token
	if isContain(noVerify, c.Request.RequestURI) {
		return
	}
	MyToken := c.GetHeader("token")
	fmt.Println(MyToken)
	if MyToken == "" {
		c.String(http.StatusForbidden, "禁止访问")
		return
	}
	//验证token，并储存在请求中
	user := ParseToken(MyToken)
	if !IsExist(user.Username) {
		fmt.Println("用户不存在")
		c.String(http.StatusForbidden, "禁止访问")
		return
	}
	c.Set("user", user)
}

//解析token
func ParseToken(tokenString string) *JWTClaims {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		log.Println(err.Error())
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		panic("token is valid")
	}
	return claims
}

// 更新token
func Refresh(tokenString string) string {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		panic("token is valid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
	return GenerateToken(claims)
}

func IsExist(userName string) bool {
	if userName == "" {
		return false
	}
	var user = User{}
	result := GlobalConn.Table("user").Where("name=?", user.Name).First(&user)
	return result.Error == nil
}

func isContain(strArr []string, s string) bool {
	for _, str := range strArr {
		if str == s {
			return true
		}
	}
	return false
}

func JsontoString(h gin.H) string {
	jsonByte, _ := json.Marshal(h)
	return string(jsonByte)
}

func GetSystemTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
