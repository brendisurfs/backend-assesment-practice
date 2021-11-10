package handlers

// Recipe struct - structure of our recipe data.
type Recipe struct {
}

// Response struct - a type to return for each handler.
type Response struct {
	Body   Recipe
	Status uint
}

// GetAllRecipes - returns all recipes to the request.
func GetAllRecipes() ([]Recipe, error) {

	return nil, nil
}

// GetSingleRecipe - retrieves a single recipe from our data, using a string param.
func GetSingleRecipe(recipe string) (Response, error) {

	return Response{}, nil
}

// AddRecipe - adds an additional recipe to the backend.
func AddRecipe(recipe Recipe) (Response, error) {

	return Response{}, nil
}

// UpdateRecipe - updates an existing recipe
func UpdateRecipe(recipe Recipe) (Response, error) {

	return Response{}, nil
}
