package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaseelaali/passwordgenerator_unique/database"
	"github.com/jaseelaali/passwordgenerator_unique/model"
	"github.com/sfreiberg/gotwilio"
	"golang.org/x/crypto/bcrypt"
)

var OTP int
var Email string
var twilio = gotwilio.NewTwilioClient("AC3037e122f46a35ae97b5a48f7413be56", "0bc43f0b4e4a492d46e26bf093c0fc40")

func ChangePassword(r *gin.Context) {
	const otpValidityDuration = 1 * time.Minute
	// user_id, _ := r.Get("user_id")
	// User_Id, _ := strconv.Atoi(fmt.Sprint(user_id))
	Email = r.Query("email")
	var mobilenumber string
	err := database.DB.Raw("SELECT phone_number FROM users WHERE email=$1;", Email).Scan(&mobilenumber)
	if err.Error != nil {
		r.JSON(400, gin.H{
			"message": err.Error,
		})
		return
	}
	otp, Error := sendOTP(mobilenumber)
	OTP = otp
	if Error != nil {
		r.JSON(400, gin.H{
			"message": "failed to send message",
			"error":   Error,
		})
		return
	}
	r.JSON(200, gin.H{
		"message": "successfully send the otp",
		"data":    OTP,
	})
}
func sendOTP(phoneNumber string) (int, error) {
	otpCode := generateOTP()
	fmt.Println(otpCode)
	message := "Your OTP code is " + strconv.Itoa(otpCode)
	_, _, err := twilio.SendSMS("+15302036484", "+91"+phoneNumber, message, "", "")
	if err != nil {
		return 0, err
	}
	return otpCode, nil
}
func generateOTP() int {
	// Generate a random 4-digit OTP code
	otp := rand.Intn(8899) + 1000
	fmt.Println(otp)
	return otp
}
func VerifyOtp(r *gin.Context) {
	Otp, err := strconv.Atoi(r.Query("otp"))
	length, err := strconv.Atoi(r.Query("length"))
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "error in query...",
		})
		return
	}

	isValid := VerifyOTP(Otp, OTP) // Should be "Otp" (lowercase) instead of "OTP"
	if !isValid {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid OTP",
		})
		return
	}

	if length <= 0 {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "must provide a password length greater than 0",
		})
		return
	}

	character := "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	result := make([]byte, length)
	for i := range result {
		index := rand.Intn(len(character))
		result[i] = character[index]
	}
	results := string(result)

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(results), 11)
	if err != nil {
		r.JSON(400, gin.H{
			"message": "unable to generate password",
		})
		return
	}

	// Assuming you have the "Email" variable defined somewhere in the code
	//qw := database.DB.Exec("UPDATE Users SET password=$1 WHERE email=$2;", encryptedPassword, Email)
	//.Scan(&model.Users{})

	// if qw != nil {
	// 	r.JSON(400, gin.H{
	// 		"message": qw,
	// 	})
	// 	return
	// }
	database.DB.Exec("UPDATE Users SET password=$1 WHERE email=$2;", encryptedPassword, Email).Scan(&model.Users{})

	r.JSON(200, gin.H{
		"message": "your new password is: " + results,
	})
}

func VerifyOTP(otpCode, expectedCode int) bool {
	if otpCode == expectedCode {
		return true
	}
	return false
}
