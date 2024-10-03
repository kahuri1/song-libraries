package repository

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) SongID(title string) (int64, error) {

	var id int64
	checkSql, checkArgs, err := sq.
		Select("id").
		From("songs").
		Where("title = ?", title).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to check query songs: %w", err)
	}
	err = r.db.QueryRow(checkSql, checkArgs...).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to check query song_detail: %w", err)
	}

	return id, nil
}
