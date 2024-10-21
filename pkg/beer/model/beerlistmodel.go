package model


type (
	Beer struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Type        string  `json:"type"`
		Description string `json:"description"`
		Picture     string `json:"picture"`
	}

	BeerFilter struct {
		Name 		string `query:"name"`
		Description string `query:"description"`
		Paginate
	}

	Paginate struct {
		Page int64 `query:"page"`
		Size int64 `query:"size"`
	}

	Result struct {
		Beer    []*Beer         `json:"beer"`
		Paginate PaginateResult `json:"paginate"`
	}

	PaginateResult struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
	}
)	


