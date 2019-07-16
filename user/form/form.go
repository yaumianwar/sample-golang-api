package form

type UserForm struct {
	Name     string `json:"name"`
	SkinType string `json:"skinType"`
	Age      uint64 `json:"age"`
}

type SkincareForm struct {
	Brand       string `json:"brand"`
	Name        string `json:"name"`
	ProductType string `json:"productType"`
	SkinConcern string `json:"skinConcern"`
}
