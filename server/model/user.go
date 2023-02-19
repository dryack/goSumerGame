package model

import (
	"goSumerGame/server/crypto"
	"goSumerGame/server/database"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Games    []Game
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	params := crypto.Params{
		Memory:      256 * 1024,
		Iterations:  5,
		Parallelism: 2,
		SaltLength:  32,
		KeyLength:   32,
	}
	passwordHash, err := crypto.GenerateFromPassword(user.Password, &params)
	if err != nil {
		return err
	}
	user.Password = passwordHash
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) (bool, error) {
	return crypto.ComparePasswordAndHash(password, user.Password)
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Preload("Games").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
