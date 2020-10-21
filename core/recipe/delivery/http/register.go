package recipeHttp

import (
	"GOSecretProject/core/recipe/interfaces"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHTTPEndpoints(router *mux.Router, useCase recipeInterfaces.RecipeUseCase) {
	h := NewRecipeHandler(useCase)

	router.Handle("/recipe", http.HandlerFunc(h.CreateRecipe)).Methods("POST")
	router.Handle("/recipe/{id:[0-9]+}", http.HandlerFunc(h.GetRecipe)).Methods("GET")
	router.Handle("/users/{id:[0-9]+}/recipes", http.HandlerFunc(h.GetRecipes)).Methods("GET")
}
