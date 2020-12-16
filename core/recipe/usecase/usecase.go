package usecase

import (
	"GOSecretProject/core/model/base"
	"GOSecretProject/core/recipe/interfaces"
)

type recipeUseCase struct {
	repository recipeInterfaces.RecipeRepository
}

func NewRecipeUseCase(repository recipeInterfaces.RecipeRepository) *recipeUseCase {
	return &recipeUseCase{repository: repository}
}

func (u *recipeUseCase) CreateRecipe(recipe *baseModels.Recipe) (err error) {
	return u.repository.CreateRecipe(recipe)
}

func (u *recipeUseCase) GetRecipe(id uint64) (recipe *baseModels.Recipe, err error) {
	return u.repository.GetRecipe(id)
}

func (u *recipeUseCase) GetRecipes(authorId uint64) (recipes []baseModels.Recipe, err error) {
	return u.repository.GetRecipes(authorId)
}

func (u *recipeUseCase) AddToFavorites(userId, recipeId uint64) (err error) {
	return u.repository.AddToFavorites(userId, recipeId)
}

func (u *recipeUseCase) DeleteFromFavorites(userId, recipeId uint64) (err error) {
	return u.repository.DeleteFromFavorites(userId, recipeId)
}

func (u *recipeUseCase) GetFavorites(userId uint64) (recipes []baseModels.Recipe, err error) {
	return u.repository.GetFavorites(userId)
}

func (u *recipeUseCase) VoteRecipe(userId, recipeId, stars uint64) (rating float64, err error) {
	return u.repository.VoteRecipe(userId, recipeId, stars)
}

func (u *recipeUseCase) FindRecipes(searchString string, page, userId uint64) (result *baseModels.SearchResult, err error) {
	params := baseModels.SearchParams{
		Text: searchString,
		Page: page,
	}

	return u.repository.FindRecipes(params, userId)
}
