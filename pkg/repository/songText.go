package repository

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song-libraries/pkg/model"
)

func (r *Repository) SongText(g model.GetSongText) (string, error) {
	var text string
	checkSql, checkArgs, err := sq.
		Select("text").
		From("song_detail").
		Where("song_id = ?", g.SongID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", fmt.Errorf("failed to check query songs: %w", err)
	}
	err = r.db.QueryRow(checkSql, checkArgs...).Scan(&text)
	if err != nil && err != sql.ErrNoRows {
		return "", fmt.Errorf("failed to check query song_detail: %w", err)
	}

	return text, nil

}
