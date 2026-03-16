package repository

import (
	"errors"

	"github.com/InkSha/Resummer/internal/globals"
	"github.com/InkSha/Resummer/internal/service/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrUserExists = errors.New("account already exists")

func UserRegister(account string, password string) (*model.User, error) {
	db := globals.GetDB()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := model.User{
		Account:  account,
		Password: string(hash),
	}

	if err := db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrUserExists
		}

		return nil, err
	}

	return &user, nil
}
