package models

type Parameter struct {
	ApiKey string `json:"apikey"`
	ImdbID string `json:"imdbID"`
	Title  string `json:"title"`
	Page   string `json:"page"`
}
