package persistence

import (
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jmoiron/sqlx"
)

var (
	db                 *sqlx.DB
	maxIdleConnections = 25
	maxOpenConnections = 100
	connMaxLifetime    = 200 * time.Millisecond
)

func InitDb() {
	var err error
	var dbType string
	var connectionString string

	dbType = "sqlite3"
	connectionString = "file::memory:?cache=shared"

	db, err = sqlx.Open(dbType, connectionString)
	if err != nil {
		panic("Cannot open DB")
	}

	sqlStmt := `
	CREATE TABLE admissions (id integer not null primary key, user_id varchar(250), brand_id varchar(250));`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("ERRORRR")
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO admissions(user_id, brand_id) values(?, ?)")
	if err != nil {
		fmt.Println("ERRORRR")
		log.Fatal(err)
	}
	_, err = stmt.Exec("10", "1")
	if err != nil {
		fmt.Println("ERROR")
	}
	defer stmt.Close()
	tx.Commit()
}

func DB() (*sqlx.DB, error) {
	err := db.Ping()

	if err != nil {
		return db, err
	}
	return db, nil
}
