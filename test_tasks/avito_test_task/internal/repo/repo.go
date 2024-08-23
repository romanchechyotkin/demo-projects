package repo

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/romanchechyotkin/avito_test_task/internal/entity"
	"github.com/romanchechyotkin/avito_test_task/internal/repo/flat"
	"github.com/romanchechyotkin/avito_test_task/internal/repo/house"
	"github.com/romanchechyotkin/avito_test_task/internal/repo/user"
	"github.com/romanchechyotkin/avito_test_task/pkg/postgresql"
)

type User interface {
	CreateUser(ctx context.Context, user *entity.User) (string, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetById(ctx context.Context, id int) (*entity.User, error)
}

type House interface {
	CreateHouse(ctx context.Context, house *entity.House) (*entity.House, error)
	CreateSubscription(ctx context.Context, houseID, userID string) error
	GetHouseSubscriptions(ctx context.Context, houseID uint) ([]string, error)
}

type Flat interface {
	CreateFlat(ctx context.Context, flat *entity.Flat) (*entity.Flat, error)
	GetStatus(ctx context.Context, id uint) (string, error)
	UpdateStatus(ctx context.Context, flat *entity.Flat, userID sql.NullString) (*entity.Flat, error)
	GetHouseFlats(ctx context.Context, houseID, userType string) ([]*entity.Flat, error)
}

type Repositories struct {
	User
	House
	Flat
}

func NewRepositories(log *slog.Logger, pg *postgresql.Postgres) *Repositories {
	return &Repositories{
		User:  user.NewRepo(log, pg),
		House: house.NewRepo(log, pg),
		Flat:  flat.NewRepo(log, pg),
	}
}
