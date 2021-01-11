package proxyflow

import (
	"io/ioutil"
	"log"
	"os"
)

// ReadToken reads token from a txt file
func ReadToken(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	tocken := string(b)
	return tocken
}
