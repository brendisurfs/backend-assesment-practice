package handlers

import (
	"github.com/gin-gonic/gin"
)

// Recipe struct - structure of our recipe data.
type Recipe struct {
	Name         string            `json:"name"`
	Ingredients  map[string]string `json:"ingredients"`
	Instructions map[string]string `json:"instructions"`
}

// Response struct - a type to return for each handler.
type Response struct {
	Body   Recipe
	Status uint
}

// GetAllRecipes - returns all recipes to the request.
func GetAllRecipes(c *gin.Context) (*[]Recipe, error) {

	return nil, nil
}

// GetSingleRecipe - retrieves a single recipe from our data, using a string param.
func GetSingleRecipe(recipe string, c *gin.Context) (*Response, error) {

	return &Response{}, nil
}

// AddRecipe - adds an additional recipe to the backend.
func AddRecipe(recipe Recipe, c *gin.Context) (*Response, error) {

	return &Response{}, nil
}

// UpdateRecipe - updates an existing recipe
func UpdateRecipe(recipe Recipe, c *gin.Context) (*Response, error) {

	return &Response{}, nil
}
