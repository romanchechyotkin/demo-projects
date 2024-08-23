package clickhouse

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mailru/go-clickhouse"
)

func NewClient(cfg *Config) *sql.DB {
	connStr := fmt.Sprintf("http://%s:%s/?user=%s&password=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password)
	driver := "clickhouse"
	connect, err := sql.Open(driver, connStr)
	if err != nil {
		log.Fatalf("Open >> %v", err)
	}

	if err := connect.Ping(); err != nil {
		log.Fatalf("Ping >> %v", err)
	}

	var dbName string
	err = connect.QueryRow("SELECT currentDatabase()").Scan(&dbName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("clickhouse current database:", dbName)

	return connect
}
