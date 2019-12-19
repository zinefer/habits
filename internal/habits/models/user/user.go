package user

import (
	"context"

	"database/sql"

	"github.com/zinefer/habits/internal/habits/middlewares"
)

// User model
type User struct {
	Name     string
	NickName string
	Email    string
	Provider string
}

// New Creates a User model
func New(name string, nickname string, email string, provider string) *User {
	return &User{
		Name:     name,
		NickName: nickname,
		Email:    email,
		Provider: provider,
	}
}

// Save a User to the database
func (u *User) Save(ctx context.Context) (sql.Result, error) {
	db := middlewares.GetDbFromContext(ctx)
	return db.NamedExec("INSERT INTO users (name, nickname, email, provider) VALUES (:name, :nickname, :email, :provider)", u)
}
