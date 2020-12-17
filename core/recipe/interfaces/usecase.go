package recipeInterfaces

import (
	"GOSecretProject/core/model/base"
	"mime/multipart"
)

type RecipeUseCase interface {
	CreateRecipe(recipe *baseModels.Recipe) (recipeId uint64, err error)
	UploadPhoto(form *multipart.Form, authorId uint64, recipeId uint64) error
	GetRecipe(id uint64) (recipe *baseModels.Recipe, err error)
	GetRecipes(authorId uint64) (recipes []baseModels.Recipe, err error)
	AddToFavorites(userId, recipeId uint64) (err error)
	DeleteFromFavorites(userId, recipeId uint64) (err error)
	GetFavorites(userId uint64) (recipes []baseModels.Recipe, err error)
	VoteRecipe(userId, recipeId, stars uint64) (rating float64, err error)
	FindRecipes(searchString string, page, userId uint64) (result *baseModels.SearchResult, err error)
}
