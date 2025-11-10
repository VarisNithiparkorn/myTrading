package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
type Mytoken interface{
	getUsername() string
}
type EmailTokenClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
func (t *EmailTokenClaims) getUsername() string{
	return t.getUsername()
}
type EmailTokenClaimsFields map[string]string

type AccessToken struct{
	Username string `json:"username"`
	Role string `json:"role"`
}
func (t *AccessToken) getUsername() string{
	return t.getUsername()
}
type AcceessTokenClaimsFields map[string]string
var signKey *rsa.PrivateKey
var verifyKey *rsa.PublicKey

func init() {

	var err error
	signKey, err = rsa.GenerateKey(rand.Reader,2048)
	if err != nil {
		fmt.Printf("FATAL ERROR: Failed to parse RSA private key: %v\n", err)
		return
	}
	verifyKey = &signKey.PublicKey
	fmt.Println("RSA Private Key loaded successfully.")
}

func generateClaim(tokenType string){
	if(tokenType == "accesstoken"){

	}else if tokenType == "refreshtoken" {
		
	}else if tokenType == "emailtoken"{
		claims := &EmailTokenClaims{
			UserID:   userID,
			Username: username,
			RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "Bitkai_Auth_Service",
		},
		
		}
	}
}

func GenerateAccessToken(userID, username string, t time.Duration, tokenType string) (string, error) {
	// Check if the key was loaded successfully
	if signKey == nil {
		return "", fmt.Errorf("RSA private key is not initialized")
	}

	expirationTime := time.Now().Add(t)



	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf("could not sign the token with RSA: %w", err)
	}

	return tokenString, nil
}
func VerifyAccessToken(tokenString string) (*EmailTokenClaims, error) {
	
	if verifyKey == nil {
		return nil, fmt.Errorf("RSA public key is not initialized")
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&EmailTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			
			return verifyKey, nil
		},
	)

	
	if err != nil {
		return nil, fmt.Errorf("token parsing/signature error: %w", err)
	}

	
	claims, ok := token.Claims.(*EmailTokenClaims)
	if !ok || !token.Valid {
		
		return nil, fmt.Errorf("invalid or expired token claims")
	}
    
    
    if claims.Issuer != "Bitkai_Auth_Service" {
        return nil, fmt.Errorf("token issuer is invalid")
    }

	
	return claims, nil
}

func IsExpired(claims jwt.Claims) bool{
	expireTime, ex :=claims.GetIssuedAt()
	createTime, ec := claims.GetIssuedAt()

	if(ex != nil){
		fmt.Errorf("invalid date %w",ex)
	}
	if(ec != nil){
		fmt.Errorf("invalid date %w",ec)
	}
	
	return createTime.Time.After(expireTime.Time)
}