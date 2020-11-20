package connector

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Start() *sql.DB {
	fmt.Println("Booting")
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err.Error)
	}

	fmt.Println("Connected shit")

	return db
}

func CreateUserTable(db *sql.DB) {
	_, err := db.Query("create table if not exists users(id int auto_increment,name varchar(255) not null,constraint users_pk primary key (id));")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("[MySQL] Asserted Users table.")
}

func CreateUser(qu string, db *sql.DB) {
	_, err := db.Query(qu)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("[MySQL] Inserted 1 row to users table.")
}

func CreateFriendsTable(db *sql.DB) {
	_, err := db.Query("create table if not exists friends(id int auto_increment,user_id int not null,friend_id int not null,constraint friends_pk primary key (id),constraint friends_users_id_fk foreign key (user_id) references users (id) on delete cascade, constraint friends_users_id_fk_2 foreign key (friend_id) references users (id) on delete cascade);")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("[MySQL] Asserted Friends table.")
}

func Close(db *sql.DB) {
	db.Close()
	fmt.Println("Worked")
}

func GetUsers(db *sql.DB) *sql.Rows {
	res, err := db.Query("SELECT * from users;")

	if err != nil {
		panic(err.Error())
	}

	return res
}

func GetFriends(db *sql.DB, id int) *sql.Rows {
	res, err := db.Query("SELECT * from friends WHERE user_id = ?;", id)

	if err != nil {
		panic(err.Error())
	}

	return res
}
