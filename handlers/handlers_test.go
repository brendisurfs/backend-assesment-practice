package handlers

import (
	"brendisurfs/go-practice-api/recipes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var TestRecipeData map[string]recipes.Recipe

func GetJSONDataForTest() {
	content, err := os.Open("data.json")
	if err != nil {
		log.Fatal("could not read data.json, application will not work.", err)
	}
	defer content.Close()

	decoder := json.NewDecoder(content)
	decoder.Token()
	// bring in our map from recipe data
	for decoder.More() {
		decoder.Decode(&TestRecipeData)
	}
	fmt.Println("[OK] Data loaded.")
}

// Test for Get all recipes route | Should return: Body = all recipes, Status 200.
func Test_GetAllRecipes_ShouldReturnAllRecipes(t *testing.T) {

	GetJSONDataForTest()

	handler := func(c *gin.Context) {
		c.JSON(http.StatusOK, TestRecipeData)
	}

	router := gin.Default()
	router.GET("/recipes", handler)

	req, err := http.NewRequest("GET", "/recipes", nil)
	if err != nil {
		log.Fatal("could not get request /recipes", err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	// this should equal the full json recipe.
	assert.Equal(t, TestRecipeData, res.Body.String())
}

// test getting a single recipe from the server | Should return Recipe body, Status 200
func Test_GetSingleRecipe_ShouldReturnRecipeAndStatus200(t *testing.T) {

}

// test if the Update route, when correct input format | Should return: No body, status 204.
func Test_UpdateRecipe_ShouldReturnStatus204(t *testing.T) {

}
