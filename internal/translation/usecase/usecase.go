package usecase

import (
	"context"

	"github.com/amagkn/translabor/internal/translation/dto"
	"github.com/amagkn/translabor/internal/translation/entity"
)

type LingvaAPI interface {
	Translate(input dto.TranslateInput) (string, error)
}

type Postgres interface {
	SelectWord(ctx context.Context, word string) (entity.WordWithTranslation, error)
	InsertWord(ctx context.Context, input dto.SaveWordInput) (entity.WordWithTranslation, error)
}

type UseCase struct {
	lingvaAPI LingvaAPI
	postgres  Postgres
}

func New(l LingvaAPI, p Postgres) *UseCase {
	return &UseCase{
		lingvaAPI: l,
		postgres:  p,
	}
}
