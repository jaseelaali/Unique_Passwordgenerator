package handlers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jaseelaali/passwordgenerator_unique/model"
	"github.com/jaseelaali/passwordgenerator_unique/repository"
)

func SignUpPassword(r *gin.Context) {
	body := model.Signup{}
	if err := r.Bind(&body); err != nil {
		r.JSON(400, gin.H{
			"message": errors.New("error in binding data"),
		})
		return
	}
	result, err := repository.SignupPassword(body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	data := "User successfully addedd..." + "your password is " + result
	r.JSON(400, gin.H{
		"message": data,
	})

}
func SignUpNumber(r *gin.Context) {
	body := model.Signup{}
	if err := r.Bind(&body); err != nil {
		r.JSON(400, gin.H{
			"message": errors.New("error in binding data"),
		})
		return
	}
	result, err := repository.SignupNumber(body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := "User successfully addedd..." + "your password is " + result
	r.JSON(400, gin.H{
		"message": data,
	})

}
func SignUpLowerCase(r *gin.Context) {
	body := model.Signup{}
	if err := r.Bind(&body); err != nil {
		r.JSON(400, gin.H{
			"message": errors.New("error in binding data"),
		})
		return
	}
	result, err := repository.SignupLowerCase(body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println(result)
	data := "User successfully addedd..." + "your password is " + result
	r.JSON(400, gin.H{
		"message": data,
	})

}
func SignUpUpperCase(r *gin.Context) {
	body := model.Signup{}
	if err := r.Bind(&body); err != nil {
		r.JSON(400, gin.H{
			"message": errors.New("error in binding data"),
		})
		return
	}
	result, err := repository.SignupupperCase(body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	data := "User successfully addedd..." + "your password is " + result
	r.JSON(400, gin.H{
		"message": data,
	})

}
func SignUpSpecialCharacter(r *gin.Context) {
	body := model.Signup{}
	if err := r.Bind(&body); err != nil {
		r.JSON(400, gin.H{
			"message": errors.New("error in binding data"),
		})
		return
	}
	result, err := repository.SignupSpecialcharacter(body)
	if err != nil {
		r.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	data := "User successfully addedd..." + "your password is " + result
	r.JSON(400, gin.H{
		"message": data,
	})

}
