package models

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/mail"
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

type UpdateUserRequest struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}

func (req *UpdateUserRequest) ToBson() bson.M {
	m := bson.M{}
	if len(req.FirstName) > minFirstNameLen {
		m["firstName"] = req.FirstName
	}
	if len(req.LastName) > minLastNameLen {
		m["lastName"] = req.LastName
	}
	return m
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
	//standard secure way to check validity of email
	if _, err := mail.ParseAddress(params.Email); err != nil {
		return err
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
