package model

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Song struct {
	ID        int    `json:"id"`
	GroupID   int    `json:"group_id"`
	NameGroup string `json:"group"`
	Title     string `json:"title"`
	Date      string `json:"releaseDate"`
}

type SongDetail struct {
	ID     int    `json:"id"`
	SongID int    `json:"songId"`
	Text   string `json:"text"`
	Link   string `json:"link"`
}

type GetSongText struct {
	ID     int64  `json:"id"`
	SongID int64  `json:"songId"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Link   string `json:"link"`
}

type SongsResponse struct {
	SongLibs []SongLibs `json:"SongLibs"`
}

type LibraryConfig struct {
	SortTitle string `json:"sortTitle"`
	SortOrder string `json:"sortOrder"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
}

type SongLibs struct {
	GroupName   string `db:"group_name" json:"group_name"`
	SongTitle   string `db:"song_title" json:"song_title"`
	ReleaseDate string `db:"release_date" json:"release_date"`
	SongText    string `db:"song_text" json:"song_text"`
	SongLink    string `db:"song_link" json:"song_link"`
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

type BodyPagination struct {
	Lines []string `json:"lines"`
	Page  int      `json:"page"`
	Limit int      `json:"limit"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SongDetailInfo struct {
	ReleaseDate string `json:"releaseDate" db:"release_date"`
	Text        string `json:"text" db:"text"`
	Link        string `json:"link" db:"link"`
}

type RequestInfo struct {
	Group string `json:"group"`
	Name  string `json:"name"`
}
