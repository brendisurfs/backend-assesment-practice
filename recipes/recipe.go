package recipes

// Recipe struct - structure of our recipe data.
type Recipe struct {
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

type DataFile struct {
	Key []Recipe `json:"recipes"`
}

var RecipeData DataFile
