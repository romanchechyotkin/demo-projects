package service

import (
	"context"
	"log/slog"
	"time"

	"github.com/romanchechyotkin/avito_test_task/internal/entity"
	"github.com/romanchechyotkin/avito_test_task/internal/repo"
)

type Sender interface {
	Send() chan<- uint
}

type AuthCreateUserInput struct {
	Email    string
	Password string
	UserType string
}

type AuthGenerateTokenInput struct {
	Email    string
	Password string
}

type Auth interface {
	CreateUser(ctx context.Context, input *AuthCreateUserInput) (string, error)

	GenerateToken(ctx context.Context, input *AuthGenerateTokenInput) (string, error)

	ParseToken(accessToken string) (*TokenClaims, error)
}

type HouseCreateInput struct {
	Address   string
	Year      uint
	Developer string
}

type GetHouseFlatsInput struct {
	HouseID  string
	UserType string
}

type CreateSubscriptionInput struct {
	HouseID string
	UserID  string
}

type House interface {
	CreateHouse(ctx context.Context, input *HouseCreateInput) (*entity.House, error)
	GetHouseFlats(ctx context.Context, input *GetHouseFlatsInput) ([]*entity.Flat, error)
	CreateSubscription(ctx context.Context, input *CreateSubscriptionInput) error
}

type FlatCreateInput struct {
	Number      uint
	HouseID     uint
	Price       uint
	RoomsAmount uint
}

type FlatUpdateInput struct {
	ID          uint
	Status      string
	ModeratorID string
}

type Flat interface {
	CreateFlat(ctx context.Context, input *FlatCreateInput) (*entity.Flat, error)
	UpdateFlat(ctx context.Context, input *FlatUpdateInput) (*entity.Flat, error)
}

type Dependencies struct {
	Log   *slog.Logger
	Repos *repo.Repositories

	SignKey  string
	TokenTTL time.Duration
}

type Services struct {
	Auth   Auth
	House  House
	Flat   Flat
	Sender Sender
}

func NewServices(deps *Dependencies) *Services {
	sender := NewSenderService(deps.Log, deps.Repos.House)

	return &Services{
		Auth:   NewAuthService(deps.Log, deps.Repos.User, deps.SignKey, deps.TokenTTL),
		House:  NewHouseService(deps.Log, deps.Repos.House, deps.Repos.Flat),
		Flat:   NewFlatService(deps.Log, sender, deps.Repos.Flat),
		Sender: sender,
	}
}
