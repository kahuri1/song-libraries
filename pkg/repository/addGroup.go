package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song-libraries/pkg/model"
)

func (r *Repository) AddGroup(g model.Group) (int64, error) {
	var id int64

	sql, args, err := sq.
		Insert("groups").
		Columns("name").
		Values(g.Name).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to create name group creation request: %w", err)
	}

	err = r.db.QueryRow(sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to proccess news create query: %w", err)
	}

	return id, nil
}
