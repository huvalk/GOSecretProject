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
}
