package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song-libraries/pkg/model"
)

func (r *Repository) UpdateSongDetails(detail model.SongDetail) (bool, error) {
	//var existingSongDetail model.SongDetail
	//err := r.db.Get(&existingSongDetail, "SELECT * FROM song_detail WHERE id=$1", detail.ID)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		return false, fmt.Errorf("song detail not found") // Ошибка, если задача не найдена
	//	}
	//	return false, err // Если возникла другая ошибка
	//}
	sql, args, err := sq.Update("song_detail").
		Set("text", detail.Text).
		Set("link", detail.Link).
		Where(sq.Eq{"id": detail.ID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return false, fmt.Errorf("failed to create update query: %w", err)
	}
	_, err = r.db.Exec(sql, args...)
	if err != nil {
		return false, fmt.Errorf("failed to execute update query: %w", err)
	}

	return true, nil
}
