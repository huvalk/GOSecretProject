package baseModels
coi
type Recipe struct {
	Id          uint64   `json:"id"`
	Author      uint64   `json:"author"`
	Title       string   `json:"title"`
	CookingTime uint64   `json:"cooking_time"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}
