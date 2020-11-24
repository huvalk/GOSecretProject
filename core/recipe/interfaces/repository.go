package recipeInterfaces

import "GOSecretProject/core/model/base"

type RecipeRepository interface {
	CreateRecipe(recipe *baseModels.Recipe) (err error)
	GetRecipe(id uint64) (recipe *baseModels.Recipe, err error)
	GetRecipes(authorId uint64) (recipes []baseModels.Recipe, err error)
	AddToFavorites(userId, recipeId uint64) (err error)
	DeleteFromFavorites(userId, recipeId uint64) (err error)
	GetFavorites(userId uint64) (recipes []baseModels.Recipe, err error)
	VoteRecipe(userId, recipeId, stars uint64) (rating float64, err error)
	FindRecipes(searchString string) (recipes []baseModels.Recipe, err error)
}
