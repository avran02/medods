package models

type User struct {
	ID           int
	Username     string
	PasswordHash string
	Email        string
}

type RefreshToken struct {
	UserID        int
	TokenHash     string
	AccessTokenID string
}
