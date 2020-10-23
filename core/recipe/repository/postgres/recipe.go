package postgres

import (
	baseModels "GOSecretProject/core/model/base"
	"database/sql"
	"github.com/kataras/golog"
	"github.com/lib/pq"
)

type recipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *recipeRepository {
	return &recipeRepository{db: db}
}

func (r *recipeRepository) CreateRecipe(recipe *baseModels.Recipe) (err error) {
	var id uint64

	query := "INSERT INTO recipe (user_id, title, cooking_time, ingredients, steps)" +
		"VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err = r.db.QueryRow(query, &recipe.Author, &recipe.Title, &recipe.CookingTime,
		pq.Array(recipe.Ingredients), pq.Array(&recipe.Steps)).Scan(&id)
	if err != nil {
		return err
	}

	golog.Infof("Created recipe with id %d", id)
	return nil
}

func (r *recipeRepository) GetRecipe(id uint64) (*baseModels.Recipe, error) {
	var recipe baseModels.Recipe

	query := "SELECT id, user_id, title, cooking_time, ingredients, steps FROM recipe WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&recipe.Id, &recipe.Author, &recipe.Title, &recipe.CookingTime,
		pq.Array(&recipe.Ingredients), pq.Array(&recipe.Steps))
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (r *recipeRepository) GetRecipes(authorId uint64) (recipes []baseModels.Recipe, err error) {
	query := "SELECT id, user_id, title, cooking_time, ingredients, steps FROM recipe WHERE user_id = $1"
	rows, err := r.db.Query(query, authorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe baseModels.Recipe

		err = rows.Scan(&recipe.Id, &recipe.Author, &recipe.Title, &recipe.CookingTime,
			pq.Array(&recipe.Ingredients), pq.Array(&recipe.Steps))
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (r *recipeRepository) AddToFavorites(userId, recipeId uint64) (err error) {
	query := "INSERT INTO favorites (user_id, recipe_id) VALUES ($1, $2)"
	_, err = r.db.Exec(query, userId, recipeId)
	if err != nil {
		return err
	}

	return nil
}

func (r *recipeRepository) GetFavorites(userId uint64) (recipes []baseModels.Recipe, err error) {
	query := "SELECT r.id, r.user_id, r.title, r.cooking_time, r.ingredients, r.steps FROM favorites f" +
		"LEFT JOIN recipe r ON f.recipe_id = r.id WHERE f.user_id = $1"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe baseModels.Recipe

		err = rows.Scan(&recipe.Id, &recipe.Author, &recipe.Title, &recipe.CookingTime,
			pq.Array(&recipe.Ingredients), pq.Array(&recipe.Steps))
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}
