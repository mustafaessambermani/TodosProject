package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a new JWT for a specific user ID
func GenerateToken(userID uint, secret string) (string, error) {
	// 1. Define the claims (The payload)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":     time.Now().Unix(),                     // Issued at
	}

	// 2. Create the token using the HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Sign the token with the secret key
	return token.SignedString([]byte(secret))
}

// ParseToken validates the token string and returns the claims
func ParseToken(tokenString string, secret string) (jwt.MapClaims, error) {
	// 1. Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	// 2. Verify claims and validity
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
