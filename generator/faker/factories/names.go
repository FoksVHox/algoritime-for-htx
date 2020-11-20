package names

import (
	"github.com/alok87/goutils/pkg/random"
	"io/ioutil"
	"log"
	"strings"
)

var lastnames []string
var firstnames []string

func Generate() string {
	firstNamesFile, fnerr := ioutil.ReadFile("data/first-names.txt")
	lastNamesFile, lnerr := ioutil.ReadFile("data/last-names.txt")
	if fnerr != nil {
		log.Fatal(fnerr)
	}
	if lnerr != nil {
		log.Fatal(lnerr)
	}

	firstnames := fileExtract(string(firstNamesFile))
	//for _, name := range firstnames {
	//	fmt.Println("New name: " + name + " has been registered")
	//}

	lastnames := fileExtract(string(lastNamesFile))
	//for _, name := range lastnames {
	//	fmt.Println("New lastname: " + name + " has been registered")
	//}

	return firstnames[random.RangeInt(0, len(firstnames)-1, 1)[0]] + " " + lastnames[random.RangeInt(0, len(lastnames)-1, 1)[0]]
}

func fileExtract(content string) []string {
	return strings.Split(content, "\r\n")
}
