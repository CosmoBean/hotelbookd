package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	minPasswordLen  = 7
	minFirstNameLen = 2
	minLastNameLen  = 2
)

type User struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` //bson:omitempty will auto create id in mongo
	FirstName     string             `bson:"firstName" json:"firstName"`
	LastName      string             `bson:"lastName" json:"lastName"`
	Email         string             `bson:"email" json:"email"`
	EncryptedPass string             `bson:"EncryptedPass" json:"-"`
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params *CreateUserRequest) Validate() error {
	if len(params.FirstName) < minFirstNameLen {
		return fmt.Errorf("firstName length should be at least %d charecter", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		return fmt.Errorf("lastName length should be at least %d charecter", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		return fmt.Errorf("password length should be at least %d charecter", minPasswordLen)
	}
	if !isValidEmail(params.Email) {
		return fmt.Errorf("invalid user email")
	}
	return nil
}

func isValidEmail(s string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$`)
	return emailRegex.MatchString(s)
}

func NewUserFromParams(params CreateUserRequest) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:     params.FirstName,
		LastName:      params.LastName,
		Email:         params.Email,
		EncryptedPass: string(encpw),
	}, nil
}
