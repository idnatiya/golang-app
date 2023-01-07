package types

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct {
	jwt.StandardClaims
	ID        *uint   `json:"ID"`
	Email     *string `json:"Email"`
	FirstName *string `json:"FirstName"`
	LastName  *string `json:"LastName"`
}
