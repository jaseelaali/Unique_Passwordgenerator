package repository

import (
	"errors"
	"math/rand"
	"time"

	"github.com/jaseelaali/passwordgenerator_unique/database"
	"github.com/jaseelaali/passwordgenerator_unique/model"
	"golang.org/x/crypto/bcrypt"
)

func SignupPassword(body model.Signup) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var password string
	err := database.DB.Raw("SELECT password FROM users WHERE email=$1 OR phone_number=$2;", body.Email, body.PhoneNumber).Scan(&password).Error
	if err != nil {
		return "", err
	}
	if password != "" {
		return "", errors.New("Provided Email or phone number already occurred")
	}

	// Insert user into the database
	err = database.DB.Exec("INSERT INTO Users(name,email,phone_number) Values($1,$2,$3);", body.Name, body.Email, body.PhoneNumber).Error
	if err != nil {
		return "", err // Handle the error appropriately
	}

	// Generate password
	l := body.Lengthofpassword
	if l <= 0 {
		return "", errors.New("must provide the length of password greater than 0")
	}
	character := "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%&*"
	result := make([]byte, l)
	for i := range result {
		index := rand.Intn(len(character))
		result[i] = character[index]
	}
	results := string(result)

	// Hash the password
	encryptedpassword, err := bcrypt.GenerateFromPassword([]byte(results), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Handle the error appropriately
	}

	err = database.DB.Exec("UPDATE Users SET password=$1 WHERE email=$2;", encryptedpassword, body.Email).Error
	if err != nil {
		return "", err // Handle the error appropriately
	}

	return results, nil
}
func SignupNumber(body model.Signup) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var password string
	err := database.DB.Raw("SELECT password FROM users WHERE email=$1 OR phone_number=$2;", body.Email, body.PhoneNumber).Scan(&password).Error
	if err != nil {
		return "", err
	}
	if password != "" {
		return "", errors.New("Provided Email or phone number already occurred")
	}

	err = database.DB.Exec("INSERT INTO Users(name,email,phone_number) Values($1,$2,$3);", body.Name, body.Email, body.PhoneNumber).Error
	if err != nil {
		return "", err
	}

	l := body.Lengthofpassword
	if l <= 0 {
		return "", errors.New("must provide the length of password greater than 0")
	}
	character := "123456789"
	result := make([]byte, l)
	for i := range result {
		index := rand.Intn(len(character))
		result[i] = character[index]
	}
	results := string(result)

	encryptedpassword, err := bcrypt.GenerateFromPassword([]byte(results), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	err = database.DB.Exec("UPDATE Users SET password=$1 WHERE email=$2;", encryptedpassword, body.Email).Error
	if err != nil {
		return "", err
	}

	return results, nil
}

func SignupLowerCase(body model.Signup) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var password string
	err := database.DB.Raw("SELECT password FROM users WHERE email=$1 OR phone_number=$2;", body.Email, body.PhoneNumber).Scan(&password).Error
	if err != nil {
		return "", err
	}
	if password != "" {
		return "", errors.New("Provided Email or phone number already occurred")
	}

	err = database.DB.Exec("INSERT INTO Users(name,email,phone_number) Values($1,$2,$3);", body.Name, body.Email, body.PhoneNumber).Error
	if err != nil {
		return "", err // Handle the error appropriately
	}

	l := body.Lengthofpassword
	if l <= 0 {
		return "", errors.New("must provide the length of password greater than 0")
	}
	character := "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, l)
	for i := range result {
		index := rand.Intn(len(character))
		result[i] = character[index]
	}
	results := string(result)

	// Hash the password
	encryptedpassword, err := bcrypt.GenerateFromPassword([]byte(results), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Handle the error appropriately
	}

	err = database.DB.Exec("UPDATE Users SET password=$1 WHERE email=$2;", encryptedpassword, body.Email).Error
	if err != nil {
		return "", err // Handle the error appropriately
	}

	return results, nil
}
func SignupupperCase(body model.Signup) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var password string
	err := database.DB.Raw("SELECT password FROM users WHERE email=$1 OR phone_number=$2;", body.Email, body.PhoneNumber).Scan(&password).Error
	if err != nil {
		return "", err
	}
	if password != "" {
		return "", errors.New("Provided Email or phone number already occurred")
	}

	// Insert user into the database
	err = database.DB.Exec("INSERT INTO Users(name,email,phone_number) Values($1,$2,$3);", body.Name, body.Email, body.PhoneNumber).Error
	if err != nil {
		return "", err // Handle the error appropriately
	}

	// Generate password
	l := body.Lengthofpassword
	if l <= 0 {
		return "", errors.New("must provide the length of password greater than 0")
	}
	character := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, l)
	for i := range result {
		index := rand.Intn(len(character))
		result[i] = character[index]
	}
	results := string(result)

	// Hash the password
	encryptedpassword, err := bcrypt.GenerateFromPassword([]byte(results), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Handle the error appropriately
	}

	err = database.DB.Exec("UPDATE Users SET password=$1 WHERE email=$2;", encryptedpassword, body.Email).Error
	if err != nil {
		return "", err // Handle the error appropriately
	}

	return results, nil
}
func SignupSpecialcharacter(body model.Signup) (string, error) {
	rand.Seed(time.Now().UnixNano())
	var password string
	err := database.DB.Raw("SELECT password FROM users WHERE email=$1 OR phone_number=$2;", body.Email, body.PhoneNumber).Scan(&password).Error
	if err != nil {
		return "", err
	}
	if password != "" {
		return "", errors.New("Provided Email or phone number already occurred")
	}

	// Insert user into the database
	err = database.DB.Exec("INSERT INTO Users(name,email,phone_number) Values($1,$2,$3);", body.Name, body.Email, body.PhoneNumber).Error
	if err != nil {
		return "", err // Handle the error appropriately
	}

	// Generate password
	l := body.Lengthofpassword
	if l <= 0 {
		return "", errors.New("must provide the length of password greater than 0")
	}
	character := "@#$%&*"
	result := make([]byte, l)
	for i := range result {
		index := rand.Intn(len(character))
		result[i] = character[index]
	}
	results := string(result)

	// Hash the password
	encryptedpassword, err := bcrypt.GenerateFromPassword([]byte(results), bcrypt.DefaultCost)
	if err != nil {
		return "", err // Handle the error appropriately
	}

	err = database.DB.Exec("UPDATE Users SET password=$1 WHERE email=$2;", encryptedpassword, body.Email).Error
	if err != nil {
		return "", err // Handle the error appropriately
	}

	return results, nil
}
