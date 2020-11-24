package recipeHttp

import (
	"GOSecretProject/core/middleware"
	"GOSecretProject/core/recipe/interfaces"
	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoints(router *mux.Router, m *middleware.AuthMiddlewareHandler, useCase recipeInterfaces.RecipeUseCase) {
	h := NewRecipeHandler(useCase)

	router.HandleFunc("/recipe", m.UserRequired(h.CreateRecipe)).Methods("POST")
	router.HandleFunc("/recipe/{id:[0-9]+}", h.GetRecipe).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}/recipes", h.GetRecipes).Methods("GET")
	router.HandleFunc("/favorites/{id:[0-9]+}/add", m.UserRequired(h.AddToFavorites)).Methods("POST")
	router.HandleFunc("/favorites/{id:[0-9]+}/delete", m.UserRequired(h.DeleteFromFavorites)).Methods("POST")
	router.HandleFunc("/favorites", m.UserRequired(h.GetFavorites)).Methods("GET")
	router.HandleFunc("/recipe/{id:[0-9]+}/vote", m.UserRequired(h.VoteRecipe)).Methods("POST")
	router.HandleFunc("/search", m.UserRequired(h.FindRecipes)).Methods("GET")
}
