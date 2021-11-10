package main

import (
	"io/ioutil"
	"log"
)

func init() {
	// "db" should be loaded here, before anything else.
	_, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal("could not read data.json, application will not work.")
	}
}

func main() {
	// routes go here

}
