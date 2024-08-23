package postgresql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateProject(pool *pgxpool.Pool, ctx context.Context, name string) error {
	exec, err := pool.Exec(ctx, "insert into projects (name) values ($1)", name)
	if err != nil {
		return err
	}

	log.Println("rows affected", exec.RowsAffected())

	return nil
}

func newClient(dsn string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("database connected")

	return pool
}
