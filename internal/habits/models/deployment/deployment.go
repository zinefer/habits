package deployment

import (
	"context"
	"time"

	"github.com/zinefer/habits/internal/habits/middlewares/database"
)

// Deployment model
type Deployment struct {
	Version string
	Created time.Time
}

// New Creates a Deployment model
func New(version string) *Deployment {
	return &Deployment{
		Version: version,
	}
}

// Save a Deployment to the database
func (d *Deployment) Save(ctx context.Context) error {
	db := database.GetDbFromContext(ctx)
	stmt, err := db.PrepareNamed("INSERT INTO deployments (version) VALUES (:version) RETURNING created")
	if err != nil {
		return err
	}
	return stmt.Get(&d.Created, d)
}

// FindByVersion returns a Deployment by it's version
func FindByVersion(ctx context.Context, version string) (*Deployment, error) {
	deployment := &Deployment{}
	db := database.GetDbFromContext(ctx)
	err := db.Get(deployment, "SELECT * FROM deployments WHERE version = $1 LIMIT 1", version)
	return deployment, err
}

// VersionPresent checks to see if a version is in the deployment table
func VersionPresent(ctx context.Context, version string) (bool, error) {
	db := database.GetDbFromContext(ctx)
	result := []Deployment{}
	err := db.Select(&result, "SELECT * FROM deployments WHERE version = $1 LIMIT 1", version)
	return len(result) == 1, err
}
