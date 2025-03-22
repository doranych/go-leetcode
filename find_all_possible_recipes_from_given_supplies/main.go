package find_all_possible_recipes_from_given_supplies

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	can := make(map[string]bool)
	toIdx := make(map[string]int)
	for _, supply := range supplies {
		can[supply] = true
	}
	for idx, recipe := range recipes {
		toIdx[recipe] = idx
	}

	var checkRecipe func(r string, visited map[string]bool, ingredients [][]string) bool
	checkRecipe = func(r string, visited map[string]bool, ingredients [][]string) bool {
		if can[r] {
			return true
		}
		if _, exists := toIdx[r]; !exists || visited[r] {
			return false
		}

		visited[r] = true
		canBeMade := true

		for _, ingredient := range ingredients[toIdx[r]] {
			if !checkRecipe(ingredient, visited, ingredients) {
				canBeMade = false
				break
			}
		}

		can[r] = canBeMade
		return can[r]
	}

	var res []string
	for _, recipe := range recipes {
		if checkRecipe(recipe, make(map[string]bool), ingredients) {
			res = append(res, recipe)
		}
	}

	return res
}
