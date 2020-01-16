package user

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/zinefer/habits/internal/habits/middlewares/database"

	"github.com/zinefer/habits/internal/habits/models/habit"
)

const cleanNameRegex = "[^a-z0-9-_]+"

// User model
type User struct {
	ID         int64
	ProviderID string `db:"provider_id"`
	Provider   string
	Name       string
	RealName   string `db:"real_name"`
	Email      string
	Created    time.Time
}

// New Creates a User model
func New(providerID string, provider string, name string, realname string, email string) *User {
	cleanName := strings.ToLower(name)
	cleanName = strings.ReplaceAll(cleanName, " ", "-")
	reg := regexp.MustCompile(cleanNameRegex)
	cleanName = reg.ReplaceAllString(cleanName, "")
	return &User{
		ProviderID: providerID,
		Provider:   provider,
		Name:       cleanName,
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

// GetHabits returns a list of habits for a user
func (u *User) GetHabits(ctx context.Context) ([]*habit.Habit, error) {
	return habit.FindAllByUser(ctx, u.ID)
}

// FindByProviderID returns a user by it's provider and providerID
func FindByProviderID(ctx context.Context, provider string, providerID string) (*User, error) {
	user := &User{}
	db := database.GetDbFromContext(ctx)
	err := db.Get(user, "SELECT * FROM users WHERE provider = $1 AND provider_ID = $2 LIMIT 1", provider, providerID)
	return user, err
}

// FindByID returns a user by it's ID
func FindByID(ctx context.Context, userID int64) (*User, error) {
	user := &User{}
	db := database.GetDbFromContext(ctx)
	err := db.Get(user, "SELECT * FROM users WHERE id = $1 LIMIT 1", userID)
	return user, err
}

// FindByName returns a user by it's Name
func FindByName(ctx context.Context, userName string) (*User, error) {
	user := &User{}
	db := database.GetDbFromContext(ctx)
	err := db.Get(user, "SELECT * FROM users WHERE name = $1 LIMIT 1", userName)
	return user, err
}

// IsNameAvailable checks to see if a username is already in use
func IsNameAvailable(ctx context.Context, name string) (bool, error) {
	db := database.GetDbFromContext(ctx)
	result := []User{}
	err := db.Select(&result, "SELECT * FROM users WHERE name = $1 LIMIT 1", name)
	return len(result) != 1, err
}
