package database

import (
	"database/sql"
	"fmt"
	"reflect"
)

func AddUsers() {
	database, _ := sql.Open("sqlite3", "./database/users.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (email TEXT PRIMARY KEY, username TEXT, password TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO users (email, username, password) VALUES (?, ?, ?)")
	insertRow(statement, userInfo{"nesrine@", "nesrineTHEBEST", "password"})
	rows, _ := database.Query("SELECT * FROM users WHERE username = 'nesrineTHEBEST'")

	for rows.Next() {
		fmt.Println(rows)
	}
}

func insertRow(db *sql.Stmt, row interface{}) {
	rv := reflect.ValueOf(row)
	var args []interface{}
	for i := 0; i < rv.NumField(); i++ {
		args = append(args, rv.Field(i).Interface())
	}
	db.Exec(args...)
}
