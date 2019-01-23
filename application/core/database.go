package core

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/source/file"

	_ "github.com/lib/pq"
)

var (
	dbc Database
)

func init() {
	GetConnection()
}

func GetConnection() *sql.DB {

	dbc = GetDatabaseConfiguration()

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.Name)
	db, err := sql.Open("postgres", dbInfo)
	db.SetMaxOpenConns(10)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfully connected to database")
	execStartupDBScripts(db)
	return db
}

func execStartupDBScripts(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err == nil {
		m, err := migrate.NewWithDatabaseInstance(
			"file://home-task-tracker/migrations/",
			"htt_db", driver)
		if err == nil {
			err := m.Steps(1)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
	log.Print("db initialized")
}
