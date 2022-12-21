package main

import (
	"github.com/ridhompra/crud-fiber/models"
	"github.com/ridhompra/crud-fiber/router"
)

func main() {
	models.ConnectDB()
	router.Router()
}
