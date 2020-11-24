package baseModels

type Recipe struct {
	Id          uint64   `json:"id"`
	Author      uint64   `json:"author"`
	Title       string   `json:"title"`
	CookingTime uint64   `json:"cookingTime"`
	Rating      float64  `json:"rating"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
	IsFavorites bool     `json:"isFavorites"`
}
