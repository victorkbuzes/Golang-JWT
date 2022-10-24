package models

import (
	"gojwt/utils/token"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func VerifyPassword(password, harashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(harashedPassword), []byte(password))

}

func LoginCheck(username string, password string) (string, error) {

	var err error

	user := User{}

	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err

	}
	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err

	}
	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}
	return token, nil

}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(*gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
