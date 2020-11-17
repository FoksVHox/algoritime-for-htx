package names

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Get() []string {
	var a []string

	// Read the files names
	firstNamesFile, fnerr := ioutil.ReadFile("data/first-names.txt")
	lastNamesFile, lnerr := ioutil.ReadFile("data/last-names.txt")
	if fnerr != nil {
		log.Fatal(fnerr)
	}
	if lnerr != nil {
		log.Fatal(lnerr)
	}

	firstNames := fileExtract(string(firstNamesFile))
	for _, name := range firstNames {
		fmt.Println("New name: " + name + " has been registered")
	}

	lastNames := fileExtract(string(lastNamesFile))
	for _, name := range lastNames {
		fmt.Println("New name: " + name + " has been registered")
	}

	return a
}

func fileExtract(content string) []string {
	return strings.Split(content, "\r\n")
}
