package simpleauth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var Issuer string = "simpleauth"
var Audience string = Issuer
var TokenValidDuration time.Duration = 7 * 24 * time.Hour

var SigningMethod jwt.SigningMethod = jwt.SigningMethodHS256

func GenerateJWT(userID string, secretKey []byte) (string, error) {

	token := jwt.NewWithClaims(SigningMethod, jwt.MapClaims{
		"sub": userID,                                    // subject
		"iss": Issuer,                                    // issuer
		"aud": Audience,                                  // audience
		"exp": time.Now().Add(TokenValidDuration).Unix(), // expiration
		"nbf": time.Now().Unix(),                         // not valid before
		"iat": time.Now().Unix(),                         // issued at
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["aud"] != Audience {
			return nil, fmt.Errorf("invalid audience")
		}
	}

	return token, nil
}
