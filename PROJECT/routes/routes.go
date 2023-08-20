package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jaseelaali/passwordgenerator_unique/handlers"
	"github.com/jaseelaali/passwordgenerator_unique/middleware"
)

func User(r *gin.Engine) {
	r.POST("/password", handlers.SignUpPassword)
	r.POST("/usersignupwithnumberpassword", handlers.SignUpNumber)
	r.POST("/usersignupwithlowercasealphabetpassword", handlers.SignUpLowerCase)
	r.POST("/usersignupwithuppercasealphabetpassword", handlers.SignUpUpperCase)
	r.POST("/usersignupwithspeacialcharacterpassword", handlers.SignUpSpecialCharacter)
	r.PATCH("/changepassword", handlers.ChangePassword)
	r.GET("/verifyotp", handlers.VerifyOtp)
	r.POST("/login", handlers.Login)
	r.GET("/profile",middleware.RequiredAuthentication,handlers.Profile)
	r.DELETE("/deleteprofile",middleware.RequiredAuthentication,handlers.DeleteProfile)


}
