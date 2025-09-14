package usecase

import (
	"github.com/amagkn/translabor/internal/translation/dto"
	"github.com/amagkn/translabor/pkg/base_errors"
)

func (u *UseCase) Translate(input dto.TranslateInput) (string, error) {
	translation, err := u.lingvaAPI.Translate(input)
	if err != nil {
		return "", base_errors.WithPath("u.lingvaAPI.Translate", err)
	}

	return translation, nil
}
