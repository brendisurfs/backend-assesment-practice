package recipes

// Recipe struct - structure of our recipe data.
type Recipe struct {
	Name         string            `json:"name"`
	Ingredients  map[string]string `json:"ingredients"`
	Instructions map[string]string `json:"instructions"`
}

// I dont like this interface usage, but it will have to do for now.
var RecipeData map[string]interface{}
