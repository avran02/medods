package models

import "github.com/golang-jwt/jwt/v5"

type AccessTokenClaims struct {
	UserIP string `json:"user_ip"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	AccessTokenID string `json:"access_token_id"`
	UserIP        string `json:"user_ip"`
	jwt.RegisteredClaims
}
