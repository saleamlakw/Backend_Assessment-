package tokenutil

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/saleamlakw/LoanTracker/domain/entities"
)

func CreateVerificationToken(user *entities.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &entities.JwtCustomVerifyClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)

	if !ok && !token.Valid {
		return false, fmt.Errorf("invalid Token")
	}

	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

	if time.Now().After(expirationTime) {
		return false, fmt.Errorf("token has expired")
	}
	return true, nil

}

func CreateAccessToken(user *entities.User, secret string, expiry int,refreshDataID string) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &entities.JwtCustomClaims{
		Role:    user.Role,
		ID:      user.ID.Hex(),
		RefreshDataId: refreshDataID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *entities.User, secret string, expiry int,refreshDataID string) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claimsRefresh := &entities.JwtCustomRefreshClaims{
		ID: user.ID.Hex(),
		RefreshDataId: refreshDataID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err
}

func ExtractUserClaimsFromToken(requestToken string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return jwt.MapClaims{}, fmt.Errorf("Invalid Token")
	}

	return claims, nil
}