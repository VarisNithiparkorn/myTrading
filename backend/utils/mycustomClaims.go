package utils

import "github.com/golang-jwt/jwt/v5"

type CustomClaims interface {
	GetUsername() string
	jwt.Claims
}
type EmailClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
type AccessClaims struct {
	Role string `json:"role"`
	EmailClaims
}
type RefreshClaims struct {
	AccessClaims
}
func (t *EmailClaims) GetUsername() string {
	return t.Username
}
func (t *AccessClaims) GetUsername() string {
	return t.Username
}
func (t *RefreshClaims) GetUsername() string {
	return t.Username
}

type EmailClaimsFields map[string]string
type AcceessClaimsFields map[string]string
type RefreshClaimsFields map[string]string