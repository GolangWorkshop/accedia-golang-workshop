package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JwtInfo struct {
	Token     string
	ExpiresAt time.Time
}

var JwtKey = []byte(os.Getenv("JWT_KEY"))
