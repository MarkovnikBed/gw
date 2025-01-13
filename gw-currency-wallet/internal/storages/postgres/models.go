package postgres

import (
	"database/sql"

	"github.com/dgrijalva/jwt-go"
)

type Repository struct {
	DB *sql.DB
}

type Claims struct {
	Username string
	jwt.StandardClaims
}
