package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test for Get all recipes route | Should return: Body = all recipes, Status 200.
func Test_GetAllRecipes_ShouldReturnAllRecipes(t *testing.T) {

	testingServer := httptest.NewServer(GetAllRecipes())

	defer testingServer.Close()

	// need to make a request to our server with the right route
	res, getErr := http.Get("/recipes")
	if getErr != nil {
		t.Fatalf("got an error trying to test /recipes: %v", getErr)
	}
	if res.StatusCode != 200 {
		t.Fatalf("expectect status code 200, got %v", res.StatusCode)
	}

	val, ok := res.Header["Content-Type"]
	if !ok {
		t.Fatalf("expected content-type header to be set")
	}
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("expected application/json in content-type, got %s", val[0])
	}

}

// test getting a single recipe from the server | Should return Recipe body, Status 200
func Test_GetSingleRecipe_ShouldReturnRecipeAndStatus200(t *testing.T) {

}

// test if the Update route, when correct input format | Should return: No body, status 204.
func Test_UpdateRecipe_ShouldReturnStatus204(t *testing.T) {

}
