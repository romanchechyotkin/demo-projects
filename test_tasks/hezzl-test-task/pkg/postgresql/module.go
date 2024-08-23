package postgresql

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/romanchechyotkin/hezzl-test-task/pkg/postgresql/schema"
)

func New(cfg *Config) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	log.Println("dsn", dsn)
	
	migrate(&schema.DB, dsn)

	return newClient(dsn)
}
