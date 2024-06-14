package user

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
)

type Repo interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id uint) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

var (
	ErrUserNotFound    = errors.New("User not found")
	ErrInvalidPassword = errors.New("Invalid user password")
)

type UserRole uint8

func (ur UserRole) String() string {
	switch ur {
	case UserRoleUser:
		return "user"
	case UserRoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
}

const (
	UserRoleUser UserRole = iota + 1
	UserRoleAdmin
)

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      UserRole
}

func (u *User) ValidatePassword() error {
	return nil
}

func (u *User) PasswordIsValid(pass string) bool {
	passSha256 := sha256.Sum256([]byte(pass))
	return bytes.Equal(passSha256[:], []byte(u.Password))
}
