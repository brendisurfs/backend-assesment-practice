package main

import (
	"brendisurfs/go-practice-api/handlers"
	"brendisurfs/go-practice-api/recipes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Load our json file into RecipeData before anything else runs.
func init() {
	LoadJSONData()
}

// LoadJSONData - takes our data.json file and marshals it to a map.
func LoadJSONData() {
	content, err := os.ReadFile("data.json")
	if err != nil {
		log.Fatal("could not read data.json, application will not work.", err)
	}
	fmt.Println("loading json")

	err = json.Unmarshal(content, &recipes.RecipeData)
	if err != nil {
		log.Fatal("could not unmarshal json")
	}

	log.Println(content)

	fmt.Println("[OK] Data loaded.")
}

func main() {
	// routes go here
	router := gin.Default()

	router.GET("/recipes", handlers.GetAllRecipes)
	router.POST("/recipe/details/:id", handlers.GetSingleRecipe)

	// finally, run our server.

	log.Println("server running on localhost:8080/")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("could not run server")
	}
}
