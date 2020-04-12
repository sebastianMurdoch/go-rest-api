package persistence

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

/* Configure Postgres DB */
var (
	err                error
	db                 *sqlx.DB
	maxIdleConnections = 25
	maxOpenConnections = 100
	connMaxLifetime    = 200 * time.Millisecond
)

func NewDB() *sqlx.DB {
	db, err = sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err.Error())
		db = sqlx.MustConnect("sqlite3", "file::memory:?cache=shared")
	}

	schema := `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY ,
    userName VARCHAR (255) UNIQUE NOT NULL);`

	// execute a query on the server
	db.MustExec(schema)

	return db
}
