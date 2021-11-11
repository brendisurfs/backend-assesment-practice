package handlers

import (
	"brendisurfs/go-practice-api/recipes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response struct - a type to return for each handler.
type Response struct {
	Body   recipes.DataFile
	Status uint
}

// GetAllRecipes - returns all recipes to the request.
func GetAllRecipes(c *gin.Context) {

	res := &Response{
		Body:   recipes.RecipeData,
		Status: http.StatusOK,
	}

	c.JSON(http.StatusOK, res)
}

// GetSingleRecipe [GET] - retrieves a single recipe from our data, using a string param.
func GetSingleRecipe(c *gin.Context) {

	id := c.Param("name")
	recipeData := recipes.RecipeData.Key

	for _, rec := range recipeData {
		if rec.Name == id {
			c.JSON(http.StatusOK, rec)
		}
	}
}

// AddRecipe [POST] - adds an additional recipe to the backend.
// NOTE: if handle if the recipe already exists.
func AddRecipe(c *gin.Context) {
	var addedRecipe recipes.Recipe

	// bind to take the json and bind it to the file.
	if err := c.BindJSON(&addedRecipe); err != nil {
		c.JSON(400, c.Errors)
	}

	// if all goes well, add the json.
	recipes.RecipeData.Key = append(recipes.RecipeData.Key, addedRecipe)
	c.JSON(http.StatusOK, recipes.RecipeData)
}

// UpdateRecipe [PUT] - updates an existing recipe
func UpdateRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
