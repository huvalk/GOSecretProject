package postgres

import (
	baseModels "GOSecretProject/core/model/base"
	"database/sql"
	"github.com/lib/pq"
)

type recipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *recipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) CreateRecipe(recipe *baseModels.Recipe) (err error) {
	query := "INSERT INTO recipe (user_id, title, cooking_time, ingredients, steps) VALUES ($1, $2, $3, $4, $5)"
	_, err = r.db.Exec(query, &recipe.Author, &recipe.Title, recipe.CookingTime,
		pq.Array(recipe.Ingredients), pq.Array(&recipe.Steps))
	if err != nil {
		return err
	}

	return nil
}

func (r *recipeRepository) GetRecipe(id uint64) (*baseModels.Recipe, error) {
	var recipe baseModels.Recipe

	query := "SELECT id, user_id, title, cooking_time, ingredients, steps FROM recipe WHERE id = $1"
	err := r.db.QueryRow(query, id).
		Scan(&recipe.Id, &recipe.Author, &recipe.Title, &recipe.CookingTime, &recipe.Ingredients, &recipe.Steps)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}
