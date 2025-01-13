package postgres

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func (rep *Repository) Login(username string, password string) (string, error) {
	var valid bool

	rep.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM wallets WHERE username=$1 AND password=$2)", username, password).Scan(&valid)
	if valid {
		token := getToken(username)
		rep.DB.Exec("UPDATE wallets SET token=$1 WHERE username=$2 AND password=$3", token, username, password)

		return token, nil

	} else {
		return "", fmt.Errorf("неверный логин или пароль")
	}
}

func getToken(usename string) string {
	claims := Claims{
		Username:       usename,
		StandardClaims: jwt.StandardClaims{},
	}
	jwtTokenUns := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtTokenS, _ := jwtTokenUns.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return jwtTokenS
}
