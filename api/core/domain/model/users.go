package model

import (
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ExternalUserID identifier.ExternalUserID
	Name           string
	Password       string
	MailAddress    string
	Comments       string
	UserDetail     UserDetail
}

type UserDetail struct {
	UserID               *identifier.UserID
	Name                 *string
	NameKana             *string
	Profession           string // 職業
	Occupation           string // 職種
	CountryOfCitizenship *uint8 // 国籍
	Age                  *uint8
	BirthDate            *string
	Gender               *uint8
	Height               uint16
	BirthPlace           string
	Residence            string // 居住地
	WorkLocation         string
	AnnualIncome         uint64
	TitleComment         string
	ProfileComment       string
}

func NewUser(
	externalUserID string,
	name string,
	password string,
	mailAddress string,
	comments string,
) (*User, error) {
	//password生成
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		ExternalUserID: identifier.ExternalUserID(externalUserID),
		Name:           name,
		Password:       string(pass),
		MailAddress:    mailAddress,
		Comments:       comments,
	}

	return &user, nil
}
