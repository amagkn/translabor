package usecase

import "github.com/amagkn/translabor/internal/translation/dto"

type LingvaAPI interface {
	Translate(input dto.TranslateInput) (string, error)
}

type UseCase struct {
	lingvaAPI LingvaAPI
}

func New(l LingvaAPI) *UseCase {
	return &UseCase{lingvaAPI: l}
}
