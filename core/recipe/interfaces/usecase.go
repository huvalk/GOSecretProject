package recipeInterfaces

import (
	"GOSecretProject/core/model/base"
)

type RecipeUseCase interface {
	CreateRecipe(recipe *baseModels.Recipe) (recipeId uint64, photoPath string, err error)
	GetRecipe(id uint64) (recipe *baseModels.Recipe, err error)
	DeleteRecipe(id, userId uint64) (err error)
	GetRecipes(authorId uint64) (recipes []baseModels.Recipe, err error)
	AddToFavorites(userId, recipeId uint64) (err error)
	DeleteFromFavorites(userId, recipeId uint64) (err error)
	GetFavorites(userId uint64) (recipes []baseModels.Recipe, err error)
	VoteRecipe(userId, recipeId, stars uint64) (rating float64, err error)
	FindRecipes(searchString string, page, userId uint64) (result *baseModels.SearchResult, err error)
}
