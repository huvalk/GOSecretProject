package recipeInterfaces

import baseModels "GOSecretProject/core/model/base"

type RecipeUseCase interface {
	CreateRecipe(recipe *baseModels.Recipe) (err error)
	GetRecipe(id uint64) (recipe *baseModels.Recipe, err error)
	GetRecipes(authorId uint64) (recipes []baseModels.Recipe, err error)
	AddToFavorites(userId, recipeId uint64) (err error)
}
