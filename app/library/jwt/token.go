package jwt

import (
	"blog/app/model/user"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

var (
	secrets = g.Cfg().Get("token.secret")
)

func CreateToken(user *user.Entity) (string, error) {
	//自定义claim
	claim := jwt.MapClaims{
		"id":       user.Id,
		"username": user.Account,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	var accessToken, err = token.SignedString(gconv.Bytes(secrets))
	if err != nil {
		fmt.Println("Error while signing the token")
		return "", err
	}
	return accessToken, err

}

func ParseToken(requestToken string) (userEntity user.Entity, err error) {
	token, err2 := jwt.Parse(requestToken, func(*jwt.Token) (interface{}, error) {
		return gconv.Bytes(secrets), nil
	})
	if err2 != nil {
		fmt.Println("err1", err2)
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapClaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	var entity user.Entity
	entity.Id = uint(claim["id"].(float64))
	entity.Account = claim["username"].(string)
	return entity, nil
}
