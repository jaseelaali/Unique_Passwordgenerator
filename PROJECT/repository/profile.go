package repository

import (
	"github.com/jaseelaali/passwordgenerator_unique/database"
	"github.com/jaseelaali/passwordgenerator_unique/model"
)

func ViewProfile(user_id int) (model.Users, error) {
	profile := model.Users{}
	err := database.DB.Raw("SELECT * FROM users WHERE id=$1", user_id).Scan(&profile)
	if err != nil {
		return profile, err.Error
	}
	return profile, nil
}

func Deleteprofile(user_id int) error {
	err := database.DB.Raw("DELETE FROM users WHERE id=$1", user_id).Scan(&model.Users{})
	if err != nil {
		return err.Error
	}
	return nil
}
