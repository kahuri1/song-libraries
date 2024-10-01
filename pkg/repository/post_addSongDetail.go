package repository

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song-libraries/pkg/model"
)

func (r *Repository) PostSongDetails(g model.SongDetail) (int64, error) {
	var id int64

	sql, args, err := sq.
		Insert("song_detail").
		Columns("song_id", "text", "link").
		Values(g.SongID, g.Text, g.Link).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to create song creation request: %w", err)
	}

	err = r.db.QueryRow(sql, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to proccess news create query: %w", err)
	}

	return id, nil
}

func (r *Repository) CheckDuplicateSongDetail(g model.SongDetail) (int64, error) {
	var id int64
	checkSql, checkArgs, err := sq.
		Select("id").
		From("song_detail").
		Where("song_id = ?", g.SongID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to check query song_detail: %w", err)
	}

	err = r.db.QueryRow(checkSql, checkArgs...).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to check query song_detail: %w", err)
	}
	if id != 0 {
		return id, nil
	}
	return 0, fmt.Errorf("the song_detail has already been created")
}
