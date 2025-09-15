package usecase

import (
	"context"
	"errors"

	"github.com/amagkn/translabor/internal/translation/dto"
	"github.com/amagkn/translabor/internal/translation/entity"
	"github.com/amagkn/translabor/pkg/base_errors"
)

func (u *UseCase) Translate(ctx context.Context, input dto.TranslateInput) (dto.TranslateOutput, error) {
	var output dto.TranslateOutput

	wordWithTranslation, err := u.postgres.SelectWord(ctx, input.Query)
	if err == nil {
		output.Translation = wordWithTranslation.Translation
		return output, nil
	} else if !errors.Is(err, entity.ErrWordDoesNotExist) {
		return output, base_errors.WithPath("u.postgres.SelectOneWord", err)
	}

	translation, err := u.lingvaAPI.Translate(input)
	if err != nil {
		return output, base_errors.WithPath("u.lingvaAPI.Translate", err)
	}

	saveWordInput := dto.SaveWordInput{
		Word:        input.Query,
		Translation: translation,
	}
	wordWithTranslation, err = u.postgres.InsertWord(ctx, saveWordInput)
	if err != nil {
		return output, base_errors.WithPath("u.postgres.InsertWord", err)
	}

	output.Translation = wordWithTranslation.Translation

	return output, nil
}
