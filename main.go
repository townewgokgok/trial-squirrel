package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//sq "github.com/Masterminds/squirrel"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	

	//rows, err := db.Query("SELECT * FROM users")
	//if err != nil {
	//	panic(err.Error())
	//}
}
