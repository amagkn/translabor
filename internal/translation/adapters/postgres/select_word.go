package postgres

import (
	"context"
	"errors"

	"github.com/amagkn/translabor/internal/translation/entity"
	"github.com/amagkn/translabor/pkg/base_errors"
	"github.com/amagkn/translabor/pkg/postgres"
	"github.com/doug-martin/goqu/v9"
)

func (p *Postgres) SelectWord(ctx context.Context, word string) (entity.WordWithTranslation, error) {
	var output entity.WordWithTranslation

	ds := goqu.
		From("word").
		Select("word", "translation").
		Where(goqu.Ex{"word": word}).
		Limit(1)

	sql, _, err := ds.ToSQL()
	if err != nil {
		return output, base_errors.WithPath("ds.ToSql", err)
	}

	err = p.pool.QueryRow(ctx, sql).Scan(&output.Word, &output.Translation)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return output, entity.ErrWordDoesNotExist
		}

		return output, base_errors.WithPath("p.pool.QueryRow.Scan", err)
	}

	return output, nil
}
