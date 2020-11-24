package recipeHttp

import (
	"GOSecretProject/core/model/base"
	"GOSecretProject/core/recipe/interfaces"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"io/ioutil"
	"net/http"
	"strconv"
)

type recipeHandler struct {
	useCase recipeInterfaces.RecipeUseCase
}

func NewRecipeHandler(useCase recipeInterfaces.RecipeUseCase) *recipeHandler {
	return &recipeHandler{useCase: useCase}
}

func (h *recipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	authorId := r.Context().Value("userID").(uint64)

	recipeByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var recipe baseModels.Recipe
	err = json.Unmarshal(recipeByte, &recipe)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	recipe.Author = authorId

	err = h.useCase.CreateRecipe(&recipe)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *recipeHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	recipe, err := h.useCase.GetRecipe(id)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipeJson, err := json.Marshal(recipe)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(recipeJson)
}

func (h *recipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	authorIdString := mux.Vars(r)["id"]
	authorId, err := strconv.ParseUint(authorIdString, 10, 64)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	recipes, err := h.useCase.GetRecipes(authorId)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipesJson, err := json.Marshal(recipes)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(recipesJson)
}

func (h *recipeHandler) AddToFavorites(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(uint64)

	recipeIdString := mux.Vars(r)["id"]
	recipeId, err := strconv.ParseUint(recipeIdString, 10, 64)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = h.useCase.AddToFavorites(userId, recipeId)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *recipeHandler) DeleteFromFavorites(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(uint64)

	recipeIdString := mux.Vars(r)["id"]
	recipeId, err := strconv.ParseUint(recipeIdString, 10, 64)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = h.useCase.DeleteFromFavorites(userId, recipeId)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *recipeHandler) GetFavorites(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(uint64)

	recipes, err := h.useCase.GetFavorites(userId)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipesJson, err := json.Marshal(recipes)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(recipesJson)
}

func (h *recipeHandler) VoteRecipe(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(uint64)

	recipeIdString := mux.Vars(r)["id"]
	recipeId, err := strconv.ParseUint(recipeIdString, 10, 64)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	voteByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var starsMap map[string]uint64
	err = json.Unmarshal(voteByte, &starsMap)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	stars := starsMap["stars"]

	rating, err := h.useCase.VoteRecipe(userId, recipeId, stars)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	ratingJson, err := json.Marshal(rating)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(ratingJson)
}

func (h *recipeHandler) FindRecipes(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(uint64)

	searchString := r.FormValue("text")
	golog.Infof("text: %s", searchString)

	recipes, err := h.useCase.FindRecipes(searchString, userId)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	recipesJson, err := json.Marshal(recipes)
	if err != nil {
		golog.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(recipesJson)
}
