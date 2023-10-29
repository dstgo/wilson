package jwtx

import "github.com/golang-jwt/jwt/v4"

type Jwt struct {
	Token     *jwt.Token
	Claims    jwt.Claims
	SignedJwt string
}

func NewJwt(secret string, method jwt.SigningMethod, claims jwt.Claims) (Jwt, error) {
	var jwtToken Jwt
	// create token struct
	token := jwt.NewWithClaims(method, claims)
	jwtToken.Token = token

	// singed with secret
	signedStr, err := token.SignedString([]byte(secret))
	if err != nil {
		return jwtToken, err
	}

	jwtToken.SignedJwt = signedStr
	jwtToken.Claims = token.Claims

	return jwtToken, nil
}

func ParseJwt(tokenStr, secret string, method jwt.SigningMethod, claims jwt.Claims) (Jwt, error) {
	var jwtToken Jwt

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}, jwt.WithValidMethods([]string{method.Alg()}))

	if token != nil {
		jwtToken.Token = token
		jwtToken.Claims = token.Claims
	}

	return jwtToken, err
}
