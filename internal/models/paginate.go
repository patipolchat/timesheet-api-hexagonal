package models

type Paginate struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
