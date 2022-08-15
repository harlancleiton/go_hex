package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        string
	firstName string
	lastName  string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(id string, firstName string, lastName string, email string, password string, createdAt time.Time, updatedAt time.Time) *User {
	if id == "" {
		id = uuid.New().String()
	}

	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) ChangeEmail(newEmail string) error {
	if u.email == newEmail {
		return errors.New("Email is the same")
	}

	u.email = newEmail

	return nil
}

func (u *User) ChangePassword(newPassword string) error {
	if u.password == newPassword {
		return errors.New("Password is the same")
	}

	u.password = newPassword

	return nil
}

func (u *User) Id() string {
	return u.id
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}
