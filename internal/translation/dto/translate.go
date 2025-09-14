package dto

type TranslateInput struct {
	Query  string `json:"query" validate:"required"`
	Source string `json:"source" validate:"required,min=2"`
	Target string `json:"target" validate:"required,min=2"`
}

type TranslateOutput struct {
	Translation string `json:"translation"`
}
