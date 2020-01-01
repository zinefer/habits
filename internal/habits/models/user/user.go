package user

import (
	"context"

	"github.com/zinefer/habits/internal/habits/middlewares/database"
)

// User model
type User struct {
	ID         int64
	ProviderID string `db:"provider_id"`
	Provider   string
	Name       string
	RealName   string `db:"real_name"`
	Email      string
}

// New Creates a User model
func New(providerID string, provider string, name string, realname string, email string) *User {
	return &User{
		ProviderID: providerID,
		Provider:   provider,
		Name:       name,
		RealName:   realname,
		Email:      email,
	}
}

// Save a User to the database
func (u *User) Save(ctx context.Context) error {
	db := database.GetDbFromContext(ctx)
	stmt, err := db.PrepareNamed("INSERT INTO users (provider_id, provider, name, real_name, email) VALUES (:provider_id, :provider, :name, :real_name, :email) RETURNING id")
	if err != nil {
		return err
	}
	return stmt.Get(&u.ID, u)
}

// FindByID returns a user by it's ID
func FindByID(ctx context.Context, userID int64) (*User, error) {
	user := &User{}
	db := database.GetDbFromContext(ctx)
	err := db.Get(user, "SELECT * FROM users WHERE id = $1 LIMIT 1", userID)
	return user, err
}

// IsNameAvailable checks to see if a username is already in use
func IsNameAvailable(ctx context.Context, name string) (bool, error) {
	db := database.GetDbFromContext(ctx)
	result := []User{}
	err := db.Select(&result, "SELECT * FROM users WHERE name = $1 LIMIT 1", name)
	return len(result) != 1, err
}
