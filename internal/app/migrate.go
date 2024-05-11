package app

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)



func RunMigrate(db *sql.DB) {
	databaseUrl, ok := os.LookupEnv("URL_DB")
	fmt.Println(databaseUrl)
	if !ok || len(databaseUrl) == 0 {
		log.Fatal("environtment variable not declare: URL_DB")
	}

	

	databaseUrl += "?sslmode=disable"

	var(
		err error
		m *migrate.Migrate
	)

	driver, err  := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Migrate -  mysql.WithInstance: %v", err)
	}

    m, err = migrate.NewWithDatabaseInstance(
        "file://migrations",
        "mysql", 
        driver,
    )

	if err != nil {
		log.Fatalf("Migrate - migrate.NewWithDatabaseInstance: %v", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}