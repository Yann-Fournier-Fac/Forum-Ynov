package main

import (
	"main/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.AddUsers()
}
