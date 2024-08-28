package entities

import "github.com/golang-jwt/jwt/v4"

type JwtCustomVerifyClaims struct {
	Role    string `json:"role"`
	ID      string `json:"id"`
	jwt.RegisteredClaims
}
type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	RefreshDataId string `json:"refresh_data_id"`
	jwt.RegisteredClaims
}
type JwtCustomClaims struct {
	Role    string `json:"role"`
	ID      string `json:"id"`
	RefreshDataId string `json:"refresh_data_id"`
	jwt.RegisteredClaims
}