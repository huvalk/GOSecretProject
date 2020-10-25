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

func (u *recipeUseCase) GetFavorites(userId uint64) (recipes []baseModels.Recipe, err error) {
	return u.repository.GetFavorites(userId)
}

func (u *recipeUseCase) VoteRecipe(userId, recipeId, stars uint64) (err error) {
	return u.repository.VoteRecipe(userId, recipeId, stars)
}

func (u *recipeUseCase) GetRating(recipeId uint64) (stars float64, err error) {
	return u.repository.GetRating(recipeId)
}
