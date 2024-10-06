package repository

import (
	"github.com/kahuri1/song-libraries/pkg/model"
)

func (r *Repository) SongInfo(group string, song string) (model.SongDetailInfo, error) {
	var g model.SongDetailInfo

	query := `
        SELECT s.release_date, sd.text, sd.link
        FROM songs s
        JOIN groups g ON s.group_id = g.id
        JOIN song_detail sd ON s.id = sd.song_id
        WHERE g.name = $1 AND s.title = $2
    `
	err := r.db.Get(&g, query, group, song)
	if err != nil {
		return model.SongDetailInfo{}, err // Возвращаем пустую структуру в случае ошибки
	}
	return g, nil

}
