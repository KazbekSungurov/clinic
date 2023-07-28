package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost:5432/clinic?sslmode=enable")
	if err != nil {
		log.Fatal(err)
	}
	//driver, err := postgres.WithInstance(db, &postgres.Config{})
	//m, err := migrate.NewWithDatabaseInstance(
	//	"schema",
	//	"clinic", driver)
	//err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	//if err != nil {
	//	fmt.Print(err)
	//}

	err = Migrate(db)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	migrate_path = "file:///Users//m0the//GolandProjects//clinic//clinic//schema"
	sqlite3_path = `postgres://postgres:postgres@localhost:5432/clinic?sslmode=disable`
)

func Migrate(db *sql.DB) error {

	m, err := migrate.New(migrate_path,
		sqlite3_path)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		fmt.Println(err)

	} else if err != nil {
		return err
	}

	return nil
}
