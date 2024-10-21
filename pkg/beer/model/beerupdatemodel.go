package model

type (
	BeerUpdategRequest struct {
		ID     		string `json:"id"`
		Name 		string `json:"name"`
		Type  		string  `json:"type"`
		Description string `json:"description"`
		Picture     string `json:"picture"`
	}

)
