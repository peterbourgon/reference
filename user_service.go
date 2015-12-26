package main

import (
	"errors"
	"time"
)

// UserService provides basic user and session management.
type UserService interface {
	Create(emailAddress, realName, plaintextPassword string) error
	Get(emailAddress string) (User, error)
	//Update(emailAddress, realName string) error
	//ChangePassword(emailAddress, oldPassword, newPassword string) error
	Delete(emailAddress string) error
	//Login(emailAddress, password string) (Session, error)
	//Logout(sessionID string) error
}

// User is a user object. The email address is the primary key.
type User struct {
	EmailAddress string
	RealName     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Session is a session object.
//type Session struct {
//	ID           string
//	EmailAddress string
//	CreatedAt    time.Time
//	ExpiresAt    time.Time
//}

var (
	// ErrEmailAddressAlreadyExists is returned when a duplicate user is
	// created.
	ErrEmailAddressAlreadyExists = errors.New("email address already exists")

	// ErrEmailAddressNotFound is returned when mutations are applied to a
	// user account that doesn't exist.
	ErrEmailAddressNotFound = errors.New("email address not found")
)
