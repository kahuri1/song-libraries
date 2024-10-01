package repository

import (
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/kahuri1/song-libraries/pkg/model"
)

func (r *Repository) PostAddSong(g model.Song) (int64, error) {
	var id int64

	sql, args, err := sq.
		Insert("songs").
		Columns("group_id", "title", "release_date").
		Values(g.GroupID, g.Title, g.Date).
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

func (r *Repository) CheckAddSong(g *model.Song) {
	var id int

	sql, args, err := sq.
		Insert("groups").
		Columns("name").
		Values(g.NameGroup).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		fmt.Errorf("failed to create name group creation request: %w", err)
	}

	err = r.db.QueryRow(sql, args...).Scan(&id)
	if err != nil {
		fmt.Errorf("failed to proccess news create query: %w", err)
	}
	g.GroupID = id
}

func (r *Repository) CheckDuplicateSong(g *model.Song) (int64, error) {
	var id int64
	checkSql, checkArgs, err := sq.
		Select("id").
		From("songs").
		Where("title = ?", g.Title).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, fmt.Errorf("failed to check query song: %w", err)
	}

	err = r.db.QueryRow(checkSql, checkArgs...).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to check query song: %w", err)
	}
	if id != 0 {
		return id, nil
	}
	return 0, fmt.Errorf("the group has already been created")
}
