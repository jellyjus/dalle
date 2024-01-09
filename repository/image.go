package repository

import (
	"context"

	"dalle/entity"

	"github.com/doug-martin/goqu/v9"
)

func (r *Repository) AddImage(ctx context.Context, image entity.Image) error {
	q, args, err := r.builder.
		Insert("images").
		Rows(goqu.Record{
			"prompt":    image.Prompt,
			"image_url": image.URL,
			"status":    "created",
		}).
		ToSQL()
	if err != nil {
		return err
	}
	if _, err := r.conn.Exec(ctx, q, args...); err != nil {
		return err
	}
	return nil
}
