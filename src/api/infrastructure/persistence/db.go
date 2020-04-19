package persistence

import (
	"fmt"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

/* Returns a postgres db with all tables already configured */
func NewDB() *sqlx.DB {
	db, err = sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))

	/* Used for localhost tests */
	if err != nil {
		db = sqlx.MustConnect("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname))
	}

	schema := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username VARCHAR (255) NOT NULL);`

	// execute a query on the server
	db.MustExec(schema)

	return db
}
