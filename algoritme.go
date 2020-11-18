package main

import (
	"fmt"
	"github.com/FoksVHox/algoritime-for-htx/connector"
	"github.com/FoksVHox/algoritime-for-htx/generator"
)

func main() {
	fmt.Println("============================")
	fmt.Println("= WELCOME TO HTX ALGORITHM =")
	fmt.Println("= Version: beta            =")
	fmt.Println("= Created by: Jimmi Hansen =")
	fmt.Println("============================")
	fmt.Println("Ensuring all tables is in the database.")
	connector.Start()
	fmt.Println("Calling Generator")
	generator.Start()
}
