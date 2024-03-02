package midware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

var jwtKey = []byte("daitoue_secret_key") // 密钥

func CreateToken(id int64) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //终止时间，7天后
	claims := &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "degim",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}

func TokenParse() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenStr := context.Query("token")
		if tokenStr == "" {
			tokenStr = context.PostForm("token")
		}
		if tokenStr == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 401,
				"status_msg":  "没有权限",
			})
		}
		token, claims, err := ParseToken(tokenStr)
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status_code": 401,
				"status_msg":  "没有权限",
			})
			context.Abort()
			return
		}
		context.Set("userId", claims.UserId)
	}
}
