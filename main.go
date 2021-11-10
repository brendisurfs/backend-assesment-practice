package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type recipeData struct {
	Name         string            `json:"name"`
	Ingredients  map[string]string `json:"ingredients"`
	Instructions map[string]string `json:"instructions"`
}

var DataFile map[string]recipeData

// init will load our json file into RecipeData before anything else runs.
func init() {
	// "db" should be loaded here, before anything else.
	content, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal("could not read data.json, application will not work.")
	}

	err = json.Unmarshal(content, &DataFile)
	if err != nil {
		log.Fatal("error: could not Unmarshal json to recipeData: ", err)
	}
}

func main() {
	// routes go here

}
