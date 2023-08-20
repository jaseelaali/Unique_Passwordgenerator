package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jaseelaali/passwordgenerator_unique/database"
	"github.com/jaseelaali/passwordgenerator_unique/model"
	"github.com/jaseelaali/passwordgenerator_unique/repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(r *gin.Context) {
	var login struct {
		Email    string
		Password string
	}
	if err := r.Bind(&login); err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"message": "error in binding data"})
		return
	}

	user := &model.Users{}
	result := database.DB.Where(&model.Users{Email: login.Email}).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		r.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user",
		})
		return
	}

	password := user.Password

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(login.Password))
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"message": "passwords are not matching"})
		return
	}

	token, err := repository.Token(login.Email)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"message": "unable to create token"})
		return
	}

	//tokenString, err := token.SignedString([]byte("AVBBVVXSJBJKQNKBNSNBX"))
	tokestring := []byte(token)
	tokenString := string(tokestring)
	if err != nil {
		r.JSON(http.StatusBadRequest, gin.H{"message": "unable to create token"})
		return
	}

	r.SetSameSite(http.SameSiteLaxMode)
	r.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	r.JSON(http.StatusOK, gin.H{
		"message": "login successfully... token: " + tokenString,
	})
}
func Profile(r *gin.Context) {
	user_id := GetId(r)
	profile, err := repository.ViewProfile(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	r.JSON(200, gin.H{
		"INFO": profile,
	})

}
func GetId(r *gin.Context) int {
	temp := fmt.Sprint(r.Get("user_id"))
	id := strings.Split(temp, " ")
	Id, _ := strconv.Atoi(id[0])
	return Id
}
func DeleteProfile(r *gin.Context) {
	user_id := GetId(r)
	err := repository.Deleteprofile(user_id)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return

	}
	r.JSON(200, gin.H{
		"message": "profile deleted",
	})
}
