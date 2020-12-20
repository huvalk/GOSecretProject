package usecase

import (
	"GOSecretProject/core/model/base"
	"GOSecretProject/core/recipe/interfaces"
	"GOSecretProject/core/utils/sss"
	"encoding/base64"
)

type recipeUseCase struct {
	repository recipeInterfaces.RecipeRepository
}

func NewRecipeUseCase(repository recipeInterfaces.RecipeRepository) *recipeUseCase {
	return &recipeUseCase{repository: repository}
}

func (u *recipeUseCase) CreateRecipe(recipe *baseModels.Recipe) (recipeId uint64, photoPath string, err error) {
	b64Image := recipe.Photo
	recipe.Photo = "default"
	recipeId, err = u.repository.CreateRecipe(recipe)
	if err != nil {
		return 0, "", err
	}

	recipe.Id = recipeId
	unBased, err := base64.StdEncoding.DecodeString(b64Image)
	if err != nil {
		panic("Cannot decode b64")
	}
	recipe.Photo, err = sss.UploadPhoto(unBased, recipe.Id)
	if err != nil {
		return 0, "", err
	}
	err = u.repository.SavePhotoLink(recipe.Photo, recipe.Id)

	return recipe.Id, recipe.Photo, err
}

func (u *recipeUseCase) GetRecipe(id uint64) (recipe *baseModels.Recipe, err error) {
	return u.repository.GetRecipe(id)
}

func (u *recipeUseCase) DeleteRecipe(id, userId uint64) (err error) {
	return u.repository.DeleteRecipe(id, userId)
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
