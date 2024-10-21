package model

type BeerCreatRequest struct {
	Name        string `json:"name"`
	Type        string  `json:"type"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}