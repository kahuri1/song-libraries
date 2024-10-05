package repository

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song-libraries/pkg/model"
)

func (r *Repository) Library(c model.LibraryConfig) (model.SongsResponse, error) {
	var songs model.SongsResponse

	query := sq.
		Select("g.name AS group_name",
			"s.title AS song_title",
			"s.release_date",
			"sd.text AS song_text",
			"sd.link AS song_link").
		From("groups g").
		Join("songs s ON g.id = s.group_id").
		LeftJoin("song_detail sd ON s.id = sd.song_id")

	if c.SortTitle != "" {
		order := "ASC"
		if c.SortOrder == "desc" {
			order = "DESC"
		}
		query = query.OrderBy(c.SortTitle + " " + order)
	}
	offset := (c.Page - 1) * c.PageSize
	query = query.Limit(uint64(c.PageSize)).Offset(uint64(offset))

	checkSql, checkArgs, err := query.ToSql()
	if err != nil {
		return songs, fmt.Errorf("failed to build SQL query: %w", err)
	}
	rows, err := r.db.Query(checkSql, checkArgs...)
	if err != nil {
		return songs, fmt.Errorf("failed to execute query songs: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var song model.SongLibs
		if err := rows.Scan(&song.GroupName, &song.SongTitle, &song.ReleaseDate, &song.SongText, &song.SongLink); err != nil {
			return songs, fmt.Errorf("failed to scan song: %w", err)
		}
		songs.SongLibs = append(songs.SongLibs, song)
	}

	return songs, nil
}
