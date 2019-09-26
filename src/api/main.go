package main

import (
	"github.com/smurdoch/my_simple_rest_api/src/api/config/wire"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := wire.InitializeApp()
	app.Start()
}
