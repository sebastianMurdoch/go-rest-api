package persistence

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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
	db = sqlx.MustConnect("postgres", os.Getenv("DATABASE_URL"))

	schema := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username VARCHAR (255) NOT NULL);`

	// execute a query on the server
	db.MustExec(schema)

	return db
}
