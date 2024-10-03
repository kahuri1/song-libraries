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

type SongResponse struct {
	Song    Song         `json:"song"`
	Group   Group        `json:"group"`
	Details []SongDetail `json:"details"`
}

type SongsResponse struct {
	Songs []Song `json:"songs"`
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
