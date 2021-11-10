package main

import (
	"brendisurfs/go-practice-api/handlers"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

var DataFile map[string]handlers.Recipe

// Load our json file into RecipeData before anything else runs.
func init() {
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
	router := gin.Default()

	router.GET("/recipes")

	// finally, run our server.
	router.Run(":8080")
}
