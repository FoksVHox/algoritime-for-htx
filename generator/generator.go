package generator

import (
	"fmt"
	"github.com/FoksVHox/algoritime-for-htx/generator/faker"
)

func Start() {
	fmt.Println("Called")
	People(5)
}

func People(amount int) []string {
	var people []string
	for i := 0; i < amount; i++ {
		people = append(people, faker.GetNames())
	}
	return people
}
