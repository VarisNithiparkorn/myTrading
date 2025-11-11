package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)





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

func generateClaim(tokenType string, field map[string]string) CustomClaims {
    if tokenType == "accesstoken" {
        return &AccessClaims{
            Role: field["role"],
            EmailClaims: EmailClaims{
                UserID:   field["userId"],
                Username: field["username"],
                RegisteredClaims: jwt.RegisteredClaims{
                    IssuedAt:  jwt.NewNumericDate(time.Now()),
                    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
                    Issuer:    "Bitkai_Auth_Service",
                }, 
            }, 
        } 
    } else if tokenType == "refreshtoken" {
        return &RefreshClaims{
			AccessClaims: AccessClaims{
				EmailClaims: EmailClaims{
					UserID: field["userId"],
					Username: field["username"],
					RegisteredClaims: jwt.RegisteredClaims{
					IssuedAt:  jwt.NewNumericDate(time.Now()),
                    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
                    Issuer:    "Bitkai_Auth_Service",
					},
				},
			Role:field["role"] ,
			},	
		}
    } else if tokenType == "emailtoken" {
        return &EmailClaims{
            UserID:   field["userId"],
            Username: field["username"],
            RegisteredClaims: jwt.RegisteredClaims{
                IssuedAt:  jwt.NewNumericDate(time.Now()),
                ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 60)),
                Issuer:    "Bitkai_Auth_Service",
            }, 
        }
    }
    return nil 
}

func GenerateToken(fields map[string]string, tokenType string) (string, error) {
	if signKey == nil {
		return "", fmt.Errorf("RSA private key is not initialized")
	}
	claims := generateClaim(tokenType,fields)
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf("could not sign the token with RSA: %w", err)
	}
	return tokenString, nil
}

func VerifyToken(tokenString string, tokentype string) (CustomClaims, error) {
	if verifyKey == nil {
		return nil, fmt.Errorf("RSA public key is not initialized")
	}
	var claimType CustomClaims
	if tokentype == "emailtoken"{
		claimType = &EmailClaims{}
	}else if tokentype == "accesstoken" {
		claimType = &AccessClaims{}
	}else if tokentype == "refreshtoken"{
	}
	_, err := jwt.ParseWithClaims(
		tokenString,
		claimType,
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
	iss,err2 := claimType.GetIssuer();
    if  err2 != nil{
		fmt.Errorf("can't get issuer")
	}
    if  iss != "Bitkai_Auth_Service" {
        return nil, fmt.Errorf("token issuer is invalid")
    }
	return claimType, nil
}

func IsExpired(claims CustomClaims) bool{
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