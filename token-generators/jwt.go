package token_generators

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"crypto/rsa"
)

func IssueJwt(clientId string, key *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": clientId,
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString(key)
	return tokenStr, err
}
