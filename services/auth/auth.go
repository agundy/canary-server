package auth

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/agundy/canary-server/database"
)

func CheckAuthToken(u *models.User) (tokenString string) {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["email"] = u.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the JWT with the server secret
	tokenString, _ = token.SignedString(ApiSecret)

	return tokenString
}

func GetAuthToken() (tokenString string) {
	jwtString := req.Header.Get("auth")
	token, err := jwt.Parse(jwtString, ApiSecret)
	if err == nil && token.Valid {
		// token parsed, exp/nbf checks out, signature verified, Valid is true
		return token, nil
	}
	return nil, jwt.ErrNoTokenInRequest
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims["email"] = u.Email
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the JWT with the server secret
	tokenString, _ = token.SignedString(ApiSecret)

	return tokenString
}
