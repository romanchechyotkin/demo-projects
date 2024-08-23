package postgresql_test

import (
	"context"
	"os"
	"testing"

	"github.com/romanchechyotkin/hezzl-test-task/pkg/postgresql"
)

func TestPostgres(t *testing.T) {
	cfg := postgresql.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_DATABASE"),
	}

	pgClient := postgresql.New(&cfg)

	t.Run("create project success", func(t *testing.T) {
		err := postgresql.CreateProject(pgClient, context.Background(), "test 1")
		if err != nil {
			t.Fatal()
		}
	})

}
