package handlers

import (
	"brendisurfs/go-practice-api/recipes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response struct - a type to return for each handler.
type Response struct {
	Body   map[string]interface{}
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

	c.JSON(http.StatusOK, recipes.RecipeData["garlicPasta"])

}

// AddRecipe [POST] - adds an additional recipe to the backend.
func AddRecipe(c *gin.Context) {
	var addedRecipe recipes.Recipe
	// bind to recieve json and bind it to the file.
	if err := c.BindJSON(&addedRecipe); err != nil {
		return
	}

	// if all goes well, add the json.
	c.JSON(http.StatusOK, addedRecipe)
}

// UpdateRecipe [PUT] - updates an existing recipe
func UpdateRecipe(recipe *recipes.Recipe, c *gin.Context) (*Response, error) {

	return &Response{}, nil
}
