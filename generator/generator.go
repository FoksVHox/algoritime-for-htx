package generator

import (
	"fmt"
	"github.com/FoksVHox/algoritime-for-htx/generator/faker"
	"io/ioutil"
)

func Start() {
	fmt.Println("Called")
	people(21)
}

func people(amount uint) {
	ioutil.WriteFile("test.txt", []byte("helplessly"), 0777)
	faker.GetNames()
}
