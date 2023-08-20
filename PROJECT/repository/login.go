package repository

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// import (
// 	"errors"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/jaseelaali/passwordgenerator_unique/database"
// 	"golang.org/x/crypto/bcrypt"
// )

// func LOgin(Email, Password string) error {
// 	var password string
// 	err := database.DB.Raw("SELECT password FROM users WHERE email=$1 ;", Email).Scan(&password)
// 	if err != nil {
// 		return err.Error
// 	}
// 	if password == "" {
// 		return errors.New("Invalid user")
// 	}
// 	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(Password)); err != nil {
// 		return errors.New("Email or password is not matching")
// 	}
// 	return nil
// }

//package handlers

func Token(Email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// Replace "your-secret-key" with your actual secret key for token signing
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
