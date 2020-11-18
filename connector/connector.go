package connector

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Start() {
	fmt.Println("Booting")
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error)
	}

	defer db.Close()

	fmt.Println("Connected shit")
}
