package main

import (
	"github.com/smurdoch/go-rest-api/src/api/config/wire"
	"github.com/smurdoch/go-rest-api/src/api/persistence"
)

func main() {
	persistence.InitDb()
	app := wire.InitializeApp()
	app.Start()
}
