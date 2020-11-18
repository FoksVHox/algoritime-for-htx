package generator

import (
	"fmt"
	"github.com/FoksVHox/algoritime-for-htx/generator/faker"
)

func Start() {
	fmt.Println("Called")
	people(5)
}

func people(amount int) {
	var people string = ""
	for i := 0; i < amount; i++ {
		people += faker.GetNames() + ","
	}
	fmt.Println(people)
}
