package postgres

import (
	"context"

	"github.com/amagkn/translabor/internal/translation/dto"
	"github.com/amagkn/translabor/internal/translation/entity"
	"github.com/amagkn/translabor/pkg/base_errors"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

func (p *Postgres) InsertWord(ctx context.Context, input dto.SaveWordInput) (entity.WordWithTranslation, error) {
	var output entity.WordWithTranslation

	ds := goqu.Insert("word").
		Rows(goqu.Record{
			"id":          uuid.New(),
			"word":        input.Word,
			"translation": input.Translation,
		}).
		Returning("word", "translation")

	sql, args, err := ds.ToSQL()
	if err != nil {
		return output, base_errors.WithPath("ds.ToSQL", err)
	}

	row := p.pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&output.Word, &output.Translation)
	if err != nil {
		return output, base_errors.WithPath("row.Scan", err)
	}

	return output, nil
}
