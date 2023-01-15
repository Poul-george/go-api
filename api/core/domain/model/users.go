package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Name        string
	Password    string
	MailAddress string
	Comments    string
}

func NewUser(
	name string,
	password string,
	mailAddress string,
	comments string,
) (*User, error) {
	//password生成
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		Name:        name,
		Password:    string(pass),
		MailAddress: mailAddress,
		Comments:    comments,
	}

	return &user, nil
}
