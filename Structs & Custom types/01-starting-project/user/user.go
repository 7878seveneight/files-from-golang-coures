package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	birthdate   string
	createdTime time.Time
}

type Admin struct {
	email       string
	password    string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "---",
			createdTime: time.Now(),
		},
	}
}

// constractor func (usually name of the constractor starts with "New...")
func New(firstName, lastName, birthdate string) (*User, error) {
	//validation
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate is required")
	}

	return &User{
		firstName:   firstName,
		lastName:    lastName,
		birthdate:   birthdate,
		createdTime: time.Now(),
	}, nil
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
