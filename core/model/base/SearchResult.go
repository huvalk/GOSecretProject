package baseModels

type SearchResult struct {
	Recipes     []Recipe `json:"recipes"`
	HasNextPage bool     `json:"hasNextPage"`
}
