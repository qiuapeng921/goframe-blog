package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

var (
	jwtSecret = g.Cfg().Get("token.secret")
)

type MapClaims struct {
	Id       uint   `json:"id"`
	Account  string `json:"account"`
	Category string `json:"category"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(id uint, account, category string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := MapClaims{
		id,
		account,
		category,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(gconv.Bytes(jwtSecret))

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*MapClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return gconv.Bytes(jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MapClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
