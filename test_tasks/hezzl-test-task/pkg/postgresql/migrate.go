package postgresql

import (
	"embed"
	"log"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func migrate(fs *embed.FS, dbUrl string) {
	source, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Printf("Failed to read migrations source: %v\n", err)

		return
	}

	instance, err := gomigrate.NewWithSourceInstance("iofs", source, dbUrl)
	if err != nil {
		log.Printf("Failed to initialization the migrate instance: %v\n", err)

		return
	}

	err = instance.Up()

	switch err {
	case nil:
		log.Println("The migration schema: The schema successfully upgraded!")
	case gomigrate.ErrNoChange:
		log.Println("The migration schema: The schema not changed")
	default:
		log.Printf("Could not apply the migration schema: %v\n", err)
	}
}
