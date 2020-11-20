package main

import (
	"database/sql"
	"fmt"
	"github.com/FoksVHox/algoritime-for-htx/connector"
	"github.com/FoksVHox/algoritime-for-htx/generator"
	"time"
)

var db *sql.DB

func main() {
	fmt.Println("============================")
	fmt.Println("= WELCOME TO HTX ALGORITHM =")
	fmt.Println("= Version: beta            =")
	fmt.Println("= Created by: Jimmi Hansen =")
	fmt.Println("============================")
	fmt.Println("Ensuring all tables is in the database.")
	db = connector.Start()
	connector.CreateUserTable(db)
	connector.CreateFriendsTable(db)
	fmt.Println("Calling Generator")
	generator.Start()
	fmt.Println("Close Database")
	algorithm()
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func algorithm() {
	var person = generator.People(1)[0]
	connector.CreateUser("INSERT INTO users (name) VALUES ('"+person+"');", db)

	// Make friends? Yes
	var users = connector.GetUsers(db)
	for users.Next() {
		var user User
		// for each row, scan the result into our tag composite object
		err := users.Scan(&user.ID, &user.Name)
		if err != nil {
			panic(err.Error())
		}
		// and then print out the tag's Name attribute
		var friends = connector.GetFriends(db, user.ID)

		if checkCount(friends) != 0 {
			fmt.Println("here")
		} else {
			for i := 0; i < checkCount(friends); i++ {
				var friend User
				err := friends.Scan(&friend.ID, &friend.Name)
				if err != nil {
					panic(err.Error())
				}
				fmt.Println(friend)
			}

			//random.RangeInt(0, checkCount(friends)-1, 1)[0]
			//connector.AddFriend(user.ID, friends[].ID, db)
		}
	}

	time.Sleep(time.Second / 2)
	//algorithm()
}

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		panic(err.Error())
	}
	return count
}
