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
